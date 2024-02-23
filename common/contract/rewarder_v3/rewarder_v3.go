// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewarder_v3

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

// BoostedMultiRewarderRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type BoostedMultiRewarderRewardInfo struct {
	RewardToken            common.Address
	TokenPerSec            *big.Int
	AccTokenPerShare       *big.Int
	AccTokenPerFactorShare *big.Int
	DistributedAmount      *big.Int
	ClaimedAmount          *big.Int
	LastRewardTimestamp    *big.Int
}

// RewarderV3MetaData contains all meta data concerning the RewarderV3 contract.
var RewarderV3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDeprecated\",\"type\":\"bool\"}],\"name\":\"IsDeprecatedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OnReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"RewardRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newStartTime\",\"type\":\"uint40\"}],\"name\":\"StartTimeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACC_TOKEN_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REWARD_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOKEN_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLE_OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_PARTITION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"_startTimestampOrNow\",\"type\":\"uint40\"},{\"internalType\":\"uint96\",\"name\":\"_tokenPerSec\",\"type\":\"uint96\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basePartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bribeFactory\",\"outputs\":[{\"internalType\":\"contractIBribeRewarderFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyClaimReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"emergencyTokenWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBribeRewarderFactory\",\"name\":\"_bribeFactory\",\"type\":\"address\"},{\"internalType\":\"contractIBoostedMasterWombat\",\"name\":\"_masterWombat\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_tokenPerSec\",\"type\":\"uint96\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmissionActive\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"isActive_\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterWombat\",\"outputs\":[{\"internalType\":\"contractIBoostedMasterWombat\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newLpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newFactor\",\"type\":\"uint256\"}],\"name\":\"onReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newFactor\",\"type\":\"uint256\"}],\"name\":\"onUpdateFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"rewardInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"tokenPerSec\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"accTokenPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accTokenPerFactorShare\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"distributedAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimedAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint40\"}],\"internalType\":\"structBoostedMultiRewarder.RewardInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInfos\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"tokenPerSec\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"accTokenPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"accTokenPerFactorShare\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"distributedAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimedAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokenSurpluses\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"surpluses_\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToDistribute\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"runoutTimestamps\",\"outputs\":[{\"internalType\":\"uint40[]\",\"name\":\"timestamps_\",\"type\":\"uint40[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isDeprecated\",\"type\":\"bool\"}],\"name\":\"setIsDeprecated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_tokenPerSec\",\"type\":\"uint96\"},{\"internalType\":\"uint40\",\"name\":\"_startTimestampToOverride\",\"type\":\"uint40\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBalanceInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"factor\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"rewardDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unpaidRewards\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RewarderV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use RewarderV3MetaData.ABI instead.
var RewarderV3ABI = RewarderV3MetaData.ABI

// RewarderV3 is an auto generated Go binding around an Ethereum contract.
type RewarderV3 struct {
	RewarderV3Caller     // Read-only binding to the contract
	RewarderV3Transactor // Write-only binding to the contract
	RewarderV3Filterer   // Log filterer for contract events
}

// RewarderV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type RewarderV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RewarderV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewarderV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewarderV3Session struct {
	Contract     *RewarderV3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewarderV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewarderV3CallerSession struct {
	Contract *RewarderV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RewarderV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewarderV3TransactorSession struct {
	Contract     *RewarderV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RewarderV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type RewarderV3Raw struct {
	Contract *RewarderV3 // Generic contract binding to access the raw methods on
}

// RewarderV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewarderV3CallerRaw struct {
	Contract *RewarderV3Caller // Generic read-only contract binding to access the raw methods on
}

// RewarderV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewarderV3TransactorRaw struct {
	Contract *RewarderV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRewarderV3 creates a new instance of RewarderV3, bound to a specific deployed contract.
func NewRewarderV3(address common.Address, backend bind.ContractBackend) (*RewarderV3, error) {
	contract, err := bindRewarderV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewarderV3{RewarderV3Caller: RewarderV3Caller{contract: contract}, RewarderV3Transactor: RewarderV3Transactor{contract: contract}, RewarderV3Filterer: RewarderV3Filterer{contract: contract}}, nil
}

// NewRewarderV3Caller creates a new read-only instance of RewarderV3, bound to a specific deployed contract.
func NewRewarderV3Caller(address common.Address, caller bind.ContractCaller) (*RewarderV3Caller, error) {
	contract, err := bindRewarderV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderV3Caller{contract: contract}, nil
}

// NewRewarderV3Transactor creates a new write-only instance of RewarderV3, bound to a specific deployed contract.
func NewRewarderV3Transactor(address common.Address, transactor bind.ContractTransactor) (*RewarderV3Transactor, error) {
	contract, err := bindRewarderV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderV3Transactor{contract: contract}, nil
}

// NewRewarderV3Filterer creates a new log filterer instance of RewarderV3, bound to a specific deployed contract.
func NewRewarderV3Filterer(address common.Address, filterer bind.ContractFilterer) (*RewarderV3Filterer, error) {
	contract, err := bindRewarderV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewarderV3Filterer{contract: contract}, nil
}

// bindRewarderV3 binds a generic wrapper to an already deployed contract.
func bindRewarderV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewarderV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewarderV3 *RewarderV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewarderV3.Contract.RewarderV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewarderV3 *RewarderV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.Contract.RewarderV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewarderV3 *RewarderV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewarderV3.Contract.RewarderV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewarderV3 *RewarderV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewarderV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewarderV3 *RewarderV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewarderV3 *RewarderV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewarderV3.Contract.contract.Transact(opts, method, params...)
}

// ACCTOKENPRECISION is a free data retrieval call binding the contract method 0xeea01604.
//
// Solidity: function ACC_TOKEN_PRECISION() view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) ACCTOKENPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "ACC_TOKEN_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ACCTOKENPRECISION is a free data retrieval call binding the contract method 0xeea01604.
//
// Solidity: function ACC_TOKEN_PRECISION() view returns(uint256)
func (_RewarderV3 *RewarderV3Session) ACCTOKENPRECISION() (*big.Int, error) {
	return _RewarderV3.Contract.ACCTOKENPRECISION(&_RewarderV3.CallOpts)
}

// ACCTOKENPRECISION is a free data retrieval call binding the contract method 0xeea01604.
//
// Solidity: function ACC_TOKEN_PRECISION() view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) ACCTOKENPRECISION() (*big.Int, error) {
	return _RewarderV3.Contract.ACCTOKENPRECISION(&_RewarderV3.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewarderV3 *RewarderV3Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewarderV3 *RewarderV3Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewarderV3.Contract.DEFAULTADMINROLE(&_RewarderV3.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewarderV3 *RewarderV3CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewarderV3.Contract.DEFAULTADMINROLE(&_RewarderV3.CallOpts)
}

// MAXREWARDTOKENS is a free data retrieval call binding the contract method 0x5d0cde97.
//
// Solidity: function MAX_REWARD_TOKENS() view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) MAXREWARDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "MAX_REWARD_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREWARDTOKENS is a free data retrieval call binding the contract method 0x5d0cde97.
//
// Solidity: function MAX_REWARD_TOKENS() view returns(uint256)
func (_RewarderV3 *RewarderV3Session) MAXREWARDTOKENS() (*big.Int, error) {
	return _RewarderV3.Contract.MAXREWARDTOKENS(&_RewarderV3.CallOpts)
}

// MAXREWARDTOKENS is a free data retrieval call binding the contract method 0x5d0cde97.
//
// Solidity: function MAX_REWARD_TOKENS() view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) MAXREWARDTOKENS() (*big.Int, error) {
	return _RewarderV3.Contract.MAXREWARDTOKENS(&_RewarderV3.CallOpts)
}

