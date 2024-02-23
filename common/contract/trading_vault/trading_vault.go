// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trading_vault

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"AllowedToInteractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCurrentBalanceUSDT\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCurrentBalanceUSDT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DistributeRewardUSDT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"USDTAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultFeeUSDT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCurrentBalanceUSDT\",\"type\":\"uint256\"}],\"name\":\"ReceivedFromTrader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"esnar\",\"type\":\"address\"}],\"name\":\"RewardTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"RewardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCurrentBalanceUSDT\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pairInfo\",\"type\":\"address\"}],\"name\":\"SetPairInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalanceUSDT\",\"type\":\"uint256\"}],\"name\":\"ToClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"name\":\"UsdtSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vester\",\"type\":\"address\"}],\"name\":\"VesterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCurrentBalanceUSDT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"USDTToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Vester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"_lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"_unLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToInteract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"averageStakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBalanceUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_send\",\"type\":\"bool\"}],\"name\":\"distributeRewardUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getUserTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairInfo\",\"outputs\":[{\"internalType\":\"contractIPairInfos\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vaultFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_send\",\"type\":\"bool\"}],\"name\":\"receiveUSDTFromTrader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendUSDTToTrader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setAllowedToInteract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairInfo\",\"type\":\"address\"}],\"name\":\"setPairInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_esToken\",\"type\":\"address\"}],\"name\":\"setRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vester\",\"type\":\"address\"}],\"name\":\"setVester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawTimelock\",\"type\":\"uint256\"}],\"name\":\"setWithdrawTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardsDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInLockup\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawTimelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_Contract *ContractCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_Contract *ContractSession) USDT() (common.Address, error) {
	return _Contract.Contract.USDT(&_Contract.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_Contract *ContractCallerSession) USDT() (common.Address, error) {
	return _Contract.Contract.USDT(&_Contract.CallOpts)
}

// USDTToClaim is a free data retrieval call binding the contract method 0xeb21f625.
//
// Solidity: function USDTToClaim(address ) view returns(uint256)
func (_Contract *ContractCaller) USDTToClaim(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "USDTToClaim", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDTToClaim is a free data retrieval call binding the contract method 0xeb21f625.
//
// Solidity: function USDTToClaim(address ) view returns(uint256)
func (_Contract *ContractSession) USDTToClaim(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.USDTToClaim(&_Contract.CallOpts, arg0)
}

// USDTToClaim is a free data retrieval call binding the contract method 0xeb21f625.
//
// Solidity: function USDTToClaim(address ) view returns(uint256)
func (_Contract *ContractCallerSession) USDTToClaim(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.USDTToClaim(&_Contract.CallOpts, arg0)
}

// Vester is a free data retrieval call binding the contract method 0x32a92a12.
//
// Solidity: function Vester() view returns(address)
func (_Contract *ContractCaller) Vester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "Vester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vester is a free data retrieval call binding the contract method 0x32a92a12.
//
// Solidity: function Vester() view returns(address)
func (_Contract *ContractSession) Vester() (common.Address, error) {
	return _Contract.Contract.Vester(&_Contract.CallOpts)
}

// Vester is a free data retrieval call binding the contract method 0x32a92a12.
//
// Solidity: function Vester() view returns(address)
func (_Contract *ContractCallerSession) Vester() (common.Address, error) {
	return _Contract.Contract.Vester(&_Contract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contract *ContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, spender)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_Contract *ContractCaller) Allowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_Contract *ContractSession) Allowed(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Allowed(&_Contract.CallOpts, arg0)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_Contract *ContractCallerSession) Allowed(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Allowed(&_Contract.CallOpts, arg0)
}

// AllowedToInteract is a free data retrieval call binding the contract method 0x41200844.
//
// Solidity: function allowedToInteract(address ) view returns(bool)
func (_Contract *ContractCaller) AllowedToInteract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowedToInteract", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToInteract is a free data retrieval call binding the contract method 0x41200844.
//
// Solidity: function allowedToInteract(address ) view returns(bool)
func (_Contract *ContractSession) AllowedToInteract(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AllowedToInteract(&_Contract.CallOpts, arg0)
}

// AllowedToInteract is a free data retrieval call binding the contract method 0x41200844.
//
// Solidity: function allowedToInteract(address ) view returns(bool)
func (_Contract *ContractCallerSession) AllowedToInteract(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AllowedToInteract(&_Contract.CallOpts, arg0)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_Contract *ContractCaller) AverageStakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "averageStakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_Contract *ContractSession) AverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AverageStakedAmounts(&_Contract.CallOpts, arg0)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_Contract *ContractCallerSession) AverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AverageStakedAmounts(&_Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_Contract *ContractCaller) CumulativeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "cumulativeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_Contract *ContractSession) CumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.CumulativeRewards(&_Contract.CallOpts, arg0)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_Contract *ContractCallerSession) CumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.CumulativeRewards(&_Contract.CallOpts, arg0)
}

// CurrentBalanceUSDT is a free data retrieval call binding the contract method 0xde2133f9.
//
// Solidity: function currentBalanceUSDT() view returns(uint256)
func (_Contract *ContractCaller) CurrentBalanceUSDT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentBalanceUSDT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBalanceUSDT is a free data retrieval call binding the contract method 0xde2133f9.
//
// Solidity: function currentBalanceUSDT() view returns(uint256)
func (_Contract *ContractSession) CurrentBalanceUSDT() (*big.Int, error) {
	return _Contract.Contract.CurrentBalanceUSDT(&_Contract.CallOpts)
}

