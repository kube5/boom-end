package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func UnpackLog(contractAbi *abi.ABI, out interface{}, event string, log types.Log) error {
	// Anonymous events are not supported.
	if len(log.Topics) == 0 {
		return fmt.Errorf("abi: cannot find event signature in log")
	}
	if log.Topics[0] != contractAbi.Events[event].ID {
		return fmt.Errorf("abi: wrong event signature in log")
	}

	if len(log.Data) > 0 {
		if err := contractAbi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range contractAbi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
