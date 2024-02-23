// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tinder_trading

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

// OpenTrade is an auto generated low-level Go binding around an user-defined struct.
type OpenTrade struct {
	Base           TradeBase
	OpenPrice      *big.Int
	LastUpdateTime *big.Int
}

// TinderTradingCloseParams is an auto generated low-level Go binding around an user-defined struct.
type TinderTradingCloseParams struct {
	OrderId    *big.Int
	UpdateData [][]byte
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

// TinderTradingMetaData contains all meta data concerning the TinderTrading contract.
var TinderTradingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"resut\",\"type\":\"bool\"}],\"name\":\"Callback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_closeMargin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingFee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rolloverFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closeFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCloseType\",\"name\":\"s\",\"type\":\"uint8\"}],\"name\":\"Close\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"base\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOpenTrade\",\"name\":\"t\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Open\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_feeHelper\",\"type\":\"address\"}],\"name\":\"SetFeeHelper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"SetOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pairInfo\",\"type\":\"address\"}],\"name\":\"SetPairInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidationP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_spreadReductionP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_callbackGasLimit\",\"type\":\"uint256\"}],\"name\":\"SetParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"SetReserve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_standardToken\",\"type\":\"address\"}],\"name\":\"SetStandardToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"SetVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callbackGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"internalType\":\"structTinderTrading.CloseParams[]\",\"name\":\"_params\",\"type\":\"tuple[]\"}],\"name\":\"closeTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeHelper\",\"outputs\":[{\"internalType\":\"contractIFeeHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getOpenTrade\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"base\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"}],\"internalType\":\"structOpenTrade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"name\":\"getPriceAfterImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getTradeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_standardToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"base\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"openTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairInfo\",\"outputs\":[{\"internalType\":\"contractIPairInfos\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeHelper\",\"type\":\"address\"}],\"name\":\"setFeeHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairInfo\",\"type\":\"address\"}],\"name\":\"setPairInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidationP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_callbackGasLimit\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_standardToken\",\"type\":\"address\"}],\"name\":\"setStandardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadReductionP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"enumCloseType\",\"name\":\"_stop\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"stopTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TinderTradingABI is the input ABI used to generate the binding from.
// Deprecated: Use TinderTradingMetaData.ABI instead.
var TinderTradingABI = TinderTradingMetaData.ABI

// TinderTrading is an auto generated Go binding around an Ethereum contract.
type TinderTrading struct {
	TinderTradingCaller     // Read-only binding to the contract
	TinderTradingTransactor // Write-only binding to the contract
	TinderTradingFilterer   // Log filterer for contract events
}

// TinderTradingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TinderTradingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinderTradingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TinderTradingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinderTradingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TinderTradingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinderTradingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TinderTradingSession struct {
	Contract     *TinderTrading    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TinderTradingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TinderTradingCallerSession struct {
	Contract *TinderTradingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TinderTradingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TinderTradingTransactorSession struct {
	Contract     *TinderTradingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TinderTradingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TinderTradingRaw struct {
	Contract *TinderTrading // Generic contract binding to access the raw methods on
}

// TinderTradingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TinderTradingCallerRaw struct {
	Contract *TinderTradingCaller // Generic read-only contract binding to access the raw methods on
}

// TinderTradingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TinderTradingTransactorRaw struct {
	Contract *TinderTradingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTinderTrading creates a new instance of TinderTrading, bound to a specific deployed contract.
func NewTinderTrading(address common.Address, backend bind.ContractBackend) (*TinderTrading, error) {
	contract, err := bindTinderTrading(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TinderTrading{TinderTradingCaller: TinderTradingCaller{contract: contract}, TinderTradingTransactor: TinderTradingTransactor{contract: contract}, TinderTradingFilterer: TinderTradingFilterer{contract: contract}}, nil
}

// NewTinderTradingCaller creates a new read-only instance of TinderTrading, bound to a specific deployed contract.
func NewTinderTradingCaller(address common.Address, caller bind.ContractCaller) (*TinderTradingCaller, error) {
	contract, err := bindTinderTrading(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TinderTradingCaller{contract: contract}, nil
}

// NewTinderTradingTransactor creates a new write-only instance of TinderTrading, bound to a specific deployed contract.
func NewTinderTradingTransactor(address common.Address, transactor bind.ContractTransactor) (*TinderTradingTransactor, error) {
	contract, err := bindTinderTrading(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TinderTradingTransactor{contract: contract}, nil
}

// NewTinderTradingFilterer creates a new log filterer instance of TinderTrading, bound to a specific deployed contract.
func NewTinderTradingFilterer(address common.Address, filterer bind.ContractFilterer) (*TinderTradingFilterer, error) {
	contract, err := bindTinderTrading(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TinderTradingFilterer{contract: contract}, nil
}

// bindTinderTrading binds a generic wrapper to an already deployed contract.
func bindTinderTrading(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TinderTradingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TinderTrading *TinderTradingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TinderTrading.Contract.TinderTradingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TinderTrading *TinderTradingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TinderTrading.Contract.TinderTradingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TinderTrading *TinderTradingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TinderTrading.Contract.TinderTradingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TinderTrading *TinderTradingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TinderTrading.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TinderTrading *TinderTradingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TinderTrading.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TinderTrading *TinderTradingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TinderTrading.Contract.contract.Transact(opts, method, params...)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_TinderTrading *TinderTradingCaller) CallbackGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "callbackGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_TinderTrading *TinderTradingSession) CallbackGasLimit() (*big.Int, error) {
	return _TinderTrading.Contract.CallbackGasLimit(&_TinderTrading.CallOpts)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_TinderTrading *TinderTradingCallerSession) CallbackGasLimit() (*big.Int, error) {
	return _TinderTrading.Contract.CallbackGasLimit(&_TinderTrading.CallOpts)
}

// FeeHelper is a free data retrieval call binding the contract method 0x18897bb7.
//
// Solidity: function feeHelper() view returns(address)
func (_TinderTrading *TinderTradingCaller) FeeHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "feeHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeHelper is a free data retrieval call binding the contract method 0x18897bb7.
//
// Solidity: function feeHelper() view returns(address)
func (_TinderTrading *TinderTradingSession) FeeHelper() (common.Address, error) {
	return _TinderTrading.Contract.FeeHelper(&_TinderTrading.CallOpts)
}

// FeeHelper is a free data retrieval call binding the contract method 0x18897bb7.
//
// Solidity: function feeHelper() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) FeeHelper() (common.Address, error) {
	return _TinderTrading.Contract.FeeHelper(&_TinderTrading.CallOpts)
}

// GetOpenTrade is a free data retrieval call binding the contract method 0x7404eede.
//
// Solidity: function getOpenTrade(uint256 orderId) view returns(((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256))
func (_TinderTrading *TinderTradingCaller) GetOpenTrade(opts *bind.CallOpts, orderId *big.Int) (OpenTrade, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "getOpenTrade", orderId)

	if err != nil {
		return *new(OpenTrade), err
	}

	out0 := *abi.ConvertType(out[0], new(OpenTrade)).(*OpenTrade)

	return out0, err

}

// GetOpenTrade is a free data retrieval call binding the contract method 0x7404eede.
//
// Solidity: function getOpenTrade(uint256 orderId) view returns(((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256))
func (_TinderTrading *TinderTradingSession) GetOpenTrade(orderId *big.Int) (OpenTrade, error) {
	return _TinderTrading.Contract.GetOpenTrade(&_TinderTrading.CallOpts, orderId)
}

// GetOpenTrade is a free data retrieval call binding the contract method 0x7404eede.
//
// Solidity: function getOpenTrade(uint256 orderId) view returns(((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256))
func (_TinderTrading *TinderTradingCallerSession) GetOpenTrade(orderId *big.Int) (OpenTrade, error) {
	return _TinderTrading.Contract.GetOpenTrade(&_TinderTrading.CallOpts, orderId)
}

// GetPriceAfterImpact is a free data retrieval call binding the contract method 0xa53fcde8.
//
// Solidity: function getPriceAfterImpact(uint256 price, uint256 pairIndex, uint256 position, bool isLong) view returns(uint256)
func (_TinderTrading *TinderTradingCaller) GetPriceAfterImpact(opts *bind.CallOpts, price *big.Int, pairIndex *big.Int, position *big.Int, isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "getPriceAfterImpact", price, pairIndex, position, isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceAfterImpact is a free data retrieval call binding the contract method 0xa53fcde8.
//
// Solidity: function getPriceAfterImpact(uint256 price, uint256 pairIndex, uint256 position, bool isLong) view returns(uint256)
func (_TinderTrading *TinderTradingSession) GetPriceAfterImpact(price *big.Int, pairIndex *big.Int, position *big.Int, isLong bool) (*big.Int, error) {
	return _TinderTrading.Contract.GetPriceAfterImpact(&_TinderTrading.CallOpts, price, pairIndex, position, isLong)
}

// GetPriceAfterImpact is a free data retrieval call binding the contract method 0xa53fcde8.
//
// Solidity: function getPriceAfterImpact(uint256 price, uint256 pairIndex, uint256 position, bool isLong) view returns(uint256)
func (_TinderTrading *TinderTradingCallerSession) GetPriceAfterImpact(price *big.Int, pairIndex *big.Int, position *big.Int, isLong bool) (*big.Int, error) {
	return _TinderTrading.Contract.GetPriceAfterImpact(&_TinderTrading.CallOpts, price, pairIndex, position, isLong)
}

// GetTradeValue is a free data retrieval call binding the contract method 0x72534890.
//
// Solidity: function getTradeValue(uint256 orderId, uint256 price) view returns(uint256)
func (_TinderTrading *TinderTradingCaller) GetTradeValue(opts *bind.CallOpts, orderId *big.Int, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "getTradeValue", orderId, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeValue is a free data retrieval call binding the contract method 0x72534890.
//
// Solidity: function getTradeValue(uint256 orderId, uint256 price) view returns(uint256)
func (_TinderTrading *TinderTradingSession) GetTradeValue(orderId *big.Int, price *big.Int) (*big.Int, error) {
	return _TinderTrading.Contract.GetTradeValue(&_TinderTrading.CallOpts, orderId, price)
}

// GetTradeValue is a free data retrieval call binding the contract method 0x72534890.
//
// Solidity: function getTradeValue(uint256 orderId, uint256 price) view returns(uint256)
func (_TinderTrading *TinderTradingCallerSession) GetTradeValue(orderId *big.Int, price *big.Int) (*big.Int, error) {
	return _TinderTrading.Contract.GetTradeValue(&_TinderTrading.CallOpts, orderId, price)
}

// LiquidationP is a free data retrieval call binding the contract method 0x615728f8.
//
// Solidity: function liquidationP() view returns(uint256)
func (_TinderTrading *TinderTradingCaller) LiquidationP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "liquidationP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationP is a free data retrieval call binding the contract method 0x615728f8.
//
// Solidity: function liquidationP() view returns(uint256)
func (_TinderTrading *TinderTradingSession) LiquidationP() (*big.Int, error) {
	return _TinderTrading.Contract.LiquidationP(&_TinderTrading.CallOpts)
}

// LiquidationP is a free data retrieval call binding the contract method 0x615728f8.
//
// Solidity: function liquidationP() view returns(uint256)
func (_TinderTrading *TinderTradingCallerSession) LiquidationP() (*big.Int, error) {
	return _TinderTrading.Contract.LiquidationP(&_TinderTrading.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_TinderTrading *TinderTradingCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_TinderTrading *TinderTradingSession) Oracle() (common.Address, error) {
	return _TinderTrading.Contract.Oracle(&_TinderTrading.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) Oracle() (common.Address, error) {
	return _TinderTrading.Contract.Oracle(&_TinderTrading.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TinderTrading *TinderTradingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TinderTrading *TinderTradingSession) Owner() (common.Address, error) {
	return _TinderTrading.Contract.Owner(&_TinderTrading.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) Owner() (common.Address, error) {
	return _TinderTrading.Contract.Owner(&_TinderTrading.CallOpts)
}

// PairInfo is a free data retrieval call binding the contract method 0xfc2a5b1d.
//
// Solidity: function pairInfo() view returns(address)
func (_TinderTrading *TinderTradingCaller) PairInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "pairInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairInfo is a free data retrieval call binding the contract method 0xfc2a5b1d.
//
// Solidity: function pairInfo() view returns(address)
func (_TinderTrading *TinderTradingSession) PairInfo() (common.Address, error) {
	return _TinderTrading.Contract.PairInfo(&_TinderTrading.CallOpts)
}

// PairInfo is a free data retrieval call binding the contract method 0xfc2a5b1d.
//
// Solidity: function pairInfo() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) PairInfo() (common.Address, error) {
	return _TinderTrading.Contract.PairInfo(&_TinderTrading.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TinderTrading *TinderTradingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TinderTrading *TinderTradingSession) Paused() (bool, error) {
	return _TinderTrading.Contract.Paused(&_TinderTrading.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TinderTrading *TinderTradingCallerSession) Paused() (bool, error) {
	return _TinderTrading.Contract.Paused(&_TinderTrading.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TinderTrading *TinderTradingCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TinderTrading *TinderTradingSession) PendingOwner() (common.Address, error) {
	return _TinderTrading.Contract.PendingOwner(&_TinderTrading.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) PendingOwner() (common.Address, error) {
	return _TinderTrading.Contract.PendingOwner(&_TinderTrading.CallOpts)
}

// ReserveRate is a free data retrieval call binding the contract method 0x58d7bf80.
//
// Solidity: function reserveRate() view returns(uint256)
func (_TinderTrading *TinderTradingCaller) ReserveRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "reserveRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveRate is a free data retrieval call binding the contract method 0x58d7bf80.
//
// Solidity: function reserveRate() view returns(uint256)
func (_TinderTrading *TinderTradingSession) ReserveRate() (*big.Int, error) {
	return _TinderTrading.Contract.ReserveRate(&_TinderTrading.CallOpts)
}

// ReserveRate is a free data retrieval call binding the contract method 0x58d7bf80.
//
// Solidity: function reserveRate() view returns(uint256)
func (_TinderTrading *TinderTradingCallerSession) ReserveRate() (*big.Int, error) {
	return _TinderTrading.Contract.ReserveRate(&_TinderTrading.CallOpts)
}

// ReserveReceiver is a free data retrieval call binding the contract method 0xbb1c7c2a.
//
// Solidity: function reserveReceiver() view returns(address)
func (_TinderTrading *TinderTradingCaller) ReserveReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "reserveReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveReceiver is a free data retrieval call binding the contract method 0xbb1c7c2a.
//
// Solidity: function reserveReceiver() view returns(address)
func (_TinderTrading *TinderTradingSession) ReserveReceiver() (common.Address, error) {
	return _TinderTrading.Contract.ReserveReceiver(&_TinderTrading.CallOpts)
}

// ReserveReceiver is a free data retrieval call binding the contract method 0xbb1c7c2a.
//
// Solidity: function reserveReceiver() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) ReserveReceiver() (common.Address, error) {
	return _TinderTrading.Contract.ReserveReceiver(&_TinderTrading.CallOpts)
}

// SpreadReductionP is a free data retrieval call binding the contract method 0xefbc5fc0.
//
// Solidity: function spreadReductionP() view returns(uint256)
func (_TinderTrading *TinderTradingCaller) SpreadReductionP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "spreadReductionP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadReductionP is a free data retrieval call binding the contract method 0xefbc5fc0.
//
// Solidity: function spreadReductionP() view returns(uint256)
func (_TinderTrading *TinderTradingSession) SpreadReductionP() (*big.Int, error) {
	return _TinderTrading.Contract.SpreadReductionP(&_TinderTrading.CallOpts)
}

// SpreadReductionP is a free data retrieval call binding the contract method 0xefbc5fc0.
//
// Solidity: function spreadReductionP() view returns(uint256)
func (_TinderTrading *TinderTradingCallerSession) SpreadReductionP() (*big.Int, error) {
	return _TinderTrading.Contract.SpreadReductionP(&_TinderTrading.CallOpts)
}

// StandardToken is a free data retrieval call binding the contract method 0x4c7d7d3f.
//
// Solidity: function standardToken() view returns(address)
func (_TinderTrading *TinderTradingCaller) StandardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "standardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StandardToken is a free data retrieval call binding the contract method 0x4c7d7d3f.
//
// Solidity: function standardToken() view returns(address)
func (_TinderTrading *TinderTradingSession) StandardToken() (common.Address, error) {
	return _TinderTrading.Contract.StandardToken(&_TinderTrading.CallOpts)
}

// StandardToken is a free data retrieval call binding the contract method 0x4c7d7d3f.
//
// Solidity: function standardToken() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) StandardToken() (common.Address, error) {
	return _TinderTrading.Contract.StandardToken(&_TinderTrading.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_TinderTrading *TinderTradingCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TinderTrading.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_TinderTrading *TinderTradingSession) Vault() (common.Address, error) {
	return _TinderTrading.Contract.Vault(&_TinderTrading.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_TinderTrading *TinderTradingCallerSession) Vault() (common.Address, error) {
	return _TinderTrading.Contract.Vault(&_TinderTrading.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TinderTrading *TinderTradingTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TinderTrading *TinderTradingSession) AcceptOwnership() (*types.Transaction, error) {
	return _TinderTrading.Contract.AcceptOwnership(&_TinderTrading.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TinderTrading *TinderTradingTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TinderTrading.Contract.AcceptOwnership(&_TinderTrading.TransactOpts)
}

// CloseTrade is a paid mutator transaction binding the contract method 0xbb308bba.
//
// Solidity: function closeTrade((uint256,bytes[])[] _params) returns()
func (_TinderTrading *TinderTradingTransactor) CloseTrade(opts *bind.TransactOpts, _params []TinderTradingCloseParams) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "closeTrade", _params)
}

// CloseTrade is a paid mutator transaction binding the contract method 0xbb308bba.
//
// Solidity: function closeTrade((uint256,bytes[])[] _params) returns()
func (_TinderTrading *TinderTradingSession) CloseTrade(_params []TinderTradingCloseParams) (*types.Transaction, error) {
	return _TinderTrading.Contract.CloseTrade(&_TinderTrading.TransactOpts, _params)
}

// CloseTrade is a paid mutator transaction binding the contract method 0xbb308bba.
//
// Solidity: function closeTrade((uint256,bytes[])[] _params) returns()
func (_TinderTrading *TinderTradingTransactorSession) CloseTrade(_params []TinderTradingCloseParams) (*types.Transaction, error) {
	return _TinderTrading.Contract.CloseTrade(&_TinderTrading.TransactOpts, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _standardToken) returns()
func (_TinderTrading *TinderTradingTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _standardToken common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "initialize", owner, _standardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _standardToken) returns()
func (_TinderTrading *TinderTradingSession) Initialize(owner common.Address, _standardToken common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.Initialize(&_TinderTrading.TransactOpts, owner, _standardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _standardToken) returns()
func (_TinderTrading *TinderTradingTransactorSession) Initialize(owner common.Address, _standardToken common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.Initialize(&_TinderTrading.TransactOpts, owner, _standardToken)
}

// OpenTrade is a paid mutator transaction binding the contract method 0x3aa24f85.
//
// Solidity: function openTrade(uint256 _orderId, (address,uint256,uint256,bool,uint256,uint256,uint256) base, bytes[] updateData, uint256 slippage, address callTarget, uint256 fee) returns()
func (_TinderTrading *TinderTradingTransactor) OpenTrade(opts *bind.TransactOpts, _orderId *big.Int, base TradeBase, updateData [][]byte, slippage *big.Int, callTarget common.Address, fee *big.Int) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "openTrade", _orderId, base, updateData, slippage, callTarget, fee)
}

// OpenTrade is a paid mutator transaction binding the contract method 0x3aa24f85.
//
// Solidity: function openTrade(uint256 _orderId, (address,uint256,uint256,bool,uint256,uint256,uint256) base, bytes[] updateData, uint256 slippage, address callTarget, uint256 fee) returns()
func (_TinderTrading *TinderTradingSession) OpenTrade(_orderId *big.Int, base TradeBase, updateData [][]byte, slippage *big.Int, callTarget common.Address, fee *big.Int) (*types.Transaction, error) {
	return _TinderTrading.Contract.OpenTrade(&_TinderTrading.TransactOpts, _orderId, base, updateData, slippage, callTarget, fee)
}

// OpenTrade is a paid mutator transaction binding the contract method 0x3aa24f85.
//
// Solidity: function openTrade(uint256 _orderId, (address,uint256,uint256,bool,uint256,uint256,uint256) base, bytes[] updateData, uint256 slippage, address callTarget, uint256 fee) returns()
func (_TinderTrading *TinderTradingTransactorSession) OpenTrade(_orderId *big.Int, base TradeBase, updateData [][]byte, slippage *big.Int, callTarget common.Address, fee *big.Int) (*types.Transaction, error) {
	return _TinderTrading.Contract.OpenTrade(&_TinderTrading.TransactOpts, _orderId, base, updateData, slippage, callTarget, fee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TinderTrading *TinderTradingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TinderTrading *TinderTradingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TinderTrading.Contract.RenounceOwnership(&_TinderTrading.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TinderTrading *TinderTradingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TinderTrading.Contract.RenounceOwnership(&_TinderTrading.TransactOpts)
}

// SetFeeHelper is a paid mutator transaction binding the contract method 0xcbc0bf0e.
//
// Solidity: function setFeeHelper(address _feeHelper) returns()
func (_TinderTrading *TinderTradingTransactor) SetFeeHelper(opts *bind.TransactOpts, _feeHelper common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setFeeHelper", _feeHelper)
}

// SetFeeHelper is a paid mutator transaction binding the contract method 0xcbc0bf0e.
//
// Solidity: function setFeeHelper(address _feeHelper) returns()
func (_TinderTrading *TinderTradingSession) SetFeeHelper(_feeHelper common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetFeeHelper(&_TinderTrading.TransactOpts, _feeHelper)
}

// SetFeeHelper is a paid mutator transaction binding the contract method 0xcbc0bf0e.
//
// Solidity: function setFeeHelper(address _feeHelper) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetFeeHelper(_feeHelper common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetFeeHelper(&_TinderTrading.TransactOpts, _feeHelper)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_TinderTrading *TinderTradingTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_TinderTrading *TinderTradingSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetOracle(&_TinderTrading.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetOracle(&_TinderTrading.TransactOpts, _oracle)
}

// SetPairInfo is a paid mutator transaction binding the contract method 0xf49a9b8a.
//
// Solidity: function setPairInfo(address _pairInfo) returns()
func (_TinderTrading *TinderTradingTransactor) SetPairInfo(opts *bind.TransactOpts, _pairInfo common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setPairInfo", _pairInfo)
}

// SetPairInfo is a paid mutator transaction binding the contract method 0xf49a9b8a.
//
// Solidity: function setPairInfo(address _pairInfo) returns()
func (_TinderTrading *TinderTradingSession) SetPairInfo(_pairInfo common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetPairInfo(&_TinderTrading.TransactOpts, _pairInfo)
}

// SetPairInfo is a paid mutator transaction binding the contract method 0xf49a9b8a.
//
// Solidity: function setPairInfo(address _pairInfo) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetPairInfo(_pairInfo common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetPairInfo(&_TinderTrading.TransactOpts, _pairInfo)
}

// SetParams is a paid mutator transaction binding the contract method 0x5a0ce676.
//
// Solidity: function setParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _callbackGasLimit) returns()
func (_TinderTrading *TinderTradingTransactor) SetParams(opts *bind.TransactOpts, _liquidationP *big.Int, _spreadReductionP *big.Int, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setParams", _liquidationP, _spreadReductionP, _callbackGasLimit)
}

// SetParams is a paid mutator transaction binding the contract method 0x5a0ce676.
//
// Solidity: function setParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _callbackGasLimit) returns()
func (_TinderTrading *TinderTradingSession) SetParams(_liquidationP *big.Int, _spreadReductionP *big.Int, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetParams(&_TinderTrading.TransactOpts, _liquidationP, _spreadReductionP, _callbackGasLimit)
}

// SetParams is a paid mutator transaction binding the contract method 0x5a0ce676.
//
// Solidity: function setParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _callbackGasLimit) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetParams(_liquidationP *big.Int, _spreadReductionP *big.Int, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetParams(&_TinderTrading.TransactOpts, _liquidationP, _spreadReductionP, _callbackGasLimit)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address _receiver, uint256 _rate) returns()
func (_TinderTrading *TinderTradingTransactor) SetReserve(opts *bind.TransactOpts, _receiver common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setReserve", _receiver, _rate)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address _receiver, uint256 _rate) returns()
func (_TinderTrading *TinderTradingSession) SetReserve(_receiver common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetReserve(&_TinderTrading.TransactOpts, _receiver, _rate)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address _receiver, uint256 _rate) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetReserve(_receiver common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetReserve(&_TinderTrading.TransactOpts, _receiver, _rate)
}

// SetStandardToken is a paid mutator transaction binding the contract method 0xdd6d012e.
//
// Solidity: function setStandardToken(address _standardToken) returns()
func (_TinderTrading *TinderTradingTransactor) SetStandardToken(opts *bind.TransactOpts, _standardToken common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setStandardToken", _standardToken)
}

// SetStandardToken is a paid mutator transaction binding the contract method 0xdd6d012e.
//
// Solidity: function setStandardToken(address _standardToken) returns()
func (_TinderTrading *TinderTradingSession) SetStandardToken(_standardToken common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetStandardToken(&_TinderTrading.TransactOpts, _standardToken)
}

// SetStandardToken is a paid mutator transaction binding the contract method 0xdd6d012e.
//
// Solidity: function setStandardToken(address _standardToken) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetStandardToken(_standardToken common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetStandardToken(&_TinderTrading.TransactOpts, _standardToken)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_TinderTrading *TinderTradingTransactor) SetVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "setVault", _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_TinderTrading *TinderTradingSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetVault(&_TinderTrading.TransactOpts, _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_TinderTrading *TinderTradingTransactorSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.SetVault(&_TinderTrading.TransactOpts, _vault)
}

// StopTrade is a paid mutator transaction binding the contract method 0xcc83bfd2.
//
// Solidity: function stopTrade(uint256 _orderId, uint8 _stop, bytes[] updateData) returns()
func (_TinderTrading *TinderTradingTransactor) StopTrade(opts *bind.TransactOpts, _orderId *big.Int, _stop uint8, updateData [][]byte) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "stopTrade", _orderId, _stop, updateData)
}

// StopTrade is a paid mutator transaction binding the contract method 0xcc83bfd2.
//
// Solidity: function stopTrade(uint256 _orderId, uint8 _stop, bytes[] updateData) returns()
func (_TinderTrading *TinderTradingSession) StopTrade(_orderId *big.Int, _stop uint8, updateData [][]byte) (*types.Transaction, error) {
	return _TinderTrading.Contract.StopTrade(&_TinderTrading.TransactOpts, _orderId, _stop, updateData)
}

// StopTrade is a paid mutator transaction binding the contract method 0xcc83bfd2.
//
// Solidity: function stopTrade(uint256 _orderId, uint8 _stop, bytes[] updateData) returns()
func (_TinderTrading *TinderTradingTransactorSession) StopTrade(_orderId *big.Int, _stop uint8, updateData [][]byte) (*types.Transaction, error) {
	return _TinderTrading.Contract.StopTrade(&_TinderTrading.TransactOpts, _orderId, _stop, updateData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TinderTrading *TinderTradingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TinderTrading.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TinderTrading *TinderTradingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.TransferOwnership(&_TinderTrading.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TinderTrading *TinderTradingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TinderTrading.Contract.TransferOwnership(&_TinderTrading.TransactOpts, newOwner)
}

// TinderTradingCallbackIterator is returned from FilterCallback and is used to iterate over the raw logs and unpacked data for Callback events raised by the TinderTrading contract.
type TinderTradingCallbackIterator struct {
	Event *TinderTradingCallback // Event containing the contract specifics and raw log

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
func (it *TinderTradingCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingCallback)
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
		it.Event = new(TinderTradingCallback)
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
func (it *TinderTradingCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingCallback represents a Callback event raised by the TinderTrading contract.
type TinderTradingCallback struct {
	Target common.Address
	Resut  bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallback is a free log retrieval operation binding the contract event 0x46ddbd62fc1a7626fe9c43026cb0694aec0b031fe81ac66fb4cfe9381dc6fe72.
//
// Solidity: event Callback(address target, bool resut)
func (_TinderTrading *TinderTradingFilterer) FilterCallback(opts *bind.FilterOpts) (*TinderTradingCallbackIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return &TinderTradingCallbackIterator{contract: _TinderTrading.contract, event: "Callback", logs: logs, sub: sub}, nil
}

// WatchCallback is a free log subscription operation binding the contract event 0x46ddbd62fc1a7626fe9c43026cb0694aec0b031fe81ac66fb4cfe9381dc6fe72.
//
// Solidity: event Callback(address target, bool resut)
func (_TinderTrading *TinderTradingFilterer) WatchCallback(opts *bind.WatchOpts, sink chan<- *TinderTradingCallback) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingCallback)
				if err := _TinderTrading.contract.UnpackLog(event, "Callback", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseCallback(log types.Log) (*TinderTradingCallback, error) {
	event := new(TinderTradingCallback)
	if err := _TinderTrading.contract.UnpackLog(event, "Callback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingCloseIterator is returned from FilterClose and is used to iterate over the raw logs and unpacked data for Close events raised by the TinderTrading contract.
type TinderTradingCloseIterator struct {
	Event *TinderTradingClose // Event containing the contract specifics and raw log

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
func (it *TinderTradingCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingClose)
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
		it.Event = new(TinderTradingClose)
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
func (it *TinderTradingCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingClose represents a Close event raised by the TinderTrading contract.
type TinderTradingClose struct {
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
func (_TinderTrading *TinderTradingFilterer) FilterClose(opts *bind.FilterOpts) (*TinderTradingCloseIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return &TinderTradingCloseIterator{contract: _TinderTrading.contract, event: "Close", logs: logs, sub: sub}, nil
}

// WatchClose is a free log subscription operation binding the contract event 0x24b27bc2a65d006b06d5ee8450abfa71760f2aac880fcf5555af1e692cfaf539.
//
// Solidity: event Close(uint256 orderId, uint256 closePrice, uint256 _closeMargin, int256 fundingFee, uint256 rolloverFee, uint256 closeFee, uint256 afterFee, uint8 s)
func (_TinderTrading *TinderTradingFilterer) WatchClose(opts *bind.WatchOpts, sink chan<- *TinderTradingClose) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingClose)
				if err := _TinderTrading.contract.UnpackLog(event, "Close", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseClose(log types.Log) (*TinderTradingClose, error) {
	event := new(TinderTradingClose)
	if err := _TinderTrading.contract.UnpackLog(event, "Close", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TinderTrading contract.
type TinderTradingInitializedIterator struct {
	Event *TinderTradingInitialized // Event containing the contract specifics and raw log

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
func (it *TinderTradingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingInitialized)
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
		it.Event = new(TinderTradingInitialized)
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
func (it *TinderTradingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingInitialized represents a Initialized event raised by the TinderTrading contract.
type TinderTradingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TinderTrading *TinderTradingFilterer) FilterInitialized(opts *bind.FilterOpts) (*TinderTradingInitializedIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TinderTradingInitializedIterator{contract: _TinderTrading.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TinderTrading *TinderTradingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TinderTradingInitialized) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingInitialized)
				if err := _TinderTrading.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseInitialized(log types.Log) (*TinderTradingInitialized, error) {
	event := new(TinderTradingInitialized)
	if err := _TinderTrading.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingOpenIterator is returned from FilterOpen and is used to iterate over the raw logs and unpacked data for Open events raised by the TinderTrading contract.
type TinderTradingOpenIterator struct {
	Event *TinderTradingOpen // Event containing the contract specifics and raw log

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
func (it *TinderTradingOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingOpen)
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
		it.Event = new(TinderTradingOpen)
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
func (it *TinderTradingOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingOpen represents a Open event raised by the TinderTrading contract.
type TinderTradingOpen struct {
	OrderId *big.Int
	T       OpenTrade
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpen is a free log retrieval operation binding the contract event 0x09c087e430ee6913868eafe26e6d9da87f95984c61bde74e50e28ecd09da890a.
//
// Solidity: event Open(uint256 orderId, ((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256) t, uint256 fee)
func (_TinderTrading *TinderTradingFilterer) FilterOpen(opts *bind.FilterOpts) (*TinderTradingOpenIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "Open")
	if err != nil {
		return nil, err
	}
	return &TinderTradingOpenIterator{contract: _TinderTrading.contract, event: "Open", logs: logs, sub: sub}, nil
}

// WatchOpen is a free log subscription operation binding the contract event 0x09c087e430ee6913868eafe26e6d9da87f95984c61bde74e50e28ecd09da890a.
//
// Solidity: event Open(uint256 orderId, ((address,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256) t, uint256 fee)
func (_TinderTrading *TinderTradingFilterer) WatchOpen(opts *bind.WatchOpts, sink chan<- *TinderTradingOpen) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "Open")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingOpen)
				if err := _TinderTrading.contract.UnpackLog(event, "Open", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseOpen(log types.Log) (*TinderTradingOpen, error) {
	event := new(TinderTradingOpen)
	if err := _TinderTrading.contract.UnpackLog(event, "Open", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the TinderTrading contract.
type TinderTradingOwnershipTransferStartedIterator struct {
	Event *TinderTradingOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *TinderTradingOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingOwnershipTransferStarted)
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
		it.Event = new(TinderTradingOwnershipTransferStarted)
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
func (it *TinderTradingOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the TinderTrading contract.
type TinderTradingOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TinderTrading *TinderTradingFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TinderTradingOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TinderTradingOwnershipTransferStartedIterator{contract: _TinderTrading.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TinderTrading *TinderTradingFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *TinderTradingOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingOwnershipTransferStarted)
				if err := _TinderTrading.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseOwnershipTransferStarted(log types.Log) (*TinderTradingOwnershipTransferStarted, error) {
	event := new(TinderTradingOwnershipTransferStarted)
	if err := _TinderTrading.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TinderTrading contract.
type TinderTradingOwnershipTransferredIterator struct {
	Event *TinderTradingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TinderTradingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingOwnershipTransferred)
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
		it.Event = new(TinderTradingOwnershipTransferred)
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
func (it *TinderTradingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingOwnershipTransferred represents a OwnershipTransferred event raised by the TinderTrading contract.
type TinderTradingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TinderTrading *TinderTradingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TinderTradingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TinderTradingOwnershipTransferredIterator{contract: _TinderTrading.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TinderTrading *TinderTradingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TinderTradingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingOwnershipTransferred)
				if err := _TinderTrading.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseOwnershipTransferred(log types.Log) (*TinderTradingOwnershipTransferred, error) {
	event := new(TinderTradingOwnershipTransferred)
	if err := _TinderTrading.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TinderTrading contract.
type TinderTradingPausedIterator struct {
	Event *TinderTradingPaused // Event containing the contract specifics and raw log

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
func (it *TinderTradingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingPaused)
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
		it.Event = new(TinderTradingPaused)
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
func (it *TinderTradingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingPaused represents a Paused event raised by the TinderTrading contract.
type TinderTradingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TinderTrading *TinderTradingFilterer) FilterPaused(opts *bind.FilterOpts) (*TinderTradingPausedIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TinderTradingPausedIterator{contract: _TinderTrading.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TinderTrading *TinderTradingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TinderTradingPaused) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingPaused)
				if err := _TinderTrading.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParsePaused(log types.Log) (*TinderTradingPaused, error) {
	event := new(TinderTradingPaused)
	if err := _TinderTrading.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetFeeHelperIterator is returned from FilterSetFeeHelper and is used to iterate over the raw logs and unpacked data for SetFeeHelper events raised by the TinderTrading contract.
type TinderTradingSetFeeHelperIterator struct {
	Event *TinderTradingSetFeeHelper // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetFeeHelperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetFeeHelper)
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
		it.Event = new(TinderTradingSetFeeHelper)
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
func (it *TinderTradingSetFeeHelperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetFeeHelperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetFeeHelper represents a SetFeeHelper event raised by the TinderTrading contract.
type TinderTradingSetFeeHelper struct {
	FeeHelper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetFeeHelper is a free log retrieval operation binding the contract event 0x052c40aee979e6c1f16304b00f8268ca1466d6f9acdadfa68c0d5b0f7aa1751c.
//
// Solidity: event SetFeeHelper(address _feeHelper)
func (_TinderTrading *TinderTradingFilterer) FilterSetFeeHelper(opts *bind.FilterOpts) (*TinderTradingSetFeeHelperIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetFeeHelper")
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetFeeHelperIterator{contract: _TinderTrading.contract, event: "SetFeeHelper", logs: logs, sub: sub}, nil
}

// WatchSetFeeHelper is a free log subscription operation binding the contract event 0x052c40aee979e6c1f16304b00f8268ca1466d6f9acdadfa68c0d5b0f7aa1751c.
//
// Solidity: event SetFeeHelper(address _feeHelper)
func (_TinderTrading *TinderTradingFilterer) WatchSetFeeHelper(opts *bind.WatchOpts, sink chan<- *TinderTradingSetFeeHelper) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetFeeHelper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetFeeHelper)
				if err := _TinderTrading.contract.UnpackLog(event, "SetFeeHelper", log); err != nil {
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

// ParseSetFeeHelper is a log parse operation binding the contract event 0x052c40aee979e6c1f16304b00f8268ca1466d6f9acdadfa68c0d5b0f7aa1751c.
//
// Solidity: event SetFeeHelper(address _feeHelper)
func (_TinderTrading *TinderTradingFilterer) ParseSetFeeHelper(log types.Log) (*TinderTradingSetFeeHelper, error) {
	event := new(TinderTradingSetFeeHelper)
	if err := _TinderTrading.contract.UnpackLog(event, "SetFeeHelper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetOracleIterator is returned from FilterSetOracle and is used to iterate over the raw logs and unpacked data for SetOracle events raised by the TinderTrading contract.
type TinderTradingSetOracleIterator struct {
	Event *TinderTradingSetOracle // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetOracle)
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
		it.Event = new(TinderTradingSetOracle)
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
func (it *TinderTradingSetOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetOracle represents a SetOracle event raised by the TinderTrading contract.
type TinderTradingSetOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOracle is a free log retrieval operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address _oracle)
func (_TinderTrading *TinderTradingFilterer) FilterSetOracle(opts *bind.FilterOpts) (*TinderTradingSetOracleIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetOracle")
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetOracleIterator{contract: _TinderTrading.contract, event: "SetOracle", logs: logs, sub: sub}, nil
}

// WatchSetOracle is a free log subscription operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address _oracle)
func (_TinderTrading *TinderTradingFilterer) WatchSetOracle(opts *bind.WatchOpts, sink chan<- *TinderTradingSetOracle) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetOracle)
				if err := _TinderTrading.contract.UnpackLog(event, "SetOracle", log); err != nil {
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

// ParseSetOracle is a log parse operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address _oracle)
func (_TinderTrading *TinderTradingFilterer) ParseSetOracle(log types.Log) (*TinderTradingSetOracle, error) {
	event := new(TinderTradingSetOracle)
	if err := _TinderTrading.contract.UnpackLog(event, "SetOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetPairInfoIterator is returned from FilterSetPairInfo and is used to iterate over the raw logs and unpacked data for SetPairInfo events raised by the TinderTrading contract.
type TinderTradingSetPairInfoIterator struct {
	Event *TinderTradingSetPairInfo // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetPairInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetPairInfo)
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
		it.Event = new(TinderTradingSetPairInfo)
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
func (it *TinderTradingSetPairInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetPairInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetPairInfo represents a SetPairInfo event raised by the TinderTrading contract.
type TinderTradingSetPairInfo struct {
	PairInfo common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPairInfo is a free log retrieval operation binding the contract event 0x498f61784682d2a34a04afb8c518b7d38e0e59907bd66601a0a41ed6fb79ff9e.
//
// Solidity: event SetPairInfo(address _pairInfo)
func (_TinderTrading *TinderTradingFilterer) FilterSetPairInfo(opts *bind.FilterOpts) (*TinderTradingSetPairInfoIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetPairInfo")
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetPairInfoIterator{contract: _TinderTrading.contract, event: "SetPairInfo", logs: logs, sub: sub}, nil
}

// WatchSetPairInfo is a free log subscription operation binding the contract event 0x498f61784682d2a34a04afb8c518b7d38e0e59907bd66601a0a41ed6fb79ff9e.
//
// Solidity: event SetPairInfo(address _pairInfo)
func (_TinderTrading *TinderTradingFilterer) WatchSetPairInfo(opts *bind.WatchOpts, sink chan<- *TinderTradingSetPairInfo) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetPairInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetPairInfo)
				if err := _TinderTrading.contract.UnpackLog(event, "SetPairInfo", log); err != nil {
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

// ParseSetPairInfo is a log parse operation binding the contract event 0x498f61784682d2a34a04afb8c518b7d38e0e59907bd66601a0a41ed6fb79ff9e.
//
// Solidity: event SetPairInfo(address _pairInfo)
func (_TinderTrading *TinderTradingFilterer) ParseSetPairInfo(log types.Log) (*TinderTradingSetPairInfo, error) {
	event := new(TinderTradingSetPairInfo)
	if err := _TinderTrading.contract.UnpackLog(event, "SetPairInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetParamsIterator is returned from FilterSetParams and is used to iterate over the raw logs and unpacked data for SetParams events raised by the TinderTrading contract.
type TinderTradingSetParamsIterator struct {
	Event *TinderTradingSetParams // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetParams)
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
		it.Event = new(TinderTradingSetParams)
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
func (it *TinderTradingSetParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetParams represents a SetParams event raised by the TinderTrading contract.
type TinderTradingSetParams struct {
	LiquidationP     *big.Int
	SpreadReductionP *big.Int
	CallbackGasLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetParams is a free log retrieval operation binding the contract event 0x5d787faa8b0e4d4e36b3fd360d484228b7dbaac8dfcb7a57fa4366e6d65456a2.
//
// Solidity: event SetParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _callbackGasLimit)
func (_TinderTrading *TinderTradingFilterer) FilterSetParams(opts *bind.FilterOpts) (*TinderTradingSetParamsIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetParams")
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetParamsIterator{contract: _TinderTrading.contract, event: "SetParams", logs: logs, sub: sub}, nil
}

// WatchSetParams is a free log subscription operation binding the contract event 0x5d787faa8b0e4d4e36b3fd360d484228b7dbaac8dfcb7a57fa4366e6d65456a2.
//
// Solidity: event SetParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _callbackGasLimit)
func (_TinderTrading *TinderTradingFilterer) WatchSetParams(opts *bind.WatchOpts, sink chan<- *TinderTradingSetParams) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetParams)
				if err := _TinderTrading.contract.UnpackLog(event, "SetParams", log); err != nil {
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

// ParseSetParams is a log parse operation binding the contract event 0x5d787faa8b0e4d4e36b3fd360d484228b7dbaac8dfcb7a57fa4366e6d65456a2.
//
// Solidity: event SetParams(uint256 _liquidationP, uint256 _spreadReductionP, uint256 _callbackGasLimit)
func (_TinderTrading *TinderTradingFilterer) ParseSetParams(log types.Log) (*TinderTradingSetParams, error) {
	event := new(TinderTradingSetParams)
	if err := _TinderTrading.contract.UnpackLog(event, "SetParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetReserveIterator is returned from FilterSetReserve and is used to iterate over the raw logs and unpacked data for SetReserve events raised by the TinderTrading contract.
type TinderTradingSetReserveIterator struct {
	Event *TinderTradingSetReserve // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetReserve)
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
		it.Event = new(TinderTradingSetReserve)
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
func (it *TinderTradingSetReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetReserve represents a SetReserve event raised by the TinderTrading contract.
type TinderTradingSetReserve struct {
	Receiver common.Address
	Rate     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetReserve is a free log retrieval operation binding the contract event 0x9241e21d4fa6dd55661c8694538bfc11a7d44d81d5766840312d2513f55fa67d.
//
// Solidity: event SetReserve(address indexed _receiver, uint256 _rate)
func (_TinderTrading *TinderTradingFilterer) FilterSetReserve(opts *bind.FilterOpts, _receiver []common.Address) (*TinderTradingSetReserveIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetReserve", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetReserveIterator{contract: _TinderTrading.contract, event: "SetReserve", logs: logs, sub: sub}, nil
}

// WatchSetReserve is a free log subscription operation binding the contract event 0x9241e21d4fa6dd55661c8694538bfc11a7d44d81d5766840312d2513f55fa67d.
//
// Solidity: event SetReserve(address indexed _receiver, uint256 _rate)
func (_TinderTrading *TinderTradingFilterer) WatchSetReserve(opts *bind.WatchOpts, sink chan<- *TinderTradingSetReserve, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetReserve", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetReserve)
				if err := _TinderTrading.contract.UnpackLog(event, "SetReserve", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseSetReserve(log types.Log) (*TinderTradingSetReserve, error) {
	event := new(TinderTradingSetReserve)
	if err := _TinderTrading.contract.UnpackLog(event, "SetReserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetStandardTokenIterator is returned from FilterSetStandardToken and is used to iterate over the raw logs and unpacked data for SetStandardToken events raised by the TinderTrading contract.
type TinderTradingSetStandardTokenIterator struct {
	Event *TinderTradingSetStandardToken // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetStandardTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetStandardToken)
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
		it.Event = new(TinderTradingSetStandardToken)
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
func (it *TinderTradingSetStandardTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetStandardTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetStandardToken represents a SetStandardToken event raised by the TinderTrading contract.
type TinderTradingSetStandardToken struct {
	StandardToken common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetStandardToken is a free log retrieval operation binding the contract event 0x47cf141b4cb8862f22d0e50186bae1ff22cbb8727b5641f32211840beeb21004.
//
// Solidity: event SetStandardToken(address _standardToken)
func (_TinderTrading *TinderTradingFilterer) FilterSetStandardToken(opts *bind.FilterOpts) (*TinderTradingSetStandardTokenIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetStandardToken")
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetStandardTokenIterator{contract: _TinderTrading.contract, event: "SetStandardToken", logs: logs, sub: sub}, nil
}

// WatchSetStandardToken is a free log subscription operation binding the contract event 0x47cf141b4cb8862f22d0e50186bae1ff22cbb8727b5641f32211840beeb21004.
//
// Solidity: event SetStandardToken(address _standardToken)
func (_TinderTrading *TinderTradingFilterer) WatchSetStandardToken(opts *bind.WatchOpts, sink chan<- *TinderTradingSetStandardToken) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetStandardToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetStandardToken)
				if err := _TinderTrading.contract.UnpackLog(event, "SetStandardToken", log); err != nil {
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

// ParseSetStandardToken is a log parse operation binding the contract event 0x47cf141b4cb8862f22d0e50186bae1ff22cbb8727b5641f32211840beeb21004.
//
// Solidity: event SetStandardToken(address _standardToken)
func (_TinderTrading *TinderTradingFilterer) ParseSetStandardToken(log types.Log) (*TinderTradingSetStandardToken, error) {
	event := new(TinderTradingSetStandardToken)
	if err := _TinderTrading.contract.UnpackLog(event, "SetStandardToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingSetVaultIterator is returned from FilterSetVault and is used to iterate over the raw logs and unpacked data for SetVault events raised by the TinderTrading contract.
type TinderTradingSetVaultIterator struct {
	Event *TinderTradingSetVault // Event containing the contract specifics and raw log

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
func (it *TinderTradingSetVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingSetVault)
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
		it.Event = new(TinderTradingSetVault)
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
func (it *TinderTradingSetVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingSetVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingSetVault represents a SetVault event raised by the TinderTrading contract.
type TinderTradingSetVault struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetVault is a free log retrieval operation binding the contract event 0xd459c7242e23d490831b5676a611c4342d899d28f342d89ae80793e56a930f30.
//
// Solidity: event SetVault(address _vault)
func (_TinderTrading *TinderTradingFilterer) FilterSetVault(opts *bind.FilterOpts) (*TinderTradingSetVaultIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "SetVault")
	if err != nil {
		return nil, err
	}
	return &TinderTradingSetVaultIterator{contract: _TinderTrading.contract, event: "SetVault", logs: logs, sub: sub}, nil
}

// WatchSetVault is a free log subscription operation binding the contract event 0xd459c7242e23d490831b5676a611c4342d899d28f342d89ae80793e56a930f30.
//
// Solidity: event SetVault(address _vault)
func (_TinderTrading *TinderTradingFilterer) WatchSetVault(opts *bind.WatchOpts, sink chan<- *TinderTradingSetVault) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "SetVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingSetVault)
				if err := _TinderTrading.contract.UnpackLog(event, "SetVault", log); err != nil {
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

// ParseSetVault is a log parse operation binding the contract event 0xd459c7242e23d490831b5676a611c4342d899d28f342d89ae80793e56a930f30.
//
// Solidity: event SetVault(address _vault)
func (_TinderTrading *TinderTradingFilterer) ParseSetVault(log types.Log) (*TinderTradingSetVault, error) {
	event := new(TinderTradingSetVault)
	if err := _TinderTrading.contract.UnpackLog(event, "SetVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinderTradingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TinderTrading contract.
type TinderTradingUnpausedIterator struct {
	Event *TinderTradingUnpaused // Event containing the contract specifics and raw log

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
func (it *TinderTradingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinderTradingUnpaused)
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
		it.Event = new(TinderTradingUnpaused)
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
func (it *TinderTradingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinderTradingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinderTradingUnpaused represents a Unpaused event raised by the TinderTrading contract.
type TinderTradingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TinderTrading *TinderTradingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TinderTradingUnpausedIterator, error) {

	logs, sub, err := _TinderTrading.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TinderTradingUnpausedIterator{contract: _TinderTrading.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TinderTrading *TinderTradingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TinderTradingUnpaused) (event.Subscription, error) {

	logs, sub, err := _TinderTrading.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinderTradingUnpaused)
				if err := _TinderTrading.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TinderTrading *TinderTradingFilterer) ParseUnpaused(log types.Log) (*TinderTradingUnpaused, error) {
	event := new(TinderTradingUnpaused)
	if err := _TinderTrading.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
