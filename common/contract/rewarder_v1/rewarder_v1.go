// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewarder_v1

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

// RewarderV1MetaData contains all meta data concerning the RewarderV1 contract.
var RewarderV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_master\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_tokenPerSec\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OnReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"RewardRateUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACC_TOKEN_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_tokenPerSec\",\"type\":\"uint96\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"emergencyTokenWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lpAmount\",\"type\":\"uint256\"}],\"name\":\"onReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"tokenPerSec\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"accTokenPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"distributedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_tokenPerSec\",\"type\":\"uint96\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rewardDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"unpaidRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RewarderV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use RewarderV1MetaData.ABI instead.
var RewarderV1ABI = RewarderV1MetaData.ABI

// RewarderV1 is an auto generated Go binding around an Ethereum contract.
type RewarderV1 struct {
	RewarderV1Caller     // Read-only binding to the contract
	RewarderV1Transactor // Write-only binding to the contract
	RewarderV1Filterer   // Log filterer for contract events
}

// RewarderV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type RewarderV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RewarderV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewarderV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewarderV1Session struct {
	Contract     *RewarderV1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewarderV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewarderV1CallerSession struct {
	Contract *RewarderV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RewarderV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewarderV1TransactorSession struct {
	Contract     *RewarderV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RewarderV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type RewarderV1Raw struct {
	Contract *RewarderV1 // Generic contract binding to access the raw methods on
}

// RewarderV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewarderV1CallerRaw struct {
	Contract *RewarderV1Caller // Generic read-only contract binding to access the raw methods on
}

// RewarderV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewarderV1TransactorRaw struct {
	Contract *RewarderV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRewarderV1 creates a new instance of RewarderV1, bound to a specific deployed contract.
func NewRewarderV1(address common.Address, backend bind.ContractBackend) (*RewarderV1, error) {
	contract, err := bindRewarderV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewarderV1{RewarderV1Caller: RewarderV1Caller{contract: contract}, RewarderV1Transactor: RewarderV1Transactor{contract: contract}, RewarderV1Filterer: RewarderV1Filterer{contract: contract}}, nil
}

// NewRewarderV1Caller creates a new read-only instance of RewarderV1, bound to a specific deployed contract.
func NewRewarderV1Caller(address common.Address, caller bind.ContractCaller) (*RewarderV1Caller, error) {
	contract, err := bindRewarderV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderV1Caller{contract: contract}, nil
}

// NewRewarderV1Transactor creates a new write-only instance of RewarderV1, bound to a specific deployed contract.
func NewRewarderV1Transactor(address common.Address, transactor bind.ContractTransactor) (*RewarderV1Transactor, error) {
	contract, err := bindRewarderV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderV1Transactor{contract: contract}, nil
}

// NewRewarderV1Filterer creates a new log filterer instance of RewarderV1, bound to a specific deployed contract.
func NewRewarderV1Filterer(address common.Address, filterer bind.ContractFilterer) (*RewarderV1Filterer, error) {
	contract, err := bindRewarderV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewarderV1Filterer{contract: contract}, nil
}

// bindRewarderV1 binds a generic wrapper to an already deployed contract.
func bindRewarderV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewarderV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewarderV1 *RewarderV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewarderV1.Contract.RewarderV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewarderV1 *RewarderV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV1.Contract.RewarderV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewarderV1 *RewarderV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewarderV1.Contract.RewarderV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewarderV1 *RewarderV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewarderV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewarderV1 *RewarderV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewarderV1 *RewarderV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewarderV1.Contract.contract.Transact(opts, method, params...)
}

// ACCTOKENPRECISION is a free data retrieval call binding the contract method 0xeea01604.
//
// Solidity: function ACC_TOKEN_PRECISION() view returns(uint256)
func (_RewarderV1 *RewarderV1Caller) ACCTOKENPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "ACC_TOKEN_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ACCTOKENPRECISION is a free data retrieval call binding the contract method 0xeea01604.
//
// Solidity: function ACC_TOKEN_PRECISION() view returns(uint256)
func (_RewarderV1 *RewarderV1Session) ACCTOKENPRECISION() (*big.Int, error) {
	return _RewarderV1.Contract.ACCTOKENPRECISION(&_RewarderV1.CallOpts)
}

