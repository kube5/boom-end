package service

import (
	"context"
	"math/big"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/game/handler/cache"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/Mu-Exchange/Mu-End/common/contract/staking"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	contract2 "github.com/Mu-Exchange/Mu-End/game/handler/contract"
	"github.com/Mu-Exchange/Mu-End/game/handler/dao"
)

type StakingService interface {
	UpdateStakingNextStartHeight(height uint64) error
	GetStakingNextStartHeight() (uint64, error)
}

var _ StakingService = (*Staking)(nil)

type Staking struct {
	BaseService
	repo.CommonComponents
	userService     proto.UserService
	staking         *contract2.StakingWrapper
	contractDao     dao.GameContractDao
	recordDao       dao.GameMonitorRecordDao
	gCache          cache.GameCache
	blockchain      *config.Blockchain
	contractConfig  *config.ContractConfig
	contractToProjM map[string]string
	ctx             context.Context
	cancel          context.CancelFunc
}

func NewStakingService(
	lc fx.Lifecycle,
	sd fx.Shutdowner,
	cfg repo.CommonComponents,
	userService proto.UserService,
	service BaseService,
	gCache cache.GameCache,
	contractDao dao.GameContractDao,
	recordDao dao.GameMonitorRecordDao,
) (StakingService, error) {
	StakingWrapper, err := contract2.NewStakingWrapper(cfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	staking := &Staking{
		BaseService:      service,
		staking:          StakingWrapper,
		CommonComponents: cfg,
		userService:      userService,
		recordDao:        recordDao,
		gCache:           gCache,
		contractDao:      contractDao,
		blockchain:       cfg.Cfg.Blockchains[0],
		contractToProjM:  make(map[string]string),
		ctx:              ctx,
		cancel:           cancel,
	}
	for _, contractCfg := range cfg.Cfg.Contracts {
		if contractCfg.ContractType != proto.ContractType_CtStaking {
			continue
		}
		staking.contractConfig = contractCfg
	}
	basic.RegisterHook(lc, sd, staking)
	return staking, nil
}

func (f *Staking) Start(sd fx.Shutdowner) error {
	go f.monitorStakingEvent()
	return nil
}

func (f *Staking) Stop(sd fx.Shutdowner) error {
	f.cancel()
	return nil
}

func (f *Staking) UpdateStakingNextStartHeight(height uint64) error {
	monitor := &dto.GameContract{
		ChainID:   f.staking.ChainId(),
		Contract:  common.HexToAddress(f.contractConfig.Address).String(),
		NextStart: height,
	}
	return f.contractDao.Update(monitor)
}

func (f *Staking) GetStakingNextStartHeight() (uint64, error) {
	chainId := f.staking.ChainId()
	address := common.HexToAddress(f.contractConfig.Address).String()
	monitor, err := f.contractDao.QueryByChainIdAndContract(chainId, address)
	if err != nil {
		return 0, err
	}
	if monitor == nil {
		monitor = &dto.GameContract{
			ChainID:   chainId,
			Contract:  address,
			NextStart: f.contractConfig.StartBlock,
		}
		if err := f.contractDao.Create(monitor); err != nil {
			f.Logger.Errorf("new Staking monitor %v error: %v", monitor, err)
			return 0, err
		}
	}

	return monitor.NextStart, nil
}

func (f *Staking) monitorStakingEvent() {
	f.Logger.Infof("start to monitor Staking contract %s on chain %d", f.staking.Address(), f.staking.ChainId())

	for {
		nextStart, err := f.filterEvents()
		if err != nil {
			f.Logger.Errorf("filterEvents err: %v, Staking contract %s", err, f.contractConfig.Address)
			continue
		}

		if err := f.subscribeEvents(nextStart, time.Minute*5); err != nil {
			f.Logger.Errorf("subscribeEvents err: %v, Staking contract %s", err, f.contractConfig.Address)
			continue
		}
	}
}

func (f *Staking) filterEvents() (uint64, error) {
	start, err := f.GetStakingNextStartHeight()
	if err != nil {
		f.Logger.Errorf("GetStakingNextStartHeight err: %v, contract %s", err, f.contractConfig.Address)
		return 0, err
	}
	latest := f.staking.BlockNumber(f.ctx) - f.blockchain.ConfirmBlock
	f.Logger.Infof("Staking contract %s  latest block num : %d", f.contractConfig.Address, latest)
	for {
		end := start + f.blockchain.FilterStep - 1
		if end > latest {
			return start, nil
		}
		//f.Logger.Infof("start to filter Staking contract %s events from block %d, end is %d", f.Cfg.Staking.Contract, start, end)
		if err := f.handleStakeSnapshotEvents(start, end); err != nil {
			f.Logger.Errorf("handleStakeSnapshotEvents, start: %d, end: %d, err: %v", start, end, err)
			time.Sleep(30 * time.Second)
			continue
		}

		if err := f.handleMintDiceEvents(start, end); err != nil {
			f.Logger.Errorf("handleMintDiceEvents, start: %d, end: %d, err: %v", start, end, err)
			time.Sleep(30 * time.Second)
			continue
		}

		f.Logger.WithFields(logrus.Fields{"start": start, "end": end}).Infof("filter staking events")

		if err := f.UpdateStakingNextStartHeight(end + 1); err != nil {
			f.Logger.Errorf("UpdateStakingNextStartHeight: %d err: %v", end+1, err)
			time.Sleep(30 * time.Second)
			continue
		}

		start = end + 1
	}
}

func (f *Staking) subscribeEvents(start uint64, duration time.Duration) error {
	f.Logger.Infof("start to subscribe Staking contract %s from block %d for %s", f.contractConfig.Address, start, duration.String())

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	openEvent := make(chan *staking.StakingStakeSnapshot)
	diceEvent := make(chan *staking.StakingMintDice)
	subTransfer, err := f.staking.WatchStakeSnapshot(&bind.WatchOpts{Start: &start, Context: f.ctx}, openEvent)
	sub2Transfer, err := f.staking.WatchMintDice(&bind.WatchOpts{Start: &start, Context: f.ctx}, diceEvent)
	if err != nil {
		f.Logger.Errorf("Staking.WatchTransfer err: %v", err)
		return err
	}

	for {
		select {
		case event := <-openEvent:
			if err := f.handleStakeSnapshotEvent(event); err != nil {
				f.Logger.Errorf("handleStakeSnapshotEvent err: %v", err)
				time.Sleep(30 * time.Second)
				return err
			}
		case event := <-diceEvent:
			if err := f.handleMintDiceEvent(event); err != nil {
				f.Logger.Errorf("handleMintDiceEvent err: %v", err)
				time.Sleep(30 * time.Second)
				return err
			}
		case <-f.ctx.Done():
			f.Logger.Infof("cancel subscribe Staking contract %s", f.contractConfig.Address)
			subTransfer.Unsubscribe()
			sub2Transfer.Unsubscribe()
			return nil
		case <-ticker.C:
			f.Logger.Infof("unsubscribe Staking contract %s", f.contractConfig.Address)
			subTransfer.Unsubscribe()
			sub2Transfer.Unsubscribe()
			return nil
		}
	}
}

func (f *Staking) handleStakeSnapshotEvents(start, end uint64) error {
	transferIterator := f.staking.FilterStakeSnapshot(&bind.FilterOpts{Start: start, End: &end, Context: f.ctx})
	for transferIterator.Next() {
		if err := f.handleStakeSnapshotEvent(transferIterator.Event); err != nil {
			return err
		}
	}
	return nil
}
func (f *Staking) handleMintDiceEvents(start, end uint64) error {
	transferIterator := f.staking.FilterMintDice(&bind.FilterOpts{Start: start, End: &end, Context: f.ctx})
	for transferIterator.Next() {
		if err := f.handleMintDiceEvent(transferIterator.Event); err != nil {
			return err
		}
	}
	return nil
}

func (f *Staking) handleStakeSnapshotEvent(event *staking.StakingStakeSnapshot) error {
	processed, err := isEventProcessed(f.recordDao, f.staking.ChainId(), &event.Raw)
	if err != nil {
		f.Logger.Errorf("handleStakeSnapshotEvent, isEventProcessed err: %v", err)
		return nil
	}
	if processed {
		f.Logger.Infof("handleStakeSnapshotEvent, tx: %s, index: %d has been processed",
			event.Raw.TxHash.Hex(), event.Raw.Index)
		return nil
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil)
	ethDice := new(big.Int).Div(new(big.Int).Mul(event.TotalEthAmount, big.NewInt(30)), divisor)
	usdbDice := new(big.Int).Div(event.TotalUsdbAmount, divisor)
	dice := ethDice.Int64() + usdbDice.Int64()
	f.Logger.Infof("handleStakeSnapshotEvent, tx: %s, trader: %s,token:%s, amount:%s, totalEth:%s,totalUsdb:%s, dice: %d",
		event.Raw.TxHash.Hex(), event.Staker.Hex(), event.Token.Hex(), event.Amount.String(), event.TotalEthAmount.String(), event.TotalUsdbAmount.String(), dice)

	_, err = f.userService.UpdateUser(context.Background(), &proto.UpdateUserReq{
		Wallet:          event.Staker.Hex(),
		EthStakeAmount:  event.TotalEthAmount.String(),
		UsdbStakeAmount: event.TotalUsdbAmount.String(),
		Dice:            dice,
	})
	if err != nil {
		f.Logger.Errorf("handleStakeSnapshotEventErr:%v", err)
		return nil
	}
	err = f.gCache.SetUserStake(event.Staker.Hex())
	if err != nil {
		f.Logger.Errorf("handleStakeSnapshotEventErr:%v", err)
		return nil
	}
	record := genRecordByEvent(f.staking.ChainId(), &event.Raw)
	if err := f.recordDao.Create(record); err != nil {
		return err
	}
	return nil
}

func (f *Staking) handleMintDiceEvent(event *staking.StakingMintDice) error {
	processed, err := isEventProcessed(f.recordDao, f.staking.ChainId(), &event.Raw)
	if err != nil {
		f.Logger.Errorf("handleStakeSnapshotEvent, isEventProcessed err: %v", err)
		return nil
	}
	if processed {
		f.Logger.Infof("handleStakeSnapshotEvent, tx: %s, index: %d has been processed",
			event.Raw.TxHash.Hex(), event.Raw.Index)
		return nil
	}

	f.Logger.Infof("handleMintDiceEvent, tx: %s, trader: %s", event.Raw.TxHash.Hex(), event.Staker.Hex())
	_, err = f.userService.UpdateUser(context.Background(), &proto.UpdateUserReq{
		Wallet:   event.Staker.Hex(),
		MintDice: true,
	})
	if err != nil {
		f.Logger.Errorf("handleMintDiceEventErr:%v", err)
	}
	record := genRecordByEvent(f.staking.ChainId(), &event.Raw)
	if err := f.recordDao.Create(record); err != nil {
		return err
	}
	return nil
}