// CurrentBalanceUSDT is a free data retrieval call binding the contract method 0xde2133f9.
//
// Solidity: function currentBalanceUSDT() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentBalanceUSDT() (*big.Int, error) {
	return _Contract.Contract.CurrentBalanceUSDT(&_Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCallerSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Contract *ContractCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Contract *ContractSession) Earned(account common.Address) (*big.Int, error) {
	return _Contract.Contract.Earned(&_Contract.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Contract *ContractCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Contract.Contract.Earned(&_Contract.CallOpts, account)
}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_Contract *ContractCaller) EsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "esToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_Contract *ContractSession) EsToken() (common.Address, error) {
	return _Contract.Contract.EsToken(&_Contract.CallOpts)
}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_Contract *ContractCallerSession) EsToken() (common.Address, error) {
	return _Contract.Contract.EsToken(&_Contract.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Contract *ContractCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Contract *ContractSession) GetPricePerFullShare() (*big.Int, error) {
	return _Contract.Contract.GetPricePerFullShare(&_Contract.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Contract *ContractCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _Contract.Contract.GetPricePerFullShare(&_Contract.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Contract *ContractCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Contract *ContractSession) GetRewardForDuration() (*big.Int, error) {
	return _Contract.Contract.GetRewardForDuration(&_Contract.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Contract *ContractCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _Contract.Contract.GetRewardForDuration(&_Contract.CallOpts)
}

// GetUserTokenBalance is a free data retrieval call binding the contract method 0xf5aec88c.
//
// Solidity: function getUserTokenBalance(address _account) view returns(uint256)
func (_Contract *ContractCaller) GetUserTokenBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUserTokenBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTokenBalance is a free data retrieval call binding the contract method 0xf5aec88c.
//
// Solidity: function getUserTokenBalance(address _account) view returns(uint256)
func (_Contract *ContractSession) GetUserTokenBalance(_account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetUserTokenBalance(&_Contract.CallOpts, _account)
}

// GetUserTokenBalance is a free data retrieval call binding the contract method 0xf5aec88c.
//
// Solidity: function getUserTokenBalance(address _account) view returns(uint256)
func (_Contract *ContractCallerSession) GetUserTokenBalance(_account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetUserTokenBalance(&_Contract.CallOpts, _account)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Contract *ContractCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Contract *ContractSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Contract.Contract.LastTimeRewardApplicable(&_Contract.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Contract *ContractCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Contract.Contract.LastTimeRewardApplicable(&_Contract.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Contract *ContractCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Contract *ContractSession) LastUpdateTime() (*big.Int, error) {
	return _Contract.Contract.LastUpdateTime(&_Contract.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Contract *ContractCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Contract.Contract.LastUpdateTime(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
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

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Contract *ContractCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Contract *ContractSession) PeriodFinish() (*big.Int, error) {
	return _Contract.Contract.PeriodFinish(&_Contract.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Contract *ContractCallerSession) PeriodFinish() (*big.Int, error) {
	return _Contract.Contract.PeriodFinish(&_Contract.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Contract *ContractCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Contract *ContractSession) RewardPerToken() (*big.Int, error) {
	return _Contract.Contract.RewardPerToken(&_Contract.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Contract *ContractCallerSession) RewardPerToken() (*big.Int, error) {
	return _Contract.Contract.RewardPerToken(&_Contract.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Contract *ContractCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Contract *ContractSession) RewardPerTokenStored() (*big.Int, error) {
	return _Contract.Contract.RewardPerTokenStored(&_Contract.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Contract *ContractCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Contract.Contract.RewardPerTokenStored(&_Contract.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Contract *ContractCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Contract *ContractSession) RewardRate() (*big.Int, error) {
	return _Contract.Contract.RewardRate(&_Contract.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Contract *ContractCallerSession) RewardRate() (*big.Int, error) {
	return _Contract.Contract.RewardRate(&_Contract.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Contract *ContractCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Contract *ContractSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Rewards(&_Contract.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Rewards(&_Contract.CallOpts, arg0)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Contract *ContractCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Contract *ContractSession) RewardsDuration() (*big.Int, error) {
	return _Contract.Contract.RewardsDuration(&_Contract.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Contract *ContractCallerSession) RewardsDuration() (*big.Int, error) {
	return _Contract.Contract.RewardsDuration(&_Contract.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(uint256)
func (_Contract *ContractCaller) RewardsToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(uint256)
func (_Contract *ContractSession) RewardsToken() (*big.Int, error) {
	return _Contract.Contract.RewardsToken(&_Contract.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(uint256)
func (_Contract *ContractCallerSession) RewardsToken() (*big.Int, error) {
	return _Contract.Contract.RewardsToken(&_Contract.CallOpts)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_Contract *ContractCaller) StakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_Contract *ContractSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.StakedAmounts(&_Contract.CallOpts, arg0)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_Contract *ContractCallerSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.StakedAmounts(&_Contract.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TotalRewardsDistributed is a free data retrieval call binding the contract method 0xee172546.
//
// Solidity: function totalRewardsDistributed() view returns(uint256)
func (_Contract *ContractCaller) TotalRewardsDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalRewardsDistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsDistributed is a free data retrieval call binding the contract method 0xee172546.
//
// Solidity: function totalRewardsDistributed() view returns(uint256)
func (_Contract *ContractSession) TotalRewardsDistributed() (*big.Int, error) {
	return _Contract.Contract.TotalRewardsDistributed(&_Contract.CallOpts)
}

// TotalRewardsDistributed is a free data retrieval call binding the contract method 0xee172546.
//
// Solidity: function totalRewardsDistributed() view returns(uint256)
func (_Contract *ContractCallerSession) TotalRewardsDistributed() (*big.Int, error) {
	return _Contract.Contract.TotalRewardsDistributed(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Contract *ContractCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Contract *ContractSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserRewardPerTokenPaid(&_Contract.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Contract *ContractCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserRewardPerTokenPaid(&_Contract.CallOpts, arg0)
}

// UserTotalBalance is a free data retrieval call binding the contract method 0x821a3e67.
//
// Solidity: function userTotalBalance(address _user) view returns(uint256)
func (_Contract *ContractCaller) UserTotalBalance(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userTotalBalance", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTotalBalance is a free data retrieval call binding the contract method 0x821a3e67.
//
// Solidity: function userTotalBalance(address _user) view returns(uint256)
func (_Contract *ContractSession) UserTotalBalance(_user common.Address) (*big.Int, error) {
	return _Contract.Contract.UserTotalBalance(&_Contract.CallOpts, _user)
}

// UserTotalBalance is a free data retrieval call binding the contract method 0x821a3e67.
//
// Solidity: function userTotalBalance(address _user) view returns(uint256)
func (_Contract *ContractCallerSession) UserTotalBalance(_user common.Address) (*big.Int, error) {
	return _Contract.Contract.UserTotalBalance(&_Contract.CallOpts, _user)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 depositTimestamp, uint256 debtToken, uint256 amountInLockup)
func (_Contract *ContractCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	DepositTimestamp *big.Int
	DebtToken        *big.Int
	AmountInLockup   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		DepositTimestamp *big.Int
		DebtToken        *big.Int
		AmountInLockup   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DebtToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountInLockup = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 depositTimestamp, uint256 debtToken, uint256 amountInLockup)
func (_Contract *ContractSession) Users(arg0 common.Address) (struct {
	DepositTimestamp *big.Int
	DebtToken        *big.Int
	AmountInLockup   *big.Int
}, error) {
	return _Contract.Contract.Users(&_Contract.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 depositTimestamp, uint256 debtToken, uint256 amountInLockup)
func (_Contract *ContractCallerSession) Users(arg0 common.Address) (struct {
	DepositTimestamp *big.Int
	DebtToken        *big.Int
	AmountInLockup   *big.Int
}, error) {
	return _Contract.Contract.Users(&_Contract.CallOpts, arg0)
}

// WithdrawTimelock is a free data retrieval call binding the contract method 0x0661c6ce.
//
// Solidity: function withdrawTimelock() view returns(uint256)
func (_Contract *ContractCaller) WithdrawTimelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawTimelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawTimelock is a free data retrieval call binding the contract method 0x0661c6ce.
//
// Solidity: function withdrawTimelock() view returns(uint256)
func (_Contract *ContractSession) WithdrawTimelock() (*big.Int, error) {
	return _Contract.Contract.WithdrawTimelock(&_Contract.CallOpts)
}

// WithdrawTimelock is a free data retrieval call binding the contract method 0x0661c6ce.
//
// Solidity: function withdrawTimelock() view returns(uint256)
func (_Contract *ContractCallerSession) WithdrawTimelock() (*big.Int, error) {
	return _Contract.Contract.WithdrawTimelock(&_Contract.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x1f73668d.
//
// Solidity: function _lock(address _account, uint256 _amount) returns()
func (_Contract *ContractTransactor) Lock(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_lock", _account, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x1f73668d.
//
// Solidity: function _lock(address _account, uint256 _amount) returns()
func (_Contract *ContractSession) Lock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Lock(&_Contract.TransactOpts, _account, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x1f73668d.
//
// Solidity: function _lock(address _account, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) Lock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Lock(&_Contract.TransactOpts, _account, _amount)
}

// UnLock is a paid mutator transaction binding the contract method 0x95809c33.
//
// Solidity: function _unLock(address _account, uint256 _amount) returns()
func (_Contract *ContractTransactor) UnLock(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_unLock", _account, _amount)
}

// UnLock is a paid mutator transaction binding the contract method 0x95809c33.
//
// Solidity: function _unLock(address _account, uint256 _amount) returns()
func (_Contract *ContractSession) UnLock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnLock(&_Contract.TransactOpts, _account, _amount)
}

// UnLock is a paid mutator transaction binding the contract method 0x95809c33.
//
// Solidity: function _unLock(address _account, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) UnLock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnLock(&_Contract.TransactOpts, _account, _amount)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contract *ContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, amount)
}

// ClaimUSDT is a paid mutator transaction binding the contract method 0x7564ac38.
//
// Solidity: function claimUSDT() returns()
func (_Contract *ContractTransactor) ClaimUSDT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimUSDT")
}

// ClaimUSDT is a paid mutator transaction binding the contract method 0x7564ac38.
//
// Solidity: function claimUSDT() returns()
func (_Contract *ContractSession) ClaimUSDT() (*types.Transaction, error) {
	return _Contract.Contract.ClaimUSDT(&_Contract.TransactOpts)
}

// ClaimUSDT is a paid mutator transaction binding the contract method 0x7564ac38.
//
// Solidity: function claimUSDT() returns()
func (_Contract *ContractTransactorSession) ClaimUSDT() (*types.Transaction, error) {
	return _Contract.Contract.ClaimUSDT(&_Contract.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contract *ContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contract *ContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DecreaseAllowance(&_Contract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contract *ContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DecreaseAllowance(&_Contract.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _user) returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", _amount, _user)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _user) returns()
func (_Contract *ContractSession) Deposit(_amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, _amount, _user)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _user) returns()
func (_Contract *ContractTransactorSession) Deposit(_amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, _amount, _user)
}

// DepositAll is a paid mutator transaction binding the contract method 0x9f0d5f27.
//
// Solidity: function depositAll(address _user) returns()
func (_Contract *ContractTransactor) DepositAll(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositAll", _user)
}

// DepositAll is a paid mutator transaction binding the contract method 0x9f0d5f27.
//
// Solidity: function depositAll(address _user) returns()
func (_Contract *ContractSession) DepositAll(_user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DepositAll(&_Contract.TransactOpts, _user)
}

// DepositAll is a paid mutator transaction binding the contract method 0x9f0d5f27.
//
// Solidity: function depositAll(address _user) returns()
func (_Contract *ContractTransactorSession) DepositAll(_user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DepositAll(&_Contract.TransactOpts, _user)
}

// DistributeRewardUSDT is a paid mutator transaction binding the contract method 0xdd20b628.
//
// Solidity: function distributeRewardUSDT(uint256 _amount, bool _send) returns()
func (_Contract *ContractTransactor) DistributeRewardUSDT(opts *bind.TransactOpts, _amount *big.Int, _send bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "distributeRewardUSDT", _amount, _send)
}

// DistributeRewardUSDT is a paid mutator transaction binding the contract method 0xdd20b628.
//
// Solidity: function distributeRewardUSDT(uint256 _amount, bool _send) returns()
func (_Contract *ContractSession) DistributeRewardUSDT(_amount *big.Int, _send bool) (*types.Transaction, error) {
	return _Contract.Contract.DistributeRewardUSDT(&_Contract.TransactOpts, _amount, _send)
}

// DistributeRewardUSDT is a paid mutator transaction binding the contract method 0xdd20b628.
//
// Solidity: function distributeRewardUSDT(uint256 _amount, bool _send) returns()
func (_Contract *ContractTransactorSession) DistributeRewardUSDT(_amount *big.Int, _send bool) (*types.Transaction, error) {
	return _Contract.Contract.DistributeRewardUSDT(&_Contract.TransactOpts, _amount, _send)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address _user) returns()
func (_Contract *ContractTransactor) Harvest(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "harvest", _user)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address _user) returns()
func (_Contract *ContractSession) Harvest(_user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Harvest(&_Contract.TransactOpts, _user)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address _user) returns()
func (_Contract *ContractTransactorSession) Harvest(_user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Harvest(&_Contract.TransactOpts, _user)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contract *ContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contract *ContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contract *ContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, spender, addedValue)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Contract *ContractTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Contract *ContractSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.NotifyRewardAmount(&_Contract.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Contract *ContractTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.NotifyRewardAmount(&_Contract.TransactOpts, reward)
}

// ReceiveUSDTFromTrader is a paid mutator transaction binding the contract method 0x55981981.
//
// Solidity: function receiveUSDTFromTrader(address _trader, uint256 _amount, uint256 _vaultFee, bool _send) returns()
func (_Contract *ContractTransactor) ReceiveUSDTFromTrader(opts *bind.TransactOpts, _trader common.Address, _amount *big.Int, _vaultFee *big.Int, _send bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "receiveUSDTFromTrader", _trader, _amount, _vaultFee, _send)
}

// ReceiveUSDTFromTrader is a paid mutator transaction binding the contract method 0x55981981.
//
// Solidity: function receiveUSDTFromTrader(address _trader, uint256 _amount, uint256 _vaultFee, bool _send) returns()
func (_Contract *ContractSession) ReceiveUSDTFromTrader(_trader common.Address, _amount *big.Int, _vaultFee *big.Int, _send bool) (*types.Transaction, error) {
	return _Contract.Contract.ReceiveUSDTFromTrader(&_Contract.TransactOpts, _trader, _amount, _vaultFee, _send)
}

// ReceiveUSDTFromTrader is a paid mutator transaction binding the contract method 0x55981981.
//
// Solidity: function receiveUSDTFromTrader(address _trader, uint256 _amount, uint256 _vaultFee, bool _send) returns()
func (_Contract *ContractTransactorSession) ReceiveUSDTFromTrader(_trader common.Address, _amount *big.Int, _vaultFee *big.Int, _send bool) (*types.Transaction, error) {
	return _Contract.Contract.ReceiveUSDTFromTrader(&_Contract.TransactOpts, _trader, _amount, _vaultFee, _send)
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

// SendUSDTToTrader is a paid mutator transaction binding the contract method 0x5e60c9c6.
//
// Solidity: function sendUSDTToTrader(address _trader, uint256 _amount) returns()
func (_Contract *ContractTransactor) SendUSDTToTrader(opts *bind.TransactOpts, _trader common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendUSDTToTrader", _trader, _amount)
}

// SendUSDTToTrader is a paid mutator transaction binding the contract method 0x5e60c9c6.
//
// Solidity: function sendUSDTToTrader(address _trader, uint256 _amount) returns()
func (_Contract *ContractSession) SendUSDTToTrader(_trader common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SendUSDTToTrader(&_Contract.TransactOpts, _trader, _amount)
}

// SendUSDTToTrader is a paid mutator transaction binding the contract method 0x5e60c9c6.
//
// Solidity: function sendUSDTToTrader(address _trader, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) SendUSDTToTrader(_trader common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SendUSDTToTrader(&_Contract.TransactOpts, _trader, _amount)
}

// SetAllowed is a paid mutator transaction binding the contract method 0x4697f05d.
//
// Solidity: function setAllowed(address _sender, bool _status) returns()
func (_Contract *ContractTransactor) SetAllowed(opts *bind.TransactOpts, _sender common.Address, _status bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAllowed", _sender, _status)
}

// SetAllowed is a paid mutator transaction binding the contract method 0x4697f05d.
//
// Solidity: function setAllowed(address _sender, bool _status) returns()
func (_Contract *ContractSession) SetAllowed(_sender common.Address, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.SetAllowed(&_Contract.TransactOpts, _sender, _status)
}

// SetAllowed is a paid mutator transaction binding the contract method 0x4697f05d.
//
// Solidity: function setAllowed(address _sender, bool _status) returns()
func (_Contract *ContractTransactorSession) SetAllowed(_sender common.Address, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.SetAllowed(&_Contract.TransactOpts, _sender, _status)
}

// SetAllowedToInteract is a paid mutator transaction binding the contract method 0x9d1765f6.
//
// Solidity: function setAllowedToInteract(address _sender, bool _status) returns()
func (_Contract *ContractTransactor) SetAllowedToInteract(opts *bind.TransactOpts, _sender common.Address, _status bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAllowedToInteract", _sender, _status)
}

// SetAllowedToInteract is a paid mutator transaction binding the contract method 0x9d1765f6.
//
// Solidity: function setAllowedToInteract(address _sender, bool _status) returns()
func (_Contract *ContractSession) SetAllowedToInteract(_sender common.Address, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.SetAllowedToInteract(&_Contract.TransactOpts, _sender, _status)
}

// SetAllowedToInteract is a paid mutator transaction binding the contract method 0x9d1765f6.
//
// Solidity: function setAllowedToInteract(address _sender, bool _status) returns()
func (_Contract *ContractTransactorSession) SetAllowedToInteract(_sender common.Address, _status bool) (*types.Transaction, error) {
	return _Contract.Contract.SetAllowedToInteract(&_Contract.TransactOpts, _sender, _status)
}

// SetPairInfo is a paid mutator transaction binding the contract method 0xf49a9b8a.
//
// Solidity: function setPairInfo(address _pairInfo) returns()
func (_Contract *ContractTransactor) SetPairInfo(opts *bind.TransactOpts, _pairInfo common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPairInfo", _pairInfo)
}

// SetPairInfo is a paid mutator transaction binding the contract method 0xf49a9b8a.
//
// Solidity: function setPairInfo(address _pairInfo) returns()
func (_Contract *ContractSession) SetPairInfo(_pairInfo common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetPairInfo(&_Contract.TransactOpts, _pairInfo)
}

// SetPairInfo is a paid mutator transaction binding the contract method 0xf49a9b8a.
//
// Solidity: function setPairInfo(address _pairInfo) returns()
func (_Contract *ContractTransactorSession) SetPairInfo(_pairInfo common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetPairInfo(&_Contract.TransactOpts, _pairInfo)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _esToken) returns()
func (_Contract *ContractTransactor) SetRewardToken(opts *bind.TransactOpts, _esToken common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRewardToken", _esToken)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _esToken) returns()
func (_Contract *ContractSession) SetRewardToken(_esToken common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRewardToken(&_Contract.TransactOpts, _esToken)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _esToken) returns()
func (_Contract *ContractTransactorSession) SetRewardToken(_esToken common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRewardToken(&_Contract.TransactOpts, _esToken)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_Contract *ContractTransactor) SetRewardsDuration(opts *bind.TransactOpts, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRewardsDuration", _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_Contract *ContractSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetRewardsDuration(&_Contract.TransactOpts, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_Contract *ContractTransactorSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetRewardsDuration(&_Contract.TransactOpts, _rewardsDuration)
}

// SetVester is a paid mutator transaction binding the contract method 0x421c40b6.
//
// Solidity: function setVester(address _vester) returns()
func (_Contract *ContractTransactor) SetVester(opts *bind.TransactOpts, _vester common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setVester", _vester)
}

// SetVester is a paid mutator transaction binding the contract method 0x421c40b6.
//
// Solidity: function setVester(address _vester) returns()
func (_Contract *ContractSession) SetVester(_vester common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVester(&_Contract.TransactOpts, _vester)
}

// SetVester is a paid mutator transaction binding the contract method 0x421c40b6.
//
// Solidity: function setVester(address _vester) returns()
func (_Contract *ContractTransactorSession) SetVester(_vester common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVester(&_Contract.TransactOpts, _vester)
}

// SetWithdrawTimelock is a paid mutator transaction binding the contract method 0x1eefa740.
//
// Solidity: function setWithdrawTimelock(uint256 _withdrawTimelock) returns()
func (_Contract *ContractTransactor) SetWithdrawTimelock(opts *bind.TransactOpts, _withdrawTimelock *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setWithdrawTimelock", _withdrawTimelock)
}

// SetWithdrawTimelock is a paid mutator transaction binding the contract method 0x1eefa740.
//
// Solidity: function setWithdrawTimelock(uint256 _withdrawTimelock) returns()
func (_Contract *ContractSession) SetWithdrawTimelock(_withdrawTimelock *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetWithdrawTimelock(&_Contract.TransactOpts, _withdrawTimelock)
}

// SetWithdrawTimelock is a paid mutator transaction binding the contract method 0x1eefa740.
//
// Solidity: function setWithdrawTimelock(uint256 _withdrawTimelock) returns()
func (_Contract *ContractTransactorSession) SetWithdrawTimelock(_withdrawTimelock *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetWithdrawTimelock(&_Contract.TransactOpts, _withdrawTimelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contract *ContractSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, amount)
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Contract *ContractSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Contract *ContractTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Contract *ContractTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Contract *ContractSession) WithdrawAll() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAll(&_Contract.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Contract *ContractTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAll(&_Contract.TransactOpts)
}

// ContractAllowedToInteractSetIterator is returned from FilterAllowedToInteractSet and is used to iterate over the raw logs and unpacked data for AllowedToInteractSet events raised by the Contract contract.
type ContractAllowedToInteractSetIterator struct {
	Event *ContractAllowedToInteractSet // Event containing the contract specifics and raw log

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
func (it *ContractAllowedToInteractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllowedToInteractSet)
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
		it.Event = new(ContractAllowedToInteractSet)
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
func (it *ContractAllowedToInteractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllowedToInteractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllowedToInteractSet represents a AllowedToInteractSet event raised by the Contract contract.
type ContractAllowedToInteractSet struct {
	Sender common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllowedToInteractSet is a free log retrieval operation binding the contract event 0x77711358779674597f238fdb43e9f28618347bc0e217ae42c982656c0b7f6be8.
//
// Solidity: event AllowedToInteractSet(address indexed sender, bool status)
func (_Contract *ContractFilterer) FilterAllowedToInteractSet(opts *bind.FilterOpts, sender []common.Address) (*ContractAllowedToInteractSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AllowedToInteractSet", senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllowedToInteractSetIterator{contract: _Contract.contract, event: "AllowedToInteractSet", logs: logs, sub: sub}, nil
}

// WatchAllowedToInteractSet is a free log subscription operation binding the contract event 0x77711358779674597f238fdb43e9f28618347bc0e217ae42c982656c0b7f6be8.
//
// Solidity: event AllowedToInteractSet(address indexed sender, bool status)
func (_Contract *ContractFilterer) WatchAllowedToInteractSet(opts *bind.WatchOpts, sink chan<- *ContractAllowedToInteractSet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AllowedToInteractSet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllowedToInteractSet)
				if err := _Contract.contract.UnpackLog(event, "AllowedToInteractSet", log); err != nil {
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
func (_Contract *ContractFilterer) ParseAllowedToInteractSet(log types.Log) (*ContractAllowedToInteractSet, error) {
	event := new(ContractAllowedToInteractSet)
	if err := _Contract.contract.UnpackLog(event, "AllowedToInteractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Contract contract.
type ContractClaimedIterator struct {
	Event *ContractClaimed // Event containing the contract specifics and raw log

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
func (it *ContractClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimed)
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
		it.Event = new(ContractClaimed)
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
func (it *ContractClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimed represents a Claimed event raised by the Contract contract.
type ContractClaimed struct {
	Trader                common.Address
	Amount                *big.Int
	NewCurrentBalanceUSDT *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address trader, uint256 amount, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) FilterClaimed(opts *bind.FilterOpts) (*ContractClaimedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &ContractClaimedIterator{contract: _Contract.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address trader, uint256 amount, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ContractClaimed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimed)
				if err := _Contract.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address trader, uint256 amount, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) ParseClaimed(log types.Log) (*ContractClaimed, error) {
	event := new(ContractClaimed)
	if err := _Contract.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Contract contract.
type ContractDepositedIterator struct {
	Event *ContractDeposited // Event containing the contract specifics and raw log

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
func (it *ContractDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeposited)
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
		it.Event = new(ContractDeposited)
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
func (it *ContractDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeposited represents a Deposited event raised by the Contract contract.
type ContractDeposited struct {
	Caller                common.Address
	Amount                *big.Int
	NewCurrentBalanceUSDT *big.Int
	Shares                *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3.
//
// Solidity: event Deposited(address caller, uint256 amount, uint256 newCurrentBalanceUSDT, uint256 shares)
func (_Contract *ContractFilterer) FilterDeposited(opts *bind.FilterOpts) (*ContractDepositedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &ContractDepositedIterator{contract: _Contract.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3.
//
// Solidity: event Deposited(address caller, uint256 amount, uint256 newCurrentBalanceUSDT, uint256 shares)
func (_Contract *ContractFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ContractDeposited) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeposited)
				if err := _Contract.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3.
//
// Solidity: event Deposited(address caller, uint256 amount, uint256 newCurrentBalanceUSDT, uint256 shares)
func (_Contract *ContractFilterer) ParseDeposited(log types.Log) (*ContractDeposited, error) {
	event := new(ContractDeposited)
	if err := _Contract.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDistributeRewardUSDTIterator is returned from FilterDistributeRewardUSDT and is used to iterate over the raw logs and unpacked data for DistributeRewardUSDT events raised by the Contract contract.
type ContractDistributeRewardUSDTIterator struct {
	Event *ContractDistributeRewardUSDT // Event containing the contract specifics and raw log

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
func (it *ContractDistributeRewardUSDTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDistributeRewardUSDT)
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
		it.Event = new(ContractDistributeRewardUSDT)
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
func (it *ContractDistributeRewardUSDTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDistributeRewardUSDTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDistributeRewardUSDT represents a DistributeRewardUSDT event raised by the Contract contract.
type ContractDistributeRewardUSDT struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewardUSDT is a free log retrieval operation binding the contract event 0xf80cd5baafb1b5a4ff6e2e1fe7a2b2aeab8da41da29572905f978e6f9ad8e338.
//
// Solidity: event DistributeRewardUSDT(uint256 _amount)
func (_Contract *ContractFilterer) FilterDistributeRewardUSDT(opts *bind.FilterOpts) (*ContractDistributeRewardUSDTIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DistributeRewardUSDT")
	if err != nil {
		return nil, err
	}
	return &ContractDistributeRewardUSDTIterator{contract: _Contract.contract, event: "DistributeRewardUSDT", logs: logs, sub: sub}, nil
}

// WatchDistributeRewardUSDT is a free log subscription operation binding the contract event 0xf80cd5baafb1b5a4ff6e2e1fe7a2b2aeab8da41da29572905f978e6f9ad8e338.
//
// Solidity: event DistributeRewardUSDT(uint256 _amount)
func (_Contract *ContractFilterer) WatchDistributeRewardUSDT(opts *bind.WatchOpts, sink chan<- *ContractDistributeRewardUSDT) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DistributeRewardUSDT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDistributeRewardUSDT)
				if err := _Contract.contract.UnpackLog(event, "DistributeRewardUSDT", log); err != nil {
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

// ParseDistributeRewardUSDT is a log parse operation binding the contract event 0xf80cd5baafb1b5a4ff6e2e1fe7a2b2aeab8da41da29572905f978e6f9ad8e338.
//
// Solidity: event DistributeRewardUSDT(uint256 _amount)
func (_Contract *ContractFilterer) ParseDistributeRewardUSDT(log types.Log) (*ContractDistributeRewardUSDT, error) {
	event := new(ContractDistributeRewardUSDT)
	if err := _Contract.contract.UnpackLog(event, "DistributeRewardUSDT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNumberUpdatedIterator is returned from FilterNumberUpdated and is used to iterate over the raw logs and unpacked data for NumberUpdated events raised by the Contract contract.
type ContractNumberUpdatedIterator struct {
	Event *ContractNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNumberUpdated)
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
		it.Event = new(ContractNumberUpdated)
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
func (it *ContractNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNumberUpdated represents a NumberUpdated event raised by the Contract contract.
type ContractNumberUpdated struct {
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNumberUpdated is a free log retrieval operation binding the contract event 0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab.
//
// Solidity: event NumberUpdated(string name, uint256 value)
func (_Contract *ContractFilterer) FilterNumberUpdated(opts *bind.FilterOpts) (*ContractNumberUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NumberUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractNumberUpdatedIterator{contract: _Contract.contract, event: "NumberUpdated", logs: logs, sub: sub}, nil
}

// WatchNumberUpdated is a free log subscription operation binding the contract event 0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab.
//
// Solidity: event NumberUpdated(string name, uint256 value)
func (_Contract *ContractFilterer) WatchNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNumberUpdated)
				if err := _Contract.contract.UnpackLog(event, "NumberUpdated", log); err != nil {
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

// ParseNumberUpdated is a log parse operation binding the contract event 0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab.
//
// Solidity: event NumberUpdated(string name, uint256 value)
func (_Contract *ContractFilterer) ParseNumberUpdated(log types.Log) (*ContractNumberUpdated, error) {
	event := new(ContractNumberUpdated)
	if err := _Contract.contract.UnpackLog(event, "NumberUpdated", log); err != nil {
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

// ContractReceivedFromTraderIterator is returned from FilterReceivedFromTrader and is used to iterate over the raw logs and unpacked data for ReceivedFromTrader events raised by the Contract contract.
type ContractReceivedFromTraderIterator struct {
	Event *ContractReceivedFromTrader // Event containing the contract specifics and raw log

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
func (it *ContractReceivedFromTraderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReceivedFromTrader)
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
		it.Event = new(ContractReceivedFromTrader)
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
func (it *ContractReceivedFromTraderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReceivedFromTraderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReceivedFromTrader represents a ReceivedFromTrader event raised by the Contract contract.
type ContractReceivedFromTrader struct {
	Caller                common.Address
	Trader                common.Address
	USDTAmount            *big.Int
	VaultFeeUSDT          *big.Int
	NewCurrentBalanceUSDT *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterReceivedFromTrader is a free log retrieval operation binding the contract event 0x6e22aa31121e8218696ecd7dc44c424a92f3d6f62814106f18b9311159fefcb2.
//
// Solidity: event ReceivedFromTrader(address caller, address trader, uint256 USDTAmount, uint256 vaultFeeUSDT, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) FilterReceivedFromTrader(opts *bind.FilterOpts) (*ContractReceivedFromTraderIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ReceivedFromTrader")
	if err != nil {
		return nil, err
	}
	return &ContractReceivedFromTraderIterator{contract: _Contract.contract, event: "ReceivedFromTrader", logs: logs, sub: sub}, nil
}

// WatchReceivedFromTrader is a free log subscription operation binding the contract event 0x6e22aa31121e8218696ecd7dc44c424a92f3d6f62814106f18b9311159fefcb2.
//
// Solidity: event ReceivedFromTrader(address caller, address trader, uint256 USDTAmount, uint256 vaultFeeUSDT, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) WatchReceivedFromTrader(opts *bind.WatchOpts, sink chan<- *ContractReceivedFromTrader) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ReceivedFromTrader")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReceivedFromTrader)
				if err := _Contract.contract.UnpackLog(event, "ReceivedFromTrader", log); err != nil {
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

// ParseReceivedFromTrader is a log parse operation binding the contract event 0x6e22aa31121e8218696ecd7dc44c424a92f3d6f62814106f18b9311159fefcb2.
//
// Solidity: event ReceivedFromTrader(address caller, address trader, uint256 USDTAmount, uint256 vaultFeeUSDT, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) ParseReceivedFromTrader(log types.Log) (*ContractReceivedFromTrader, error) {
	event := new(ContractReceivedFromTrader)
	if err := _Contract.contract.UnpackLog(event, "ReceivedFromTrader", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardTokenSetIterator is returned from FilterRewardTokenSet and is used to iterate over the raw logs and unpacked data for RewardTokenSet events raised by the Contract contract.
type ContractRewardTokenSetIterator struct {
	Event *ContractRewardTokenSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardTokenSet)
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
		it.Event = new(ContractRewardTokenSet)
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
func (it *ContractRewardTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardTokenSet represents a RewardTokenSet event raised by the Contract contract.
type ContractRewardTokenSet struct {
	Esnar common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardTokenSet is a free log retrieval operation binding the contract event 0x2d6b04df9b7d358407d1a014f1114b064add34c19d63d395db155a7e533e967a.
//
// Solidity: event RewardTokenSet(address indexed esnar)
func (_Contract *ContractFilterer) FilterRewardTokenSet(opts *bind.FilterOpts, esnar []common.Address) (*ContractRewardTokenSetIterator, error) {

	var esnarRule []interface{}
	for _, esnarItem := range esnar {
		esnarRule = append(esnarRule, esnarItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardTokenSet", esnarRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardTokenSetIterator{contract: _Contract.contract, event: "RewardTokenSet", logs: logs, sub: sub}, nil
}

// WatchRewardTokenSet is a free log subscription operation binding the contract event 0x2d6b04df9b7d358407d1a014f1114b064add34c19d63d395db155a7e533e967a.
//
// Solidity: event RewardTokenSet(address indexed esnar)
func (_Contract *ContractFilterer) WatchRewardTokenSet(opts *bind.WatchOpts, sink chan<- *ContractRewardTokenSet, esnar []common.Address) (event.Subscription, error) {

	var esnarRule []interface{}
	for _, esnarItem := range esnar {
		esnarRule = append(esnarRule, esnarItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardTokenSet", esnarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardTokenSet)
				if err := _Contract.contract.UnpackLog(event, "RewardTokenSet", log); err != nil {
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

// ParseRewardTokenSet is a log parse operation binding the contract event 0x2d6b04df9b7d358407d1a014f1114b064add34c19d63d395db155a7e533e967a.
//
// Solidity: event RewardTokenSet(address indexed esnar)
func (_Contract *ContractFilterer) ParseRewardTokenSet(log types.Log) (*ContractRewardTokenSet, error) {
	event := new(ContractRewardTokenSet)
	if err := _Contract.contract.UnpackLog(event, "RewardTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsDurationSetIterator is returned from FilterRewardsDurationSet and is used to iterate over the raw logs and unpacked data for RewardsDurationSet events raised by the Contract contract.
type ContractRewardsDurationSetIterator struct {
	Event *ContractRewardsDurationSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsDurationSet)
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
		it.Event = new(ContractRewardsDurationSet)
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
func (it *ContractRewardsDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsDurationSet represents a RewardsDurationSet event raised by the Contract contract.
type ContractRewardsDurationSet struct {
	RewardsDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationSet is a free log retrieval operation binding the contract event 0x47744a1ac202df7b343bfda5f2866a65a670c3817fd5b93f61d76ba1fdde953c.
//
// Solidity: event RewardsDurationSet(uint256 rewardsDuration)
func (_Contract *ContractFilterer) FilterRewardsDurationSet(opts *bind.FilterOpts) (*ContractRewardsDurationSetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardsDurationSet")
	if err != nil {
		return nil, err
	}
	return &ContractRewardsDurationSetIterator{contract: _Contract.contract, event: "RewardsDurationSet", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationSet is a free log subscription operation binding the contract event 0x47744a1ac202df7b343bfda5f2866a65a670c3817fd5b93f61d76ba1fdde953c.
//
// Solidity: event RewardsDurationSet(uint256 rewardsDuration)
func (_Contract *ContractFilterer) WatchRewardsDurationSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsDurationSet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardsDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsDurationSet)
				if err := _Contract.contract.UnpackLog(event, "RewardsDurationSet", log); err != nil {
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

// ParseRewardsDurationSet is a log parse operation binding the contract event 0x47744a1ac202df7b343bfda5f2866a65a670c3817fd5b93f61d76ba1fdde953c.
//
// Solidity: event RewardsDurationSet(uint256 rewardsDuration)
func (_Contract *ContractFilterer) ParseRewardsDurationSet(log types.Log) (*ContractRewardsDurationSet, error) {
	event := new(ContractRewardsDurationSet)
	if err := _Contract.contract.UnpackLog(event, "RewardsDurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsSetIterator is returned from FilterRewardsSet and is used to iterate over the raw logs and unpacked data for RewardsSet events raised by the Contract contract.
type ContractRewardsSetIterator struct {
	Event *ContractRewardsSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsSet)
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
		it.Event = new(ContractRewardsSet)
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
func (it *ContractRewardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsSet represents a RewardsSet event raised by the Contract contract.
type ContractRewardsSet struct {
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsSet is a free log retrieval operation binding the contract event 0x3a9c73865b863f6b3112a48ba71b6cc75b1b5440abddbcba0135c64a4056508b.
//
// Solidity: event RewardsSet(uint256 rewards)
func (_Contract *ContractFilterer) FilterRewardsSet(opts *bind.FilterOpts) (*ContractRewardsSetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardsSet")
	if err != nil {
		return nil, err
	}
	return &ContractRewardsSetIterator{contract: _Contract.contract, event: "RewardsSet", logs: logs, sub: sub}, nil
}

// WatchRewardsSet is a free log subscription operation binding the contract event 0x3a9c73865b863f6b3112a48ba71b6cc75b1b5440abddbcba0135c64a4056508b.
//
// Solidity: event RewardsSet(uint256 rewards)
func (_Contract *ContractFilterer) WatchRewardsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsSet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsSet)
				if err := _Contract.contract.UnpackLog(event, "RewardsSet", log); err != nil {
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

// ParseRewardsSet is a log parse operation binding the contract event 0x3a9c73865b863f6b3112a48ba71b6cc75b1b5440abddbcba0135c64a4056508b.
//
// Solidity: event RewardsSet(uint256 rewards)
func (_Contract *ContractFilterer) ParseRewardsSet(log types.Log) (*ContractRewardsSet, error) {
	event := new(ContractRewardsSet)
	if err := _Contract.contract.UnpackLog(event, "RewardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Contract contract.
type ContractSentIterator struct {
	Event *ContractSent // Event containing the contract specifics and raw log

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
func (it *ContractSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSent)
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
		it.Event = new(ContractSent)
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
func (it *ContractSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSent represents a Sent event raised by the Contract contract.
type ContractSent struct {
	Caller                common.Address
	Trader                common.Address
	Amount                *big.Int
	NewCurrentBalanceUSDT *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x34355b4c5dff25f21b90975d65f648edf2c50bea228323bb74333bfe5f015f3c.
//
// Solidity: event Sent(address caller, address trader, uint256 amount, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) FilterSent(opts *bind.FilterOpts) (*ContractSentIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &ContractSentIterator{contract: _Contract.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x34355b4c5dff25f21b90975d65f648edf2c50bea228323bb74333bfe5f015f3c.
//
// Solidity: event Sent(address caller, address trader, uint256 amount, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *ContractSent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSent)
				if err := _Contract.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x34355b4c5dff25f21b90975d65f648edf2c50bea228323bb74333bfe5f015f3c.
//
// Solidity: event Sent(address caller, address trader, uint256 amount, uint256 newCurrentBalanceUSDT)
func (_Contract *ContractFilterer) ParseSent(log types.Log) (*ContractSent, error) {
	event := new(ContractSent)
	if err := _Contract.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetPairInfoIterator is returned from FilterSetPairInfo and is used to iterate over the raw logs and unpacked data for SetPairInfo events raised by the Contract contract.
type ContractSetPairInfoIterator struct {
	Event *ContractSetPairInfo // Event containing the contract specifics and raw log

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
func (it *ContractSetPairInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetPairInfo)
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
		it.Event = new(ContractSetPairInfo)
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
func (it *ContractSetPairInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetPairInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetPairInfo represents a SetPairInfo event raised by the Contract contract.
type ContractSetPairInfo struct {
	PairInfo common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPairInfo is a free log retrieval operation binding the contract event 0x498f61784682d2a34a04afb8c518b7d38e0e59907bd66601a0a41ed6fb79ff9e.
//
// Solidity: event SetPairInfo(address _pairInfo)
func (_Contract *ContractFilterer) FilterSetPairInfo(opts *bind.FilterOpts) (*ContractSetPairInfoIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetPairInfo")
	if err != nil {
		return nil, err
	}
	return &ContractSetPairInfoIterator{contract: _Contract.contract, event: "SetPairInfo", logs: logs, sub: sub}, nil
}

// WatchSetPairInfo is a free log subscription operation binding the contract event 0x498f61784682d2a34a04afb8c518b7d38e0e59907bd66601a0a41ed6fb79ff9e.
//
// Solidity: event SetPairInfo(address _pairInfo)
func (_Contract *ContractFilterer) WatchSetPairInfo(opts *bind.WatchOpts, sink chan<- *ContractSetPairInfo) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetPairInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetPairInfo)
				if err := _Contract.contract.UnpackLog(event, "SetPairInfo", log); err != nil {
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
func (_Contract *ContractFilterer) ParseSetPairInfo(log types.Log) (*ContractSetPairInfo, error) {
	event := new(ContractSetPairInfo)
	if err := _Contract.contract.UnpackLog(event, "SetPairInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractToClaimIterator is returned from FilterToClaim and is used to iterate over the raw logs and unpacked data for ToClaim events raised by the Contract contract.
type ContractToClaimIterator struct {
	Event *ContractToClaim // Event containing the contract specifics and raw log

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
func (it *ContractToClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractToClaim)
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
		it.Event = new(ContractToClaim)
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
func (it *ContractToClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractToClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractToClaim represents a ToClaim event raised by the Contract contract.
type ContractToClaim struct {
	Caller             common.Address
	Trader             common.Address
	Amount             *big.Int
	CurrentBalanceUSDT *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterToClaim is a free log retrieval operation binding the contract event 0x5084298b6a649fa60b059555adae31e4f916067c54a48c56b76a106266f9136f.
//
// Solidity: event ToClaim(address caller, address trader, uint256 amount, uint256 currentBalanceUSDT)
func (_Contract *ContractFilterer) FilterToClaim(opts *bind.FilterOpts) (*ContractToClaimIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ToClaim")
	if err != nil {
		return nil, err
	}
	return &ContractToClaimIterator{contract: _Contract.contract, event: "ToClaim", logs: logs, sub: sub}, nil
}

// WatchToClaim is a free log subscription operation binding the contract event 0x5084298b6a649fa60b059555adae31e4f916067c54a48c56b76a106266f9136f.
//
// Solidity: event ToClaim(address caller, address trader, uint256 amount, uint256 currentBalanceUSDT)
func (_Contract *ContractFilterer) WatchToClaim(opts *bind.WatchOpts, sink chan<- *ContractToClaim) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ToClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractToClaim)
				if err := _Contract.contract.UnpackLog(event, "ToClaim", log); err != nil {
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

// ParseToClaim is a log parse operation binding the contract event 0x5084298b6a649fa60b059555adae31e4f916067c54a48c56b76a106266f9136f.
//
// Solidity: event ToClaim(address caller, address trader, uint256 amount, uint256 currentBalanceUSDT)
func (_Contract *ContractFilterer) ParseToClaim(log types.Log) (*ContractToClaim, error) {
	event := new(ContractToClaim)
	if err := _Contract.contract.UnpackLog(event, "ToClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUsdtSetIterator is returned from FilterUsdtSet and is used to iterate over the raw logs and unpacked data for UsdtSet events raised by the Contract contract.
type ContractUsdtSetIterator struct {
	Event *ContractUsdtSet // Event containing the contract specifics and raw log

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
func (it *ContractUsdtSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUsdtSet)
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
		it.Event = new(ContractUsdtSet)
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
func (it *ContractUsdtSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUsdtSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUsdtSet represents a UsdtSet event raised by the Contract contract.
type ContractUsdtSet struct {
	Usdt common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUsdtSet is a free log retrieval operation binding the contract event 0xa44f293dfa9228916345a6016220f304fd4e10c2f25ef62c896b4946926a70f4.
//
// Solidity: event UsdtSet(address indexed usdt)
func (_Contract *ContractFilterer) FilterUsdtSet(opts *bind.FilterOpts, usdt []common.Address) (*ContractUsdtSetIterator, error) {

	var usdtRule []interface{}
	for _, usdtItem := range usdt {
		usdtRule = append(usdtRule, usdtItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UsdtSet", usdtRule)
	if err != nil {
		return nil, err
	}
	return &ContractUsdtSetIterator{contract: _Contract.contract, event: "UsdtSet", logs: logs, sub: sub}, nil
}

// WatchUsdtSet is a free log subscription operation binding the contract event 0xa44f293dfa9228916345a6016220f304fd4e10c2f25ef62c896b4946926a70f4.
//
// Solidity: event UsdtSet(address indexed usdt)
func (_Contract *ContractFilterer) WatchUsdtSet(opts *bind.WatchOpts, sink chan<- *ContractUsdtSet, usdt []common.Address) (event.Subscription, error) {

	var usdtRule []interface{}
	for _, usdtItem := range usdt {
		usdtRule = append(usdtRule, usdtItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UsdtSet", usdtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUsdtSet)
				if err := _Contract.contract.UnpackLog(event, "UsdtSet", log); err != nil {
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

// ParseUsdtSet is a log parse operation binding the contract event 0xa44f293dfa9228916345a6016220f304fd4e10c2f25ef62c896b4946926a70f4.
//
// Solidity: event UsdtSet(address indexed usdt)
func (_Contract *ContractFilterer) ParseUsdtSet(log types.Log) (*ContractUsdtSet, error) {
	event := new(ContractUsdtSet)
	if err := _Contract.contract.UnpackLog(event, "UsdtSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVesterSetIterator is returned from FilterVesterSet and is used to iterate over the raw logs and unpacked data for VesterSet events raised by the Contract contract.
type ContractVesterSetIterator struct {
	Event *ContractVesterSet // Event containing the contract specifics and raw log

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
func (it *ContractVesterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVesterSet)
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
		it.Event = new(ContractVesterSet)
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
func (it *ContractVesterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVesterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVesterSet represents a VesterSet event raised by the Contract contract.
type ContractVesterSet struct {
	Vester common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVesterSet is a free log retrieval operation binding the contract event 0xf6a1c111d2ee5305a0fc35f0a12282991fa815e20e87da570dcb0de2ee73ed40.
//
// Solidity: event VesterSet(address indexed vester)
func (_Contract *ContractFilterer) FilterVesterSet(opts *bind.FilterOpts, vester []common.Address) (*ContractVesterSetIterator, error) {

	var vesterRule []interface{}
	for _, vesterItem := range vester {
		vesterRule = append(vesterRule, vesterItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VesterSet", vesterRule)
	if err != nil {
		return nil, err
	}
	return &ContractVesterSetIterator{contract: _Contract.contract, event: "VesterSet", logs: logs, sub: sub}, nil
}

// WatchVesterSet is a free log subscription operation binding the contract event 0xf6a1c111d2ee5305a0fc35f0a12282991fa815e20e87da570dcb0de2ee73ed40.
//
// Solidity: event VesterSet(address indexed vester)
func (_Contract *ContractFilterer) WatchVesterSet(opts *bind.WatchOpts, sink chan<- *ContractVesterSet, vester []common.Address) (event.Subscription, error) {

	var vesterRule []interface{}
	for _, vesterItem := range vester {
		vesterRule = append(vesterRule, vesterItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VesterSet", vesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVesterSet)
				if err := _Contract.contract.UnpackLog(event, "VesterSet", log); err != nil {
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

// ParseVesterSet is a log parse operation binding the contract event 0xf6a1c111d2ee5305a0fc35f0a12282991fa815e20e87da570dcb0de2ee73ed40.
//
// Solidity: event VesterSet(address indexed vester)
func (_Contract *ContractFilterer) ParseVesterSet(log types.Log) (*ContractVesterSet, error) {
	event := new(ContractVesterSet)
	if err := _Contract.contract.UnpackLog(event, "VesterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contract contract.
type ContractWithdrawnIterator struct {
	Event *ContractWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawn)
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
		it.Event = new(ContractWithdrawn)
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
func (it *ContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdrawn represents a Withdrawn event raised by the Contract contract.
type ContractWithdrawn struct {
	Caller                common.Address
	Amount                *big.Int
	NewCurrentBalanceUSDT *big.Int
	Shares                *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address caller, uint256 amount, uint256 newCurrentBalanceUSDT, uint256 shares)
func (_Contract *ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*ContractWithdrawnIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnIterator{contract: _Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address caller, uint256 amount, uint256 newCurrentBalanceUSDT, uint256 shares)
func (_Contract *ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address caller, uint256 amount, uint256 newCurrentBalanceUSDT, uint256 shares)
func (_Contract *ContractFilterer) ParseWithdrawn(log types.Log) (*ContractWithdrawn, error) {
	event := new(ContractWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
