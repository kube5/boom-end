// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package referrals

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ReferralsConfig is an auto generated low-level Go binding around an user-defined struct.
type ReferralsConfig struct {
	Direct    *big.Int
	Indirect  *big.Int
	Discounts [3]*big.Int
}

// ReferralsABI is the input ABI used to generate the binding from.
const ReferralsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"AllowedToInteractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DirectRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tradeSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rebase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"name\":\"DistributeRefReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IndirectRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"compounded\",\"type\":\"bool\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"direct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indirect\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"discounts\",\"type\":\"uint256[3]\"}],\"indexed\":false,\"internalType\":\"structReferrals.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"SetConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIReferral.DiscountGrade\",\"name\":\"_grade\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumIReferral.DiscountGrade\",\"name\":\"_referrerGrade\",\"type\":\"uint8\"}],\"name\":\"SetDiscountGrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_pubilcRegister\",\"type\":\"bool\"}],\"name\":\"SetPubilcRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_standardToken\",\"type\":\"address\"}],\"name\":\"SetStandardToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tradingVault\",\"type\":\"address\"}],\"name\":\"SetTradingVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"UpdateAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToInteract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_compound\",\"type\":\"bool\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tradeSize\",\"type\":\"uint256\"}],\"name\":\"distributeRefReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rebase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tradingVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_standardToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubilcRegister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referrerDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"directVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indirectVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"enumIReferral.DiscountGrade\",\"name\":\"grade\",\"type\":\"uint8\"},{\"internalType\":\"enumIReferral.DiscountGrade\",\"name\":\"referrerGrade\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registerAdmins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setAllowedToInteract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"direct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indirect\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"discounts\",\"type\":\"uint256[3]\"}],\"internalType\":\"structReferrals.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"enumIReferral.DiscountGrade\",\"name\":\"_grade\",\"type\":\"uint8\"},{\"internalType\":\"enumIReferral.DiscountGrade\",\"name\":\"_referrerGrade\",\"type\":\"uint8\"}],\"name\":\"setDiscountGrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_pubilcRegister\",\"type\":\"bool\"}],\"name\":\"setPubilcRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_standardToken\",\"type\":\"address\"}],\"name\":\"setStandardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tradingVault\",\"type\":\"address\"}],\"name\":\"setTradingVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"signUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Referrals is an auto generated Go binding around an Ethereum contract.
type Referrals struct {
	ReferralsCaller     // Read-only binding to the contract
	ReferralsTransactor // Write-only binding to the contract
	ReferralsFilterer   // Log filterer for contract events
}

// ReferralsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReferralsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReferralsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReferralsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReferralsSession struct {
	Contract     *Referrals        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReferralsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReferralsCallerSession struct {
	Contract *ReferralsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ReferralsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReferralsTransactorSession struct {
	Contract     *ReferralsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ReferralsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReferralsRaw struct {
	Contract *Referrals // Generic contract binding to access the raw methods on
}

// ReferralsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReferralsCallerRaw struct {
	Contract *ReferralsCaller // Generic read-only contract binding to access the raw methods on
}

// ReferralsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReferralsTransactorRaw struct {
	Contract *ReferralsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReferrals creates a new instance of Referrals, bound to a specific deployed contract.
func NewReferrals(address common.Address, backend bind.ContractBackend) (*Referrals, error) {
	contract, err := bindReferrals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Referrals{ReferralsCaller: ReferralsCaller{contract: contract}, ReferralsTransactor: ReferralsTransactor{contract: contract}, ReferralsFilterer: ReferralsFilterer{contract: contract}}, nil
}

// NewReferralsCaller creates a new read-only instance of Referrals, bound to a specific deployed contract.
func NewReferralsCaller(address common.Address, caller bind.ContractCaller) (*ReferralsCaller, error) {
	contract, err := bindReferrals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralsCaller{contract: contract}, nil
}

// NewReferralsTransactor creates a new write-only instance of Referrals, bound to a specific deployed contract.
func NewReferralsTransactor(address common.Address, transactor bind.ContractTransactor) (*ReferralsTransactor, error) {
	contract, err := bindReferrals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralsTransactor{contract: contract}, nil
}

// NewReferralsFilterer creates a new log filterer instance of Referrals, bound to a specific deployed contract.
func NewReferralsFilterer(address common.Address, filterer bind.ContractFilterer) (*ReferralsFilterer, error) {
	contract, err := bindReferrals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReferralsFilterer{contract: contract}, nil
}

// bindReferrals binds a generic wrapper to an already deployed contract.
func bindReferrals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReferralsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Referrals *ReferralsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Referrals.Contract.ReferralsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Referrals *ReferralsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referrals.Contract.ReferralsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Referrals *ReferralsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Referrals.Contract.ReferralsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Referrals *ReferralsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Referrals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Referrals *ReferralsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referrals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Referrals *ReferralsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Referrals.Contract.contract.Transact(opts, method, params...)
}

// AllowedToInteract is a free data retrieval call binding the contract method 0x41200844.
//
// Solidity: function allowedToInteract(address ) view returns(bool)
func (_Referrals *ReferralsCaller) AllowedToInteract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "allowedToInteract", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToInteract is a free data retrieval call binding the contract method 0x41200844.
//
// Solidity: function allowedToInteract(address ) view returns(bool)
func (_Referrals *ReferralsSession) AllowedToInteract(arg0 common.Address) (bool, error) {
	return _Referrals.Contract.AllowedToInteract(&_Referrals.CallOpts, arg0)
}

