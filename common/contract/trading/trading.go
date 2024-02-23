// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trading

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OpenRequest is an auto generated low-level Go binding around an user-defined struct.
type OpenRequest struct {
	TradeType  uint8
	Base       TradeBase
	CallTarget common.Address
	MaxPrice   *big.Int
	MinPrice   *big.Int
	Time       *big.Int
}

// OpenTrade is an auto generated low-level Go binding around an user-defined struct.
type OpenTrade struct {
	Base           TradeBase
	OpenPrice      *big.Int
	LastUpdateTime *big.Int
}

// TradeBase is an auto generated low-level Go binding around an user-defined struct.
type TradeBase struct {
	Trader    common.Address
	PairIndex *big.Int
	Margin    *big.Int
	Long      bool
	Leverage  *big.Int
	Tp        *big.Int
	Sl        *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"resut\",\"type\":\"bool\"}],\"name\":\"Callback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"CancelOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_closeMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rolloverFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closeFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCloseType\",\"name\":\"s\",\"type\":\"uint8\"}],\"name\":\"Close\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ExecRequestClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ExecRequestOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"base\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOpenTrade\",\"name\":\"t\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Open\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_closeMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requstTime\",\"type\":\"uint256\"}],\"name\":\"RequestClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumTradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"base\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOpenRequest\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"RequestOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tradingStorage\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"name\":\"SetContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_forKeeper\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_forCallback\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"SetNativeFeeForKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidationP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spreadReductionP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minAcceptanceDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tradeSwitch\",\"type\":\"uint256\"}],\"name\":\"SetParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"SetReserve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"TradeClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"}],\"name\":\"UpdateMargin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"name\":\"UpdateOpenRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"name\":\"UpdateTPAndSL\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callbackGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOpen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"closeTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeHelper\",\"outputs\":[{\"internalType\":\"contractIFeeHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"name\":\"getPriceAfterImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getTradeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractITradingStorage\",\"name\":\"_storageT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"link\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAcceptanceDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeForCallback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeForKeeper\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"openTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairInfo\",\"outputs\":[{\"internalType\":\"contractIPairInfos\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_closeMargin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"}],\"name\":\"requestClose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTradeType\",\"name\":\"_t\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"_base\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"}],\"name\":\"requestOpenLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"_base\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"}],\"name\":\"requestOpenMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tradingStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_forKeeper\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_forCallback\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"setNativeFeeForKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidationP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAcceptanceDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tradeSwitch\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadReductionP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"enumCloseType\",\"name\":\"_stop\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"stopTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageT\",\"outputs\":[{\"internalType\":\"contractITradingStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeSwitch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isAdd\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"updateMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"name\":\"updateOpenRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"updateTPAndSL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_Contract *ContractCaller) CallbackGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "callbackGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_Contract *ContractSession) CallbackGasLimit() (*big.Int, error) {
	return _Contract.Contract.CallbackGasLimit(&_Contract.CallOpts)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_Contract *ContractCallerSession) CallbackGasLimit() (*big.Int, error) {
	return _Contract.Contract.CallbackGasLimit(&_Contract.CallOpts)
}

// FeeHelper is a free data retrieval call binding the contract method 0x18897bb7.
//
// Solidity: function feeHelper() view returns(address)
func (_Contract *ContractCaller) FeeHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "feeHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeHelper is a free data retrieval call binding the contract method 0x18897bb7.
//
// Solidity: function feeHelper() view returns(address)
func (_Contract *ContractSession) FeeHelper() (common.Address, error) {
	return _Contract.Contract.FeeHelper(&_Contract.CallOpts)
}

// FeeHelper is a free data retrieval call binding the contract method 0x18897bb7.
//
// Solidity: function feeHelper() view returns(address)
func (_Contract *ContractCallerSession) FeeHelper() (common.Address, error) {
	return _Contract.Contract.FeeHelper(&_Contract.CallOpts)
}

