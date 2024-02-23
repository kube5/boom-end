package contract

import "github.com/ethereum/go-ethereum/core/types"

type BlockLogs struct {
	ChainId  uint64      `json:"chain_id"`
	Block    uint64      `json:"block"`
	Contacts []string    `json:"contacts"`
	Logs     []types.Log `json:"logs"` // may have logs in multiple blocks
}