// MAXTOKENRATE is a free data retrieval call binding the contract method 0xd6a084f3.
//
// Solidity: function MAX_TOKEN_RATE() view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) MAXTOKENRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "MAX_TOKEN_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKENRATE is a free data retrieval call binding the contract method 0xd6a084f3.
//
// Solidity: function MAX_TOKEN_RATE() view returns(uint256)
func (_RewarderV3 *RewarderV3Session) MAXTOKENRATE() (*big.Int, error) {
	return _RewarderV3.Contract.MAXTOKENRATE(&_RewarderV3.CallOpts)
}

// MAXTOKENRATE is a free data retrieval call binding the contract method 0xd6a084f3.
//
// Solidity: function MAX_TOKEN_RATE() view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) MAXTOKENRATE() (*big.Int, error) {
	return _RewarderV3.Contract.MAXTOKENRATE(&_RewarderV3.CallOpts)
}

// ROLEOPERATOR is a free data retrieval call binding the contract method 0x98a1b397.
//
// Solidity: function ROLE_OPERATOR() view returns(bytes32)
func (_RewarderV3 *RewarderV3Caller) ROLEOPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "ROLE_OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROLEOPERATOR is a free data retrieval call binding the contract method 0x98a1b397.
//
// Solidity: function ROLE_OPERATOR() view returns(bytes32)
func (_RewarderV3 *RewarderV3Session) ROLEOPERATOR() ([32]byte, error) {
	return _RewarderV3.Contract.ROLEOPERATOR(&_RewarderV3.CallOpts)
}

// ROLEOPERATOR is a free data retrieval call binding the contract method 0x98a1b397.
//
// Solidity: function ROLE_OPERATOR() view returns(bytes32)
func (_RewarderV3 *RewarderV3CallerSession) ROLEOPERATOR() ([32]byte, error) {
	return _RewarderV3.Contract.ROLEOPERATOR(&_RewarderV3.CallOpts)
}

// TOTALPARTITION is a free data retrieval call binding the contract method 0x441b502c.
//
// Solidity: function TOTAL_PARTITION() view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) TOTALPARTITION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "TOTAL_PARTITION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALPARTITION is a free data retrieval call binding the contract method 0x441b502c.
//
// Solidity: function TOTAL_PARTITION() view returns(uint256)
func (_RewarderV3 *RewarderV3Session) TOTALPARTITION() (*big.Int, error) {
	return _RewarderV3.Contract.TOTALPARTITION(&_RewarderV3.CallOpts)
}

// TOTALPARTITION is a free data retrieval call binding the contract method 0x441b502c.
//
// Solidity: function TOTAL_PARTITION() view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) TOTALPARTITION() (*big.Int, error) {
	return _RewarderV3.Contract.TOTALPARTITION(&_RewarderV3.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256[] balances_)
func (_RewarderV3 *RewarderV3Caller) Balances(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "balances")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256[] balances_)
func (_RewarderV3 *RewarderV3Session) Balances() ([]*big.Int, error) {
	return _RewarderV3.Contract.Balances(&_RewarderV3.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256[] balances_)
func (_RewarderV3 *RewarderV3CallerSession) Balances() ([]*big.Int, error) {
	return _RewarderV3.Contract.Balances(&_RewarderV3.CallOpts)
}

// BasePartition is a free data retrieval call binding the contract method 0xabfef111.
//
// Solidity: function basePartition() view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) BasePartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "basePartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePartition is a free data retrieval call binding the contract method 0xabfef111.
//
// Solidity: function basePartition() view returns(uint256)
func (_RewarderV3 *RewarderV3Session) BasePartition() (*big.Int, error) {
	return _RewarderV3.Contract.BasePartition(&_RewarderV3.CallOpts)
}

// BasePartition is a free data retrieval call binding the contract method 0xabfef111.
//
// Solidity: function basePartition() view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) BasePartition() (*big.Int, error) {
	return _RewarderV3.Contract.BasePartition(&_RewarderV3.CallOpts)
}

// BribeFactory is a free data retrieval call binding the contract method 0xeb4a78e0.
//
// Solidity: function bribeFactory() view returns(address)
func (_RewarderV3 *RewarderV3Caller) BribeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "bribeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BribeFactory is a free data retrieval call binding the contract method 0xeb4a78e0.
//
// Solidity: function bribeFactory() view returns(address)
func (_RewarderV3 *RewarderV3Session) BribeFactory() (common.Address, error) {
	return _RewarderV3.Contract.BribeFactory(&_RewarderV3.CallOpts)
}