// GetPriceAfterImpact is a free data retrieval call binding the contract method 0xa53fcde8.
//
// Solidity: function getPriceAfterImpact(uint256 price, uint256 pairIndex, uint256 position, bool isLong) view returns(uint256)
func (_Contract *ContractCaller) GetPriceAfterImpact(opts *bind.CallOpts, price *big.Int, pairIndex *big.Int, position *big.Int, isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPriceAfterImpact", price, pairIndex, position, isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceAfterImpact is a free data retrieval call binding the contract method 0xa53fcde8.
//
// Solidity: function getPriceAfterImpact(uint256 price, uint256 pairIndex, uint256 position, bool isLong) view returns(uint256)
func (_Contract *ContractSession) GetPriceAfterImpact(price *big.Int, pairIndex *big.Int, position *big.Int, isLong bool) (*big.Int, error) {
	return _Contract.Contract.GetPriceAfterImpact(&_Contract.CallOpts, price, pairIndex, position, isLong)
}

// GetPriceAfterImpact is a free data retrieval call binding the contract method 0xa53fcde8.
//
// Solidity: function getPriceAfterImpact(uint256 price, uint256 pairIndex, uint256 position, bool isLong) view returns(uint256)
func (_Contract *ContractCallerSession) GetPriceAfterImpact(price *big.Int, pairIndex *big.Int, position *big.Int, isLong bool) (*big.Int, error) {
	return _Contract.Contract.GetPriceAfterImpact(&_Contract.CallOpts, price, pairIndex, position, isLong)
}

// GetTradeValue is a free data retrieval call binding the contract method 0x72534890.
//
// Solidity: function getTradeValue(uint256 orderId, uint256 price) view returns(uint256)
func (_Contract *ContractCaller) GetTradeValue(opts *bind.CallOpts, orderId *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTradeValue", orderId, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeValue is a free data retrieval call binding the contract method 0x72534890.
//
// Solidity: function getTradeValue(uint256 orderId, uint256 price) view returns(uint256)
func (_Contract *ContractSession) GetTradeValue(orderId *big.Int, price *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetTradeValue(&_Contract.CallOpts, orderId, price)
}

// GetTradeValue is a free data retrieval call binding the contract method 0x72534890.
//
// Solidity: function getTradeValue(uint256 orderId, uint256 price) view returns(uint256)
func (_Contract *ContractCallerSession) GetTradeValue(orderId *big.Int, price *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetTradeValue(&_Contract.CallOpts, orderId, price)
}

// LiquidationP is a free data retrieval call binding the contract method 0x615728f8.
//
// Solidity: function liquidationP() view returns(uint256)
func (_Contract *ContractCaller) LiquidationP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "liquidationP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationP is a free data retrieval call binding the contract method 0x615728f8.
//
// Solidity: function liquidationP() view returns(uint256)
func (_Contract *ContractSession) LiquidationP() (*big.Int, error) {
	return _Contract.Contract.LiquidationP(&_Contract.CallOpts)
}

// LiquidationP is a free data retrieval call binding the contract method 0x615728f8.
//
// Solidity: function liquidationP() view returns(uint256)
func (_Contract *ContractCallerSession) LiquidationP() (*big.Int, error) {
	return _Contract.Contract.LiquidationP(&_Contract.CallOpts)
}

// MinAcceptanceDelay is a free data retrieval call binding the contract method 0xf485ccf0.
//
// Solidity: function minAcceptanceDelay() view returns(uint256)
func (_Contract *ContractCaller) MinAcceptanceDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minAcceptanceDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAcceptanceDelay is a free data retrieval call binding the contract method 0xf485ccf0.
//
// Solidity: function minAcceptanceDelay() view returns(uint256)
func (_Contract *ContractSession) MinAcceptanceDelay() (*big.Int, error) {
	return _Contract.Contract.MinAcceptanceDelay(&_Contract.CallOpts)
}

// MinAcceptanceDelay is a free data retrieval call binding the contract method 0xf485ccf0.
//
// Solidity: function minAcceptanceDelay() view returns(uint256)
func (_Contract *ContractCallerSession) MinAcceptanceDelay() (*big.Int, error) {
	return _Contract.Contract.MinAcceptanceDelay(&_Contract.CallOpts)
}

// NativeForCallback is a free data retrieval call binding the contract method 0x6bc81aa2.
//
// Solidity: function nativeForCallback() view returns(uint256)
func (_Contract *ContractCaller) NativeForCallback(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nativeForCallback")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NativeForCallback is a free data retrieval call binding the contract method 0x6bc81aa2.
//
// Solidity: function nativeForCallback() view returns(uint256)
func (_Contract *ContractSession) NativeForCallback() (*big.Int, error) {
	return _Contract.Contract.NativeForCallback(&_Contract.CallOpts)
}

// NativeForCallback is a free data retrieval call binding the contract method 0x6bc81aa2.
//
// Solidity: function nativeForCallback() view returns(uint256)
func (_Contract *ContractCallerSession) NativeForCallback() (*big.Int, error) {
	return _Contract.Contract.NativeForCallback(&_Contract.CallOpts)
}

// NativeForKeeper is a free data retrieval call binding the contract method 0xb0e4630d.
//
// Solidity: function nativeForKeeper() view returns(uint256)
func (_Contract *ContractCaller) NativeForKeeper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nativeForKeeper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NativeForKeeper is a free data retrieval call binding the contract method 0xb0e4630d.
//
// Solidity: function nativeForKeeper() view returns(uint256)
func (_Contract *ContractSession) NativeForKeeper() (*big.Int, error) {
	return _Contract.Contract.NativeForKeeper(&_Contract.CallOpts)
}

// NativeForKeeper is a free data retrieval call binding the contract method 0xb0e4630d.
//
// Solidity: function nativeForKeeper() view returns(uint256)
func (_Contract *ContractCallerSession) NativeForKeeper() (*big.Int, error) {
	return _Contract.Contract.NativeForKeeper(&_Contract.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contract *ContractCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contract *ContractSession) Oracle() (common.Address, error) {
	return _Contract.Contract.Oracle(&_Contract.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contract *ContractCallerSession) Oracle() (common.Address, error) {
	return _Contract.Contract.Oracle(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PairInfo is a free data retrieval call binding the contract method 0xfc2a5b1d.
//
// Solidity: function pairInfo() view returns(address)
func (_Contract *ContractCaller) PairInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pairInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairInfo is a free data retrieval call binding the contract method 0xfc2a5b1d.
//
// Solidity: function pairInfo() view returns(address)
func (_Contract *ContractSession) PairInfo() (common.Address, error) {
	return _Contract.Contract.PairInfo(&_Contract.CallOpts)
}

// PairInfo is a free data retrieval call binding the contract method 0xfc2a5b1d.
//
// Solidity: function pairInfo() view returns(address)
func (_Contract *ContractCallerSession) PairInfo() (common.Address, error) {
	return _Contract.Contract.PairInfo(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCallerSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Contract *ContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Contract *ContractSession) PendingOwner() (common.Address, error) {
	return _Contract.Contract.PendingOwner(&_Contract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Contract *ContractCallerSession) PendingOwner() (common.Address, error) {
	return _Contract.Contract.PendingOwner(&_Contract.CallOpts)
}

// ReserveRate is a free data retrieval call binding the contract method 0x58d7bf80.
//
// Solidity: function reserveRate() view returns(uint256)
func (_Contract *ContractCaller) ReserveRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reserveRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveRate is a free data retrieval call binding the contract method 0x58d7bf80.
//
// Solidity: function reserveRate() view returns(uint256)
func (_Contract *ContractSession) ReserveRate() (*big.Int, error) {
	return _Contract.Contract.ReserveRate(&_Contract.CallOpts)
}

// ReserveRate is a free data retrieval call binding the contract method 0x58d7bf80.
//
// Solidity: function reserveRate() view returns(uint256)
func (_Contract *ContractCallerSession) ReserveRate() (*big.Int, error) {
	return _Contract.Contract.ReserveRate(&_Contract.CallOpts)
}

// ReserveReceiver is a free data retrieval call binding the contract method 0xbb1c7c2a.
//
// Solidity: function reserveReceiver() view returns(address)
func (_Contract *ContractCaller) ReserveReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reserveReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveReceiver is a free data retrieval call binding the contract method 0xbb1c7c2a.
//
// Solidity: function reserveReceiver() view returns(address)
func (_Contract *ContractSession) ReserveReceiver() (common.Address, error) {
	return _Contract.Contract.ReserveReceiver(&_Contract.CallOpts)
}

// ReserveReceiver is a free data retrieval call binding the contract method 0xbb1c7c2a.
//
// Solidity: function reserveReceiver() view returns(address)
func (_Contract *ContractCallerSession) ReserveReceiver() (common.Address, error) {
	return _Contract.Contract.ReserveReceiver(&_Contract.CallOpts)
}

// SpreadReductionP is a free data retrieval call binding the contract method 0xefbc5fc0.
//
// Solidity: function spreadReductionP() view returns(uint256)
func (_Contract *ContractCaller) SpreadReductionP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "spreadReductionP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadReductionP is a free data retrieval call binding the contract method 0xefbc5fc0.
//
// Solidity: function spreadReductionP() view returns(uint256)
func (_Contract *ContractSession) SpreadReductionP() (*big.Int, error) {
	return _Contract.Contract.SpreadReductionP(&_Contract.CallOpts)
}

// SpreadReductionP is a free data retrieval call binding the contract method 0xefbc5fc0.
//
// Solidity: function spreadReductionP() view returns(uint256)
func (_Contract *ContractCallerSession) SpreadReductionP() (*big.Int, error) {
	return _Contract.Contract.SpreadReductionP(&_Contract.CallOpts)
}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_Contract *ContractCaller) StorageT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "storageT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_Contract *ContractSession) StorageT() (common.Address, error) {
	return _Contract.Contract.StorageT(&_Contract.CallOpts)
}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_Contract *ContractCallerSession) StorageT() (common.Address, error) {
	return _Contract.Contract.StorageT(&_Contract.CallOpts)
}

// TradeSwitch is a free data retrieval call binding the contract method 0xa5c34ba7.
//
// Solidity: function tradeSwitch() view returns(uint256)
func (_Contract *ContractCaller) TradeSwitch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tradeSwitch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradeSwitch is a free data retrieval call binding the contract method 0xa5c34ba7.
//
// Solidity: function tradeSwitch() view returns(uint256)
func (_Contract *ContractSession) TradeSwitch() (*big.Int, error) {
	return _Contract.Contract.TradeSwitch(&_Contract.CallOpts)
}

// TradeSwitch is a free data retrieval call binding the contract method 0xa5c34ba7.
//
// Solidity: function tradeSwitch() view returns(uint256)
func (_Contract *ContractCallerSession) TradeSwitch() (*big.Int, error) {
	return _Contract.Contract.TradeSwitch(&_Contract.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Contract *ContractCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Contract *ContractSession) Usdt() (common.Address, error) {
	return _Contract.Contract.Usdt(&_Contract.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Contract *ContractCallerSession) Usdt() (common.Address, error) {
	return _Contract.Contract.Usdt(&_Contract.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Contract *ContractCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Contract *ContractSession) Vault() (common.Address, error) {
	return _Contract.Contract.Vault(&_Contract.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Contract *ContractCallerSession) Vault() (common.Address, error) {
	return _Contract.Contract.Vault(&_Contract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contract *ContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contract *ContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contract.Contract.AcceptOwnership(&_Contract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contract *ContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contract.Contract.AcceptOwnership(&_Contract.TransactOpts)
}

// CancelOpen is a paid mutator transaction binding the contract method 0x806762a0.
//
// Solidity: function cancelOpen(uint256 _orderId) returns()
func (_Contract *ContractTransactor) CancelOpen(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelOpen", _orderId)
}

// CancelOpen is a paid mutator transaction binding the contract method 0x806762a0.
//
// Solidity: function cancelOpen(uint256 _orderId) returns()
func (_Contract *ContractSession) CancelOpen(_orderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOpen(&_Contract.TransactOpts, _orderId)
}

// CancelOpen is a paid mutator transaction binding the contract method 0x806762a0.
//
// Solidity: function cancelOpen(uint256 _orderId) returns()
func (_Contract *ContractTransactorSession) CancelOpen(_orderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOpen(&_Contract.TransactOpts, _orderId)
}

// CloseTrade is a paid mutator transaction binding the contract method 0x338b0133.
//
// Solidity: function closeTrade(uint256 _orderId, uint256 index, bytes[] updateData) returns()
func (_Contract *ContractTransactor) CloseTrade(opts *bind.TransactOpts, _orderId *big.Int, index *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "closeTrade", _orderId, index, updateData)
}

// CloseTrade is a paid mutator transaction binding the contract method 0x338b0133.
//
// Solidity: function closeTrade(uint256 _orderId, uint256 index, bytes[] updateData) returns()
func (_Contract *ContractSession) CloseTrade(_orderId *big.Int, index *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.CloseTrade(&_Contract.TransactOpts, _orderId, index, updateData)
}

// CloseTrade is a paid mutator transaction binding the contract method 0x338b0133.
//
// Solidity: function closeTrade(uint256 _orderId, uint256 index, bytes[] updateData) returns()
func (_Contract *ContractTransactorSession) CloseTrade(_orderId *big.Int, index *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.CloseTrade(&_Contract.TransactOpts, _orderId, index, updateData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner, address _storageT, address _usdt) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _storageT common.Address, _usdt common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", owner, _storageT, _usdt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner, address _storageT, address _usdt) returns()
func (_Contract *ContractSession) Initialize(owner common.Address, _storageT common.Address, _usdt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, owner, _storageT, _usdt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner, address _storageT, address _usdt) returns()
func (_Contract *ContractTransactorSession) Initialize(owner common.Address, _storageT common.Address, _usdt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, owner, _storageT, _usdt)
}

// Link is a paid mutator transaction binding the contract method 0x1c4695f4.
//
// Solidity: function link() returns()
func (_Contract *ContractTransactor) Link(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "link")
}

// Link is a paid mutator transaction binding the contract method 0x1c4695f4.
//
// Solidity: function link() returns()
func (_Contract *ContractSession) Link() (*types.Transaction, error) {
	return _Contract.Contract.Link(&_Contract.TransactOpts)
}

// Link is a paid mutator transaction binding the contract method 0x1c4695f4.
//
// Solidity: function link() returns()
func (_Contract *ContractTransactorSession) Link() (*types.Transaction, error) {
	return _Contract.Contract.Link(&_Contract.TransactOpts)
}

// OpenTrade is a paid mutator transaction binding the contract method 0x1c7f6f4c.
//
// Solidity: function openTrade(uint256 _orderId, bytes[] updateData) returns()
func (_Contract *ContractTransactor) OpenTrade(opts *bind.TransactOpts, _orderId *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "openTrade", _orderId, updateData)
}

// OpenTrade is a paid mutator transaction binding the contract method 0x1c7f6f4c.
//
// Solidity: function openTrade(uint256 _orderId, bytes[] updateData) returns()
func (_Contract *ContractSession) OpenTrade(_orderId *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.OpenTrade(&_Contract.TransactOpts, _orderId, updateData)
}

// OpenTrade is a paid mutator transaction binding the contract method 0x1c7f6f4c.
//
// Solidity: function openTrade(uint256 _orderId, bytes[] updateData) returns()
func (_Contract *ContractTransactorSession) OpenTrade(_orderId *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.OpenTrade(&_Contract.TransactOpts, _orderId, updateData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RequestClose is a paid mutator transaction binding the contract method 0xc04398ee.
//
// Solidity: function requestClose(uint256 _orderId, uint256 _closeMargin, address callTarget) payable returns(uint256)
func (_Contract *ContractTransactor) RequestClose(opts *bind.TransactOpts, _orderId *big.Int, _closeMargin *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestClose", _orderId, _closeMargin, callTarget)
}

// RequestClose is a paid mutator transaction binding the contract method 0xc04398ee.
//
// Solidity: function requestClose(uint256 _orderId, uint256 _closeMargin, address callTarget) payable returns(uint256)
func (_Contract *ContractSession) RequestClose(_orderId *big.Int, _closeMargin *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestClose(&_Contract.TransactOpts, _orderId, _closeMargin, callTarget)
}

// RequestClose is a paid mutator transaction binding the contract method 0xc04398ee.
//
// Solidity: function requestClose(uint256 _orderId, uint256 _closeMargin, address callTarget) payable returns(uint256)
func (_Contract *ContractTransactorSession) RequestClose(_orderId *big.Int, _closeMargin *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestClose(&_Contract.TransactOpts, _orderId, _closeMargin, callTarget)
}

// RequestOpenLimit is a paid mutator transaction binding the contract method 0xf750704f.
//
// Solidity: function requestOpenLimit(uint8 _t, (address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 limit, address callTarget) payable returns(uint256)
func (_Contract *ContractTransactor) RequestOpenLimit(opts *bind.TransactOpts, _t uint8, _base TradeBase, updateData [][]byte, limit *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestOpenLimit", _t, _base, updateData, limit, callTarget)
}

// RequestOpenLimit is a paid mutator transaction binding the contract method 0xf750704f.
//
// Solidity: function requestOpenLimit(uint8 _t, (address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 limit, address callTarget) payable returns(uint256)
func (_Contract *ContractSession) RequestOpenLimit(_t uint8, _base TradeBase, updateData [][]byte, limit *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenLimit(&_Contract.TransactOpts, _t, _base, updateData, limit, callTarget)
}

// RequestOpenLimit is a paid mutator transaction binding the contract method 0xf750704f.
//
// Solidity: function requestOpenLimit(uint8 _t, (address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 limit, address callTarget) payable returns(uint256)
func (_Contract *ContractTransactorSession) RequestOpenLimit(_t uint8, _base TradeBase, updateData [][]byte, limit *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenLimit(&_Contract.TransactOpts, _t, _base, updateData, limit, callTarget)
}

// RequestOpenMarket is a paid mutator transaction binding the contract method 0x8f92dab8.
//
// Solidity: function requestOpenMarket((address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 slippage, address callTarget) payable returns(uint256)
func (_Contract *ContractTransactor) RequestOpenMarket(opts *bind.TransactOpts, _base TradeBase, updateData [][]byte, slippage *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestOpenMarket", _base, updateData, slippage, callTarget)
}

// RequestOpenMarket is a paid mutator transaction binding the contract method 0x8f92dab8.
//
// Solidity: function requestOpenMarket((address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 slippage, address callTarget) payable returns(uint256)
func (_Contract *ContractSession) RequestOpenMarket(_base TradeBase, updateData [][]byte, slippage *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenMarket(&_Contract.TransactOpts, _base, updateData, slippage, callTarget)
}

// RequestOpenMarket is a paid mutator transaction binding the contract method 0x8f92dab8.
//
// Solidity: function requestOpenMarket((address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 slippage, address callTarget) payable returns(uint256)
func (_Contract *ContractTransactorSession) RequestOpenMarket(_base TradeBase, updateData [][]byte, slippage *big.Int, callTarget common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenMarket(&_Contract.TransactOpts, _base, updateData, slippage, callTarget)
}

// SetContract is a paid mutator transaction binding the contract method 0x2bf6e0a5.
//
// Solidity: function setContract(address _tradingStorage, address _usdt) returns()
func (_Contract *ContractTransactor) SetContract(opts *bind.TransactOpts, _tradingStorage common.Address, _usdt common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setContract", _tradingStorage, _usdt)
}

// SetContract is a paid mutator transaction binding the contract method 0x2bf6e0a5.
//
// Solidity: function setContract(address _tradingStorage, address _usdt) returns()
func (_Contract *ContractSession) SetContract(_tradingStorage common.Address, _usdt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetContract(&_Contract.TransactOpts, _tradingStorage, _usdt)
}

// SetContract is a paid mutator transaction binding the contract method 0x2bf6e0a5.
//
// Solidity: function setContract(address _tradingStorage, address _usdt) returns()
func (_Contract *ContractTransactorSession) SetContract(_tradingStorage common.Address, _usdt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetContract(&_Contract.TransactOpts, _tradingStorage, _usdt)
}

// SetNativeFeeForKeeper is a paid mutator transaction binding the contract method 0xb3f6309e.
//
// Solidity: function setNativeFeeForKeeper(uint256 _forKeeper, uint256 _forCallback, uint256 _gasLimit) returns()
func (_Contract *ContractTransactor) SetNativeFeeForKeeper(opts *bind.TransactOpts, _forKeeper *big.Int, _forCallback *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNativeFeeForKeeper", _forKeeper, _forCallback, _gasLimit)
}

// SetNativeFeeForKeeper is a paid mutator transaction binding the contract method 0xb3f6309e.
//
// Solidity: function setNativeFeeForKeeper(uint256 _forKeeper, uint256 _forCallback, uint256 _gasLimit) returns()
func (_Contract *ContractSession) SetNativeFeeForKeeper(_forKeeper *big.Int, _forCallback *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetNativeFeeForKeeper(&_Contract.TransactOpts, _forKeeper, _forCallback, _gasLimit)
}

// SetNativeFeeForKeeper is a paid mutator transaction binding the contract method 0xb3f6309e.
//
// Solidity: function setNativeFeeForKeeper(uint256 _forKeeper, uint256 _forCallback, uint256 _gasLimit) returns()
func (_Contract *ContractTransactorSession) SetNativeFeeForKeeper(_forKeeper *big.Int, _forCallback *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetNativeFeeForKeeper(&_Contract.TransactOpts, _forKeeper, _forCallback, _gasLimit)
}

// SetParams is a paid mutator transaction binding the contract method 0x7ece45e8.
//
// Solidity: function setParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _minAcceptanceDelay, uint256 _tradeSwitch) returns()
func (_Contract *ContractTransactor) SetParams(opts *bind.TransactOpts, _liquidationP *big.Int, _spreadReductionP *big.Int, _minAcceptanceDelay *big.Int, _tradeSwitch *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setParams", _liquidationP, _spreadReductionP, _minAcceptanceDelay, _tradeSwitch)
}

// SetParams is a paid mutator transaction binding the contract method 0x7ece45e8.
//
// Solidity: function setParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _minAcceptanceDelay, uint256 _tradeSwitch) returns()
func (_Contract *ContractSession) SetParams(_liquidationP *big.Int, _spreadReductionP *big.Int, _minAcceptanceDelay *big.Int, _tradeSwitch *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetParams(&_Contract.TransactOpts, _liquidationP, _spreadReductionP, _minAcceptanceDelay, _tradeSwitch)
}

// SetParams is a paid mutator transaction binding the contract method 0x7ece45e8.
//
// Solidity: function setParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _minAcceptanceDelay, uint256 _tradeSwitch) returns()
func (_Contract *ContractTransactorSession) SetParams(_liquidationP *big.Int, _spreadReductionP *big.Int, _minAcceptanceDelay *big.Int, _tradeSwitch *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetParams(&_Contract.TransactOpts, _liquidationP, _spreadReductionP, _minAcceptanceDelay, _tradeSwitch)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address _receiver, uint256 _rate) returns()
func (_Contract *ContractTransactor) SetReserve(opts *bind.TransactOpts, _receiver common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setReserve", _receiver, _rate)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address _receiver, uint256 _rate) returns()
func (_Contract *ContractSession) SetReserve(_receiver common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetReserve(&_Contract.TransactOpts, _receiver, _rate)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address _receiver, uint256 _rate) returns()
func (_Contract *ContractTransactorSession) SetReserve(_receiver common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetReserve(&_Contract.TransactOpts, _receiver, _rate)
}

// StopTrade is a paid mutator transaction binding the contract method 0xcc83bfd2.
//
// Solidity: function stopTrade(uint256 _orderId, uint8 _stop, bytes[] updateData) returns()
func (_Contract *ContractTransactor) StopTrade(opts *bind.TransactOpts, _orderId *big.Int, _stop uint8, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stopTrade", _orderId, _stop, updateData)
}

// StopTrade is a paid mutator transaction binding the contract method 0xcc83bfd2.
//
// Solidity: function stopTrade(uint256 _orderId, uint8 _stop, bytes[] updateData) returns()
func (_Contract *ContractSession) StopTrade(_orderId *big.Int, _stop uint8, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.StopTrade(&_Contract.TransactOpts, _orderId, _stop, updateData)
}

// StopTrade is a paid mutator transaction binding the contract method 0xcc83bfd2.
//
// Solidity: function stopTrade(uint256 _orderId, uint8 _stop, bytes[] updateData) returns()
func (_Contract *ContractTransactorSession) StopTrade(_orderId *big.Int, _stop uint8, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.StopTrade(&_Contract.TransactOpts, _orderId, _stop, updateData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateMargin is a paid mutator transaction binding the contract method 0xfd77f87f.
//
// Solidity: function updateMargin(uint256 _orderId, uint256 _amount, bool _isAdd, bytes[] updateData) returns()
func (_Contract *ContractTransactor) UpdateMargin(opts *bind.TransactOpts, _orderId *big.Int, _amount *big.Int, _isAdd bool, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMargin", _orderId, _amount, _isAdd, updateData)
}

// UpdateMargin is a paid mutator transaction binding the contract method 0xfd77f87f.
//
// Solidity: function updateMargin(uint256 _orderId, uint256 _amount, bool _isAdd, bytes[] updateData) returns()
func (_Contract *ContractSession) UpdateMargin(_orderId *big.Int, _amount *big.Int, _isAdd bool, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMargin(&_Contract.TransactOpts, _orderId, _amount, _isAdd, updateData)
}

// UpdateMargin is a paid mutator transaction binding the contract method 0xfd77f87f.
//
// Solidity: function updateMargin(uint256 _orderId, uint256 _amount, bool _isAdd, bytes[] updateData) returns()
func (_Contract *ContractTransactorSession) UpdateMargin(_orderId *big.Int, _amount *big.Int, _isAdd bool, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMargin(&_Contract.TransactOpts, _orderId, _amount, _isAdd, updateData)
}

// UpdateOpenRequest is a paid mutator transaction binding the contract method 0xefcf0bd0.
//
// Solidity: function updateOpenRequest(uint256 orderId, bytes[] updateData, uint256 limit, uint256 tp, uint256 sl) returns()
func (_Contract *ContractTransactor) UpdateOpenRequest(opts *bind.TransactOpts, orderId *big.Int, updateData [][]byte, limit *big.Int, tp *big.Int, sl *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateOpenRequest", orderId, updateData, limit, tp, sl)
}

// UpdateOpenRequest is a paid mutator transaction binding the contract method 0xefcf0bd0.
//
// Solidity: function updateOpenRequest(uint256 orderId, bytes[] updateData, uint256 limit, uint256 tp, uint256 sl) returns()
func (_Contract *ContractSession) UpdateOpenRequest(orderId *big.Int, updateData [][]byte, limit *big.Int, tp *big.Int, sl *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOpenRequest(&_Contract.TransactOpts, orderId, updateData, limit, tp, sl)
}

// UpdateOpenRequest is a paid mutator transaction binding the contract method 0xefcf0bd0.
//
// Solidity: function updateOpenRequest(uint256 orderId, bytes[] updateData, uint256 limit, uint256 tp, uint256 sl) returns()
func (_Contract *ContractTransactorSession) UpdateOpenRequest(orderId *big.Int, updateData [][]byte, limit *big.Int, tp *big.Int, sl *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOpenRequest(&_Contract.TransactOpts, orderId, updateData, limit, tp, sl)
}

// UpdateTPAndSL is a paid mutator transaction binding the contract method 0x04c7ec71.
//
// Solidity: function updateTPAndSL(uint256 orderId, uint256 tp, uint256 sl, bytes[] updateData) returns()
func (_Contract *ContractTransactor) UpdateTPAndSL(opts *bind.TransactOpts, orderId *big.Int, tp *big.Int, sl *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateTPAndSL", orderId, tp, sl, updateData)
}

// UpdateTPAndSL is a paid mutator transaction binding the contract method 0x04c7ec71.
//
// Solidity: function updateTPAndSL(uint256 orderId, uint256 tp, uint256 sl, bytes[] updateData) returns()
func (_Contract *ContractSession) UpdateTPAndSL(orderId *big.Int, tp *big.Int, sl *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTPAndSL(&_Contract.TransactOpts, orderId, tp, sl, updateData)
}

// UpdateTPAndSL is a paid mutator transaction binding the contract method 0x04c7ec71.
//
// Solidity: function updateTPAndSL(uint256 orderId, uint256 tp, uint256 sl, bytes[] updateData) returns()
func (_Contract *ContractTransactorSession) UpdateTPAndSL(orderId *big.Int, tp *big.Int, sl *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTPAndSL(&_Contract.TransactOpts, orderId, tp, sl, updateData)
}

// ContractCallbackIterator is returned from FilterCallback and is used to iterate over the raw logs and unpacked data for Callback events raised by the Contract contract.
type ContractCallbackIterator struct {
	Event *ContractCallback // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCallback)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCallback)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCallback represents a Callback event raised by the Contract contract.
type ContractCallback struct {
	Target common.Address
	Resut  bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallback is a free log retrieval operation binding the contract event 0x46ddbd62fc1a7626fe9c43026cb0694aec0b031fe81ac66fb4cfe9381dc6fe72.
//
// Solidity: event Callback(address target, bool resut)
func (_Contract *ContractFilterer) FilterCallback(opts *bind.FilterOpts) (*ContractCallbackIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return &ContractCallbackIterator{contract: _Contract.contract, event: "Callback", logs: logs, sub: sub}, nil
}

// WatchCallback is a free log subscription operation binding the contract event 0x46ddbd62fc1a7626fe9c43026cb0694aec0b031fe81ac66fb4cfe9381dc6fe72.
//
// Solidity: event Callback(address target, bool resut)
func (_Contract *ContractFilterer) WatchCallback(opts *bind.WatchOpts, sink chan<- *ContractCallback) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCallback)
				if err := _Contract.contract.UnpackLog(event, "Callback", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallback is a log parse operation binding the contract event 0x46ddbd62fc1a7626fe9c43026cb0694aec0b031fe81ac66fb4cfe9381dc6fe72.
//
// Solidity: event Callback(address target, bool resut)
func (_Contract *ContractFilterer) ParseCallback(log types.Log) (*ContractCallback, error) {
	event := new(ContractCallback)
	if err := _Contract.contract.UnpackLog(event, "Callback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCancelOpenIterator is returned from FilterCancelOpen and is used to iterate over the raw logs and unpacked data for CancelOpen events raised by the Contract contract.
type ContractCancelOpenIterator struct {
	Event *ContractCancelOpen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCancelOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCancelOpen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCancelOpen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCancelOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCancelOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCancelOpen represents a CancelOpen event raised by the Contract contract.
type ContractCancelOpen struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelOpen is a free log retrieval operation binding the contract event 0x1dbfda3a6640539c3fac591e271a7b4ccd445cf9bc2b951e58abde22dfe73f53.
//
// Solidity: event CancelOpen(uint256 orderId)
func (_Contract *ContractFilterer) FilterCancelOpen(opts *bind.FilterOpts) (*ContractCancelOpenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CancelOpen")
	if err != nil {
		return nil, err
	}
	return &ContractCancelOpenIterator{contract: _Contract.contract, event: "CancelOpen", logs: logs, sub: sub}, nil
}

// WatchCancelOpen is a free log subscription operation binding the contract event 0x1dbfda3a6640539c3fac591e271a7b4ccd445cf9bc2b951e58abde22dfe73f53.
//
// Solidity: event CancelOpen(uint256 orderId)
func (_Contract *ContractFilterer) WatchCancelOpen(opts *bind.WatchOpts, sink chan<- *ContractCancelOpen) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CancelOpen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCancelOpen)
				if err := _Contract.contract.UnpackLog(event, "CancelOpen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelOpen is a log parse operation binding the contract event 0x1dbfda3a6640539c3fac591e271a7b4ccd445cf9bc2b951e58abde22dfe73f53.
//
// Solidity: event CancelOpen(uint256 orderId)
func (_Contract *ContractFilterer) ParseCancelOpen(log types.Log) (*ContractCancelOpen, error) {
	event := new(ContractCancelOpen)
	if err := _Contract.contract.UnpackLog(event, "CancelOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCloseIterator is returned from FilterClose and is used to iterate over the raw logs and unpacked data for Close events raised by the Contract contract.
type ContractCloseIterator struct {
	Event *ContractClose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClose represents a Close event raised by the Contract contract.
type ContractClose struct {
	OrderId     *big.Int
	ClosePrice  *big.Int
	CloseMargin *big.Int
	FundingFee  *big.Int
	RolloverFee *big.Int
	CloseFee    *big.Int
	AfterFee    *big.Int
	S           uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClose is a free log retrieval operation binding the contract event 0x24b27bc2a65d006b06d5ee8450abfa71760f2aac880fcf5555af1e692cfaf539.
//
// Solidity: event Close(uint256 orderId, uint256 closePrice, uint256 _closeMargin, int256 fundingFee, uint256 rolloverFee, uint256 closeFee, uint256 afterFee, uint8 s)
func (_Contract *ContractFilterer) FilterClose(opts *bind.FilterOpts) (*ContractCloseIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return &ContractCloseIterator{contract: _Contract.contract, event: "Close", logs: logs, sub: sub}, nil
}

// WatchClose is a free log subscription operation binding the contract event 0x24b27bc2a65d006b06d5ee8450abfa71760f2aac880fcf5555af1e692cfaf539.
//
// Solidity: event Close(uint256 orderId, uint256 closePrice, uint256 _closeMargin, int256 fundingFee, uint256 rolloverFee, uint256 closeFee, uint256 afterFee, uint8 s)
func (_Contract *ContractFilterer) WatchClose(opts *bind.WatchOpts, sink chan<- *ContractClose) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClose)
				if err := _Contract.contract.UnpackLog(event, "Close", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClose is a log parse operation binding the contract event 0x24b27bc2a65d006b06d5ee8450abfa71760f2aac880fcf5555af1e692cfaf539.
//
// Solidity: event Close(uint256 orderId, uint256 closePrice, uint256 _closeMargin, int256 fundingFee, uint256 rolloverFee, uint256 closeFee, uint256 afterFee, uint8 s)
func (_Contract *ContractFilterer) ParseClose(log types.Log) (*ContractClose, error) {
	event := new(ContractClose)
	if err := _Contract.contract.UnpackLog(event, "Close", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExecRequestCloseIterator is returned from FilterExecRequestClose and is used to iterate over the raw logs and unpacked data for ExecRequestClose events raised by the Contract contract.
type ContractExecRequestCloseIterator struct {
	Event *ContractExecRequestClose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractExecRequestCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExecRequestClose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractExecRequestClose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractExecRequestCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExecRequestCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExecRequestClose represents a ExecRequestClose event raised by the Contract contract.
type ContractExecRequestClose struct {
	OrderId *big.Int
	Index   *big.Int
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecRequestClose is a free log retrieval operation binding the contract event 0x28cf3bfcf594f6c4e9da1e79297930dbd971dcae2e08a19857abc14e759e9401.
//
// Solidity: event ExecRequestClose(uint256 orderId, uint256 index, bool status)
func (_Contract *ContractFilterer) FilterExecRequestClose(opts *bind.FilterOpts) (*ContractExecRequestCloseIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExecRequestClose")
	if err != nil {
		return nil, err
	}
	return &ContractExecRequestCloseIterator{contract: _Contract.contract, event: "ExecRequestClose", logs: logs, sub: sub}, nil
}

// WatchExecRequestClose is a free log subscription operation binding the contract event 0x28cf3bfcf594f6c4e9da1e79297930dbd971dcae2e08a19857abc14e759e9401.
//
// Solidity: event ExecRequestClose(uint256 orderId, uint256 index, bool status)
func (_Contract *ContractFilterer) WatchExecRequestClose(opts *bind.WatchOpts, sink chan<- *ContractExecRequestClose) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExecRequestClose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExecRequestClose)
				if err := _Contract.contract.UnpackLog(event, "ExecRequestClose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecRequestClose is a log parse operation binding the contract event 0x28cf3bfcf594f6c4e9da1e79297930dbd971dcae2e08a19857abc14e759e9401.
//
// Solidity: event ExecRequestClose(uint256 orderId, uint256 index, bool status)
func (_Contract *ContractFilterer) ParseExecRequestClose(log types.Log) (*ContractExecRequestClose, error) {
	event := new(ContractExecRequestClose)
	if err := _Contract.contract.UnpackLog(event, "ExecRequestClose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExecRequestOpenIterator is returned from FilterExecRequestOpen and is used to iterate over the raw logs and unpacked data for ExecRequestOpen events raised by the Contract contract.
type ContractExecRequestOpenIterator struct {
	Event *ContractExecRequestOpen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractExecRequestOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExecRequestOpen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractExecRequestOpen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractExecRequestOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExecRequestOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExecRequestOpen represents a ExecRequestOpen event raised by the Contract contract.
type ContractExecRequestOpen struct {
	OrderId *big.Int
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecRequestOpen is a free log retrieval operation binding the contract event 0x6fc494543173949b35b8ae1933d2c0f5208180b4f2aef262bb87523bb1def648.
//
// Solidity: event ExecRequestOpen(uint256 orderId, bool status)
func (_Contract *ContractFilterer) FilterExecRequestOpen(opts *bind.FilterOpts) (*ContractExecRequestOpenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExecRequestOpen")
	if err != nil {
		return nil, err
	}
	return &ContractExecRequestOpenIterator{contract: _Contract.contract, event: "ExecRequestOpen", logs: logs, sub: sub}, nil
}

// WatchExecRequestOpen is a free log subscription operation binding the contract event 0x6fc494543173949b35b8ae1933d2c0f5208180b4f2aef262bb87523bb1def648.
//
// Solidity: event ExecRequestOpen(uint256 orderId, bool status)
func (_Contract *ContractFilterer) WatchExecRequestOpen(opts *bind.WatchOpts, sink chan<- *ContractExecRequestOpen) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExecRequestOpen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExecRequestOpen)
				if err := _Contract.contract.UnpackLog(event, "ExecRequestOpen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecRequestOpen is a log parse operation binding the contract event 0x6fc494543173949b35b8ae1933d2c0f5208180b4f2aef262bb87523bb1def648.
//
// Solidity: event ExecRequestOpen(uint256 orderId, bool status)
func (_Contract *ContractFilterer) ParseExecRequestOpen(log types.Log) (*ContractExecRequestOpen, error) {
	event := new(ContractExecRequestOpen)
	if err := _Contract.contract.UnpackLog(event, "ExecRequestOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenIterator is returned from FilterOpen and is used to iterate over the raw logs and unpacked data for Open events raised by the Contract contract.
type ContractOpenIterator struct {
	Event *ContractOpen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOpen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpen represents a Open event raised by the Contract contract.
type ContractOpen struct {
	OrderId *big.Int
	T       OpenTrade
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpen is a free log retrieval operation binding the contract event 0x09c087e430ee6913868eafe26e6d9da87f95984c61bde74e50e28ecd09da890a.
//
// Solidity: event Open(uint256 orderId, ((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256) t, uint256 fee)
func (_Contract *ContractFilterer) FilterOpen(opts *bind.FilterOpts) (*ContractOpenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Open")
	if err != nil {
		return nil, err
	}
	return &ContractOpenIterator{contract: _Contract.contract, event: "Open", logs: logs, sub: sub}, nil
}

// WatchOpen is a free log subscription operation binding the contract event 0x09c087e430ee6913868eafe26e6d9da87f95984c61bde74e50e28ecd09da890a.
//
// Solidity: event Open(uint256 orderId, ((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256) t, uint256 fee)
func (_Contract *ContractFilterer) WatchOpen(opts *bind.WatchOpts, sink chan<- *ContractOpen) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Open")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpen)
				if err := _Contract.contract.UnpackLog(event, "Open", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpen is a log parse operation binding the contract event 0x09c087e430ee6913868eafe26e6d9da87f95984c61bde74e50e28ecd09da890a.
//
// Solidity: event Open(uint256 orderId, ((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256) t, uint256 fee)
func (_Contract *ContractFilterer) ParseOpen(log types.Log) (*ContractOpen, error) {
	event := new(ContractOpen)
	if err := _Contract.contract.UnpackLog(event, "Open", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Contract contract.
type ContractOwnershipTransferStartedIterator struct {
	Event *ContractOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Contract contract.
type ContractOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferStartedIterator{contract: _Contract.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferStarted)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferStarted(log types.Log) (*ContractOwnershipTransferStarted, error) {
	event := new(ContractOwnershipTransferStarted)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contract contract.
type ContractPausedIterator struct {
	Event *ContractPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPaused represents a Paused event raised by the Contract contract.
type ContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contract *ContractFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractPausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractPausedIterator{contract: _Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contract *ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPaused)
				if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contract *ContractFilterer) ParsePaused(log types.Log) (*ContractPaused, error) {
	event := new(ContractPaused)
	if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRequestCloseIterator is returned from FilterRequestClose and is used to iterate over the raw logs and unpacked data for RequestClose events raised by the Contract contract.
type ContractRequestCloseIterator struct {
	Event *ContractRequestClose // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRequestCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRequestClose)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRequestClose)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRequestCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRequestCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRequestClose represents a RequestClose event raised by the Contract contract.
type ContractRequestClose struct {
	OrderId     *big.Int
	Index       *big.Int
	CloseMargin *big.Int
	RequstTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestClose is a free log retrieval operation binding the contract event 0x7695d3ca62416d30d27e9fac1080c5322b164e257e8151db3b139aec5d81605e.
//
// Solidity: event RequestClose(uint256 indexed orderId, uint256 indexed index, uint256 _closeMargin, uint256 requstTime)
func (_Contract *ContractFilterer) FilterRequestClose(opts *bind.FilterOpts, orderId []*big.Int, index []*big.Int) (*ContractRequestCloseIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RequestClose", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &ContractRequestCloseIterator{contract: _Contract.contract, event: "RequestClose", logs: logs, sub: sub}, nil
}

// WatchRequestClose is a free log subscription operation binding the contract event 0x7695d3ca62416d30d27e9fac1080c5322b164e257e8151db3b139aec5d81605e.
//
// Solidity: event RequestClose(uint256 indexed orderId, uint256 indexed index, uint256 _closeMargin, uint256 requstTime)
func (_Contract *ContractFilterer) WatchRequestClose(opts *bind.WatchOpts, sink chan<- *ContractRequestClose, orderId []*big.Int, index []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RequestClose", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRequestClose)
				if err := _Contract.contract.UnpackLog(event, "RequestClose", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestClose is a log parse operation binding the contract event 0x7695d3ca62416d30d27e9fac1080c5322b164e257e8151db3b139aec5d81605e.
//
// Solidity: event RequestClose(uint256 indexed orderId, uint256 indexed index, uint256 _closeMargin, uint256 requstTime)
func (_Contract *ContractFilterer) ParseRequestClose(log types.Log) (*ContractRequestClose, error) {
	event := new(ContractRequestClose)
	if err := _Contract.contract.UnpackLog(event, "RequestClose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRequestOpenIterator is returned from FilterRequestOpen and is used to iterate over the raw logs and unpacked data for RequestOpen events raised by the Contract contract.
type ContractRequestOpenIterator struct {
	Event *ContractRequestOpen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRequestOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRequestOpen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRequestOpen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRequestOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRequestOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRequestOpen represents a RequestOpen event raised by the Contract contract.
type ContractRequestOpen struct {
	OrderId *big.Int
	Request OpenRequest
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequestOpen is a free log retrieval operation binding the contract event 0xc815d98967a87feaaa8e57736b69b57611867c43e87d7c0cd7ba5edf3f13c528.
//
// Solidity: event RequestOpen(uint256 indexed orderId, (uint8,(address,uint256,uint256,bool,uint256,uint256,uint256),address,uint256,uint256,uint256) _request)
func (_Contract *ContractFilterer) FilterRequestOpen(opts *bind.FilterOpts, orderId []*big.Int) (*ContractRequestOpenIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RequestOpen", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractRequestOpenIterator{contract: _Contract.contract, event: "RequestOpen", logs: logs, sub: sub}, nil
}

// WatchRequestOpen is a free log subscription operation binding the contract event 0xc815d98967a87feaaa8e57736b69b57611867c43e87d7c0cd7ba5edf3f13c528.
//
// Solidity: event RequestOpen(uint256 indexed orderId, (uint8,(address,uint256,uint256,bool,uint256,uint256,uint256),address,uint256,uint256,uint256) _request)
func (_Contract *ContractFilterer) WatchRequestOpen(opts *bind.WatchOpts, sink chan<- *ContractRequestOpen, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RequestOpen", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRequestOpen)
				if err := _Contract.contract.UnpackLog(event, "RequestOpen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestOpen is a log parse operation binding the contract event 0xc815d98967a87feaaa8e57736b69b57611867c43e87d7c0cd7ba5edf3f13c528.
//
// Solidity: event RequestOpen(uint256 indexed orderId, (uint8,(address,uint256,uint256,bool,uint256,uint256,uint256),address,uint256,uint256,uint256) _request)
func (_Contract *ContractFilterer) ParseRequestOpen(log types.Log) (*ContractRequestOpen, error) {
	event := new(ContractRequestOpen)
	if err := _Contract.contract.UnpackLog(event, "RequestOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetContractIterator is returned from FilterSetContract and is used to iterate over the raw logs and unpacked data for SetContract events raised by the Contract contract.
type ContractSetContractIterator struct {
	Event *ContractSetContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSetContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSetContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSetContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetContract represents a SetContract event raised by the Contract contract.
type ContractSetContract struct {
	TradingStorage common.Address
	Usdt           common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetContract is a free log retrieval operation binding the contract event 0x41fe2135e07f51f107aa1df22b2d51b5b6328f3a324fbfc555699785941adb59.
//
// Solidity: event SetContract(address _tradingStorage, address _usdt)
func (_Contract *ContractFilterer) FilterSetContract(opts *bind.FilterOpts) (*ContractSetContractIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetContract")
	if err != nil {
		return nil, err
	}
	return &ContractSetContractIterator{contract: _Contract.contract, event: "SetContract", logs: logs, sub: sub}, nil
}

// WatchSetContract is a free log subscription operation binding the contract event 0x41fe2135e07f51f107aa1df22b2d51b5b6328f3a324fbfc555699785941adb59.
//
// Solidity: event SetContract(address _tradingStorage, address _usdt)
func (_Contract *ContractFilterer) WatchSetContract(opts *bind.WatchOpts, sink chan<- *ContractSetContract) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetContract)
				if err := _Contract.contract.UnpackLog(event, "SetContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetContract is a log parse operation binding the contract event 0x41fe2135e07f51f107aa1df22b2d51b5b6328f3a324fbfc555699785941adb59.
//
// Solidity: event SetContract(address _tradingStorage, address _usdt)
func (_Contract *ContractFilterer) ParseSetContract(log types.Log) (*ContractSetContract, error) {
	event := new(ContractSetContract)
	if err := _Contract.contract.UnpackLog(event, "SetContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetNativeFeeForKeeperIterator is returned from FilterSetNativeFeeForKeeper and is used to iterate over the raw logs and unpacked data for SetNativeFeeForKeeper events raised by the Contract contract.
type ContractSetNativeFeeForKeeperIterator struct {
	Event *ContractSetNativeFeeForKeeper // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSetNativeFeeForKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetNativeFeeForKeeper)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSetNativeFeeForKeeper)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSetNativeFeeForKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetNativeFeeForKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetNativeFeeForKeeper represents a SetNativeFeeForKeeper event raised by the Contract contract.
type ContractSetNativeFeeForKeeper struct {
	ForKeeper   *big.Int
	ForCallback *big.Int
	GasLimit    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetNativeFeeForKeeper is a free log retrieval operation binding the contract event 0x0f07c37c0aa1f3014f3bf74dfbdd00f64787896d550636c713804ef1c43e7f29.
//
// Solidity: event SetNativeFeeForKeeper(uint256 _forKeeper, uint256 _forCallback, uint256 _gasLimit)
func (_Contract *ContractFilterer) FilterSetNativeFeeForKeeper(opts *bind.FilterOpts) (*ContractSetNativeFeeForKeeperIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetNativeFeeForKeeper")
	if err != nil {
		return nil, err
	}
	return &ContractSetNativeFeeForKeeperIterator{contract: _Contract.contract, event: "SetNativeFeeForKeeper", logs: logs, sub: sub}, nil
}

// WatchSetNativeFeeForKeeper is a free log subscription operation binding the contract event 0x0f07c37c0aa1f3014f3bf74dfbdd00f64787896d550636c713804ef1c43e7f29.
//
// Solidity: event SetNativeFeeForKeeper(uint256 _forKeeper, uint256 _forCallback, uint256 _gasLimit)
func (_Contract *ContractFilterer) WatchSetNativeFeeForKeeper(opts *bind.WatchOpts, sink chan<- *ContractSetNativeFeeForKeeper) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetNativeFeeForKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetNativeFeeForKeeper)
				if err := _Contract.contract.UnpackLog(event, "SetNativeFeeForKeeper", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetNativeFeeForKeeper is a log parse operation binding the contract event 0x0f07c37c0aa1f3014f3bf74dfbdd00f64787896d550636c713804ef1c43e7f29.
//
// Solidity: event SetNativeFeeForKeeper(uint256 _forKeeper, uint256 _forCallback, uint256 _gasLimit)
func (_Contract *ContractFilterer) ParseSetNativeFeeForKeeper(log types.Log) (*ContractSetNativeFeeForKeeper, error) {
	event := new(ContractSetNativeFeeForKeeper)
	if err := _Contract.contract.UnpackLog(event, "SetNativeFeeForKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetParamsIterator is returned from FilterSetParams and is used to iterate over the raw logs and unpacked data for SetParams events raised by the Contract contract.
type ContractSetParamsIterator struct {
	Event *ContractSetParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSetParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSetParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSetParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetParams represents a SetParams event raised by the Contract contract.
type ContractSetParams struct {
	LiquidationP       *big.Int
	SpreadReductionP   *big.Int
	MinAcceptanceDelay *big.Int
	TradeSwitch        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetParams is a free log retrieval operation binding the contract event 0xe33eb1ee60efac814379b14a1d72ec5d666301d3264613e9d3f4ad9ef38b0946.
//
// Solidity: event SetParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _minAcceptanceDelay, uint256 _tradeSwitch)
func (_Contract *ContractFilterer) FilterSetParams(opts *bind.FilterOpts) (*ContractSetParamsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetParams")
	if err != nil {
		return nil, err
	}
	return &ContractSetParamsIterator{contract: _Contract.contract, event: "SetParams", logs: logs, sub: sub}, nil
}

// WatchSetParams is a free log subscription operation binding the contract event 0xe33eb1ee60efac814379b14a1d72ec5d666301d3264613e9d3f4ad9ef38b0946.
//
// Solidity: event SetParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _minAcceptanceDelay, uint256 _tradeSwitch)
func (_Contract *ContractFilterer) WatchSetParams(opts *bind.WatchOpts, sink chan<- *ContractSetParams) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetParams)
				if err := _Contract.contract.UnpackLog(event, "SetParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParams is a log parse operation binding the contract event 0xe33eb1ee60efac814379b14a1d72ec5d666301d3264613e9d3f4ad9ef38b0946.
//
// Solidity: event SetParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _minAcceptanceDelay, uint256 _tradeSwitch)
func (_Contract *ContractFilterer) ParseSetParams(log types.Log) (*ContractSetParams, error) {
	event := new(ContractSetParams)
	if err := _Contract.contract.UnpackLog(event, "SetParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetReserveIterator is returned from FilterSetReserve and is used to iterate over the raw logs and unpacked data for SetReserve events raised by the Contract contract.
type ContractSetReserveIterator struct {
	Event *ContractSetReserve // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSetReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetReserve)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSetReserve)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSetReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetReserve represents a SetReserve event raised by the Contract contract.
type ContractSetReserve struct {
	Receiver common.Address
	Rate     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetReserve is a free log retrieval operation binding the contract event 0x9241e21d4fa6dd55661c8694538bfc11a7d44d81d5766840312d2513f55fa67d.
//
// Solidity: event SetReserve(address indexed _receiver, uint256 _rate)
func (_Contract *ContractFilterer) FilterSetReserve(opts *bind.FilterOpts, _receiver []common.Address) (*ContractSetReserveIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetReserve", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetReserveIterator{contract: _Contract.contract, event: "SetReserve", logs: logs, sub: sub}, nil
}

// WatchSetReserve is a free log subscription operation binding the contract event 0x9241e21d4fa6dd55661c8694538bfc11a7d44d81d5766840312d2513f55fa67d.
//
// Solidity: event SetReserve(address indexed _receiver, uint256 _rate)
func (_Contract *ContractFilterer) WatchSetReserve(opts *bind.WatchOpts, sink chan<- *ContractSetReserve, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetReserve", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetReserve)
				if err := _Contract.contract.UnpackLog(event, "SetReserve", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetReserve is a log parse operation binding the contract event 0x9241e21d4fa6dd55661c8694538bfc11a7d44d81d5766840312d2513f55fa67d.
//
// Solidity: event SetReserve(address indexed _receiver, uint256 _rate)
func (_Contract *ContractFilterer) ParseSetReserve(log types.Log) (*ContractSetReserve, error) {
	event := new(ContractSetReserve)
	if err := _Contract.contract.UnpackLog(event, "SetReserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTradeClosedIterator is returned from FilterTradeClosed and is used to iterate over the raw logs and unpacked data for TradeClosed events raised by the Contract contract.
type ContractTradeClosedIterator struct {
	Event *ContractTradeClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTradeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTradeClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTradeClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTradeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTradeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTradeClosed represents a TradeClosed event raised by the Contract contract.
type ContractTradeClosed struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTradeClosed is a free log retrieval operation binding the contract event 0xbe2b332e8e34afe0840746913f0e88f88fbbda69b3c05c49a197de3f55941277.
//
// Solidity: event TradeClosed(uint256 orderId)
func (_Contract *ContractFilterer) FilterTradeClosed(opts *bind.FilterOpts) (*ContractTradeClosedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TradeClosed")
	if err != nil {
		return nil, err
	}
	return &ContractTradeClosedIterator{contract: _Contract.contract, event: "TradeClosed", logs: logs, sub: sub}, nil
}

// WatchTradeClosed is a free log subscription operation binding the contract event 0xbe2b332e8e34afe0840746913f0e88f88fbbda69b3c05c49a197de3f55941277.
//
// Solidity: event TradeClosed(uint256 orderId)
func (_Contract *ContractFilterer) WatchTradeClosed(opts *bind.WatchOpts, sink chan<- *ContractTradeClosed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TradeClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTradeClosed)
				if err := _Contract.contract.UnpackLog(event, "TradeClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradeClosed is a log parse operation binding the contract event 0xbe2b332e8e34afe0840746913f0e88f88fbbda69b3c05c49a197de3f55941277.
//
// Solidity: event TradeClosed(uint256 orderId)
func (_Contract *ContractFilterer) ParseTradeClosed(log types.Log) (*ContractTradeClosed, error) {
	event := new(ContractTradeClosed)
	if err := _Contract.contract.UnpackLog(event, "TradeClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Contract contract.
type ContractUnpausedIterator struct {
	Event *ContractUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnpaused represents a Unpaused event raised by the Contract contract.
type ContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contract *ContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractUnpausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractUnpausedIterator{contract: _Contract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contract *ContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnpaused)
				if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contract *ContractFilterer) ParseUnpaused(log types.Log) (*ContractUnpaused, error) {
	event := new(ContractUnpaused)
	if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateMarginIterator is returned from FilterUpdateMargin and is used to iterate over the raw logs and unpacked data for UpdateMargin events raised by the Contract contract.
type ContractUpdateMarginIterator struct {
	Event *ContractUpdateMargin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdateMarginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateMargin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdateMargin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdateMarginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateMarginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateMargin represents a UpdateMargin event raised by the Contract contract.
type ContractUpdateMargin struct {
	OrderId  *big.Int
	Amount   *big.Int
	IsAdd    bool
	Margin   *big.Int
	Leverage *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateMargin is a free log retrieval operation binding the contract event 0xb58ece2937681ab71ee2a4257af71abb8d43386650f2fd177ee49164f8a0a4ca.
//
// Solidity: event UpdateMargin(uint256 orderId, uint256 amount, bool isAdd, uint256 margin, uint256 leverage)
func (_Contract *ContractFilterer) FilterUpdateMargin(opts *bind.FilterOpts) (*ContractUpdateMarginIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateMargin")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateMarginIterator{contract: _Contract.contract, event: "UpdateMargin", logs: logs, sub: sub}, nil
}

// WatchUpdateMargin is a free log subscription operation binding the contract event 0xb58ece2937681ab71ee2a4257af71abb8d43386650f2fd177ee49164f8a0a4ca.
//
// Solidity: event UpdateMargin(uint256 orderId, uint256 amount, bool isAdd, uint256 margin, uint256 leverage)
func (_Contract *ContractFilterer) WatchUpdateMargin(opts *bind.WatchOpts, sink chan<- *ContractUpdateMargin) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateMargin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateMargin)
				if err := _Contract.contract.UnpackLog(event, "UpdateMargin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMargin is a log parse operation binding the contract event 0xb58ece2937681ab71ee2a4257af71abb8d43386650f2fd177ee49164f8a0a4ca.
//
// Solidity: event UpdateMargin(uint256 orderId, uint256 amount, bool isAdd, uint256 margin, uint256 leverage)
func (_Contract *ContractFilterer) ParseUpdateMargin(log types.Log) (*ContractUpdateMargin, error) {
	event := new(ContractUpdateMargin)
	if err := _Contract.contract.UnpackLog(event, "UpdateMargin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateOpenRequestIterator is returned from FilterUpdateOpenRequest and is used to iterate over the raw logs and unpacked data for UpdateOpenRequest events raised by the Contract contract.
type ContractUpdateOpenRequestIterator struct {
	Event *ContractUpdateOpenRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdateOpenRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateOpenRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdateOpenRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdateOpenRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateOpenRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateOpenRequest represents a UpdateOpenRequest event raised by the Contract contract.
type ContractUpdateOpenRequest struct {
	OrderId    *big.Int
	LimitPrice *big.Int
	Tp         *big.Int
	Sl         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateOpenRequest is a free log retrieval operation binding the contract event 0xed9036017ec7fc31a1ca97298d80b8219039778a809afd38dc6480cea3964a4b.
//
// Solidity: event UpdateOpenRequest(uint256 orderId, uint256 limitPrice, uint256 tp, uint256 sl)
func (_Contract *ContractFilterer) FilterUpdateOpenRequest(opts *bind.FilterOpts) (*ContractUpdateOpenRequestIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateOpenRequest")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateOpenRequestIterator{contract: _Contract.contract, event: "UpdateOpenRequest", logs: logs, sub: sub}, nil
}

// WatchUpdateOpenRequest is a free log subscription operation binding the contract event 0xed9036017ec7fc31a1ca97298d80b8219039778a809afd38dc6480cea3964a4b.
//
// Solidity: event UpdateOpenRequest(uint256 orderId, uint256 limitPrice, uint256 tp, uint256 sl)
func (_Contract *ContractFilterer) WatchUpdateOpenRequest(opts *bind.WatchOpts, sink chan<- *ContractUpdateOpenRequest) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateOpenRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateOpenRequest)
				if err := _Contract.contract.UnpackLog(event, "UpdateOpenRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateOpenRequest is a log parse operation binding the contract event 0xed9036017ec7fc31a1ca97298d80b8219039778a809afd38dc6480cea3964a4b.
//
// Solidity: event UpdateOpenRequest(uint256 orderId, uint256 limitPrice, uint256 tp, uint256 sl)
func (_Contract *ContractFilterer) ParseUpdateOpenRequest(log types.Log) (*ContractUpdateOpenRequest, error) {
	event := new(ContractUpdateOpenRequest)
	if err := _Contract.contract.UnpackLog(event, "UpdateOpenRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateTPAndSLIterator is returned from FilterUpdateTPAndSL and is used to iterate over the raw logs and unpacked data for UpdateTPAndSL events raised by the Contract contract.
type ContractUpdateTPAndSLIterator struct {
	Event *ContractUpdateTPAndSL // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdateTPAndSLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateTPAndSL)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdateTPAndSL)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdateTPAndSLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateTPAndSLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateTPAndSL represents a UpdateTPAndSL event raised by the Contract contract.
type ContractUpdateTPAndSL struct {
	OrderId *big.Int
	Tp      *big.Int
	Sl      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTPAndSL is a free log retrieval operation binding the contract event 0x3fea1f17d88eda811ab87655f92885a6ec5fc9fd390aca775b9327f43ad00d77.
//
// Solidity: event UpdateTPAndSL(uint256 orderId, uint256 tp, uint256 sl)
func (_Contract *ContractFilterer) FilterUpdateTPAndSL(opts *bind.FilterOpts) (*ContractUpdateTPAndSLIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateTPAndSL")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateTPAndSLIterator{contract: _Contract.contract, event: "UpdateTPAndSL", logs: logs, sub: sub}, nil
}

// WatchUpdateTPAndSL is a free log subscription operation binding the contract event 0x3fea1f17d88eda811ab87655f92885a6ec5fc9fd390aca775b9327f43ad00d77.
//
// Solidity: event UpdateTPAndSL(uint256 orderId, uint256 tp, uint256 sl)
func (_Contract *ContractFilterer) WatchUpdateTPAndSL(opts *bind.WatchOpts, sink chan<- *ContractUpdateTPAndSL) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateTPAndSL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateTPAndSL)
				if err := _Contract.contract.UnpackLog(event, "UpdateTPAndSL", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateTPAndSL is a log parse operation binding the contract event 0x3fea1f17d88eda811ab87655f92885a6ec5fc9fd390aca775b9327f43ad00d77.
//
// Solidity: event UpdateTPAndSL(uint256 orderId, uint256 tp, uint256 sl)
func (_Contract *ContractFilterer) ParseUpdateTPAndSL(log types.Log) (*ContractUpdateTPAndSL, error) {
	event := new(ContractUpdateTPAndSL)
	if err := _Contract.contract.UnpackLog(event, "UpdateTPAndSL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
