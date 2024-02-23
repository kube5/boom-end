package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrBlockchainNotFound      = errors.New("error blockchain not found")
	ErrRpcUrlNotSet            = errors.New("error rpc url not set")
	ErrDialBlockchainRpc       = errors.New("error dial blockchain rpc")
	ErrInvalidContractType     = errors.New("error invalid contract type")
	ErrAllRPCFailed            = errors.New("error all rpc failed")
	ErrInvalidChainId          = errors.New("error invalid chain id")
	ErrInvalidStartBlock       = errors.New("error invalid start block")
	ErrDuplicateChannel        = errors.New("error duplicate channel")
	ErrMonitorHasPublish       = errors.New("error monitor has publish")
	ErrCanNotGetChildContract  = errors.New("error cannot get child contract")
	ErrUnknownContractType     = errors.New("error unknown contract type")
	ErrStopMonitorNotAllowed   = errors.New("error stop monitor not allowed")
	ErrMonitorStartBlockTooLow = errors.New("error monitor start block too low")
)