// AllowedToInteract is a free data retrieval call binding the contract method 0x41200844.
//
// Solidity: function allowedToInteract(address ) view returns(bool)
func (_Referrals *ReferralsCallerSession) AllowedToInteract(arg0 common.Address) (bool, error) {
	return _Referrals.Contract.AllowedToInteract(&_Referrals.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Referrals *ReferralsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Referrals *ReferralsSession) Owner() (common.Address, error) {
	return _Referrals.Contract.Owner(&_Referrals.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Referrals *ReferralsCallerSession) Owner() (common.Address, error) {
	return _Referrals.Contract.Owner(&_Referrals.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Referrals *ReferralsCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Referrals *ReferralsSession) PendingOwner() (common.Address, error) {
	return _Referrals.Contract.PendingOwner(&_Referrals.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Referrals *ReferralsCallerSession) PendingOwner() (common.Address, error) {
	return _Referrals.Contract.PendingOwner(&_Referrals.CallOpts)
}

// PubilcRegister is a free data retrieval call binding the contract method 0xaab5acbe.
//
// Solidity: function pubilcRegister() view returns(bool)
func (_Referrals *ReferralsCaller) PubilcRegister(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "pubilcRegister")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PubilcRegister is a free data retrieval call binding the contract method 0xaab5acbe.
//
// Solidity: function pubilcRegister() view returns(bool)
func (_Referrals *ReferralsSession) PubilcRegister() (bool, error) {
	return _Referrals.Contract.PubilcRegister(&_Referrals.CallOpts)
}

// PubilcRegister is a free data retrieval call binding the contract method 0xaab5acbe.
//
// Solidity: function pubilcRegister() view returns(bool)
func (_Referrals *ReferralsCallerSession) PubilcRegister() (bool, error) {
	return _Referrals.Contract.PubilcRegister(&_Referrals.CallOpts)
}

// ReferrerDetails is a free data retrieval call binding the contract method 0x6a83de79.
//
// Solidity: function referrerDetails(address ) view returns(bool registered, uint256 directVolume, uint256 indirectVolume, uint256 pendingRewards, uint256 claimedRewards, address referrer, uint8 grade, uint8 referrerGrade)
func (_Referrals *ReferralsCaller) ReferrerDetails(opts *bind.CallOpts, arg0 common.Address) (struct {
	Registered     bool
	DirectVolume   *big.Int
	IndirectVolume *big.Int
	PendingRewards *big.Int
	ClaimedRewards *big.Int
	Referrer       common.Address
	Grade          uint8
	ReferrerGrade  uint8
}, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "referrerDetails", arg0)

	outstruct := new(struct {
		Registered     bool
		DirectVolume   *big.Int
		IndirectVolume *big.Int
		PendingRewards *big.Int
		ClaimedRewards *big.Int
		Referrer       common.Address
		Grade          uint8
		ReferrerGrade  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.DirectVolume = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IndirectVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ClaimedRewards = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Referrer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Grade = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.ReferrerGrade = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// ReferrerDetails is a free data retrieval call binding the contract method 0x6a83de79.
//
// Solidity: function referrerDetails(address ) view returns(bool registered, uint256 directVolume, uint256 indirectVolume, uint256 pendingRewards, uint256 claimedRewards, address referrer, uint8 grade, uint8 referrerGrade)
func (_Referrals *ReferralsSession) ReferrerDetails(arg0 common.Address) (struct {
	Registered     bool
	DirectVolume   *big.Int
	IndirectVolume *big.Int
	PendingRewards *big.Int
	ClaimedRewards *big.Int
	Referrer       common.Address
	Grade          uint8
	ReferrerGrade  uint8
}, error) {
	return _Referrals.Contract.ReferrerDetails(&_Referrals.CallOpts, arg0)
}

// ReferrerDetails is a free data retrieval call binding the contract method 0x6a83de79.
//
// Solidity: function referrerDetails(address ) view returns(bool registered, uint256 directVolume, uint256 indirectVolume, uint256 pendingRewards, uint256 claimedRewards, address referrer, uint8 grade, uint8 referrerGrade)
func (_Referrals *ReferralsCallerSession) ReferrerDetails(arg0 common.Address) (struct {
	Registered     bool
	DirectVolume   *big.Int
	IndirectVolume *big.Int
	PendingRewards *big.Int
	ClaimedRewards *big.Int
	Referrer       common.Address
	Grade          uint8
	ReferrerGrade  uint8
}, error) {
	return _Referrals.Contract.ReferrerDetails(&_Referrals.CallOpts, arg0)
}

// RegisterAdmins is a free data retrieval call binding the contract method 0x2892bfd2.
//
// Solidity: function registerAdmins(address ) view returns(bool)
func (_Referrals *ReferralsCaller) RegisterAdmins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "registerAdmins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisterAdmins is a free data retrieval call binding the contract method 0x2892bfd2.
//
// Solidity: function registerAdmins(address ) view returns(bool)
func (_Referrals *ReferralsSession) RegisterAdmins(arg0 common.Address) (bool, error) {
	return _Referrals.Contract.RegisterAdmins(&_Referrals.CallOpts, arg0)
}

// RegisterAdmins is a free data retrieval call binding the contract method 0x2892bfd2.
//
// Solidity: function registerAdmins(address ) view returns(bool)
func (_Referrals *ReferralsCallerSession) RegisterAdmins(arg0 common.Address) (bool, error) {
	return _Referrals.Contract.RegisterAdmins(&_Referrals.CallOpts, arg0)
}

// StandardToken is a free data retrieval call binding the contract method 0x4c7d7d3f.
//
// Solidity: function standardToken() view returns(address)
func (_Referrals *ReferralsCaller) StandardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "standardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StandardToken is a free data retrieval call binding the contract method 0x4c7d7d3f.
//
// Solidity: function standardToken() view returns(address)
func (_Referrals *ReferralsSession) StandardToken() (common.Address, error) {
	return _Referrals.Contract.StandardToken(&_Referrals.CallOpts)
}

// StandardToken is a free data retrieval call binding the contract method 0x4c7d7d3f.
//
// Solidity: function standardToken() view returns(address)
func (_Referrals *ReferralsCallerSession) StandardToken() (common.Address, error) {
	return _Referrals.Contract.StandardToken(&_Referrals.CallOpts)
}

// TradingVault is a free data retrieval call binding the contract method 0xf1bd27a8.
//
// Solidity: function tradingVault() view returns(address)
func (_Referrals *ReferralsCaller) TradingVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Referrals.contract.Call(opts, &out, "tradingVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradingVault is a free data retrieval call binding the contract method 0xf1bd27a8.
//
// Solidity: function tradingVault() view returns(address)
func (_Referrals *ReferralsSession) TradingVault() (common.Address, error) {
	return _Referrals.Contract.TradingVault(&_Referrals.CallOpts)
}

// TradingVault is a free data retrieval call binding the contract method 0xf1bd27a8.
//
// Solidity: function tradingVault() view returns(address)
func (_Referrals *ReferralsCallerSession) TradingVault() (common.Address, error) {
	return _Referrals.Contract.TradingVault(&_Referrals.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Referrals *ReferralsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Referrals *ReferralsSession) AcceptOwnership() (*types.Transaction, error) {
	return _Referrals.Contract.AcceptOwnership(&_Referrals.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Referrals *ReferralsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Referrals.Contract.AcceptOwnership(&_Referrals.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0e6878a3.
//
// Solidity: function claimRewards(bool _compound) returns()
func (_Referrals *ReferralsTransactor) ClaimRewards(opts *bind.TransactOpts, _compound bool) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "claimRewards", _compound)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0e6878a3.
//
// Solidity: function claimRewards(bool _compound) returns()
func (_Referrals *ReferralsSession) ClaimRewards(_compound bool) (*types.Transaction, error) {
	return _Referrals.Contract.ClaimRewards(&_Referrals.TransactOpts, _compound)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0e6878a3.
//
// Solidity: function claimRewards(bool _compound) returns()
func (_Referrals *ReferralsTransactorSession) ClaimRewards(_compound bool) (*types.Transaction, error) {
	return _Referrals.Contract.ClaimRewards(&_Referrals.TransactOpts, _compound)
}

// DistributeRefReward is a paid mutator transaction binding the contract method 0xae8b3868.
//
// Solidity: function distributeRefReward(address _trader, uint256 _totalFee, uint256 _tradeSize) returns(uint256 rebase, uint256 discount)
func (_Referrals *ReferralsTransactor) DistributeRefReward(opts *bind.TransactOpts, _trader common.Address, _totalFee *big.Int, _tradeSize *big.Int) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "distributeRefReward", _trader, _totalFee, _tradeSize)
}

// DistributeRefReward is a paid mutator transaction binding the contract method 0xae8b3868.
//
// Solidity: function distributeRefReward(address _trader, uint256 _totalFee, uint256 _tradeSize) returns(uint256 rebase, uint256 discount)
func (_Referrals *ReferralsSession) DistributeRefReward(_trader common.Address, _totalFee *big.Int, _tradeSize *big.Int) (*types.Transaction, error) {
	return _Referrals.Contract.DistributeRefReward(&_Referrals.TransactOpts, _trader, _totalFee, _tradeSize)
}

// DistributeRefReward is a paid mutator transaction binding the contract method 0xae8b3868.
//
// Solidity: function distributeRefReward(address _trader, uint256 _totalFee, uint256 _tradeSize) returns(uint256 rebase, uint256 discount)
func (_Referrals *ReferralsTransactorSession) DistributeRefReward(_trader common.Address, _totalFee *big.Int, _tradeSize *big.Int) (*types.Transaction, error) {
	return _Referrals.Contract.DistributeRefReward(&_Referrals.TransactOpts, _trader, _totalFee, _tradeSize)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tradingVault, address _standardToken) returns()
func (_Referrals *ReferralsTransactor) Initialize(opts *bind.TransactOpts, _tradingVault common.Address, _standardToken common.Address) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "initialize", _tradingVault, _standardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tradingVault, address _standardToken) returns()
func (_Referrals *ReferralsSession) Initialize(_tradingVault common.Address, _standardToken common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.Initialize(&_Referrals.TransactOpts, _tradingVault, _standardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tradingVault, address _standardToken) returns()
func (_Referrals *ReferralsTransactorSession) Initialize(_tradingVault common.Address, _standardToken common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.Initialize(&_Referrals.TransactOpts, _tradingVault, _standardToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Referrals *ReferralsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Referrals *ReferralsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Referrals.Contract.RenounceOwnership(&_Referrals.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Referrals *ReferralsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Referrals.Contract.RenounceOwnership(&_Referrals.TransactOpts)
}

// SetAllowedToInteract is a paid mutator transaction binding the contract method 0x9d1765f6.
//
// Solidity: function setAllowedToInteract(address _sender, bool _status) returns()
func (_Referrals *ReferralsTransactor) SetAllowedToInteract(opts *bind.TransactOpts, _sender common.Address, _status bool) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "setAllowedToInteract", _sender, _status)
}

// SetAllowedToInteract is a paid mutator transaction binding the contract method 0x9d1765f6.
//
// Solidity: function setAllowedToInteract(address _sender, bool _status) returns()
func (_Referrals *ReferralsSession) SetAllowedToInteract(_sender common.Address, _status bool) (*types.Transaction, error) {
	return _Referrals.Contract.SetAllowedToInteract(&_Referrals.TransactOpts, _sender, _status)
}

// SetAllowedToInteract is a paid mutator transaction binding the contract method 0x9d1765f6.
//
// Solidity: function setAllowedToInteract(address _sender, bool _status) returns()
func (_Referrals *ReferralsTransactorSession) SetAllowedToInteract(_sender common.Address, _status bool) (*types.Transaction, error) {
	return _Referrals.Contract.SetAllowedToInteract(&_Referrals.TransactOpts, _sender, _status)
}

// SetConfig is a paid mutator transaction binding the contract method 0xfd32545d.
//
// Solidity: function setConfig((uint256,uint256,uint256[3]) _config) returns()
func (_Referrals *ReferralsTransactor) SetConfig(opts *bind.TransactOpts, _config ReferralsConfig) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "setConfig", _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xfd32545d.
//
// Solidity: function setConfig((uint256,uint256,uint256[3]) _config) returns()
func (_Referrals *ReferralsSession) SetConfig(_config ReferralsConfig) (*types.Transaction, error) {
	return _Referrals.Contract.SetConfig(&_Referrals.TransactOpts, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xfd32545d.
//
// Solidity: function setConfig((uint256,uint256,uint256[3]) _config) returns()
func (_Referrals *ReferralsTransactorSession) SetConfig(_config ReferralsConfig) (*types.Transaction, error) {
	return _Referrals.Contract.SetConfig(&_Referrals.TransactOpts, _config)
}

// SetDiscountGrade is a paid mutator transaction binding the contract method 0x0de0fdda.
//
// Solidity: function setDiscountGrade(address _user, uint8 _grade, uint8 _referrerGrade) returns()
func (_Referrals *ReferralsTransactor) SetDiscountGrade(opts *bind.TransactOpts, _user common.Address, _grade uint8, _referrerGrade uint8) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "setDiscountGrade", _user, _grade, _referrerGrade)
}

// SetDiscountGrade is a paid mutator transaction binding the contract method 0x0de0fdda.
//
// Solidity: function setDiscountGrade(address _user, uint8 _grade, uint8 _referrerGrade) returns()
func (_Referrals *ReferralsSession) SetDiscountGrade(_user common.Address, _grade uint8, _referrerGrade uint8) (*types.Transaction, error) {
	return _Referrals.Contract.SetDiscountGrade(&_Referrals.TransactOpts, _user, _grade, _referrerGrade)
}

// SetDiscountGrade is a paid mutator transaction binding the contract method 0x0de0fdda.
//
// Solidity: function setDiscountGrade(address _user, uint8 _grade, uint8 _referrerGrade) returns()
func (_Referrals *ReferralsTransactorSession) SetDiscountGrade(_user common.Address, _grade uint8, _referrerGrade uint8) (*types.Transaction, error) {
	return _Referrals.Contract.SetDiscountGrade(&_Referrals.TransactOpts, _user, _grade, _referrerGrade)
}

// SetPubilcRegister is a paid mutator transaction binding the contract method 0x00ea9df6.
//
// Solidity: function setPubilcRegister(bool _pubilcRegister) returns()
func (_Referrals *ReferralsTransactor) SetPubilcRegister(opts *bind.TransactOpts, _pubilcRegister bool) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "setPubilcRegister", _pubilcRegister)
}

// SetPubilcRegister is a paid mutator transaction binding the contract method 0x00ea9df6.
//
// Solidity: function setPubilcRegister(bool _pubilcRegister) returns()
func (_Referrals *ReferralsSession) SetPubilcRegister(_pubilcRegister bool) (*types.Transaction, error) {
	return _Referrals.Contract.SetPubilcRegister(&_Referrals.TransactOpts, _pubilcRegister)
}

// SetPubilcRegister is a paid mutator transaction binding the contract method 0x00ea9df6.
//
// Solidity: function setPubilcRegister(bool _pubilcRegister) returns()
func (_Referrals *ReferralsTransactorSession) SetPubilcRegister(_pubilcRegister bool) (*types.Transaction, error) {
	return _Referrals.Contract.SetPubilcRegister(&_Referrals.TransactOpts, _pubilcRegister)
}

// SetStandardToken is a paid mutator transaction binding the contract method 0xdd6d012e.
//
// Solidity: function setStandardToken(address _standardToken) returns()
func (_Referrals *ReferralsTransactor) SetStandardToken(opts *bind.TransactOpts, _standardToken common.Address) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "setStandardToken", _standardToken)
}

// SetStandardToken is a paid mutator transaction binding the contract method 0xdd6d012e.
//
// Solidity: function setStandardToken(address _standardToken) returns()
func (_Referrals *ReferralsSession) SetStandardToken(_standardToken common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.SetStandardToken(&_Referrals.TransactOpts, _standardToken)
}

// SetStandardToken is a paid mutator transaction binding the contract method 0xdd6d012e.
//
// Solidity: function setStandardToken(address _standardToken) returns()
func (_Referrals *ReferralsTransactorSession) SetStandardToken(_standardToken common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.SetStandardToken(&_Referrals.TransactOpts, _standardToken)
}

// SetTradingVault is a paid mutator transaction binding the contract method 0x60dc001a.
//
// Solidity: function setTradingVault(address _tradingVault) returns()
func (_Referrals *ReferralsTransactor) SetTradingVault(opts *bind.TransactOpts, _tradingVault common.Address) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "setTradingVault", _tradingVault)
}

// SetTradingVault is a paid mutator transaction binding the contract method 0x60dc001a.
//
// Solidity: function setTradingVault(address _tradingVault) returns()
func (_Referrals *ReferralsSession) SetTradingVault(_tradingVault common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.SetTradingVault(&_Referrals.TransactOpts, _tradingVault)
}

// SetTradingVault is a paid mutator transaction binding the contract method 0x60dc001a.
//
// Solidity: function setTradingVault(address _tradingVault) returns()
func (_Referrals *ReferralsTransactorSession) SetTradingVault(_tradingVault common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.SetTradingVault(&_Referrals.TransactOpts, _tradingVault)
}

// SignUp is a paid mutator transaction binding the contract method 0xb2f025d8.
//
// Solidity: function signUp(address user, address referrer) returns()
func (_Referrals *ReferralsTransactor) SignUp(opts *bind.TransactOpts, user common.Address, referrer common.Address) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "signUp", user, referrer)
}

// SignUp is a paid mutator transaction binding the contract method 0xb2f025d8.
//
// Solidity: function signUp(address user, address referrer) returns()
func (_Referrals *ReferralsSession) SignUp(user common.Address, referrer common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.SignUp(&_Referrals.TransactOpts, user, referrer)
}

// SignUp is a paid mutator transaction binding the contract method 0xb2f025d8.
//
// Solidity: function signUp(address user, address referrer) returns()
func (_Referrals *ReferralsTransactorSession) SignUp(user common.Address, referrer common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.SignUp(&_Referrals.TransactOpts, user, referrer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Referrals *ReferralsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Referrals *ReferralsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.TransferOwnership(&_Referrals.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Referrals *ReferralsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Referrals.Contract.TransferOwnership(&_Referrals.TransactOpts, newOwner)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _admin, bool _flag) returns()
func (_Referrals *ReferralsTransactor) UpdateAdmin(opts *bind.TransactOpts, _admin common.Address, _flag bool) (*types.Transaction, error) {
	return _Referrals.contract.Transact(opts, "updateAdmin", _admin, _flag)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _admin, bool _flag) returns()
func (_Referrals *ReferralsSession) UpdateAdmin(_admin common.Address, _flag bool) (*types.Transaction, error) {
	return _Referrals.Contract.UpdateAdmin(&_Referrals.TransactOpts, _admin, _flag)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _admin, bool _flag) returns()
func (_Referrals *ReferralsTransactorSession) UpdateAdmin(_admin common.Address, _flag bool) (*types.Transaction, error) {
	return _Referrals.Contract.UpdateAdmin(&_Referrals.TransactOpts, _admin, _flag)
}

// ReferralsAllowedToInteractSetIterator is returned from FilterAllowedToInteractSet and is used to iterate over the raw logs and unpacked data for AllowedToInteractSet events raised by the Referrals contract.
type ReferralsAllowedToInteractSetIterator struct {
	Event *ReferralsAllowedToInteractSet // Event containing the contract specifics and raw log

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
func (it *ReferralsAllowedToInteractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsAllowedToInteractSet)
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
		it.Event = new(ReferralsAllowedToInteractSet)
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
func (it *ReferralsAllowedToInteractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsAllowedToInteractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsAllowedToInteractSet represents a AllowedToInteractSet event raised by the Referrals contract.
type ReferralsAllowedToInteractSet struct {
	Sender common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllowedToInteractSet is a free log retrieval operation binding the contract event 0x77711358779674597f238fdb43e9f28618347bc0e217ae42c982656c0b7f6be8.
//
// Solidity: event AllowedToInteractSet(address indexed sender, bool status)
func (_Referrals *ReferralsFilterer) FilterAllowedToInteractSet(opts *bind.FilterOpts, sender []common.Address) (*ReferralsAllowedToInteractSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "AllowedToInteractSet", senderRule)
	if err != nil {
		return nil, err
	}
	return &ReferralsAllowedToInteractSetIterator{contract: _Referrals.contract, event: "AllowedToInteractSet", logs: logs, sub: sub}, nil
}

// WatchAllowedToInteractSet is a free log subscription operation binding the contract event 0x77711358779674597f238fdb43e9f28618347bc0e217ae42c982656c0b7f6be8.
//
// Solidity: event AllowedToInteractSet(address indexed sender, bool status)
func (_Referrals *ReferralsFilterer) WatchAllowedToInteractSet(opts *bind.WatchOpts, sink chan<- *ReferralsAllowedToInteractSet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "AllowedToInteractSet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsAllowedToInteractSet)
				if err := _Referrals.contract.UnpackLog(event, "AllowedToInteractSet", log); err != nil {
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

// ParseAllowedToInteractSet is a log parse operation binding the contract event 0x77711358779674597f238fdb43e9f28618347bc0e217ae42c982656c0b7f6be8.
//
// Solidity: event AllowedToInteractSet(address indexed sender, bool status)
func (_Referrals *ReferralsFilterer) ParseAllowedToInteractSet(log types.Log) (*ReferralsAllowedToInteractSet, error) {
	event := new(ReferralsAllowedToInteractSet)
	if err := _Referrals.contract.UnpackLog(event, "AllowedToInteractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsDirectRewardsIterator is returned from FilterDirectRewards and is used to iterate over the raw logs and unpacked data for DirectRewards events raised by the Referrals contract.
type ReferralsDirectRewardsIterator struct {
	Event *ReferralsDirectRewards // Event containing the contract specifics and raw log

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
func (it *ReferralsDirectRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsDirectRewards)
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
		it.Event = new(ReferralsDirectRewards)
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
func (it *ReferralsDirectRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsDirectRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsDirectRewards represents a DirectRewards event raised by the Referrals contract.
type ReferralsDirectRewards struct {
	Trader   common.Address
	Referrer common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDirectRewards is a free log retrieval operation binding the contract event 0xe1db568e575e20b3cc2e3c21788483533a8801007ad456c4b2c17780c094c5df.
//
// Solidity: event DirectRewards(address trader, address referrer, uint256 amount)
func (_Referrals *ReferralsFilterer) FilterDirectRewards(opts *bind.FilterOpts) (*ReferralsDirectRewardsIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "DirectRewards")
	if err != nil {
		return nil, err
	}
	return &ReferralsDirectRewardsIterator{contract: _Referrals.contract, event: "DirectRewards", logs: logs, sub: sub}, nil
}

// WatchDirectRewards is a free log subscription operation binding the contract event 0xe1db568e575e20b3cc2e3c21788483533a8801007ad456c4b2c17780c094c5df.
//
// Solidity: event DirectRewards(address trader, address referrer, uint256 amount)
func (_Referrals *ReferralsFilterer) WatchDirectRewards(opts *bind.WatchOpts, sink chan<- *ReferralsDirectRewards) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "DirectRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsDirectRewards)
				if err := _Referrals.contract.UnpackLog(event, "DirectRewards", log); err != nil {
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

// ParseDirectRewards is a log parse operation binding the contract event 0xe1db568e575e20b3cc2e3c21788483533a8801007ad456c4b2c17780c094c5df.
//
// Solidity: event DirectRewards(address trader, address referrer, uint256 amount)
func (_Referrals *ReferralsFilterer) ParseDirectRewards(log types.Log) (*ReferralsDirectRewards, error) {
	event := new(ReferralsDirectRewards)
	if err := _Referrals.contract.UnpackLog(event, "DirectRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsDistributeRefRewardIterator is returned from FilterDistributeRefReward and is used to iterate over the raw logs and unpacked data for DistributeRefReward events raised by the Referrals contract.
type ReferralsDistributeRefRewardIterator struct {
	Event *ReferralsDistributeRefReward // Event containing the contract specifics and raw log

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
func (it *ReferralsDistributeRefRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsDistributeRefReward)
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
		it.Event = new(ReferralsDistributeRefReward)
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
func (it *ReferralsDistributeRefRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsDistributeRefRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsDistributeRefReward represents a DistributeRefReward event raised by the Referrals contract.
type ReferralsDistributeRefReward struct {
	Trader    common.Address
	TotalFee  *big.Int
	TradeSize *big.Int
	Rebase    *big.Int
	Discount  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributeRefReward is a free log retrieval operation binding the contract event 0xe54024fd795e1d68a94ade5eae0c744e88bfae5eb3db45c469e941fec38292d9.
//
// Solidity: event DistributeRefReward(address trader, uint256 _totalFee, uint256 _tradeSize, uint256 rebase, uint256 discount)
func (_Referrals *ReferralsFilterer) FilterDistributeRefReward(opts *bind.FilterOpts) (*ReferralsDistributeRefRewardIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "DistributeRefReward")
	if err != nil {
		return nil, err
	}
	return &ReferralsDistributeRefRewardIterator{contract: _Referrals.contract, event: "DistributeRefReward", logs: logs, sub: sub}, nil
}

// WatchDistributeRefReward is a free log subscription operation binding the contract event 0xe54024fd795e1d68a94ade5eae0c744e88bfae5eb3db45c469e941fec38292d9.
//
// Solidity: event DistributeRefReward(address trader, uint256 _totalFee, uint256 _tradeSize, uint256 rebase, uint256 discount)
func (_Referrals *ReferralsFilterer) WatchDistributeRefReward(opts *bind.WatchOpts, sink chan<- *ReferralsDistributeRefReward) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "DistributeRefReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsDistributeRefReward)
				if err := _Referrals.contract.UnpackLog(event, "DistributeRefReward", log); err != nil {
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

// ParseDistributeRefReward is a log parse operation binding the contract event 0xe54024fd795e1d68a94ade5eae0c744e88bfae5eb3db45c469e941fec38292d9.
//
// Solidity: event DistributeRefReward(address trader, uint256 _totalFee, uint256 _tradeSize, uint256 rebase, uint256 discount)
func (_Referrals *ReferralsFilterer) ParseDistributeRefReward(log types.Log) (*ReferralsDistributeRefReward, error) {
	event := new(ReferralsDistributeRefReward)
	if err := _Referrals.contract.UnpackLog(event, "DistributeRefReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsIndirectRewardsIterator is returned from FilterIndirectRewards and is used to iterate over the raw logs and unpacked data for IndirectRewards events raised by the Referrals contract.
type ReferralsIndirectRewardsIterator struct {
	Event *ReferralsIndirectRewards // Event containing the contract specifics and raw log

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
func (it *ReferralsIndirectRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsIndirectRewards)
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
		it.Event = new(ReferralsIndirectRewards)
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
func (it *ReferralsIndirectRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsIndirectRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsIndirectRewards represents a IndirectRewards event raised by the Referrals contract.
type ReferralsIndirectRewards struct {
	Trader   common.Address
	Referrer common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIndirectRewards is a free log retrieval operation binding the contract event 0x18a9a3c6960f0356ff6553230f7444f126e6e29721111bf5274c13e9b8557403.
//
// Solidity: event IndirectRewards(address trader, address referrer, uint256 amount)
func (_Referrals *ReferralsFilterer) FilterIndirectRewards(opts *bind.FilterOpts) (*ReferralsIndirectRewardsIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "IndirectRewards")
	if err != nil {
		return nil, err
	}
	return &ReferralsIndirectRewardsIterator{contract: _Referrals.contract, event: "IndirectRewards", logs: logs, sub: sub}, nil
}

// WatchIndirectRewards is a free log subscription operation binding the contract event 0x18a9a3c6960f0356ff6553230f7444f126e6e29721111bf5274c13e9b8557403.
//
// Solidity: event IndirectRewards(address trader, address referrer, uint256 amount)
func (_Referrals *ReferralsFilterer) WatchIndirectRewards(opts *bind.WatchOpts, sink chan<- *ReferralsIndirectRewards) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "IndirectRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsIndirectRewards)
				if err := _Referrals.contract.UnpackLog(event, "IndirectRewards", log); err != nil {
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

// ParseIndirectRewards is a log parse operation binding the contract event 0x18a9a3c6960f0356ff6553230f7444f126e6e29721111bf5274c13e9b8557403.
//
// Solidity: event IndirectRewards(address trader, address referrer, uint256 amount)
func (_Referrals *ReferralsFilterer) ParseIndirectRewards(log types.Log) (*ReferralsIndirectRewards, error) {
	event := new(ReferralsIndirectRewards)
	if err := _Referrals.contract.UnpackLog(event, "IndirectRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Referrals contract.
type ReferralsInitializedIterator struct {
	Event *ReferralsInitialized // Event containing the contract specifics and raw log

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
func (it *ReferralsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsInitialized)
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
		it.Event = new(ReferralsInitialized)
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
func (it *ReferralsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsInitialized represents a Initialized event raised by the Referrals contract.
type ReferralsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Referrals *ReferralsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReferralsInitializedIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReferralsInitializedIterator{contract: _Referrals.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Referrals *ReferralsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReferralsInitialized) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsInitialized)
				if err := _Referrals.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Referrals *ReferralsFilterer) ParseInitialized(log types.Log) (*ReferralsInitialized, error) {
	event := new(ReferralsInitialized)
	if err := _Referrals.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Referrals contract.
type ReferralsOwnershipTransferStartedIterator struct {
	Event *ReferralsOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ReferralsOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsOwnershipTransferStarted)
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
		it.Event = new(ReferralsOwnershipTransferStarted)
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
func (it *ReferralsOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Referrals contract.
type ReferralsOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Referrals *ReferralsFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReferralsOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReferralsOwnershipTransferStartedIterator{contract: _Referrals.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Referrals *ReferralsFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ReferralsOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsOwnershipTransferStarted)
				if err := _Referrals.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Referrals *ReferralsFilterer) ParseOwnershipTransferStarted(log types.Log) (*ReferralsOwnershipTransferStarted, error) {
	event := new(ReferralsOwnershipTransferStarted)
	if err := _Referrals.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Referrals contract.
type ReferralsOwnershipTransferredIterator struct {
	Event *ReferralsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReferralsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsOwnershipTransferred)
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
		it.Event = new(ReferralsOwnershipTransferred)
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
func (it *ReferralsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsOwnershipTransferred represents a OwnershipTransferred event raised by the Referrals contract.
type ReferralsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Referrals *ReferralsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReferralsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReferralsOwnershipTransferredIterator{contract: _Referrals.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Referrals *ReferralsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReferralsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsOwnershipTransferred)
				if err := _Referrals.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Referrals *ReferralsFilterer) ParseOwnershipTransferred(log types.Log) (*ReferralsOwnershipTransferred, error) {
	event := new(ReferralsOwnershipTransferred)
	if err := _Referrals.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Referrals contract.
type ReferralsRegisterIterator struct {
	Event *ReferralsRegister // Event containing the contract specifics and raw log

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
func (it *ReferralsRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsRegister)
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
		it.Event = new(ReferralsRegister)
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
func (it *ReferralsRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsRegister represents a Register event raised by the Referrals contract.
type ReferralsRegister struct {
	User     common.Address
	Referrer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a351519.
//
// Solidity: event Register(address user, address referrer)
func (_Referrals *ReferralsFilterer) FilterRegister(opts *bind.FilterOpts) (*ReferralsRegisterIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &ReferralsRegisterIterator{contract: _Referrals.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a351519.
//
// Solidity: event Register(address user, address referrer)
func (_Referrals *ReferralsFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *ReferralsRegister) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsRegister)
				if err := _Referrals.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0x98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a351519.
//
// Solidity: event Register(address user, address referrer)
func (_Referrals *ReferralsFilterer) ParseRegister(log types.Log) (*ReferralsRegister, error) {
	event := new(ReferralsRegister)
	if err := _Referrals.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Referrals contract.
type ReferralsRewardsClaimedIterator struct {
	Event *ReferralsRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *ReferralsRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsRewardsClaimed)
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
		it.Event = new(ReferralsRewardsClaimed)
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
func (it *ReferralsRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsRewardsClaimed represents a RewardsClaimed event raised by the Referrals contract.
type ReferralsRewardsClaimed struct {
	User       common.Address
	Amount     *big.Int
	Compounded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xab3c4ae116d5285676cffaad9c32cae8acf49b49ca0fc9953b029f97acb955ef.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, bool compounded)
func (_Referrals *ReferralsFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address) (*ReferralsRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ReferralsRewardsClaimedIterator{contract: _Referrals.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xab3c4ae116d5285676cffaad9c32cae8acf49b49ca0fc9953b029f97acb955ef.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, bool compounded)
func (_Referrals *ReferralsFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ReferralsRewardsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsRewardsClaimed)
				if err := _Referrals.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xab3c4ae116d5285676cffaad9c32cae8acf49b49ca0fc9953b029f97acb955ef.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, bool compounded)
func (_Referrals *ReferralsFilterer) ParseRewardsClaimed(log types.Log) (*ReferralsRewardsClaimed, error) {
	event := new(ReferralsRewardsClaimed)
	if err := _Referrals.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsSetConfigIterator is returned from FilterSetConfig and is used to iterate over the raw logs and unpacked data for SetConfig events raised by the Referrals contract.
type ReferralsSetConfigIterator struct {
	Event *ReferralsSetConfig // Event containing the contract specifics and raw log

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
func (it *ReferralsSetConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsSetConfig)
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
		it.Event = new(ReferralsSetConfig)
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
func (it *ReferralsSetConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsSetConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsSetConfig represents a SetConfig event raised by the Referrals contract.
type ReferralsSetConfig struct {
	Config ReferralsConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetConfig is a free log retrieval operation binding the contract event 0xbeeecdc0f2c093644cb1fa6288be9385441def141b43886d61961a14e301d044.
//
// Solidity: event SetConfig((uint256,uint256,uint256[3]) _config)
func (_Referrals *ReferralsFilterer) FilterSetConfig(opts *bind.FilterOpts) (*ReferralsSetConfigIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return &ReferralsSetConfigIterator{contract: _Referrals.contract, event: "SetConfig", logs: logs, sub: sub}, nil
}

// WatchSetConfig is a free log subscription operation binding the contract event 0xbeeecdc0f2c093644cb1fa6288be9385441def141b43886d61961a14e301d044.
//
// Solidity: event SetConfig((uint256,uint256,uint256[3]) _config)
func (_Referrals *ReferralsFilterer) WatchSetConfig(opts *bind.WatchOpts, sink chan<- *ReferralsSetConfig) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "SetConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsSetConfig)
				if err := _Referrals.contract.UnpackLog(event, "SetConfig", log); err != nil {
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

// ParseSetConfig is a log parse operation binding the contract event 0xbeeecdc0f2c093644cb1fa6288be9385441def141b43886d61961a14e301d044.
//
// Solidity: event SetConfig((uint256,uint256,uint256[3]) _config)
func (_Referrals *ReferralsFilterer) ParseSetConfig(log types.Log) (*ReferralsSetConfig, error) {
	event := new(ReferralsSetConfig)
	if err := _Referrals.contract.UnpackLog(event, "SetConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsSetDiscountGradeIterator is returned from FilterSetDiscountGrade and is used to iterate over the raw logs and unpacked data for SetDiscountGrade events raised by the Referrals contract.
type ReferralsSetDiscountGradeIterator struct {
	Event *ReferralsSetDiscountGrade // Event containing the contract specifics and raw log

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
func (it *ReferralsSetDiscountGradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsSetDiscountGrade)
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
		it.Event = new(ReferralsSetDiscountGrade)
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
func (it *ReferralsSetDiscountGradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsSetDiscountGradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsSetDiscountGrade represents a SetDiscountGrade event raised by the Referrals contract.
type ReferralsSetDiscountGrade struct {
	User          common.Address
	Grade         uint8
	ReferrerGrade uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDiscountGrade is a free log retrieval operation binding the contract event 0xcf3701488755dcb741a9d9f570064b924c8d179cf685a1d5d9c76ab1ebd88aa5.
//
// Solidity: event SetDiscountGrade(address _user, uint8 _grade, uint8 _referrerGrade)
func (_Referrals *ReferralsFilterer) FilterSetDiscountGrade(opts *bind.FilterOpts) (*ReferralsSetDiscountGradeIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "SetDiscountGrade")
	if err != nil {
		return nil, err
	}
	return &ReferralsSetDiscountGradeIterator{contract: _Referrals.contract, event: "SetDiscountGrade", logs: logs, sub: sub}, nil
}

// WatchSetDiscountGrade is a free log subscription operation binding the contract event 0xcf3701488755dcb741a9d9f570064b924c8d179cf685a1d5d9c76ab1ebd88aa5.
//
// Solidity: event SetDiscountGrade(address _user, uint8 _grade, uint8 _referrerGrade)
func (_Referrals *ReferralsFilterer) WatchSetDiscountGrade(opts *bind.WatchOpts, sink chan<- *ReferralsSetDiscountGrade) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "SetDiscountGrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsSetDiscountGrade)
				if err := _Referrals.contract.UnpackLog(event, "SetDiscountGrade", log); err != nil {
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

// ParseSetDiscountGrade is a log parse operation binding the contract event 0xcf3701488755dcb741a9d9f570064b924c8d179cf685a1d5d9c76ab1ebd88aa5.
//
// Solidity: event SetDiscountGrade(address _user, uint8 _grade, uint8 _referrerGrade)
func (_Referrals *ReferralsFilterer) ParseSetDiscountGrade(log types.Log) (*ReferralsSetDiscountGrade, error) {
	event := new(ReferralsSetDiscountGrade)
	if err := _Referrals.contract.UnpackLog(event, "SetDiscountGrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsSetPubilcRegisterIterator is returned from FilterSetPubilcRegister and is used to iterate over the raw logs and unpacked data for SetPubilcRegister events raised by the Referrals contract.
type ReferralsSetPubilcRegisterIterator struct {
	Event *ReferralsSetPubilcRegister // Event containing the contract specifics and raw log

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
func (it *ReferralsSetPubilcRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsSetPubilcRegister)
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
		it.Event = new(ReferralsSetPubilcRegister)
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
func (it *ReferralsSetPubilcRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsSetPubilcRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsSetPubilcRegister represents a SetPubilcRegister event raised by the Referrals contract.
type ReferralsSetPubilcRegister struct {
	PubilcRegister bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetPubilcRegister is a free log retrieval operation binding the contract event 0x3c111f6ef63a268d35f6497d12923c5d1ef1187a3c06d2a57d15952b9f5cb3fe.
//
// Solidity: event SetPubilcRegister(bool _pubilcRegister)
func (_Referrals *ReferralsFilterer) FilterSetPubilcRegister(opts *bind.FilterOpts) (*ReferralsSetPubilcRegisterIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "SetPubilcRegister")
	if err != nil {
		return nil, err
	}
	return &ReferralsSetPubilcRegisterIterator{contract: _Referrals.contract, event: "SetPubilcRegister", logs: logs, sub: sub}, nil
}

// WatchSetPubilcRegister is a free log subscription operation binding the contract event 0x3c111f6ef63a268d35f6497d12923c5d1ef1187a3c06d2a57d15952b9f5cb3fe.
//
// Solidity: event SetPubilcRegister(bool _pubilcRegister)
func (_Referrals *ReferralsFilterer) WatchSetPubilcRegister(opts *bind.WatchOpts, sink chan<- *ReferralsSetPubilcRegister) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "SetPubilcRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsSetPubilcRegister)
				if err := _Referrals.contract.UnpackLog(event, "SetPubilcRegister", log); err != nil {
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

// ParseSetPubilcRegister is a log parse operation binding the contract event 0x3c111f6ef63a268d35f6497d12923c5d1ef1187a3c06d2a57d15952b9f5cb3fe.
//
// Solidity: event SetPubilcRegister(bool _pubilcRegister)
func (_Referrals *ReferralsFilterer) ParseSetPubilcRegister(log types.Log) (*ReferralsSetPubilcRegister, error) {
	event := new(ReferralsSetPubilcRegister)
	if err := _Referrals.contract.UnpackLog(event, "SetPubilcRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsSetStandardTokenIterator is returned from FilterSetStandardToken and is used to iterate over the raw logs and unpacked data for SetStandardToken events raised by the Referrals contract.
type ReferralsSetStandardTokenIterator struct {
	Event *ReferralsSetStandardToken // Event containing the contract specifics and raw log

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
func (it *ReferralsSetStandardTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsSetStandardToken)
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
		it.Event = new(ReferralsSetStandardToken)
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
func (it *ReferralsSetStandardTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsSetStandardTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsSetStandardToken represents a SetStandardToken event raised by the Referrals contract.
type ReferralsSetStandardToken struct {
	StandardToken common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetStandardToken is a free log retrieval operation binding the contract event 0x47cf141b4cb8862f22d0e50186bae1ff22cbb8727b5641f32211840beeb21004.
//
// Solidity: event SetStandardToken(address _standardToken)
func (_Referrals *ReferralsFilterer) FilterSetStandardToken(opts *bind.FilterOpts) (*ReferralsSetStandardTokenIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "SetStandardToken")
	if err != nil {
		return nil, err
	}
	return &ReferralsSetStandardTokenIterator{contract: _Referrals.contract, event: "SetStandardToken", logs: logs, sub: sub}, nil
}

// WatchSetStandardToken is a free log subscription operation binding the contract event 0x47cf141b4cb8862f22d0e50186bae1ff22cbb8727b5641f32211840beeb21004.
//
// Solidity: event SetStandardToken(address _standardToken)
func (_Referrals *ReferralsFilterer) WatchSetStandardToken(opts *bind.WatchOpts, sink chan<- *ReferralsSetStandardToken) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "SetStandardToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsSetStandardToken)
				if err := _Referrals.contract.UnpackLog(event, "SetStandardToken", log); err != nil {
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
func (_Referrals *ReferralsFilterer) ParseSetStandardToken(log types.Log) (*ReferralsSetStandardToken, error) {
	event := new(ReferralsSetStandardToken)
	if err := _Referrals.contract.UnpackLog(event, "SetStandardToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsSetTradingVaultIterator is returned from FilterSetTradingVault and is used to iterate over the raw logs and unpacked data for SetTradingVault events raised by the Referrals contract.
type ReferralsSetTradingVaultIterator struct {
	Event *ReferralsSetTradingVault // Event containing the contract specifics and raw log

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
func (it *ReferralsSetTradingVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsSetTradingVault)
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
		it.Event = new(ReferralsSetTradingVault)
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
func (it *ReferralsSetTradingVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsSetTradingVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsSetTradingVault represents a SetTradingVault event raised by the Referrals contract.
type ReferralsSetTradingVault struct {
	TradingVault common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetTradingVault is a free log retrieval operation binding the contract event 0x6fe680921e670dd2a8327fec07a433d938467f06126cf70eadb6d1d52ef3f7eb.
//
// Solidity: event SetTradingVault(address indexed tradingVault)
func (_Referrals *ReferralsFilterer) FilterSetTradingVault(opts *bind.FilterOpts, tradingVault []common.Address) (*ReferralsSetTradingVaultIterator, error) {

	var tradingVaultRule []interface{}
	for _, tradingVaultItem := range tradingVault {
		tradingVaultRule = append(tradingVaultRule, tradingVaultItem)
	}

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "SetTradingVault", tradingVaultRule)
	if err != nil {
		return nil, err
	}
	return &ReferralsSetTradingVaultIterator{contract: _Referrals.contract, event: "SetTradingVault", logs: logs, sub: sub}, nil
}

// WatchSetTradingVault is a free log subscription operation binding the contract event 0x6fe680921e670dd2a8327fec07a433d938467f06126cf70eadb6d1d52ef3f7eb.
//
// Solidity: event SetTradingVault(address indexed tradingVault)
func (_Referrals *ReferralsFilterer) WatchSetTradingVault(opts *bind.WatchOpts, sink chan<- *ReferralsSetTradingVault, tradingVault []common.Address) (event.Subscription, error) {

	var tradingVaultRule []interface{}
	for _, tradingVaultItem := range tradingVault {
		tradingVaultRule = append(tradingVaultRule, tradingVaultItem)
	}

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "SetTradingVault", tradingVaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsSetTradingVault)
				if err := _Referrals.contract.UnpackLog(event, "SetTradingVault", log); err != nil {
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

// ParseSetTradingVault is a log parse operation binding the contract event 0x6fe680921e670dd2a8327fec07a433d938467f06126cf70eadb6d1d52ef3f7eb.
//
// Solidity: event SetTradingVault(address indexed tradingVault)
func (_Referrals *ReferralsFilterer) ParseSetTradingVault(log types.Log) (*ReferralsSetTradingVault, error) {
	event := new(ReferralsSetTradingVault)
	if err := _Referrals.contract.UnpackLog(event, "SetTradingVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralsUpdateAdminIterator is returned from FilterUpdateAdmin and is used to iterate over the raw logs and unpacked data for UpdateAdmin events raised by the Referrals contract.
type ReferralsUpdateAdminIterator struct {
	Event *ReferralsUpdateAdmin // Event containing the contract specifics and raw log

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
func (it *ReferralsUpdateAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralsUpdateAdmin)
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
		it.Event = new(ReferralsUpdateAdmin)
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
func (it *ReferralsUpdateAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralsUpdateAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralsUpdateAdmin represents a UpdateAdmin event raised by the Referrals contract.
type ReferralsUpdateAdmin struct {
	Admin common.Address
	Flag  bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateAdmin is a free log retrieval operation binding the contract event 0xae0a768e1f5a7943e3f1bb8a4d503c6fbfea4c9bbbded6b463e48bebd28ef725.
//
// Solidity: event UpdateAdmin(address _admin, bool _flag)
func (_Referrals *ReferralsFilterer) FilterUpdateAdmin(opts *bind.FilterOpts) (*ReferralsUpdateAdminIterator, error) {

	logs, sub, err := _Referrals.contract.FilterLogs(opts, "UpdateAdmin")
	if err != nil {
		return nil, err
	}
	return &ReferralsUpdateAdminIterator{contract: _Referrals.contract, event: "UpdateAdmin", logs: logs, sub: sub}, nil
}

// WatchUpdateAdmin is a free log subscription operation binding the contract event 0xae0a768e1f5a7943e3f1bb8a4d503c6fbfea4c9bbbded6b463e48bebd28ef725.
//
// Solidity: event UpdateAdmin(address _admin, bool _flag)
func (_Referrals *ReferralsFilterer) WatchUpdateAdmin(opts *bind.WatchOpts, sink chan<- *ReferralsUpdateAdmin) (event.Subscription, error) {

	logs, sub, err := _Referrals.contract.WatchLogs(opts, "UpdateAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralsUpdateAdmin)
				if err := _Referrals.contract.UnpackLog(event, "UpdateAdmin", log); err != nil {
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

// ParseUpdateAdmin is a log parse operation binding the contract event 0xae0a768e1f5a7943e3f1bb8a4d503c6fbfea4c9bbbded6b463e48bebd28ef725.
//
// Solidity: event UpdateAdmin(address _admin, bool _flag)
func (_Referrals *ReferralsFilterer) ParseUpdateAdmin(log types.Log) (*ReferralsUpdateAdmin, error) {
	event := new(ReferralsUpdateAdmin)
	if err := _Referrals.contract.UnpackLog(event, "UpdateAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
