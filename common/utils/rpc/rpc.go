package rpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"github.com/Mu-Exchange/Mu-End/common/errors"
)

var _ EthClient = (*FallbackClient)(nil)

type EthClient interface {
	bind.ContractCaller
	FilterLogs(ctx context.Context, q *ethereum.FilterQuery) ([]types.Log, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	MustGetHeaderByNumber(ctx context.Context, number *big.Int) *types.Header
	TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error)
	MustGetTransactionReceipt(ctx context.Context, hash common.Hash) *types.Receipt
}

type FallbackClient struct {
	clients []*ethclient.Client
	current int
	logger  *logrus.Logger
	chainId uint64
}

func (fc *FallbackClient) MustGetTransactionReceipt(ctx context.Context, hash common.Hash) *types.Receipt {
	for i := 0; ; i++ {
		client := fc.clients[fc.current]

		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil {
			return receipt
		}
		fc.logger.Warnf("get transaction receipt failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)

		time.Sleep(1 * time.Second)
	}
}

func (fc *FallbackClient) MustGetHeaderByNumber(ctx context.Context, number *big.Int) *types.Header {
	for i := 0; ; i++ {
		client := fc.clients[fc.current]

		header, err := client.HeaderByNumber(ctx, number)
		if err == nil {
			return header
		}
		fc.logger.Warnf("get header by number failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)

		time.Sleep(1 * time.Second)
	}
}

func (fc *FallbackClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		code, err := client.CodeAt(ctx, contract, blockNumber)
		if err == nil {
			return code, nil
		}
		fc.logger.Warnf("get code failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		code, err := client.CallContract(ctx, call, blockNumber)
		if err == nil {
			return code, nil
		}
		fc.logger.Warnf("call contract failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func NewFallbackClient(chainId uint64, nodeURLs []string, logger *logrus.Logger) (*FallbackClient, error) {
	clients := make([]*ethclient.Client, len(nodeURLs))
	for i, url := range nodeURLs {
		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, err
		}
		clients[i] = client
	}
	return &FallbackClient{clients: clients, current: 0, chainId: chainId, logger: logger}, nil
}

func (fc *FallbackClient) FilterLogs(ctx context.Context, q *ethereum.FilterQuery) ([]types.Log, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		logs, err := client.FilterLogs(ctx, *q)
		if err == nil {
			return logs, nil
		}
		fc.logger.Warnf("filter logs failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		sub, err := client.SubscribeNewHead(ctx, ch)
		if err == nil {
			return sub, nil
		}
		fc.logger.Warnf("subscribe new head failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		sub, err := client.SubscribeFilterLogs(ctx, q, ch)
		if err == nil {
			return sub, nil
		}
		fc.logger.Warnf("subscribe filter logs failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) BlockNumber(ctx context.Context) (uint64, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		number, err := client.BlockNumber(ctx)
		if err == nil {
			return number, nil
		}
		fc.logger.Warnf("get block number failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return 0, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		block, err := client.BlockByNumber(ctx, number)
		if err == nil {
			return block, nil
		}
		fc.logger.Warnf("get block by number failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		header, err := client.HeaderByNumber(ctx, number)
		if err == nil {
			return header, nil
		}
		fc.logger.Warnf("get header by number failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed
}

func (fc *FallbackClient) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	for i := 0; i < len(fc.clients); i++ {
		client := fc.clients[fc.current]

		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil {
			return receipt, nil
		}
		fc.logger.Warnf("get transaction receipt failed, chainId: %d client: %d, err: %v", fc.chainId, fc.current, err)
		fc.current = (fc.current + 1) % len(fc.clients)
	}

	return nil, errors.ErrAllRPCFailed

}