// ACCTOKENPRECISION is a free data retrieval call binding the contract method 0xeea01604.
//
// Solidity: function ACC_TOKEN_PRECISION() view returns(uint256)
func (_RewarderV1 *RewarderV1CallerSession) ACCTOKENPRECISION() (*big.Int, error) {
	return _RewarderV1.Contract.ACCTOKENPRECISION(&_RewarderV1.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256[] balances_)
func (_RewarderV1 *RewarderV1Caller) Balances(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "balances")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256[] balances_)
func (_RewarderV1 *RewarderV1Session) Balances() ([]*big.Int, error) {
	return _RewarderV1.Contract.Balances(&_RewarderV1.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256[] balances_)
func (_RewarderV1 *RewarderV1CallerSession) Balances() ([]*big.Int, error) {
	return _RewarderV1.Contract.Balances(&_RewarderV1.CallOpts)
}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_RewarderV1 *RewarderV1Caller) LastRewardTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "lastRewardTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_RewarderV1 *RewarderV1Session) LastRewardTimestamp() (*big.Int, error) {
	return _RewarderV1.Contract.LastRewardTimestamp(&_RewarderV1.CallOpts)
}

// LastRewardTimestamp is a free data retrieval call binding the contract method 0xf8077fae.
//
// Solidity: function lastRewardTimestamp() view returns(uint256)
func (_RewarderV1 *RewarderV1CallerSession) LastRewardTimestamp() (*big.Int, error) {
	return _RewarderV1.Contract.LastRewardTimestamp(&_RewarderV1.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_RewarderV1 *RewarderV1Caller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_RewarderV1 *RewarderV1Session) LpToken() (common.Address, error) {
	return _RewarderV1.Contract.LpToken(&_RewarderV1.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_RewarderV1 *RewarderV1CallerSession) LpToken() (common.Address, error) {
	return _RewarderV1.Contract.LpToken(&_RewarderV1.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_RewarderV1 *RewarderV1Caller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_RewarderV1 *RewarderV1Session) Master() (common.Address, error) {
	return _RewarderV1.Contract.Master(&_RewarderV1.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_RewarderV1 *RewarderV1CallerSession) Master() (common.Address, error) {
	return _RewarderV1.Contract.Master(&_RewarderV1.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_RewarderV1 *RewarderV1Caller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_RewarderV1 *RewarderV1Session) Operator() (common.Address, error) {
	return _RewarderV1.Contract.Operator(&_RewarderV1.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_RewarderV1 *RewarderV1CallerSession) Operator() (common.Address, error) {
	return _RewarderV1.Contract.Operator(&_RewarderV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewarderV1 *RewarderV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewarderV1 *RewarderV1Session) Owner() (common.Address, error) {
	return _RewarderV1.Contract.Owner(&_RewarderV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewarderV1 *RewarderV1CallerSession) Owner() (common.Address, error) {
	return _RewarderV1.Contract.Owner(&_RewarderV1.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256[] rewards)
func (_RewarderV1 *RewarderV1Caller) PendingTokens(opts *bind.CallOpts, _user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "pendingTokens", _user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256[] rewards)
func (_RewarderV1 *RewarderV1Session) PendingTokens(_user common.Address) ([]*big.Int, error) {
	return _RewarderV1.Contract.PendingTokens(&_RewarderV1.CallOpts, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256[] rewards)
func (_RewarderV1 *RewarderV1CallerSession) PendingTokens(_user common.Address) ([]*big.Int, error) {
	return _RewarderV1.Contract.PendingTokens(&_RewarderV1.CallOpts, _user)
}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 ) view returns(address rewardToken, uint96 tokenPerSec, uint128 accTokenPerShare, uint128 distributedAmount)
func (_RewarderV1 *RewarderV1Caller) RewardInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RewardToken       common.Address
	TokenPerSec       *big.Int
	AccTokenPerShare  *big.Int
	DistributedAmount *big.Int
}, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "rewardInfo", arg0)

	outstruct := new(struct {
		RewardToken       common.Address
		TokenPerSec       *big.Int
		AccTokenPerShare  *big.Int
		DistributedAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenPerSec = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccTokenPerShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DistributedAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 ) view returns(address rewardToken, uint96 tokenPerSec, uint128 accTokenPerShare, uint128 distributedAmount)
func (_RewarderV1 *RewarderV1Session) RewardInfo(arg0 *big.Int) (struct {
	RewardToken       common.Address
	TokenPerSec       *big.Int
	AccTokenPerShare  *big.Int
	DistributedAmount *big.Int
}, error) {
	return _RewarderV1.Contract.RewardInfo(&_RewarderV1.CallOpts, arg0)
}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 ) view returns(address rewardToken, uint96 tokenPerSec, uint128 accTokenPerShare, uint128 distributedAmount)
func (_RewarderV1 *RewarderV1CallerSession) RewardInfo(arg0 *big.Int) (struct {
	RewardToken       common.Address
	TokenPerSec       *big.Int
	AccTokenPerShare  *big.Int
	DistributedAmount *big.Int
}, error) {
	return _RewarderV1.Contract.RewardInfo(&_RewarderV1.CallOpts, arg0)
}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_RewarderV1 *RewarderV1Caller) RewardLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "rewardLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_RewarderV1 *RewarderV1Session) RewardLength() (*big.Int, error) {
	return _RewarderV1.Contract.RewardLength(&_RewarderV1.CallOpts)
}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_RewarderV1 *RewarderV1CallerSession) RewardLength() (*big.Int, error) {
	return _RewarderV1.Contract.RewardLength(&_RewarderV1.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0xc2b18aa0.
//
// Solidity: function rewardTokens() view returns(address[] tokens)
func (_RewarderV1 *RewarderV1Caller) RewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "rewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0xc2b18aa0.
//
// Solidity: function rewardTokens() view returns(address[] tokens)
func (_RewarderV1 *RewarderV1Session) RewardTokens() ([]common.Address, error) {
	return _RewarderV1.Contract.RewardTokens(&_RewarderV1.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0xc2b18aa0.
//
// Solidity: function rewardTokens() view returns(address[] tokens)
func (_RewarderV1 *RewarderV1CallerSession) RewardTokens() ([]common.Address, error) {
	return _RewarderV1.Contract.RewardTokens(&_RewarderV1.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 rewardDebt, uint256 unpaidRewards)
func (_RewarderV1 *RewarderV1Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount        *big.Int
	RewardDebt    *big.Int
	UnpaidRewards *big.Int
}, error) {
	var out []interface{}
	err := _RewarderV1.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount        *big.Int
		RewardDebt    *big.Int
		UnpaidRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnpaidRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 rewardDebt, uint256 unpaidRewards)
func (_RewarderV1 *RewarderV1Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount        *big.Int
	RewardDebt    *big.Int
	UnpaidRewards *big.Int
}, error) {
	return _RewarderV1.Contract.UserInfo(&_RewarderV1.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 rewardDebt, uint256 unpaidRewards)
func (_RewarderV1 *RewarderV1CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount        *big.Int
	RewardDebt    *big.Int
	UnpaidRewards *big.Int
}, error) {
	return _RewarderV1.Contract.UserInfo(&_RewarderV1.CallOpts, arg0, arg1)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xedc9d772.
//
// Solidity: function addRewardToken(address _rewardToken, uint96 _tokenPerSec) returns()
func (_RewarderV1 *RewarderV1Transactor) AddRewardToken(opts *bind.TransactOpts, _rewardToken common.Address, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "addRewardToken", _rewardToken, _tokenPerSec)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xedc9d772.
//
// Solidity: function addRewardToken(address _rewardToken, uint96 _tokenPerSec) returns()
func (_RewarderV1 *RewarderV1Session) AddRewardToken(_rewardToken common.Address, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV1.Contract.AddRewardToken(&_RewarderV1.TransactOpts, _rewardToken, _tokenPerSec)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xedc9d772.
//
// Solidity: function addRewardToken(address _rewardToken, uint96 _tokenPerSec) returns()
func (_RewarderV1 *RewarderV1TransactorSession) AddRewardToken(_rewardToken common.Address, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV1.Contract.AddRewardToken(&_RewarderV1.TransactOpts, _rewardToken, _tokenPerSec)
}

// EmergencyTokenWithdraw is a paid mutator transaction binding the contract method 0xad568827.
//
// Solidity: function emergencyTokenWithdraw(address token) returns()
func (_RewarderV1 *RewarderV1Transactor) EmergencyTokenWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "emergencyTokenWithdraw", token)
}

// EmergencyTokenWithdraw is a paid mutator transaction binding the contract method 0xad568827.
//
// Solidity: function emergencyTokenWithdraw(address token) returns()
func (_RewarderV1 *RewarderV1Session) EmergencyTokenWithdraw(token common.Address) (*types.Transaction, error) {
	return _RewarderV1.Contract.EmergencyTokenWithdraw(&_RewarderV1.TransactOpts, token)
}

// EmergencyTokenWithdraw is a paid mutator transaction binding the contract method 0xad568827.
//
// Solidity: function emergencyTokenWithdraw(address token) returns()
func (_RewarderV1 *RewarderV1TransactorSession) EmergencyTokenWithdraw(token common.Address) (*types.Transaction, error) {
	return _RewarderV1.Contract.EmergencyTokenWithdraw(&_RewarderV1.TransactOpts, token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_RewarderV1 *RewarderV1Transactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_RewarderV1 *RewarderV1Session) EmergencyWithdraw() (*types.Transaction, error) {
	return _RewarderV1.Contract.EmergencyWithdraw(&_RewarderV1.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_RewarderV1 *RewarderV1TransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _RewarderV1.Contract.EmergencyWithdraw(&_RewarderV1.TransactOpts)
}

// OnReward is a paid mutator transaction binding the contract method 0xc3723288.
//
// Solidity: function onReward(address _user, uint256 _lpAmount) returns(uint256[] rewards)
func (_RewarderV1 *RewarderV1Transactor) OnReward(opts *bind.TransactOpts, _user common.Address, _lpAmount *big.Int) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "onReward", _user, _lpAmount)
}

// OnReward is a paid mutator transaction binding the contract method 0xc3723288.
//
// Solidity: function onReward(address _user, uint256 _lpAmount) returns(uint256[] rewards)
func (_RewarderV1 *RewarderV1Session) OnReward(_user common.Address, _lpAmount *big.Int) (*types.Transaction, error) {
	return _RewarderV1.Contract.OnReward(&_RewarderV1.TransactOpts, _user, _lpAmount)
}

// OnReward is a paid mutator transaction binding the contract method 0xc3723288.
//
// Solidity: function onReward(address _user, uint256 _lpAmount) returns(uint256[] rewards)
func (_RewarderV1 *RewarderV1TransactorSession) OnReward(_user common.Address, _lpAmount *big.Int) (*types.Transaction, error) {
	return _RewarderV1.Contract.OnReward(&_RewarderV1.TransactOpts, _user, _lpAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewarderV1 *RewarderV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewarderV1 *RewarderV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _RewarderV1.Contract.RenounceOwnership(&_RewarderV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewarderV1 *RewarderV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewarderV1.Contract.RenounceOwnership(&_RewarderV1.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_RewarderV1 *RewarderV1Transactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_RewarderV1 *RewarderV1Session) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _RewarderV1.Contract.SetOperator(&_RewarderV1.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_RewarderV1 *RewarderV1TransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _RewarderV1.Contract.SetOperator(&_RewarderV1.TransactOpts, _operator)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x0bc79363.
//
// Solidity: function setRewardRate(uint256 _tokenId, uint96 _tokenPerSec) returns()
func (_RewarderV1 *RewarderV1Transactor) SetRewardRate(opts *bind.TransactOpts, _tokenId *big.Int, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "setRewardRate", _tokenId, _tokenPerSec)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x0bc79363.
//
// Solidity: function setRewardRate(uint256 _tokenId, uint96 _tokenPerSec) returns()
func (_RewarderV1 *RewarderV1Session) SetRewardRate(_tokenId *big.Int, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV1.Contract.SetRewardRate(&_RewarderV1.TransactOpts, _tokenId, _tokenPerSec)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x0bc79363.
//
// Solidity: function setRewardRate(uint256 _tokenId, uint96 _tokenPerSec) returns()
func (_RewarderV1 *RewarderV1TransactorSession) SetRewardRate(_tokenId *big.Int, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV1.Contract.SetRewardRate(&_RewarderV1.TransactOpts, _tokenId, _tokenPerSec)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewarderV1 *RewarderV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewarderV1 *RewarderV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewarderV1.Contract.TransferOwnership(&_RewarderV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewarderV1 *RewarderV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewarderV1.Contract.TransferOwnership(&_RewarderV1.TransactOpts, newOwner)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_RewarderV1 *RewarderV1Transactor) UpdateReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV1.contract.Transact(opts, "updateReward")
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_RewarderV1 *RewarderV1Session) UpdateReward() (*types.Transaction, error) {
	return _RewarderV1.Contract.UpdateReward(&_RewarderV1.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_RewarderV1 *RewarderV1TransactorSession) UpdateReward() (*types.Transaction, error) {
	return _RewarderV1.Contract.UpdateReward(&_RewarderV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewarderV1 *RewarderV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewarderV1 *RewarderV1Session) Receive() (*types.Transaction, error) {
	return _RewarderV1.Contract.Receive(&_RewarderV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewarderV1 *RewarderV1TransactorSession) Receive() (*types.Transaction, error) {
	return _RewarderV1.Contract.Receive(&_RewarderV1.TransactOpts)
}

// RewarderV1OnRewardIterator is returned from FilterOnReward and is used to iterate over the raw logs and unpacked data for OnReward events raised by the RewarderV1 contract.
type RewarderV1OnRewardIterator struct {
	Event *RewarderV1OnReward // Event containing the contract specifics and raw log

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
func (it *RewarderV1OnRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV1OnReward)
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
		it.Event = new(RewarderV1OnReward)
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
func (it *RewarderV1OnRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV1OnRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV1OnReward represents a OnReward event raised by the RewarderV1 contract.
type RewarderV1OnReward struct {
	RewardToken common.Address
	User        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOnReward is a free log retrieval operation binding the contract event 0x986cbc32375de61d1fabfb01aef452f5c919f2180bb72fff0fb182126a02b527.
//
// Solidity: event OnReward(address indexed rewardToken, address indexed user, uint256 amount)
func (_RewarderV1 *RewarderV1Filterer) FilterOnReward(opts *bind.FilterOpts, rewardToken []common.Address, user []common.Address) (*RewarderV1OnRewardIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewarderV1.contract.FilterLogs(opts, "OnReward", rewardTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV1OnRewardIterator{contract: _RewarderV1.contract, event: "OnReward", logs: logs, sub: sub}, nil
}

// WatchOnReward is a free log subscription operation binding the contract event 0x986cbc32375de61d1fabfb01aef452f5c919f2180bb72fff0fb182126a02b527.
//
// Solidity: event OnReward(address indexed rewardToken, address indexed user, uint256 amount)
func (_RewarderV1 *RewarderV1Filterer) WatchOnReward(opts *bind.WatchOpts, sink chan<- *RewarderV1OnReward, rewardToken []common.Address, user []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewarderV1.contract.WatchLogs(opts, "OnReward", rewardTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV1OnReward)
				if err := _RewarderV1.contract.UnpackLog(event, "OnReward", log); err != nil {
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

// ParseOnReward is a log parse operation binding the contract event 0x986cbc32375de61d1fabfb01aef452f5c919f2180bb72fff0fb182126a02b527.
//
// Solidity: event OnReward(address indexed rewardToken, address indexed user, uint256 amount)
func (_RewarderV1 *RewarderV1Filterer) ParseOnReward(log types.Log) (*RewarderV1OnReward, error) {
	event := new(RewarderV1OnReward)
	if err := _RewarderV1.contract.UnpackLog(event, "OnReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RewarderV1 contract.
type RewarderV1OwnershipTransferredIterator struct {
	Event *RewarderV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewarderV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV1OwnershipTransferred)
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
		it.Event = new(RewarderV1OwnershipTransferred)
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
func (it *RewarderV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV1OwnershipTransferred represents a OwnershipTransferred event raised by the RewarderV1 contract.
type RewarderV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewarderV1 *RewarderV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewarderV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewarderV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV1OwnershipTransferredIterator{contract: _RewarderV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewarderV1 *RewarderV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewarderV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewarderV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV1OwnershipTransferred)
				if err := _RewarderV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RewarderV1 *RewarderV1Filterer) ParseOwnershipTransferred(log types.Log) (*RewarderV1OwnershipTransferred, error) {
	event := new(RewarderV1OwnershipTransferred)
	if err := _RewarderV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV1RewardRateUpdatedIterator is returned from FilterRewardRateUpdated and is used to iterate over the raw logs and unpacked data for RewardRateUpdated events raised by the RewarderV1 contract.
type RewarderV1RewardRateUpdatedIterator struct {
	Event *RewarderV1RewardRateUpdated // Event containing the contract specifics and raw log

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
func (it *RewarderV1RewardRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV1RewardRateUpdated)
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
		it.Event = new(RewarderV1RewardRateUpdated)
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
func (it *RewarderV1RewardRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV1RewardRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV1RewardRateUpdated represents a RewardRateUpdated event raised by the RewarderV1 contract.
type RewarderV1RewardRateUpdated struct {
	RewardToken common.Address
	OldRate     *big.Int
	NewRate     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRateUpdated is a free log retrieval operation binding the contract event 0x225033f2ea5486463cbb49ceda2823be38daddc85031ce2c637e7ad0950bc85a.
//
// Solidity: event RewardRateUpdated(address indexed rewardToken, uint256 oldRate, uint256 newRate)
func (_RewarderV1 *RewarderV1Filterer) FilterRewardRateUpdated(opts *bind.FilterOpts, rewardToken []common.Address) (*RewarderV1RewardRateUpdatedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _RewarderV1.contract.FilterLogs(opts, "RewardRateUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV1RewardRateUpdatedIterator{contract: _RewarderV1.contract, event: "RewardRateUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardRateUpdated is a free log subscription operation binding the contract event 0x225033f2ea5486463cbb49ceda2823be38daddc85031ce2c637e7ad0950bc85a.
//
// Solidity: event RewardRateUpdated(address indexed rewardToken, uint256 oldRate, uint256 newRate)
func (_RewarderV1 *RewarderV1Filterer) WatchRewardRateUpdated(opts *bind.WatchOpts, sink chan<- *RewarderV1RewardRateUpdated, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _RewarderV1.contract.WatchLogs(opts, "RewardRateUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV1RewardRateUpdated)
				if err := _RewarderV1.contract.UnpackLog(event, "RewardRateUpdated", log); err != nil {
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

// ParseRewardRateUpdated is a log parse operation binding the contract event 0x225033f2ea5486463cbb49ceda2823be38daddc85031ce2c637e7ad0950bc85a.
//
// Solidity: event RewardRateUpdated(address indexed rewardToken, uint256 oldRate, uint256 newRate)
func (_RewarderV1 *RewarderV1Filterer) ParseRewardRateUpdated(log types.Log) (*RewarderV1RewardRateUpdated, error) {
	event := new(RewarderV1RewardRateUpdated)
	if err := _RewarderV1.contract.UnpackLog(event, "RewardRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