// BribeFactory is a free data retrieval call binding the contract method 0xeb4a78e0.
//
// Solidity: function bribeFactory() view returns(address)
func (_RewarderV3 *RewarderV3CallerSession) BribeFactory() (common.Address, error) {
	return _RewarderV3.Contract.BribeFactory(&_RewarderV3.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewarderV3 *RewarderV3Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewarderV3 *RewarderV3Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewarderV3.Contract.GetRoleAdmin(&_RewarderV3.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewarderV3 *RewarderV3CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewarderV3.Contract.GetRoleAdmin(&_RewarderV3.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RewarderV3 *RewarderV3Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RewarderV3 *RewarderV3Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RewarderV3.Contract.GetRoleMember(&_RewarderV3.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RewarderV3 *RewarderV3CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RewarderV3.Contract.GetRoleMember(&_RewarderV3.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RewarderV3 *RewarderV3Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RewarderV3.Contract.GetRoleMemberCount(&_RewarderV3.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RewarderV3.Contract.GetRoleMemberCount(&_RewarderV3.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewarderV3 *RewarderV3Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewarderV3 *RewarderV3Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewarderV3.Contract.HasRole(&_RewarderV3.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewarderV3 *RewarderV3CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewarderV3.Contract.HasRole(&_RewarderV3.CallOpts, role, account)
}

// IsDeprecated is a free data retrieval call binding the contract method 0xc7178230.
//
// Solidity: function isDeprecated() view returns(bool)
func (_RewarderV3 *RewarderV3Caller) IsDeprecated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "isDeprecated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeprecated is a free data retrieval call binding the contract method 0xc7178230.
//
// Solidity: function isDeprecated() view returns(bool)
func (_RewarderV3 *RewarderV3Session) IsDeprecated() (bool, error) {
	return _RewarderV3.Contract.IsDeprecated(&_RewarderV3.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0xc7178230.
//
// Solidity: function isDeprecated() view returns(bool)
func (_RewarderV3 *RewarderV3CallerSession) IsDeprecated() (bool, error) {
	return _RewarderV3.Contract.IsDeprecated(&_RewarderV3.CallOpts)
}

// IsEmissionActive is a free data retrieval call binding the contract method 0x4e8a947e.
//
// Solidity: function isEmissionActive() view returns(bool[] isActive_)
func (_RewarderV3 *RewarderV3Caller) IsEmissionActive(opts *bind.CallOpts) ([]bool, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "isEmissionActive")

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// IsEmissionActive is a free data retrieval call binding the contract method 0x4e8a947e.
//
// Solidity: function isEmissionActive() view returns(bool[] isActive_)
func (_RewarderV3 *RewarderV3Session) IsEmissionActive() ([]bool, error) {
	return _RewarderV3.Contract.IsEmissionActive(&_RewarderV3.CallOpts)
}

// IsEmissionActive is a free data retrieval call binding the contract method 0x4e8a947e.
//
// Solidity: function isEmissionActive() view returns(bool[] isActive_)
func (_RewarderV3 *RewarderV3CallerSession) IsEmissionActive() ([]bool, error) {
	return _RewarderV3.Contract.IsEmissionActive(&_RewarderV3.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_RewarderV3 *RewarderV3Caller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_RewarderV3 *RewarderV3Session) LpToken() (common.Address, error) {
	return _RewarderV3.Contract.LpToken(&_RewarderV3.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_RewarderV3 *RewarderV3CallerSession) LpToken() (common.Address, error) {
	return _RewarderV3.Contract.LpToken(&_RewarderV3.CallOpts)
}

// MasterWombat is a free data retrieval call binding the contract method 0x3bd61ba8.
//
// Solidity: function masterWombat() view returns(address)
func (_RewarderV3 *RewarderV3Caller) MasterWombat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "masterWombat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterWombat is a free data retrieval call binding the contract method 0x3bd61ba8.
//
// Solidity: function masterWombat() view returns(address)
func (_RewarderV3 *RewarderV3Session) MasterWombat() (common.Address, error) {
	return _RewarderV3.Contract.MasterWombat(&_RewarderV3.CallOpts)
}

// MasterWombat is a free data retrieval call binding the contract method 0x3bd61ba8.
//
// Solidity: function masterWombat() view returns(address)
func (_RewarderV3 *RewarderV3CallerSession) MasterWombat() (common.Address, error) {
	return _RewarderV3.Contract.MasterWombat(&_RewarderV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewarderV3 *RewarderV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewarderV3 *RewarderV3Session) Owner() (common.Address, error) {
	return _RewarderV3.Contract.Owner(&_RewarderV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewarderV3 *RewarderV3CallerSession) Owner() (common.Address, error) {
	return _RewarderV3.Contract.Owner(&_RewarderV3.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256[] rewards_)
func (_RewarderV3 *RewarderV3Caller) PendingTokens(opts *bind.CallOpts, _user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "pendingTokens", _user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256[] rewards_)
func (_RewarderV3 *RewarderV3Session) PendingTokens(_user common.Address) ([]*big.Int, error) {
	return _RewarderV3.Contract.PendingTokens(&_RewarderV3.CallOpts, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xc031a66f.
//
// Solidity: function pendingTokens(address _user) view returns(uint256[] rewards_)
func (_RewarderV3 *RewarderV3CallerSession) PendingTokens(_user common.Address) ([]*big.Int, error) {
	return _RewarderV3.Contract.PendingTokens(&_RewarderV3.CallOpts, _user)
}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 i) view returns((address,uint96,uint128,uint128,uint128,uint128,uint40) info)
func (_RewarderV3 *RewarderV3Caller) RewardInfo(opts *bind.CallOpts, i *big.Int) (BoostedMultiRewarderRewardInfo, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "rewardInfo", i)

	if err != nil {
		return *new(BoostedMultiRewarderRewardInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(BoostedMultiRewarderRewardInfo)).(*BoostedMultiRewarderRewardInfo)

	return out0, err

}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 i) view returns((address,uint96,uint128,uint128,uint128,uint128,uint40) info)
func (_RewarderV3 *RewarderV3Session) RewardInfo(i *big.Int) (BoostedMultiRewarderRewardInfo, error) {
	return _RewarderV3.Contract.RewardInfo(&_RewarderV3.CallOpts, i)
}

// RewardInfo is a free data retrieval call binding the contract method 0x81a00f83.
//
// Solidity: function rewardInfo(uint256 i) view returns((address,uint96,uint128,uint128,uint128,uint128,uint40) info)
func (_RewarderV3 *RewarderV3CallerSession) RewardInfo(i *big.Int) (BoostedMultiRewarderRewardInfo, error) {
	return _RewarderV3.Contract.RewardInfo(&_RewarderV3.CallOpts, i)
}

// RewardInfos is a free data retrieval call binding the contract method 0x89d6517f.
//
// Solidity: function rewardInfos(uint256 ) view returns(address rewardToken, uint96 tokenPerSec, uint128 accTokenPerShare, uint128 accTokenPerFactorShare, uint128 distributedAmount, uint128 claimedAmount, uint40 lastRewardTimestamp)
func (_RewarderV3 *RewarderV3Caller) RewardInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RewardToken            common.Address
	TokenPerSec            *big.Int
	AccTokenPerShare       *big.Int
	AccTokenPerFactorShare *big.Int
	DistributedAmount      *big.Int
	ClaimedAmount          *big.Int
	LastRewardTimestamp    *big.Int
}, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "rewardInfos", arg0)

	outstruct := new(struct {
		RewardToken            common.Address
		TokenPerSec            *big.Int
		AccTokenPerShare       *big.Int
		AccTokenPerFactorShare *big.Int
		DistributedAmount      *big.Int
		ClaimedAmount          *big.Int
		LastRewardTimestamp    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenPerSec = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccTokenPerShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccTokenPerFactorShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DistributedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ClaimedAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInfos is a free data retrieval call binding the contract method 0x89d6517f.
//
// Solidity: function rewardInfos(uint256 ) view returns(address rewardToken, uint96 tokenPerSec, uint128 accTokenPerShare, uint128 accTokenPerFactorShare, uint128 distributedAmount, uint128 claimedAmount, uint40 lastRewardTimestamp)
func (_RewarderV3 *RewarderV3Session) RewardInfos(arg0 *big.Int) (struct {
	RewardToken            common.Address
	TokenPerSec            *big.Int
	AccTokenPerShare       *big.Int
	AccTokenPerFactorShare *big.Int
	DistributedAmount      *big.Int
	ClaimedAmount          *big.Int
	LastRewardTimestamp    *big.Int
}, error) {
	return _RewarderV3.Contract.RewardInfos(&_RewarderV3.CallOpts, arg0)
}

// RewardInfos is a free data retrieval call binding the contract method 0x89d6517f.
//
// Solidity: function rewardInfos(uint256 ) view returns(address rewardToken, uint96 tokenPerSec, uint128 accTokenPerShare, uint128 accTokenPerFactorShare, uint128 distributedAmount, uint128 claimedAmount, uint40 lastRewardTimestamp)
func (_RewarderV3 *RewarderV3CallerSession) RewardInfos(arg0 *big.Int) (struct {
	RewardToken            common.Address
	TokenPerSec            *big.Int
	AccTokenPerShare       *big.Int
	AccTokenPerFactorShare *big.Int
	DistributedAmount      *big.Int
	ClaimedAmount          *big.Int
	LastRewardTimestamp    *big.Int
}, error) {
	return _RewarderV3.Contract.RewardInfos(&_RewarderV3.CallOpts, arg0)
}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_RewarderV3 *RewarderV3Caller) RewardLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "rewardLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_RewarderV3 *RewarderV3Session) RewardLength() (*big.Int, error) {
	return _RewarderV3.Contract.RewardLength(&_RewarderV3.CallOpts)
}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_RewarderV3 *RewarderV3CallerSession) RewardLength() (*big.Int, error) {
	return _RewarderV3.Contract.RewardLength(&_RewarderV3.CallOpts)
}

// RewardTokenSurpluses is a free data retrieval call binding the contract method 0x8bab3786.
//
// Solidity: function rewardTokenSurpluses() view returns(int256[] surpluses_)
func (_RewarderV3 *RewarderV3Caller) RewardTokenSurpluses(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "rewardTokenSurpluses")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RewardTokenSurpluses is a free data retrieval call binding the contract method 0x8bab3786.
//
// Solidity: function rewardTokenSurpluses() view returns(int256[] surpluses_)
func (_RewarderV3 *RewarderV3Session) RewardTokenSurpluses() ([]*big.Int, error) {
	return _RewarderV3.Contract.RewardTokenSurpluses(&_RewarderV3.CallOpts)
}

// RewardTokenSurpluses is a free data retrieval call binding the contract method 0x8bab3786.
//
// Solidity: function rewardTokenSurpluses() view returns(int256[] surpluses_)
func (_RewarderV3 *RewarderV3CallerSession) RewardTokenSurpluses() ([]*big.Int, error) {
	return _RewarderV3.Contract.RewardTokenSurpluses(&_RewarderV3.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0xc2b18aa0.
//
// Solidity: function rewardTokens() view returns(address[] tokens)
func (_RewarderV3 *RewarderV3Caller) RewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "rewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0xc2b18aa0.
//
// Solidity: function rewardTokens() view returns(address[] tokens)
func (_RewarderV3 *RewarderV3Session) RewardTokens() ([]common.Address, error) {
	return _RewarderV3.Contract.RewardTokens(&_RewarderV3.CallOpts)
}

// RewardTokens is a free data retrieval call binding the contract method 0xc2b18aa0.
//
// Solidity: function rewardTokens() view returns(address[] tokens)
func (_RewarderV3 *RewarderV3CallerSession) RewardTokens() ([]common.Address, error) {
	return _RewarderV3.Contract.RewardTokens(&_RewarderV3.CallOpts)
}

// RewardsToDistribute is a free data retrieval call binding the contract method 0xd5dd2c4f.
//
// Solidity: function rewardsToDistribute() view returns(uint256[] rewards_)
func (_RewarderV3 *RewarderV3Caller) RewardsToDistribute(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "rewardsToDistribute")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RewardsToDistribute is a free data retrieval call binding the contract method 0xd5dd2c4f.
//
// Solidity: function rewardsToDistribute() view returns(uint256[] rewards_)
func (_RewarderV3 *RewarderV3Session) RewardsToDistribute() ([]*big.Int, error) {
	return _RewarderV3.Contract.RewardsToDistribute(&_RewarderV3.CallOpts)
}

// RewardsToDistribute is a free data retrieval call binding the contract method 0xd5dd2c4f.
//
// Solidity: function rewardsToDistribute() view returns(uint256[] rewards_)
func (_RewarderV3 *RewarderV3CallerSession) RewardsToDistribute() ([]*big.Int, error) {
	return _RewarderV3.Contract.RewardsToDistribute(&_RewarderV3.CallOpts)
}

// RunoutTimestamps is a free data retrieval call binding the contract method 0x858ce1a4.
//
// Solidity: function runoutTimestamps() view returns(uint40[] timestamps_)
func (_RewarderV3 *RewarderV3Caller) RunoutTimestamps(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "runoutTimestamps")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RunoutTimestamps is a free data retrieval call binding the contract method 0x858ce1a4.
//
// Solidity: function runoutTimestamps() view returns(uint40[] timestamps_)
func (_RewarderV3 *RewarderV3Session) RunoutTimestamps() ([]*big.Int, error) {
	return _RewarderV3.Contract.RunoutTimestamps(&_RewarderV3.CallOpts)
}

// RunoutTimestamps is a free data retrieval call binding the contract method 0x858ce1a4.
//
// Solidity: function runoutTimestamps() view returns(uint40[] timestamps_)
func (_RewarderV3 *RewarderV3CallerSession) RunoutTimestamps() ([]*big.Int, error) {
	return _RewarderV3.Contract.RunoutTimestamps(&_RewarderV3.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewarderV3 *RewarderV3Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewarderV3 *RewarderV3Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewarderV3.Contract.SupportsInterface(&_RewarderV3.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewarderV3 *RewarderV3CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewarderV3.Contract.SupportsInterface(&_RewarderV3.CallOpts, interfaceId)
}

// UserBalanceInfo is a free data retrieval call binding the contract method 0x112a5d58.
//
// Solidity: function userBalanceInfo(address ) view returns(uint128 amount, uint128 factor)
func (_RewarderV3 *RewarderV3Caller) UserBalanceInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Factor *big.Int
}, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "userBalanceInfo", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Factor *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Factor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserBalanceInfo is a free data retrieval call binding the contract method 0x112a5d58.
//
// Solidity: function userBalanceInfo(address ) view returns(uint128 amount, uint128 factor)
func (_RewarderV3 *RewarderV3Session) UserBalanceInfo(arg0 common.Address) (struct {
	Amount *big.Int
	Factor *big.Int
}, error) {
	return _RewarderV3.Contract.UserBalanceInfo(&_RewarderV3.CallOpts, arg0)
}

// UserBalanceInfo is a free data retrieval call binding the contract method 0x112a5d58.
//
// Solidity: function userBalanceInfo(address ) view returns(uint128 amount, uint128 factor)
func (_RewarderV3 *RewarderV3CallerSession) UserBalanceInfo(arg0 common.Address) (struct {
	Amount *big.Int
	Factor *big.Int
}, error) {
	return _RewarderV3.Contract.UserBalanceInfo(&_RewarderV3.CallOpts, arg0)
}

// UserRewardInfo is a free data retrieval call binding the contract method 0xf709f23b.
//
// Solidity: function userRewardInfo(uint256 , address ) view returns(uint128 rewardDebt, uint128 unpaidRewards)
func (_RewarderV3 *RewarderV3Caller) UserRewardInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	RewardDebt    *big.Int
	UnpaidRewards *big.Int
}, error) {
	var out []interface{}
	err := _RewarderV3.contract.Call(opts, &out, "userRewardInfo", arg0, arg1)

	outstruct := new(struct {
		RewardDebt    *big.Int
		UnpaidRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardDebt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnpaidRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserRewardInfo is a free data retrieval call binding the contract method 0xf709f23b.
//
// Solidity: function userRewardInfo(uint256 , address ) view returns(uint128 rewardDebt, uint128 unpaidRewards)
func (_RewarderV3 *RewarderV3Session) UserRewardInfo(arg0 *big.Int, arg1 common.Address) (struct {
	RewardDebt    *big.Int
	UnpaidRewards *big.Int
}, error) {
	return _RewarderV3.Contract.UserRewardInfo(&_RewarderV3.CallOpts, arg0, arg1)
}

// UserRewardInfo is a free data retrieval call binding the contract method 0xf709f23b.
//
// Solidity: function userRewardInfo(uint256 , address ) view returns(uint128 rewardDebt, uint128 unpaidRewards)
func (_RewarderV3 *RewarderV3CallerSession) UserRewardInfo(arg0 *big.Int, arg1 common.Address) (struct {
	RewardDebt    *big.Int
	UnpaidRewards *big.Int
}, error) {
	return _RewarderV3.Contract.UserRewardInfo(&_RewarderV3.CallOpts, arg0, arg1)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_RewarderV3 *RewarderV3Transactor) AddOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "addOperator", _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_RewarderV3 *RewarderV3Session) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.AddOperator(&_RewarderV3.TransactOpts, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_RewarderV3 *RewarderV3TransactorSession) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.AddOperator(&_RewarderV3.TransactOpts, _operator)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa5eceb1b.
//
// Solidity: function addRewardToken(address _rewardToken, uint40 _startTimestampOrNow, uint96 _tokenPerSec) returns()
func (_RewarderV3 *RewarderV3Transactor) AddRewardToken(opts *bind.TransactOpts, _rewardToken common.Address, _startTimestampOrNow *big.Int, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "addRewardToken", _rewardToken, _startTimestampOrNow, _tokenPerSec)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa5eceb1b.
//
// Solidity: function addRewardToken(address _rewardToken, uint40 _startTimestampOrNow, uint96 _tokenPerSec) returns()
func (_RewarderV3 *RewarderV3Session) AddRewardToken(_rewardToken common.Address, _startTimestampOrNow *big.Int, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.AddRewardToken(&_RewarderV3.TransactOpts, _rewardToken, _startTimestampOrNow, _tokenPerSec)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa5eceb1b.
//
// Solidity: function addRewardToken(address _rewardToken, uint40 _startTimestampOrNow, uint96 _tokenPerSec) returns()
func (_RewarderV3 *RewarderV3TransactorSession) AddRewardToken(_rewardToken common.Address, _startTimestampOrNow *big.Int, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.AddRewardToken(&_RewarderV3.TransactOpts, _rewardToken, _startTimestampOrNow, _tokenPerSec)
}

// EmergencyClaimReward is a paid mutator transaction binding the contract method 0x12bc7bfd.
//
// Solidity: function emergencyClaimReward() returns(uint256[] rewards)
func (_RewarderV3 *RewarderV3Transactor) EmergencyClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "emergencyClaimReward")
}

// EmergencyClaimReward is a paid mutator transaction binding the contract method 0x12bc7bfd.
//
// Solidity: function emergencyClaimReward() returns(uint256[] rewards)
func (_RewarderV3 *RewarderV3Session) EmergencyClaimReward() (*types.Transaction, error) {
	return _RewarderV3.Contract.EmergencyClaimReward(&_RewarderV3.TransactOpts)
}

// EmergencyClaimReward is a paid mutator transaction binding the contract method 0x12bc7bfd.
//
// Solidity: function emergencyClaimReward() returns(uint256[] rewards)
func (_RewarderV3 *RewarderV3TransactorSession) EmergencyClaimReward() (*types.Transaction, error) {
	return _RewarderV3.Contract.EmergencyClaimReward(&_RewarderV3.TransactOpts)
}

// EmergencyTokenWithdraw is a paid mutator transaction binding the contract method 0xad568827.
//
// Solidity: function emergencyTokenWithdraw(address token) returns()
func (_RewarderV3 *RewarderV3Transactor) EmergencyTokenWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "emergencyTokenWithdraw", token)
}

// EmergencyTokenWithdraw is a paid mutator transaction binding the contract method 0xad568827.
//
// Solidity: function emergencyTokenWithdraw(address token) returns()
func (_RewarderV3 *RewarderV3Session) EmergencyTokenWithdraw(token common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.EmergencyTokenWithdraw(&_RewarderV3.TransactOpts, token)
}

// EmergencyTokenWithdraw is a paid mutator transaction binding the contract method 0xad568827.
//
// Solidity: function emergencyTokenWithdraw(address token) returns()
func (_RewarderV3 *RewarderV3TransactorSession) EmergencyTokenWithdraw(token common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.EmergencyTokenWithdraw(&_RewarderV3.TransactOpts, token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_RewarderV3 *RewarderV3Transactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_RewarderV3 *RewarderV3Session) EmergencyWithdraw() (*types.Transaction, error) {
	return _RewarderV3.Contract.EmergencyWithdraw(&_RewarderV3.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_RewarderV3 *RewarderV3TransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _RewarderV3.Contract.EmergencyWithdraw(&_RewarderV3.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.GrantRole(&_RewarderV3.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.GrantRole(&_RewarderV3.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x731994fc.
//
// Solidity: function initialize(address _bribeFactory, address _masterWombat, address _lpToken, uint256 _startTimestamp, address _rewardToken, uint96 _tokenPerSec) returns()
func (_RewarderV3 *RewarderV3Transactor) Initialize(opts *bind.TransactOpts, _bribeFactory common.Address, _masterWombat common.Address, _lpToken common.Address, _startTimestamp *big.Int, _rewardToken common.Address, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "initialize", _bribeFactory, _masterWombat, _lpToken, _startTimestamp, _rewardToken, _tokenPerSec)
}

// Initialize is a paid mutator transaction binding the contract method 0x731994fc.
//
// Solidity: function initialize(address _bribeFactory, address _masterWombat, address _lpToken, uint256 _startTimestamp, address _rewardToken, uint96 _tokenPerSec) returns()
func (_RewarderV3 *RewarderV3Session) Initialize(_bribeFactory common.Address, _masterWombat common.Address, _lpToken common.Address, _startTimestamp *big.Int, _rewardToken common.Address, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.Initialize(&_RewarderV3.TransactOpts, _bribeFactory, _masterWombat, _lpToken, _startTimestamp, _rewardToken, _tokenPerSec)
}

// Initialize is a paid mutator transaction binding the contract method 0x731994fc.
//
// Solidity: function initialize(address _bribeFactory, address _masterWombat, address _lpToken, uint256 _startTimestamp, address _rewardToken, uint96 _tokenPerSec) returns()
func (_RewarderV3 *RewarderV3TransactorSession) Initialize(_bribeFactory common.Address, _masterWombat common.Address, _lpToken common.Address, _startTimestamp *big.Int, _rewardToken common.Address, _tokenPerSec *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.Initialize(&_RewarderV3.TransactOpts, _bribeFactory, _masterWombat, _lpToken, _startTimestamp, _rewardToken, _tokenPerSec)
}

// OnReward is a paid mutator transaction binding the contract method 0x658ec4e4.
//
// Solidity: function onReward(address _user, uint256 _newLpAmount, uint256 _newFactor) returns(uint256[] rewards)
func (_RewarderV3 *RewarderV3Transactor) OnReward(opts *bind.TransactOpts, _user common.Address, _newLpAmount *big.Int, _newFactor *big.Int) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "onReward", _user, _newLpAmount, _newFactor)
}

// OnReward is a paid mutator transaction binding the contract method 0x658ec4e4.
//
// Solidity: function onReward(address _user, uint256 _newLpAmount, uint256 _newFactor) returns(uint256[] rewards)
func (_RewarderV3 *RewarderV3Session) OnReward(_user common.Address, _newLpAmount *big.Int, _newFactor *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.OnReward(&_RewarderV3.TransactOpts, _user, _newLpAmount, _newFactor)
}

// OnReward is a paid mutator transaction binding the contract method 0x658ec4e4.
//
// Solidity: function onReward(address _user, uint256 _newLpAmount, uint256 _newFactor) returns(uint256[] rewards)
func (_RewarderV3 *RewarderV3TransactorSession) OnReward(_user common.Address, _newLpAmount *big.Int, _newFactor *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.OnReward(&_RewarderV3.TransactOpts, _user, _newLpAmount, _newFactor)
}

// OnUpdateFactor is a paid mutator transaction binding the contract method 0xb024b437.
//
// Solidity: function onUpdateFactor(address _user, uint256 _newFactor) returns()
func (_RewarderV3 *RewarderV3Transactor) OnUpdateFactor(opts *bind.TransactOpts, _user common.Address, _newFactor *big.Int) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "onUpdateFactor", _user, _newFactor)
}

// OnUpdateFactor is a paid mutator transaction binding the contract method 0xb024b437.
//
// Solidity: function onUpdateFactor(address _user, uint256 _newFactor) returns()
func (_RewarderV3 *RewarderV3Session) OnUpdateFactor(_user common.Address, _newFactor *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.OnUpdateFactor(&_RewarderV3.TransactOpts, _user, _newFactor)
}

// OnUpdateFactor is a paid mutator transaction binding the contract method 0xb024b437.
//
// Solidity: function onUpdateFactor(address _user, uint256 _newFactor) returns()
func (_RewarderV3 *RewarderV3TransactorSession) OnUpdateFactor(_user common.Address, _newFactor *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.OnUpdateFactor(&_RewarderV3.TransactOpts, _user, _newFactor)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_RewarderV3 *RewarderV3Transactor) RemoveOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "removeOperator", _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_RewarderV3 *RewarderV3Session) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.RemoveOperator(&_RewarderV3.TransactOpts, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_RewarderV3 *RewarderV3TransactorSession) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.RemoveOperator(&_RewarderV3.TransactOpts, _operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewarderV3 *RewarderV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewarderV3 *RewarderV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _RewarderV3.Contract.RenounceOwnership(&_RewarderV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewarderV3 *RewarderV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewarderV3.Contract.RenounceOwnership(&_RewarderV3.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.RenounceRole(&_RewarderV3.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.RenounceRole(&_RewarderV3.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.RevokeRole(&_RewarderV3.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewarderV3 *RewarderV3TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.RevokeRole(&_RewarderV3.TransactOpts, role, account)
}

// SetIsDeprecated is a paid mutator transaction binding the contract method 0xccbd9b75.
//
// Solidity: function setIsDeprecated(bool _isDeprecated) returns()
func (_RewarderV3 *RewarderV3Transactor) SetIsDeprecated(opts *bind.TransactOpts, _isDeprecated bool) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "setIsDeprecated", _isDeprecated)
}

// SetIsDeprecated is a paid mutator transaction binding the contract method 0xccbd9b75.
//
// Solidity: function setIsDeprecated(bool _isDeprecated) returns()
func (_RewarderV3 *RewarderV3Session) SetIsDeprecated(_isDeprecated bool) (*types.Transaction, error) {
	return _RewarderV3.Contract.SetIsDeprecated(&_RewarderV3.TransactOpts, _isDeprecated)
}

// SetIsDeprecated is a paid mutator transaction binding the contract method 0xccbd9b75.
//
// Solidity: function setIsDeprecated(bool _isDeprecated) returns()
func (_RewarderV3 *RewarderV3TransactorSession) SetIsDeprecated(_isDeprecated bool) (*types.Transaction, error) {
	return _RewarderV3.Contract.SetIsDeprecated(&_RewarderV3.TransactOpts, _isDeprecated)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x2474e181.
//
// Solidity: function setRewardRate(uint256 _tokenId, uint96 _tokenPerSec, uint40 _startTimestampToOverride) returns()
func (_RewarderV3 *RewarderV3Transactor) SetRewardRate(opts *bind.TransactOpts, _tokenId *big.Int, _tokenPerSec *big.Int, _startTimestampToOverride *big.Int) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "setRewardRate", _tokenId, _tokenPerSec, _startTimestampToOverride)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x2474e181.
//
// Solidity: function setRewardRate(uint256 _tokenId, uint96 _tokenPerSec, uint40 _startTimestampToOverride) returns()
func (_RewarderV3 *RewarderV3Session) SetRewardRate(_tokenId *big.Int, _tokenPerSec *big.Int, _startTimestampToOverride *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.SetRewardRate(&_RewarderV3.TransactOpts, _tokenId, _tokenPerSec, _startTimestampToOverride)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x2474e181.
//
// Solidity: function setRewardRate(uint256 _tokenId, uint96 _tokenPerSec, uint40 _startTimestampToOverride) returns()
func (_RewarderV3 *RewarderV3TransactorSession) SetRewardRate(_tokenId *big.Int, _tokenPerSec *big.Int, _startTimestampToOverride *big.Int) (*types.Transaction, error) {
	return _RewarderV3.Contract.SetRewardRate(&_RewarderV3.TransactOpts, _tokenId, _tokenPerSec, _startTimestampToOverride)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewarderV3 *RewarderV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewarderV3 *RewarderV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.TransferOwnership(&_RewarderV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewarderV3 *RewarderV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewarderV3.Contract.TransferOwnership(&_RewarderV3.TransactOpts, newOwner)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_RewarderV3 *RewarderV3Transactor) UpdateReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.contract.Transact(opts, "updateReward")
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_RewarderV3 *RewarderV3Session) UpdateReward() (*types.Transaction, error) {
	return _RewarderV3.Contract.UpdateReward(&_RewarderV3.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0xf36c0a72.
//
// Solidity: function updateReward() returns()
func (_RewarderV3 *RewarderV3TransactorSession) UpdateReward() (*types.Transaction, error) {
	return _RewarderV3.Contract.UpdateReward(&_RewarderV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewarderV3 *RewarderV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewarderV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewarderV3 *RewarderV3Session) Receive() (*types.Transaction, error) {
	return _RewarderV3.Contract.Receive(&_RewarderV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewarderV3 *RewarderV3TransactorSession) Receive() (*types.Transaction, error) {
	return _RewarderV3.Contract.Receive(&_RewarderV3.TransactOpts)
}

// RewarderV3InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewarderV3 contract.
type RewarderV3InitializedIterator struct {
	Event *RewarderV3Initialized // Event containing the contract specifics and raw log

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
func (it *RewarderV3InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3Initialized)
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
		it.Event = new(RewarderV3Initialized)
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
func (it *RewarderV3InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3Initialized represents a Initialized event raised by the RewarderV3 contract.
type RewarderV3Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewarderV3 *RewarderV3Filterer) FilterInitialized(opts *bind.FilterOpts) (*RewarderV3InitializedIterator, error) {

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewarderV3InitializedIterator{contract: _RewarderV3.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewarderV3 *RewarderV3Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewarderV3Initialized) (event.Subscription, error) {

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3Initialized)
				if err := _RewarderV3.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RewarderV3 *RewarderV3Filterer) ParseInitialized(log types.Log) (*RewarderV3Initialized, error) {
	event := new(RewarderV3Initialized)
	if err := _RewarderV3.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3IsDeprecatedUpdatedIterator is returned from FilterIsDeprecatedUpdated and is used to iterate over the raw logs and unpacked data for IsDeprecatedUpdated events raised by the RewarderV3 contract.
type RewarderV3IsDeprecatedUpdatedIterator struct {
	Event *RewarderV3IsDeprecatedUpdated // Event containing the contract specifics and raw log

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
func (it *RewarderV3IsDeprecatedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3IsDeprecatedUpdated)
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
		it.Event = new(RewarderV3IsDeprecatedUpdated)
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
func (it *RewarderV3IsDeprecatedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3IsDeprecatedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3IsDeprecatedUpdated represents a IsDeprecatedUpdated event raised by the RewarderV3 contract.
type RewarderV3IsDeprecatedUpdated struct {
	IsDeprecated bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIsDeprecatedUpdated is a free log retrieval operation binding the contract event 0xdf1c2ccfd21e3329e1124daf53bdfee90ba7d708a557bf15e5c4eddd57fa24c9.
//
// Solidity: event IsDeprecatedUpdated(bool isDeprecated)
func (_RewarderV3 *RewarderV3Filterer) FilterIsDeprecatedUpdated(opts *bind.FilterOpts) (*RewarderV3IsDeprecatedUpdatedIterator, error) {

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "IsDeprecatedUpdated")
	if err != nil {
		return nil, err
	}
	return &RewarderV3IsDeprecatedUpdatedIterator{contract: _RewarderV3.contract, event: "IsDeprecatedUpdated", logs: logs, sub: sub}, nil
}

// WatchIsDeprecatedUpdated is a free log subscription operation binding the contract event 0xdf1c2ccfd21e3329e1124daf53bdfee90ba7d708a557bf15e5c4eddd57fa24c9.
//
// Solidity: event IsDeprecatedUpdated(bool isDeprecated)
func (_RewarderV3 *RewarderV3Filterer) WatchIsDeprecatedUpdated(opts *bind.WatchOpts, sink chan<- *RewarderV3IsDeprecatedUpdated) (event.Subscription, error) {

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "IsDeprecatedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3IsDeprecatedUpdated)
				if err := _RewarderV3.contract.UnpackLog(event, "IsDeprecatedUpdated", log); err != nil {
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

// ParseIsDeprecatedUpdated is a log parse operation binding the contract event 0xdf1c2ccfd21e3329e1124daf53bdfee90ba7d708a557bf15e5c4eddd57fa24c9.
//
// Solidity: event IsDeprecatedUpdated(bool isDeprecated)
func (_RewarderV3 *RewarderV3Filterer) ParseIsDeprecatedUpdated(log types.Log) (*RewarderV3IsDeprecatedUpdated, error) {
	event := new(RewarderV3IsDeprecatedUpdated)
	if err := _RewarderV3.contract.UnpackLog(event, "IsDeprecatedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3OnRewardIterator is returned from FilterOnReward and is used to iterate over the raw logs and unpacked data for OnReward events raised by the RewarderV3 contract.
type RewarderV3OnRewardIterator struct {
	Event *RewarderV3OnReward // Event containing the contract specifics and raw log

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
func (it *RewarderV3OnRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3OnReward)
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
		it.Event = new(RewarderV3OnReward)
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
func (it *RewarderV3OnRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3OnRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3OnReward represents a OnReward event raised by the RewarderV3 contract.
type RewarderV3OnReward struct {
	RewardToken common.Address
	User        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOnReward is a free log retrieval operation binding the contract event 0x986cbc32375de61d1fabfb01aef452f5c919f2180bb72fff0fb182126a02b527.
//
// Solidity: event OnReward(address indexed rewardToken, address indexed user, uint256 amount)
func (_RewarderV3 *RewarderV3Filterer) FilterOnReward(opts *bind.FilterOpts, rewardToken []common.Address, user []common.Address) (*RewarderV3OnRewardIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "OnReward", rewardTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3OnRewardIterator{contract: _RewarderV3.contract, event: "OnReward", logs: logs, sub: sub}, nil
}

// WatchOnReward is a free log subscription operation binding the contract event 0x986cbc32375de61d1fabfb01aef452f5c919f2180bb72fff0fb182126a02b527.
//
// Solidity: event OnReward(address indexed rewardToken, address indexed user, uint256 amount)
func (_RewarderV3 *RewarderV3Filterer) WatchOnReward(opts *bind.WatchOpts, sink chan<- *RewarderV3OnReward, rewardToken []common.Address, user []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "OnReward", rewardTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3OnReward)
				if err := _RewarderV3.contract.UnpackLog(event, "OnReward", log); err != nil {
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
func (_RewarderV3 *RewarderV3Filterer) ParseOnReward(log types.Log) (*RewarderV3OnReward, error) {
	event := new(RewarderV3OnReward)
	if err := _RewarderV3.contract.UnpackLog(event, "OnReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RewarderV3 contract.
type RewarderV3OwnershipTransferredIterator struct {
	Event *RewarderV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewarderV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3OwnershipTransferred)
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
		it.Event = new(RewarderV3OwnershipTransferred)
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
func (it *RewarderV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3OwnershipTransferred represents a OwnershipTransferred event raised by the RewarderV3 contract.
type RewarderV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewarderV3 *RewarderV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewarderV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3OwnershipTransferredIterator{contract: _RewarderV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewarderV3 *RewarderV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewarderV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3OwnershipTransferred)
				if err := _RewarderV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RewarderV3 *RewarderV3Filterer) ParseOwnershipTransferred(log types.Log) (*RewarderV3OwnershipTransferred, error) {
	event := new(RewarderV3OwnershipTransferred)
	if err := _RewarderV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3RewardRateUpdatedIterator is returned from FilterRewardRateUpdated and is used to iterate over the raw logs and unpacked data for RewardRateUpdated events raised by the RewarderV3 contract.
type RewarderV3RewardRateUpdatedIterator struct {
	Event *RewarderV3RewardRateUpdated // Event containing the contract specifics and raw log

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
func (it *RewarderV3RewardRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3RewardRateUpdated)
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
		it.Event = new(RewarderV3RewardRateUpdated)
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
func (it *RewarderV3RewardRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3RewardRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3RewardRateUpdated represents a RewardRateUpdated event raised by the RewarderV3 contract.
type RewarderV3RewardRateUpdated struct {
	RewardToken common.Address
	OldRate     *big.Int
	NewRate     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRateUpdated is a free log retrieval operation binding the contract event 0x225033f2ea5486463cbb49ceda2823be38daddc85031ce2c637e7ad0950bc85a.
//
// Solidity: event RewardRateUpdated(address indexed rewardToken, uint256 oldRate, uint256 newRate)
func (_RewarderV3 *RewarderV3Filterer) FilterRewardRateUpdated(opts *bind.FilterOpts, rewardToken []common.Address) (*RewarderV3RewardRateUpdatedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "RewardRateUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3RewardRateUpdatedIterator{contract: _RewarderV3.contract, event: "RewardRateUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardRateUpdated is a free log subscription operation binding the contract event 0x225033f2ea5486463cbb49ceda2823be38daddc85031ce2c637e7ad0950bc85a.
//
// Solidity: event RewardRateUpdated(address indexed rewardToken, uint256 oldRate, uint256 newRate)
func (_RewarderV3 *RewarderV3Filterer) WatchRewardRateUpdated(opts *bind.WatchOpts, sink chan<- *RewarderV3RewardRateUpdated, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "RewardRateUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3RewardRateUpdated)
				if err := _RewarderV3.contract.UnpackLog(event, "RewardRateUpdated", log); err != nil {
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
func (_RewarderV3 *RewarderV3Filterer) ParseRewardRateUpdated(log types.Log) (*RewarderV3RewardRateUpdated, error) {
	event := new(RewarderV3RewardRateUpdated)
	if err := _RewarderV3.contract.UnpackLog(event, "RewardRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RewarderV3 contract.
type RewarderV3RoleAdminChangedIterator struct {
	Event *RewarderV3RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RewarderV3RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3RoleAdminChanged)
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
		it.Event = new(RewarderV3RoleAdminChanged)
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
func (it *RewarderV3RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3RoleAdminChanged represents a RoleAdminChanged event raised by the RewarderV3 contract.
type RewarderV3RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewarderV3 *RewarderV3Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RewarderV3RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3RoleAdminChangedIterator{contract: _RewarderV3.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewarderV3 *RewarderV3Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RewarderV3RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3RoleAdminChanged)
				if err := _RewarderV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewarderV3 *RewarderV3Filterer) ParseRoleAdminChanged(log types.Log) (*RewarderV3RoleAdminChanged, error) {
	event := new(RewarderV3RoleAdminChanged)
	if err := _RewarderV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RewarderV3 contract.
type RewarderV3RoleGrantedIterator struct {
	Event *RewarderV3RoleGranted // Event containing the contract specifics and raw log

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
func (it *RewarderV3RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3RoleGranted)
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
		it.Event = new(RewarderV3RoleGranted)
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
func (it *RewarderV3RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3RoleGranted represents a RoleGranted event raised by the RewarderV3 contract.
type RewarderV3RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewarderV3 *RewarderV3Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewarderV3RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3RoleGrantedIterator{contract: _RewarderV3.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewarderV3 *RewarderV3Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RewarderV3RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3RoleGranted)
				if err := _RewarderV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewarderV3 *RewarderV3Filterer) ParseRoleGranted(log types.Log) (*RewarderV3RoleGranted, error) {
	event := new(RewarderV3RoleGranted)
	if err := _RewarderV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RewarderV3 contract.
type RewarderV3RoleRevokedIterator struct {
	Event *RewarderV3RoleRevoked // Event containing the contract specifics and raw log

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
func (it *RewarderV3RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3RoleRevoked)
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
		it.Event = new(RewarderV3RoleRevoked)
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
func (it *RewarderV3RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3RoleRevoked represents a RoleRevoked event raised by the RewarderV3 contract.
type RewarderV3RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewarderV3 *RewarderV3Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewarderV3RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3RoleRevokedIterator{contract: _RewarderV3.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewarderV3 *RewarderV3Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RewarderV3RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3RoleRevoked)
				if err := _RewarderV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewarderV3 *RewarderV3Filterer) ParseRoleRevoked(log types.Log) (*RewarderV3RoleRevoked, error) {
	event := new(RewarderV3RoleRevoked)
	if err := _RewarderV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderV3StartTimeUpdatedIterator is returned from FilterStartTimeUpdated and is used to iterate over the raw logs and unpacked data for StartTimeUpdated events raised by the RewarderV3 contract.
type RewarderV3StartTimeUpdatedIterator struct {
	Event *RewarderV3StartTimeUpdated // Event containing the contract specifics and raw log

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
func (it *RewarderV3StartTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderV3StartTimeUpdated)
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
		it.Event = new(RewarderV3StartTimeUpdated)
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
func (it *RewarderV3StartTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderV3StartTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderV3StartTimeUpdated represents a StartTimeUpdated event raised by the RewarderV3 contract.
type RewarderV3StartTimeUpdated struct {
	RewardToken  common.Address
	NewStartTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStartTimeUpdated is a free log retrieval operation binding the contract event 0x99e4a80c4e66b0b3fbf1bbe586cf7708be5e406b09778228cdb64b480226c389.
//
// Solidity: event StartTimeUpdated(address indexed rewardToken, uint40 newStartTime)
func (_RewarderV3 *RewarderV3Filterer) FilterStartTimeUpdated(opts *bind.FilterOpts, rewardToken []common.Address) (*RewarderV3StartTimeUpdatedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _RewarderV3.contract.FilterLogs(opts, "StartTimeUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &RewarderV3StartTimeUpdatedIterator{contract: _RewarderV3.contract, event: "StartTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchStartTimeUpdated is a free log subscription operation binding the contract event 0x99e4a80c4e66b0b3fbf1bbe586cf7708be5e406b09778228cdb64b480226c389.
//
// Solidity: event StartTimeUpdated(address indexed rewardToken, uint40 newStartTime)
func (_RewarderV3 *RewarderV3Filterer) WatchStartTimeUpdated(opts *bind.WatchOpts, sink chan<- *RewarderV3StartTimeUpdated, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _RewarderV3.contract.WatchLogs(opts, "StartTimeUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderV3StartTimeUpdated)
				if err := _RewarderV3.contract.UnpackLog(event, "StartTimeUpdated", log); err != nil {
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

// ParseStartTimeUpdated is a log parse operation binding the contract event 0x99e4a80c4e66b0b3fbf1bbe586cf7708be5e406b09778228cdb64b480226c389.
//
// Solidity: event StartTimeUpdated(address indexed rewardToken, uint40 newStartTime)
func (_RewarderV3 *RewarderV3Filterer) ParseStartTimeUpdated(log types.Log) (*RewarderV3StartTimeUpdated, error) {
	event := new(RewarderV3StartTimeUpdated)
	if err := _RewarderV3.contract.UnpackLog(event, "StartTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
