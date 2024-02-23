package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

type Processor interface {
	// HandleLog is a function that handles the given log.
	HandleLog(*types.Log) (interface{}, interface{}, error)

	// IsHandled is a function that returns true if the given log is handled.
	IsHandled(*types.Log) bool

	// GetLastHandledBlock is a function that returns the last handled block.
	GetLastHandledBlock() uint64

	// GetChainId is a function that returns the chain id.
	GetChainId() uint64

	// GetContract is a function that returns the contract address.
	GetContract() string

	// GetAbi is a function that returns the contract abi.
	GetAbi() *abi.ABI

	// UpdateFinishedBlockAndIndex is a function that updates the finished block and index.
	UpdateFinishedBlockAndIndex(block uint64, index uint64) error
}
