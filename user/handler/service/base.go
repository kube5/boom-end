package service

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/Mu-Exchange/Mu-End/common/utils/hexutil"
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type BaseService interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt
	SendTransaction(ctx context.Context, signedTx *types.Transaction) error
	NewAdminKeyedTransactorWithChainID(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error)
	TransactionExecute(executor func(tx *gorm.DB) error) error
	SwitchToNextAddr(opts *bind.TransactOpts)
	IsNetworkError(err error) bool
	GetMeta() Meta
}

// var AdminNonce *atomic.Uint64
// var AdminAddr common.Address

type Meta struct {
	AddrIdx   int
	ChainId   *big.Int
	Config    *config.Blockchain
	RpcClient *rpc.Client
	Cli       *ethclient.Client
	Logger    *logrus.Logger
}

type Base struct {
	repo.CommonComponents
	Meta
	db     *gorm.DB
	logger *logrus.Logger
}

// Meta implements BaseService
func (base *Base) GetMeta() Meta {
	return base.Meta
}

func NewBaseService(lc fx.Lifecycle, sd fx.Shutdowner, cfg repo.CommonComponents, db *db.DB) (BaseService, error) {
	blkconfig := cfg.Cfg.Blockchains[0]
	if len(blkconfig.RpcUrls) == 0 {
		return nil, fmt.Errorf("addrs for erc session wrapper is empty")
	}

	rpcClient, err := rpc.DialContext(context.Background(), blkconfig.RpcUrls[0])
	if err != nil {
		return nil, fmt.Errorf("dial node: %w", err)
	}

	cli := ethclient.NewClient(rpcClient)
	ctx := ensureContext(nil)
	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	base := &Base{
		CommonComponents: cfg,
		db:               db.DB,
		logger:           cfg.Logger,
		Meta: Meta{
			AddrIdx:   0,
			Config:    blkconfig,
			RpcClient: rpcClient,
			Cli:       cli,
			ChainId:   chainID,
		},
	}
	basic.RegisterHook(lc, sd, base)
	return base, nil
}

func (b *Base) Start(sd fx.Shutdowner) error {
	// adminAddr, err := b.walletService.GetAdminAddress(context.Background(), &pb.Empty{})
	// if err != nil {
	// 	return err
	// }
	// AdminAddr = common.HexToAddress(adminAddr.Address)
	// if AdminNonce == nil || AdminNonce.Load() == 0 {
	// 	nonce, err := b.Cli.NonceAt(context.Background(), common.HexToAddress(adminAddr.Address), nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	AdminNonce = atomic.NewUint64(nonce - 1)
	// }
	return nil
}

func (erc *Base) Stop(sd fx.Shutdowner) error {
	return nil
}

func (b *Base) TransactionExecute(executor func(tx *gorm.DB) error) error {
	return b.db.Transaction(func(tx *gorm.DB) error {
		return executor(tx)
	})
}

func (b *Base) SendTransaction(ctx context.Context, signedTx *types.Transaction) error {
	return b.Cli.SendTransaction(ctx, signedTx)
}

func (b *Base) SwitchToNextAddr(opts *bind.TransactOpts) {
	var err error
	if opts != nil {
		opts.Context = nil
	}

	for i := 0; i < len(b.Config.RpcUrls); i++ {
		b.AddrIdx++
		if b.AddrIdx == len(b.Config.RpcUrls) {
			b.AddrIdx = 0
		}

		b.logger.Warnf("try to switch to %s", b.Config.RpcUrls[b.AddrIdx])

		b.Cli, err = ethclient.Dial(b.Config.RpcUrls[b.AddrIdx])
		if err != nil {
			continue
		}

		b.logger.Infof("switch to %s successfully", b.Config.RpcUrls[b.AddrIdx])

		return
	}

	b.logger.Errorf("all addrs are not valid")
}

func (b *Base) BlockNumber(ctx context.Context) uint64 {
	var number uint64
	var err error

	if err := retry.Retry(func(attempt uint) error {
		number, err = b.Cli.BlockNumber(ctx)
		if err != nil {
			b.logger.Warnf("BlockNumber: %s", err.Error())

			if b.IsNetworkError(err) {
				b.SwitchToNextAddr(nil)
			}
		}
		return err
	}, strategy.Wait(2*time.Second)); err != nil {
		b.logger.Panic(err)
	}

	return number
}

func (b *Base) TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
	var receipt *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		receipt, err = b.Cli.TransactionReceipt(ctx, txHash)
		if err != nil {
			b.logger.Warnf("TransactionReceipt[%s]: %s", txHash.String(), err.Error())

			if b.IsNetworkError(err) {
				b.SwitchToNextAddr(nil)
			}
		}
		return err
	}, strategy.Wait(2*time.Second)); err != nil {
		b.logger.Panic(err)
	}

	return receipt
}

func (b *Base) IsNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout") ||
		strings.Contains(err.Error(), "502") ||
		strings.Contains(err.Error(), "404")
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}

func (b *Base) NewAdminKeyedTransactorWithChainID(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	gasPrice, err := b.GetMeta().Cli.SuggestGasPrice(ctx)
	b.logger.Infof("gasPrice:%s", gasPrice.String())
	if err != nil {
		return nil, err
	}
	privKey, err := crypto.ToECDSA(hexutil.Decode(b.Cfg.Keys.Admin))
	if err != nil {
		return nil, err
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, err
	}
	transactOpts.GasPrice = new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(2))
	return transactOpts, nil
}
