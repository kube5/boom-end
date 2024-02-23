// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// DeploymentParams is an auto generated low-level Go binding around an user-defined struct.
type DeploymentParams struct {
	Vault common.Address
	Eth   common.Address
	Usdb  common.Address
	Owner common.Address
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"MintDice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEthAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUsdbAmount\",\"type\":\"uint256\"}],\"name\":\"StakeSnapshot\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"earnStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdbStaked\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isMintDice\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEarn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eth\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakedTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdb\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structDeploymentParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTokenWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"mintDice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setDisableEarn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEnableEarn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdbStaked\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isMintDice\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalEarnEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalEarnUsdb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedUsdb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdb\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewStaker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// EarnStakers is a free data retrieval call binding the contract method 0x6798de03.
//
// Solidity: function earnStakers(address ) view returns(uint256 ethStaked, uint256 usdbStaked, bool isMintDice)
func (_Staking *StakingCaller) EarnStakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	EthStaked  *big.Int
	UsdbStaked *big.Int
	IsMintDice bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "earnStakers", arg0)

	outstruct := new(struct {
		EthStaked  *big.Int
		UsdbStaked *big.Int
		IsMintDice bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthStaked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UsdbStaked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsMintDice = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// EarnStakers is a free data retrieval call binding the contract method 0x6798de03.
//
// Solidity: function earnStakers(address ) view returns(uint256 ethStaked, uint256 usdbStaked, bool isMintDice)
func (_Staking *StakingSession) EarnStakers(arg0 common.Address) (struct {
	EthStaked  *big.Int
	UsdbStaked *big.Int
	IsMintDice bool
}, error) {
	return _Staking.Contract.EarnStakers(&_Staking.CallOpts, arg0)
}

// EarnStakers is a free data retrieval call binding the contract method 0x6798de03.
//
// Solidity: function earnStakers(address ) view returns(uint256 ethStaked, uint256 usdbStaked, bool isMintDice)
func (_Staking *StakingCallerSession) EarnStakers(arg0 common.Address) (struct {
	EthStaked  *big.Int
	UsdbStaked *big.Int
	IsMintDice bool
}, error) {
	return _Staking.Contract.EarnStakers(&_Staking.CallOpts, arg0)
}

// EnableEarn is a free data retrieval call binding the contract method 0x84277754.
//
// Solidity: function enableEarn() view returns(bool)
func (_Staking *StakingCaller) EnableEarn(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "enableEarn")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnableEarn is a free data retrieval call binding the contract method 0x84277754.
//
// Solidity: function enableEarn() view returns(bool)
func (_Staking *StakingSession) EnableEarn() (bool, error) {
	return _Staking.Contract.EnableEarn(&_Staking.CallOpts)
}

// EnableEarn is a free data retrieval call binding the contract method 0x84277754.
//
// Solidity: function enableEarn() view returns(bool)
func (_Staking *StakingCallerSession) EnableEarn() (bool, error) {
	return _Staking.Contract.EnableEarn(&_Staking.CallOpts)
}

// Eth is a free data retrieval call binding the contract method 0x8c7c9e0c.
//
// Solidity: function eth() view returns(address)
func (_Staking *StakingCaller) Eth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "eth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eth is a free data retrieval call binding the contract method 0x8c7c9e0c.
//
// Solidity: function eth() view returns(address)
func (_Staking *StakingSession) Eth() (common.Address, error) {
	return _Staking.Contract.Eth(&_Staking.CallOpts)
}

// Eth is a free data retrieval call binding the contract method 0x8c7c9e0c.
//
// Solidity: function eth() view returns(address)
func (_Staking *StakingCallerSession) Eth() (common.Address, error) {
	return _Staking.Contract.Eth(&_Staking.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Staking *StakingCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDecimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Staking *StakingSession) GetDecimals(token common.Address) (uint8, error) {
	return _Staking.Contract.GetDecimals(&_Staking.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Staking *StakingCallerSession) GetDecimals(token common.Address) (uint8, error) {
	return _Staking.Contract.GetDecimals(&_Staking.CallOpts, token)
}

// GetStakedTotals is a free data retrieval call binding the contract method 0xe400fd4d.
//
// Solidity: function getStakedTotals() view returns(uint256, uint256)
func (_Staking *StakingCaller) GetStakedTotals(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStakedTotals")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakedTotals is a free data retrieval call binding the contract method 0xe400fd4d.
//
// Solidity: function getStakedTotals() view returns(uint256, uint256)
func (_Staking *StakingSession) GetStakedTotals() (*big.Int, *big.Int, error) {
	return _Staking.Contract.GetStakedTotals(&_Staking.CallOpts)
}

// GetStakedTotals is a free data retrieval call binding the contract method 0xe400fd4d.
//
// Solidity: function getStakedTotals() view returns(uint256, uint256)
func (_Staking *StakingCallerSession) GetStakedTotals() (*big.Int, *big.Int, error) {
	return _Staking.Contract.GetStakedTotals(&_Staking.CallOpts)
}

// IsTokenWhitelisted is a free data retrieval call binding the contract method 0xb5af090f.
//
// Solidity: function isTokenWhitelisted(address ) view returns(bool)
func (_Staking *StakingCaller) IsTokenWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isTokenWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenWhitelisted is a free data retrieval call binding the contract method 0xb5af090f.
//
// Solidity: function isTokenWhitelisted(address ) view returns(bool)
func (_Staking *StakingSession) IsTokenWhitelisted(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsTokenWhitelisted(&_Staking.CallOpts, arg0)
}

// IsTokenWhitelisted is a free data retrieval call binding the contract method 0xb5af090f.
//
// Solidity: function isTokenWhitelisted(address ) view returns(bool)
func (_Staking *StakingCallerSession) IsTokenWhitelisted(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsTokenWhitelisted(&_Staking.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(uint256 ethStaked, uint256 usdbStaked, bool isMintDice)
func (_Staking *StakingCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	EthStaked  *big.Int
	UsdbStaked *big.Int
	IsMintDice bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		EthStaked  *big.Int
		UsdbStaked *big.Int
		IsMintDice bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthStaked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UsdbStaked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsMintDice = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(uint256 ethStaked, uint256 usdbStaked, bool isMintDice)
func (_Staking *StakingSession) Stakers(arg0 common.Address) (struct {
	EthStaked  *big.Int
	UsdbStaked *big.Int
	IsMintDice bool
}, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(uint256 ethStaked, uint256 usdbStaked, bool isMintDice)
func (_Staking *StakingCallerSession) Stakers(arg0 common.Address) (struct {
	EthStaked  *big.Int
	UsdbStaked *big.Int
	IsMintDice bool
}, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// TotalEarnEth is a free data retrieval call binding the contract method 0xa50444d2.
//
// Solidity: function totalEarnEth() view returns(uint256)
func (_Staking *StakingCaller) TotalEarnEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalEarnEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEarnEth is a free data retrieval call binding the contract method 0xa50444d2.
//
// Solidity: function totalEarnEth() view returns(uint256)
func (_Staking *StakingSession) TotalEarnEth() (*big.Int, error) {
	return _Staking.Contract.TotalEarnEth(&_Staking.CallOpts)
}

// TotalEarnEth is a free data retrieval call binding the contract method 0xa50444d2.
//
// Solidity: function totalEarnEth() view returns(uint256)
func (_Staking *StakingCallerSession) TotalEarnEth() (*big.Int, error) {
	return _Staking.Contract.TotalEarnEth(&_Staking.CallOpts)
}

// TotalEarnUsdb is a free data retrieval call binding the contract method 0x567653a3.
//
// Solidity: function totalEarnUsdb() view returns(uint256)
func (_Staking *StakingCaller) TotalEarnUsdb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalEarnUsdb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEarnUsdb is a free data retrieval call binding the contract method 0x567653a3.
//
// Solidity: function totalEarnUsdb() view returns(uint256)
func (_Staking *StakingSession) TotalEarnUsdb() (*big.Int, error) {
	return _Staking.Contract.TotalEarnUsdb(&_Staking.CallOpts)
}

// TotalEarnUsdb is a free data retrieval call binding the contract method 0x567653a3.
//
// Solidity: function totalEarnUsdb() view returns(uint256)
func (_Staking *StakingCallerSession) TotalEarnUsdb() (*big.Int, error) {
	return _Staking.Contract.TotalEarnUsdb(&_Staking.CallOpts)
}

// TotalStakedEth is a free data retrieval call binding the contract method 0x504cdacf.
//
// Solidity: function totalStakedEth() view returns(uint256)
func (_Staking *StakingCaller) TotalStakedEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStakedEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedEth is a free data retrieval call binding the contract method 0x504cdacf.
//
// Solidity: function totalStakedEth() view returns(uint256)
func (_Staking *StakingSession) TotalStakedEth() (*big.Int, error) {
	return _Staking.Contract.TotalStakedEth(&_Staking.CallOpts)
}

// TotalStakedEth is a free data retrieval call binding the contract method 0x504cdacf.
//
// Solidity: function totalStakedEth() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStakedEth() (*big.Int, error) {
	return _Staking.Contract.TotalStakedEth(&_Staking.CallOpts)
}

// TotalStakedUsdb is a free data retrieval call binding the contract method 0xef2bf7ea.
//
// Solidity: function totalStakedUsdb() view returns(uint256)
func (_Staking *StakingCaller) TotalStakedUsdb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStakedUsdb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedUsdb is a free data retrieval call binding the contract method 0xef2bf7ea.
//
// Solidity: function totalStakedUsdb() view returns(uint256)
func (_Staking *StakingSession) TotalStakedUsdb() (*big.Int, error) {
	return _Staking.Contract.TotalStakedUsdb(&_Staking.CallOpts)
}

// TotalStakedUsdb is a free data retrieval call binding the contract method 0xef2bf7ea.
//
// Solidity: function totalStakedUsdb() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStakedUsdb() (*big.Int, error) {
	return _Staking.Contract.TotalStakedUsdb(&_Staking.CallOpts)
}

// Usdb is a free data retrieval call binding the contract method 0x006f7d00.
//
// Solidity: function usdb() view returns(address)
func (_Staking *StakingCaller) Usdb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "usdb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdb is a free data retrieval call binding the contract method 0x006f7d00.
//
// Solidity: function usdb() view returns(address)
func (_Staking *StakingSession) Usdb() (common.Address, error) {
	return _Staking.Contract.Usdb(&_Staking.CallOpts)
}

// Usdb is a free data retrieval call binding the contract method 0x006f7d00.
//
// Solidity: function usdb() view returns(address)
func (_Staking *StakingCallerSession) Usdb() (common.Address, error) {
	return _Staking.Contract.Usdb(&_Staking.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Staking *StakingCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Staking *StakingSession) Vault() (common.Address, error) {
	return _Staking.Contract.Vault(&_Staking.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Staking *StakingCallerSession) Vault() (common.Address, error) {
	return _Staking.Contract.Vault(&_Staking.CallOpts)
}

// ViewStaker is a free data retrieval call binding the contract method 0xd33ee140.
//
// Solidity: function viewStaker() view returns(uint256, uint256, bool)
func (_Staking *StakingCaller) ViewStaker(opts *bind.CallOpts) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "viewStaker")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// ViewStaker is a free data retrieval call binding the contract method 0xd33ee140.
//
// Solidity: function viewStaker() view returns(uint256, uint256, bool)
func (_Staking *StakingSession) ViewStaker() (*big.Int, *big.Int, bool, error) {
	return _Staking.Contract.ViewStaker(&_Staking.CallOpts)
}

// ViewStaker is a free data retrieval call binding the contract method 0xd33ee140.
//
// Solidity: function viewStaker() view returns(uint256, uint256, bool)
func (_Staking *StakingCallerSession) ViewStaker() (*big.Int, *big.Int, bool, error) {
	return _Staking.Contract.ViewStaker(&_Staking.CallOpts)
}

// InitState is a paid mutator transaction binding the contract method 0x57275d58.
//
// Solidity: function initState((address,address,address,address) params) returns()
func (_Staking *StakingTransactor) InitState(opts *bind.TransactOpts, params DeploymentParams) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initState", params)
}

// InitState is a paid mutator transaction binding the contract method 0x57275d58.
//
// Solidity: function initState((address,address,address,address) params) returns()
func (_Staking *StakingSession) InitState(params DeploymentParams) (*types.Transaction, error) {
	return _Staking.Contract.InitState(&_Staking.TransactOpts, params)
}

// InitState is a paid mutator transaction binding the contract method 0x57275d58.
//
// Solidity: function initState((address,address,address,address) params) returns()
func (_Staking *StakingTransactorSession) InitState(params DeploymentParams) (*types.Transaction, error) {
	return _Staking.Contract.InitState(&_Staking.TransactOpts, params)
}

// MintDice is a paid mutator transaction binding the contract method 0x12621e19.
//
// Solidity: function mintDice(address token) returns()
func (_Staking *StakingTransactor) MintDice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "mintDice", token)
}

// MintDice is a paid mutator transaction binding the contract method 0x12621e19.
//
// Solidity: function mintDice(address token) returns()
func (_Staking *StakingSession) MintDice(token common.Address) (*types.Transaction, error) {
	return _Staking.Contract.MintDice(&_Staking.TransactOpts, token)
}

// MintDice is a paid mutator transaction binding the contract method 0x12621e19.
//
// Solidity: function mintDice(address token) returns()
func (_Staking *StakingTransactorSession) MintDice(token common.Address) (*types.Transaction, error) {
	return _Staking.Contract.MintDice(&_Staking.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// SetDisableEarn is a paid mutator transaction binding the contract method 0x6f8d83bd.
//
// Solidity: function setDisableEarn() returns()
func (_Staking *StakingTransactor) SetDisableEarn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDisableEarn")
}

// SetDisableEarn is a paid mutator transaction binding the contract method 0x6f8d83bd.
//
// Solidity: function setDisableEarn() returns()
func (_Staking *StakingSession) SetDisableEarn() (*types.Transaction, error) {
	return _Staking.Contract.SetDisableEarn(&_Staking.TransactOpts)
}

// SetDisableEarn is a paid mutator transaction binding the contract method 0x6f8d83bd.
//
// Solidity: function setDisableEarn() returns()
func (_Staking *StakingTransactorSession) SetDisableEarn() (*types.Transaction, error) {
	return _Staking.Contract.SetDisableEarn(&_Staking.TransactOpts)
}

// SetEnableEarn is a paid mutator transaction binding the contract method 0x123b9f07.
//
// Solidity: function setEnableEarn() returns()
func (_Staking *StakingTransactor) SetEnableEarn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setEnableEarn")
}

// SetEnableEarn is a paid mutator transaction binding the contract method 0x123b9f07.
//
// Solidity: function setEnableEarn() returns()
func (_Staking *StakingSession) SetEnableEarn() (*types.Transaction, error) {
	return _Staking.Contract.SetEnableEarn(&_Staking.TransactOpts)
}

// SetEnableEarn is a paid mutator transaction binding the contract method 0x123b9f07.
//
// Solidity: function setEnableEarn() returns()
func (_Staking *StakingTransactorSession) SetEnableEarn() (*types.Transaction, error) {
	return _Staking.Contract.SetEnableEarn(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 amount, address _tokenAddress) returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", amount, _tokenAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 amount, address _tokenAddress) returns()
func (_Staking *StakingSession) Stake(amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, amount, _tokenAddress)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 amount, address _tokenAddress) returns()
func (_Staking *StakingTransactorSession) Stake(amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, amount, _tokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address _tokenAddress) returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw", amount, _tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address _tokenAddress) returns()
func (_Staking *StakingSession) Withdraw(amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, amount, _tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address _tokenAddress) returns()
func (_Staking *StakingTransactorSession) Withdraw(amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, amount, _tokenAddress)
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingMintDiceIterator is returned from FilterMintDice and is used to iterate over the raw logs and unpacked data for MintDice events raised by the Staking contract.
type StakingMintDiceIterator struct {
	Event *StakingMintDice // Event containing the contract specifics and raw log

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
func (it *StakingMintDiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingMintDice)
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
		it.Event = new(StakingMintDice)
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
func (it *StakingMintDiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingMintDiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingMintDice represents a MintDice event raised by the Staking contract.
type StakingMintDice struct {
	Staker common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintDice is a free log retrieval operation binding the contract event 0x3a5025f41888ac2c802907287be16311141395626f348e7470fccbe4b7a9ede1.
//
// Solidity: event MintDice(address staker, address token)
func (_Staking *StakingFilterer) FilterMintDice(opts *bind.FilterOpts) (*StakingMintDiceIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "MintDice")
	if err != nil {
		return nil, err
	}
	return &StakingMintDiceIterator{contract: _Staking.contract, event: "MintDice", logs: logs, sub: sub}, nil
}

// WatchMintDice is a free log subscription operation binding the contract event 0x3a5025f41888ac2c802907287be16311141395626f348e7470fccbe4b7a9ede1.
//
// Solidity: event MintDice(address staker, address token)
func (_Staking *StakingFilterer) WatchMintDice(opts *bind.WatchOpts, sink chan<- *StakingMintDice) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "MintDice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingMintDice)
				if err := _Staking.contract.UnpackLog(event, "MintDice", log); err != nil {
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

// ParseMintDice is a log parse operation binding the contract event 0x3a5025f41888ac2c802907287be16311141395626f348e7470fccbe4b7a9ede1.
//
// Solidity: event MintDice(address staker, address token)
func (_Staking *StakingFilterer) ParseMintDice(log types.Log) (*StakingMintDice, error) {
	event := new(StakingMintDice)
	if err := _Staking.contract.UnpackLog(event, "MintDice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeSnapshotIterator is returned from FilterStakeSnapshot and is used to iterate over the raw logs and unpacked data for StakeSnapshot events raised by the Staking contract.
type StakingStakeSnapshotIterator struct {
	Event *StakingStakeSnapshot // Event containing the contract specifics and raw log

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
func (it *StakingStakeSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeSnapshot)
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
		it.Event = new(StakingStakeSnapshot)
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
func (it *StakingStakeSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeSnapshot represents a StakeSnapshot event raised by the Staking contract.
type StakingStakeSnapshot struct {
	Staker          common.Address
	Token           common.Address
	Amount          *big.Int
	TotalEthAmount  *big.Int
	TotalUsdbAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeSnapshot is a free log retrieval operation binding the contract event 0x0a3eea017465bac40e409578b809fc513844114ba30071d3b78bdac432a42ed9.
//
// Solidity: event StakeSnapshot(address staker, address token, uint256 amount, uint256 totalEthAmount, uint256 totalUsdbAmount)
func (_Staking *StakingFilterer) FilterStakeSnapshot(opts *bind.FilterOpts) (*StakingStakeSnapshotIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeSnapshot")
	if err != nil {
		return nil, err
	}
	return &StakingStakeSnapshotIterator{contract: _Staking.contract, event: "StakeSnapshot", logs: logs, sub: sub}, nil
}

// WatchStakeSnapshot is a free log subscription operation binding the contract event 0x0a3eea017465bac40e409578b809fc513844114ba30071d3b78bdac432a42ed9.
//
// Solidity: event StakeSnapshot(address staker, address token, uint256 amount, uint256 totalEthAmount, uint256 totalUsdbAmount)
func (_Staking *StakingFilterer) WatchStakeSnapshot(opts *bind.WatchOpts, sink chan<- *StakingStakeSnapshot) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeSnapshot)
				if err := _Staking.contract.UnpackLog(event, "StakeSnapshot", log); err != nil {
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

// ParseStakeSnapshot is a log parse operation binding the contract event 0x0a3eea017465bac40e409578b809fc513844114ba30071d3b78bdac432a42ed9.
//
// Solidity: event StakeSnapshot(address staker, address token, uint256 amount, uint256 totalEthAmount, uint256 totalUsdbAmount)
func (_Staking *StakingFilterer) ParseStakeSnapshot(log types.Log) (*StakingStakeSnapshot, error) {
	event := new(StakingStakeSnapshot)
	if err := _Staking.contract.UnpackLog(event, "StakeSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
