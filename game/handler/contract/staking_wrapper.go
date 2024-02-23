package contract

import (
	"context"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/sirupsen/logrus"

	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/Mu-Exchange/Mu-End/common/contract/staking"
	"github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/repo"
)

type StakingWrapper struct {
	repo.CommonComponents
	index      int
	address    string
	chainId    uint64
	ethClient  *ethclient.Client
	blockchain *config.Blockchain
	contract   *staking.Staking
	Session    *staking.StakingSession
	logger     logrus.FieldLogger
}

func NewStakingWrapper(cfg repo.CommonComponents) (*StakingWrapper, error) {
	logger := cfg.Logger
	if len(cfg.Cfg.Blockchains[0].RpcUrls) == 0 {
		logger.Errorf("blockchain rpc url is empty")
		return nil, errors.ErrRpcUrlNotSet
	}
	address := ""
	for _, contractCfg := range cfg.Cfg.Contracts {
		if contractCfg.ContractType != proto.ContractType_CtStaking {
			continue
		}
		address = contractCfg.Address
	}
	rpcClient, index, err := dialRpcWithRetry(cfg.Cfg.Blockchains[0].WssUrls, logger)
	if err != nil {
		logger.Errorf("dial rpc url %v err: %v", cfg.Cfg.Blockchains[0].RpcUrls, err)
		return nil, err
	}

	etherCli := ethclient.NewClient(rpcClient)
	contract, err := staking.NewStaking(common.HexToAddress(address), etherCli)
	if err != nil {
		logger.Errorf("failed to instantiate Staking contract, address: %s, err: %v", address, err)
		return nil, err
	}

	session := &staking.StakingSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
	}
	chainId, err := etherCli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &StakingWrapper{
		index:      index,
		chainId:    chainId.Uint64(),
		address:    common.HexToAddress(address).String(),
		blockchain: cfg.Cfg.Blockchains[0],
		ethClient:  etherCli,
		contract:   contract,
		Session:    session,
		logger:     logger,
	}, nil
}

func (aw *StakingWrapper) ChainId() uint64 {
	return aw.chainId
}

func (aw *StakingWrapper) BlockNumber(ctx context.Context) uint64 {
	var number uint64
	var err error

	if err := retry.Retry(func(attempt uint) error {
		number, err = aw.ethClient.BlockNumber(ctx)
		if err != nil {
			aw.logger.Warnf("BlockNumber: %s", err.Error())

			if isNetworkError(err) {
				aw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		aw.logger.Panic(err)
	}

	return number
}

func (aw *StakingWrapper) FilterStakeSnapshot(opts *bind.FilterOpts) *staking.StakingStakeSnapshotIterator {
	var (
		iterator *staking.StakingStakeSnapshotIterator
		err      error
	)

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = aw.contract.FilterStakeSnapshot(opts)
		if err != nil {
			aw.logger.Warnf("FilterTransfer err: %s", err.Error())

			if isNetworkError(err) {
				aw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		aw.logger.Panic(err)
	}

	return iterator
}

func (aw *StakingWrapper) FilterMintDice(opts *bind.FilterOpts) *staking.StakingMintDiceIterator {
	var (
		iterator *staking.StakingMintDiceIterator
		err      error
	)

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = aw.contract.FilterMintDice(opts)
		if err != nil {
			aw.logger.Warnf("FilterTransfer err: %s", err.Error())

			if isNetworkError(err) {
				aw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		aw.logger.Panic(err)
	}

	return iterator
}

func (aw *StakingWrapper) WatchStakeSnapshot(opts *bind.WatchOpts, sink chan *staking.StakingStakeSnapshot) (event.Subscription, error) {
	return aw.contract.WatchStakeSnapshot(opts, sink)
}

func (aw *StakingWrapper) WatchMintDice(opts *bind.WatchOpts, sink chan *staking.StakingMintDice) (event.Subscription, error) {
	return aw.contract.WatchMintDice(opts, sink)
}

func (aw *StakingWrapper) switchToNextAddr() {
	var err error

	for i := 0; i < len(aw.blockchain.RpcUrls); i++ {
		aw.index++
		if aw.index == len(aw.blockchain.RpcUrls) {
			aw.index = 0
		}

		aw.logger.Warnf("try to switch to %s for Staking contract", aw.blockchain.RpcUrls[aw.index])

		aw.ethClient, err = ethclient.Dial(aw.blockchain.RpcUrls[aw.index])
		if err != nil {
			continue
		}

		aw.contract, err = staking.NewStaking(common.HexToAddress(aw.address), aw.ethClient)
		if err != nil {
			continue
		}

		aw.Session.Contract = aw.contract

		aw.logger.Infof("switch to %s successfully", aw.blockchain.RpcUrls[aw.index])

		return
	}

	panic("all blockchain rpc urls are not valid")
}

func (aw *StakingWrapper) Close() {
	aw.ethClient.Close()
}

func (aw *StakingWrapper) Address() string {
	return aw.address
}
