// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault_factory

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

// ICopyTradingVaultVaultInfo is an auto generated low-level Go binding around an user-defined struct.
type ICopyTradingVaultVaultInfo struct {
	StartTime              *big.Int
	ReadinessTime          *big.Int
	VaultDuration          *big.Int
	MinStartTradingAmount  *big.Int
	MaxVaultCapAmount      *big.Int
	MinUserDepositAmount   *big.Int
	MaxUserDepositAmount   *big.Int
	PerformanceFee         *big.Int
	VaultManagementFee     *big.Int
	MinManagerCommitAmount *big.Int
	MinDurationTime        *big.Int
}

// ICopyTradingVaultVaultSocialInfo is an auto generated low-level Go binding around an user-defined struct.
type ICopyTradingVaultVaultSocialInfo struct {
	VaultName     string
	VaultTwitter  string
	VaultPhoto    string
	VaultStrategy string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"CreateVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExtractVaultFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetToken\",\"type\":\"address\"}],\"name\":\"NewAssetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDepositAmountPeruser\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDepositAmountPeruser\",\"type\":\"uint256\"}],\"name\":\"NewDepositAmountPeruser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"tag\",\"type\":\"bool\"}],\"name\":\"NewManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"NewMaxPerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"}],\"name\":\"NewMaxPositionSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"managerCommitAmount\",\"type\":\"uint256\"}],\"name\":\"NewMinManagerCommitAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minReadinessTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxReadinessTime\",\"type\":\"uint256\"}],\"name\":\"NewReadinessTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradingAddress\",\"type\":\"address\"}],\"name\":\"NewTradingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"NewVaultAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultCap\",\"type\":\"uint256\"}],\"name\":\"NewVaultCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDuration\",\"type\":\"uint256\"}],\"name\":\"NewVaultDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultFee\",\"type\":\"uint256\"}],\"name\":\"NewVaultFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_twitter\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_photo\",\"type\":\"string\"}],\"name\":\"NewVaultInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_strategy\",\"type\":\"string\"}],\"name\":\"NewVaultStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTime\",\"type\":\"uint256\"}],\"name\":\"NewVaultduration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"tag\",\"type\":\"bool\"}],\"name\":\"NewWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_tags\",\"type\":\"bool[]\"}],\"name\":\"batchManagerWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"closeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_readinessTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStartTradingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxVaultCapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUserDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUserDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_performanceFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_vaultManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"vaultName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultTwitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultPhoto\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultStrategy\",\"type\":\"string\"}],\"internalType\":\"structICopyTradingVault.VaultSocialInfo\",\"name\":\"_vaultSocialInfo\",\"type\":\"tuple\"}],\"name\":\"createVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"expireCancelOpen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"expireCloseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_closeMargin\",\"type\":\"uint256\"}],\"name\":\"expireRequestClose\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"extractVaultFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_feedId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_priceRollover\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tradingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"managerList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_tag\",\"type\":\"bool\"}],\"name\":\"managerWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDepositAmountPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPerformanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPositionSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReadinessTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVaultDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDepositAmountPeruser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minManagerCommitAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReadinessTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minUserDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVaultDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_twitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_photo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetVaultInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_strategy\",\"type\":\"string\"}],\"name\":\"resetVaultStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetToken\",\"type\":\"address\"}],\"name\":\"setAssetToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDepositAmountPeruser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDepositAmountPeruser\",\"type\":\"uint256\"}],\"name\":\"setDepositAmountPeruser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_tag\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_performanceFee\",\"type\":\"uint256\"}],\"name\":\"setMaxPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_positionSize\",\"type\":\"uint256\"}],\"name\":\"setMaxPositionSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_managerCommitAmount\",\"type\":\"uint256\"}],\"name\":\"setMinManagerCommitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDuration\",\"type\":\"uint256\"}],\"name\":\"setMinVaultDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minReadinessTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxReadinessTime\",\"type\":\"uint256\"}],\"name\":\"setReadinessTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tradingAddress\",\"type\":\"address\"}],\"name\":\"setTradingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultCap\",\"type\":\"uint256\"}],\"name\":\"setVaultCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultFee\",\"type\":\"uint256\"}],\"name\":\"setVaultFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTime\",\"type\":\"uint256\"}],\"name\":\"setVaultduration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userVaultInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"readinessTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStartTradingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVaultCapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUserDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUserDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultManagementFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minManagerCommitAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDurationTime\",\"type\":\"uint256\"}],\"internalType\":\"structICopyTradingVault.VaultInfo\",\"name\":\"vaultInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"vaultName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultTwitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultPhoto\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultStrategy\",\"type\":\"string\"}],\"internalType\":\"structICopyTradingVault.VaultSocialInfo\",\"name\":\"socialInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"copyVaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userVaultNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultCloneAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManagementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AssetToken is a free data retrieval call binding the contract method 0x1083f761.
//
// Solidity: function assetToken() view returns(address)
func (_Contract *ContractCaller) AssetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "assetToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetToken is a free data retrieval call binding the contract method 0x1083f761.
//
// Solidity: function assetToken() view returns(address)
func (_Contract *ContractSession) AssetToken() (common.Address, error) {
	return _Contract.Contract.AssetToken(&_Contract.CallOpts)
}

// AssetToken is a free data retrieval call binding the contract method 0x1083f761.
//
// Solidity: function assetToken() view returns(address)
func (_Contract *ContractCallerSession) AssetToken() (common.Address, error) {
	return _Contract.Contract.AssetToken(&_Contract.CallOpts)
}

// ManagerList is a free data retrieval call binding the contract method 0xa59be4c7.
//
// Solidity: function managerList(address ) view returns(bool)
func (_Contract *ContractCaller) ManagerList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "managerList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ManagerList is a free data retrieval call binding the contract method 0xa59be4c7.
//
// Solidity: function managerList(address ) view returns(bool)
func (_Contract *ContractSession) ManagerList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.ManagerList(&_Contract.CallOpts, arg0)
}

// ManagerList is a free data retrieval call binding the contract method 0xa59be4c7.
//
// Solidity: function managerList(address ) view returns(bool)
func (_Contract *ContractCallerSession) ManagerList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.ManagerList(&_Contract.CallOpts, arg0)
}

// MaxDepositAmountPerUser is a free data retrieval call binding the contract method 0x7381ba45.
//
// Solidity: function maxDepositAmountPerUser() view returns(uint256)
func (_Contract *ContractCaller) MaxDepositAmountPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxDepositAmountPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDepositAmountPerUser is a free data retrieval call binding the contract method 0x7381ba45.
//
// Solidity: function maxDepositAmountPerUser() view returns(uint256)
func (_Contract *ContractSession) MaxDepositAmountPerUser() (*big.Int, error) {
	return _Contract.Contract.MaxDepositAmountPerUser(&_Contract.CallOpts)
}

// MaxDepositAmountPerUser is a free data retrieval call binding the contract method 0x7381ba45.
//
// Solidity: function maxDepositAmountPerUser() view returns(uint256)
func (_Contract *ContractCallerSession) MaxDepositAmountPerUser() (*big.Int, error) {
	return _Contract.Contract.MaxDepositAmountPerUser(&_Contract.CallOpts)
}

// MaxPerformanceFee is a free data retrieval call binding the contract method 0x5265ed61.
//
// Solidity: function maxPerformanceFee() view returns(uint256)
func (_Contract *ContractCaller) MaxPerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxPerformanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerformanceFee is a free data retrieval call binding the contract method 0x5265ed61.
//
// Solidity: function maxPerformanceFee() view returns(uint256)
func (_Contract *ContractSession) MaxPerformanceFee() (*big.Int, error) {
	return _Contract.Contract.MaxPerformanceFee(&_Contract.CallOpts)
}

// MaxPerformanceFee is a free data retrieval call binding the contract method 0x5265ed61.
//
// Solidity: function maxPerformanceFee() view returns(uint256)
func (_Contract *ContractCallerSession) MaxPerformanceFee() (*big.Int, error) {
	return _Contract.Contract.MaxPerformanceFee(&_Contract.CallOpts)
}

// MaxPositionSize is a free data retrieval call binding the contract method 0xa329dc57.
//
// Solidity: function maxPositionSize() view returns(uint256)
func (_Contract *ContractCaller) MaxPositionSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxPositionSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPositionSize is a free data retrieval call binding the contract method 0xa329dc57.
//
// Solidity: function maxPositionSize() view returns(uint256)
func (_Contract *ContractSession) MaxPositionSize() (*big.Int, error) {
	return _Contract.Contract.MaxPositionSize(&_Contract.CallOpts)
}

// MaxPositionSize is a free data retrieval call binding the contract method 0xa329dc57.
//
// Solidity: function maxPositionSize() view returns(uint256)
func (_Contract *ContractCallerSession) MaxPositionSize() (*big.Int, error) {
	return _Contract.Contract.MaxPositionSize(&_Contract.CallOpts)
}

// MaxReadinessTime is a free data retrieval call binding the contract method 0x834a4d2d.
//
// Solidity: function maxReadinessTime() view returns(uint256)
func (_Contract *ContractCaller) MaxReadinessTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxReadinessTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReadinessTime is a free data retrieval call binding the contract method 0x834a4d2d.
//
// Solidity: function maxReadinessTime() view returns(uint256)
func (_Contract *ContractSession) MaxReadinessTime() (*big.Int, error) {
	return _Contract.Contract.MaxReadinessTime(&_Contract.CallOpts)
}

// MaxReadinessTime is a free data retrieval call binding the contract method 0x834a4d2d.
//
// Solidity: function maxReadinessTime() view returns(uint256)
func (_Contract *ContractCallerSession) MaxReadinessTime() (*big.Int, error) {
	return _Contract.Contract.MaxReadinessTime(&_Contract.CallOpts)
}

// MaxVaultDuration is a free data retrieval call binding the contract method 0x1bac488e.
//
// Solidity: function maxVaultDuration() view returns(uint256)
func (_Contract *ContractCaller) MaxVaultDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxVaultDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVaultDuration is a free data retrieval call binding the contract method 0x1bac488e.
//
// Solidity: function maxVaultDuration() view returns(uint256)
func (_Contract *ContractSession) MaxVaultDuration() (*big.Int, error) {
	return _Contract.Contract.MaxVaultDuration(&_Contract.CallOpts)
}

// MaxVaultDuration is a free data retrieval call binding the contract method 0x1bac488e.
//
// Solidity: function maxVaultDuration() view returns(uint256)
func (_Contract *ContractCallerSession) MaxVaultDuration() (*big.Int, error) {
	return _Contract.Contract.MaxVaultDuration(&_Contract.CallOpts)
}

// MinDepositAmountPeruser is a free data retrieval call binding the contract method 0xcdb0f928.
//
// Solidity: function minDepositAmountPeruser() view returns(uint256)
func (_Contract *ContractCaller) MinDepositAmountPeruser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDepositAmountPeruser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDepositAmountPeruser is a free data retrieval call binding the contract method 0xcdb0f928.
//
// Solidity: function minDepositAmountPeruser() view returns(uint256)
func (_Contract *ContractSession) MinDepositAmountPeruser() (*big.Int, error) {
	return _Contract.Contract.MinDepositAmountPeruser(&_Contract.CallOpts)
}

// MinDepositAmountPeruser is a free data retrieval call binding the contract method 0xcdb0f928.
//
// Solidity: function minDepositAmountPeruser() view returns(uint256)
func (_Contract *ContractCallerSession) MinDepositAmountPeruser() (*big.Int, error) {
	return _Contract.Contract.MinDepositAmountPeruser(&_Contract.CallOpts)
}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Contract *ContractCaller) MinDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Contract *ContractSession) MinDuration() (*big.Int, error) {
	return _Contract.Contract.MinDuration(&_Contract.CallOpts)
}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Contract *ContractCallerSession) MinDuration() (*big.Int, error) {
	return _Contract.Contract.MinDuration(&_Contract.CallOpts)
}

// MinManagerCommitAmount is a free data retrieval call binding the contract method 0xa0f5ed85.
//
// Solidity: function minManagerCommitAmount() view returns(uint256)
func (_Contract *ContractCaller) MinManagerCommitAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minManagerCommitAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinManagerCommitAmount is a free data retrieval call binding the contract method 0xa0f5ed85.
//
// Solidity: function minManagerCommitAmount() view returns(uint256)
func (_Contract *ContractSession) MinManagerCommitAmount() (*big.Int, error) {
	return _Contract.Contract.MinManagerCommitAmount(&_Contract.CallOpts)
}

// MinManagerCommitAmount is a free data retrieval call binding the contract method 0xa0f5ed85.
//
// Solidity: function minManagerCommitAmount() view returns(uint256)
func (_Contract *ContractCallerSession) MinManagerCommitAmount() (*big.Int, error) {
	return _Contract.Contract.MinManagerCommitAmount(&_Contract.CallOpts)
}

// MinReadinessTime is a free data retrieval call binding the contract method 0xc49a6d66.
//
// Solidity: function minReadinessTime() view returns(uint256)
func (_Contract *ContractCaller) MinReadinessTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minReadinessTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReadinessTime is a free data retrieval call binding the contract method 0xc49a6d66.
//
// Solidity: function minReadinessTime() view returns(uint256)
func (_Contract *ContractSession) MinReadinessTime() (*big.Int, error) {
	return _Contract.Contract.MinReadinessTime(&_Contract.CallOpts)
}

// MinReadinessTime is a free data retrieval call binding the contract method 0xc49a6d66.
//
// Solidity: function minReadinessTime() view returns(uint256)
func (_Contract *ContractCallerSession) MinReadinessTime() (*big.Int, error) {
	return _Contract.Contract.MinReadinessTime(&_Contract.CallOpts)
}

// MinUserDeposit is a free data retrieval call binding the contract method 0xd5a2ab53.
//
// Solidity: function minUserDeposit() view returns(uint256)
func (_Contract *ContractCaller) MinUserDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minUserDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinUserDeposit is a free data retrieval call binding the contract method 0xd5a2ab53.
//
// Solidity: function minUserDeposit() view returns(uint256)
func (_Contract *ContractSession) MinUserDeposit() (*big.Int, error) {
	return _Contract.Contract.MinUserDeposit(&_Contract.CallOpts)
}

// MinUserDeposit is a free data retrieval call binding the contract method 0xd5a2ab53.
//
// Solidity: function minUserDeposit() view returns(uint256)
func (_Contract *ContractCallerSession) MinUserDeposit() (*big.Int, error) {
	return _Contract.Contract.MinUserDeposit(&_Contract.CallOpts)
}

// MinVaultDuration is a free data retrieval call binding the contract method 0x00b54b55.
//
// Solidity: function minVaultDuration() view returns(uint256)
func (_Contract *ContractCaller) MinVaultDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minVaultDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVaultDuration is a free data retrieval call binding the contract method 0x00b54b55.
//
// Solidity: function minVaultDuration() view returns(uint256)
func (_Contract *ContractSession) MinVaultDuration() (*big.Int, error) {
	return _Contract.Contract.MinVaultDuration(&_Contract.CallOpts)
}

// MinVaultDuration is a free data retrieval call binding the contract method 0x00b54b55.
//
// Solidity: function minVaultDuration() view returns(uint256)
func (_Contract *ContractCallerSession) MinVaultDuration() (*big.Int, error) {
	return _Contract.Contract.MinVaultDuration(&_Contract.CallOpts)
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

// TradingAddress is a free data retrieval call binding the contract method 0x1db787e8.
//
// Solidity: function tradingAddress() view returns(address)
func (_Contract *ContractCaller) TradingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tradingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradingAddress is a free data retrieval call binding the contract method 0x1db787e8.
//
// Solidity: function tradingAddress() view returns(address)
func (_Contract *ContractSession) TradingAddress() (common.Address, error) {
	return _Contract.Contract.TradingAddress(&_Contract.CallOpts)
}

// TradingAddress is a free data retrieval call binding the contract method 0x1db787e8.
//
// Solidity: function tradingAddress() view returns(address)
func (_Contract *ContractCallerSession) TradingAddress() (common.Address, error) {
	return _Contract.Contract.TradingAddress(&_Contract.CallOpts)
}

// UserVaultInfo is a free data retrieval call binding the contract method 0x42538a1f.
//
// Solidity: function userVaultInfo(address , uint256 ) view returns(address vaultAddress, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) vaultInfo, (string,string,string,string) socialInfo, address copyVaultAddress)
func (_Contract *ContractCaller) UserVaultInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	VaultAddress     common.Address
	VaultInfo        ICopyTradingVaultVaultInfo
	SocialInfo       ICopyTradingVaultVaultSocialInfo
	CopyVaultAddress common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userVaultInfo", arg0, arg1)

	outstruct := new(struct {
		VaultAddress     common.Address
		VaultInfo        ICopyTradingVaultVaultInfo
		SocialInfo       ICopyTradingVaultVaultSocialInfo
		CopyVaultAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VaultAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VaultInfo = *abi.ConvertType(out[1], new(ICopyTradingVaultVaultInfo)).(*ICopyTradingVaultVaultInfo)
	outstruct.SocialInfo = *abi.ConvertType(out[2], new(ICopyTradingVaultVaultSocialInfo)).(*ICopyTradingVaultVaultSocialInfo)
	outstruct.CopyVaultAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// UserVaultInfo is a free data retrieval call binding the contract method 0x42538a1f.
//
// Solidity: function userVaultInfo(address , uint256 ) view returns(address vaultAddress, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) vaultInfo, (string,string,string,string) socialInfo, address copyVaultAddress)
func (_Contract *ContractSession) UserVaultInfo(arg0 common.Address, arg1 *big.Int) (struct {
	VaultAddress     common.Address
	VaultInfo        ICopyTradingVaultVaultInfo
	SocialInfo       ICopyTradingVaultVaultSocialInfo
	CopyVaultAddress common.Address
}, error) {
	return _Contract.Contract.UserVaultInfo(&_Contract.CallOpts, arg0, arg1)
}

// UserVaultInfo is a free data retrieval call binding the contract method 0x42538a1f.
//
// Solidity: function userVaultInfo(address , uint256 ) view returns(address vaultAddress, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) vaultInfo, (string,string,string,string) socialInfo, address copyVaultAddress)
func (_Contract *ContractCallerSession) UserVaultInfo(arg0 common.Address, arg1 *big.Int) (struct {
	VaultAddress     common.Address
	VaultInfo        ICopyTradingVaultVaultInfo
	SocialInfo       ICopyTradingVaultVaultSocialInfo
	CopyVaultAddress common.Address
}, error) {
	return _Contract.Contract.UserVaultInfo(&_Contract.CallOpts, arg0, arg1)
}

// UserVaultNumber is a free data retrieval call binding the contract method 0x72c0380c.
//
// Solidity: function userVaultNumber(address ) view returns(uint256)
func (_Contract *ContractCaller) UserVaultNumber(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userVaultNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserVaultNumber is a free data retrieval call binding the contract method 0x72c0380c.
//
// Solidity: function userVaultNumber(address ) view returns(uint256)
func (_Contract *ContractSession) UserVaultNumber(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserVaultNumber(&_Contract.CallOpts, arg0)
}

// UserVaultNumber is a free data retrieval call binding the contract method 0x72c0380c.
//
// Solidity: function userVaultNumber(address ) view returns(uint256)
func (_Contract *ContractCallerSession) UserVaultNumber(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserVaultNumber(&_Contract.CallOpts, arg0)
}

// VaultCap is a free data retrieval call binding the contract method 0x3c1bda09.
//
// Solidity: function vaultCap() view returns(uint256)
func (_Contract *ContractCaller) VaultCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCap is a free data retrieval call binding the contract method 0x3c1bda09.
//
// Solidity: function vaultCap() view returns(uint256)
func (_Contract *ContractSession) VaultCap() (*big.Int, error) {
	return _Contract.Contract.VaultCap(&_Contract.CallOpts)
}

// VaultCap is a free data retrieval call binding the contract method 0x3c1bda09.
//
// Solidity: function vaultCap() view returns(uint256)
func (_Contract *ContractCallerSession) VaultCap() (*big.Int, error) {
	return _Contract.Contract.VaultCap(&_Contract.CallOpts)
}

// VaultCloneAddress is a free data retrieval call binding the contract method 0x86de517b.
//
// Solidity: function vaultCloneAddress() view returns(address)
func (_Contract *ContractCaller) VaultCloneAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultCloneAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultCloneAddress is a free data retrieval call binding the contract method 0x86de517b.
//
// Solidity: function vaultCloneAddress() view returns(address)
func (_Contract *ContractSession) VaultCloneAddress() (common.Address, error) {
	return _Contract.Contract.VaultCloneAddress(&_Contract.CallOpts)
}

// VaultCloneAddress is a free data retrieval call binding the contract method 0x86de517b.
//
// Solidity: function vaultCloneAddress() view returns(address)
func (_Contract *ContractCallerSession) VaultCloneAddress() (common.Address, error) {
	return _Contract.Contract.VaultCloneAddress(&_Contract.CallOpts)
}

// VaultManagementFee is a free data retrieval call binding the contract method 0x8b4fd385.
//
// Solidity: function vaultManagementFee() view returns(uint256)
func (_Contract *ContractCaller) VaultManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultManagementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultManagementFee is a free data retrieval call binding the contract method 0x8b4fd385.
//
// Solidity: function vaultManagementFee() view returns(uint256)
func (_Contract *ContractSession) VaultManagementFee() (*big.Int, error) {
	return _Contract.Contract.VaultManagementFee(&_Contract.CallOpts)
}

// VaultManagementFee is a free data retrieval call binding the contract method 0x8b4fd385.
//
// Solidity: function vaultManagementFee() view returns(uint256)
func (_Contract *ContractCallerSession) VaultManagementFee() (*big.Int, error) {
	return _Contract.Contract.VaultManagementFee(&_Contract.CallOpts)
}

// VaultWhiteList is a free data retrieval call binding the contract method 0xe1b09b76.
//
// Solidity: function vaultWhiteList(address ) view returns(bool)
func (_Contract *ContractCaller) VaultWhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultWhiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultWhiteList is a free data retrieval call binding the contract method 0xe1b09b76.
//
// Solidity: function vaultWhiteList(address ) view returns(bool)
func (_Contract *ContractSession) VaultWhiteList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.VaultWhiteList(&_Contract.CallOpts, arg0)
}

// VaultWhiteList is a free data retrieval call binding the contract method 0xe1b09b76.
//
// Solidity: function vaultWhiteList(address ) view returns(bool)
func (_Contract *ContractCallerSession) VaultWhiteList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.VaultWhiteList(&_Contract.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Contract *ContractCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Contract *ContractSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.WhiteList(&_Contract.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_Contract *ContractCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.WhiteList(&_Contract.CallOpts, arg0)
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

// BatchManagerWhiteList is a paid mutator transaction binding the contract method 0xb3df3a8c.
//
// Solidity: function batchManagerWhiteList(address[] _users, bool[] _tags) returns()
func (_Contract *ContractTransactor) BatchManagerWhiteList(opts *bind.TransactOpts, _users []common.Address, _tags []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batchManagerWhiteList", _users, _tags)
}

// BatchManagerWhiteList is a paid mutator transaction binding the contract method 0xb3df3a8c.
//
// Solidity: function batchManagerWhiteList(address[] _users, bool[] _tags) returns()
func (_Contract *ContractSession) BatchManagerWhiteList(_users []common.Address, _tags []bool) (*types.Transaction, error) {
	return _Contract.Contract.BatchManagerWhiteList(&_Contract.TransactOpts, _users, _tags)
}

// BatchManagerWhiteList is a paid mutator transaction binding the contract method 0xb3df3a8c.
//
// Solidity: function batchManagerWhiteList(address[] _users, bool[] _tags) returns()
func (_Contract *ContractTransactorSession) BatchManagerWhiteList(_users []common.Address, _tags []bool) (*types.Transaction, error) {
	return _Contract.Contract.BatchManagerWhiteList(&_Contract.TransactOpts, _users, _tags)
}

// CloseVault is a paid mutator transaction binding the contract method 0xc99cb2b7.
//
// Solidity: function closeVault(address _vault) returns()
func (_Contract *ContractTransactor) CloseVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "closeVault", _vault)
}

// CloseVault is a paid mutator transaction binding the contract method 0xc99cb2b7.
//
// Solidity: function closeVault(address _vault) returns()
func (_Contract *ContractSession) CloseVault(_vault common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CloseVault(&_Contract.TransactOpts, _vault)
}

// CloseVault is a paid mutator transaction binding the contract method 0xc99cb2b7.
//
// Solidity: function closeVault(address _vault) returns()
func (_Contract *ContractTransactorSession) CloseVault(_vault common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CloseVault(&_Contract.TransactOpts, _vault)
}

// CreateVault is a paid mutator transaction binding the contract method 0x8f3c1ae6.
//
// Solidity: function createVault(string _name, string _symbol, uint256 _readinessTime, uint256 _duration, uint256 _minStartTradingAmount, uint256 _maxVaultCapAmount, uint256 _minUserDepositAmount, uint256 _maxUserDepositAmount, uint256 _performanceFee, address _vaultManager, (string,string,string,string) _vaultSocialInfo) returns()
func (_Contract *ContractTransactor) CreateVault(opts *bind.TransactOpts, _name string, _symbol string, _readinessTime *big.Int, _duration *big.Int, _minStartTradingAmount *big.Int, _maxVaultCapAmount *big.Int, _minUserDepositAmount *big.Int, _maxUserDepositAmount *big.Int, _performanceFee *big.Int, _vaultManager common.Address, _vaultSocialInfo ICopyTradingVaultVaultSocialInfo) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createVault", _name, _symbol, _readinessTime, _duration, _minStartTradingAmount, _maxVaultCapAmount, _minUserDepositAmount, _maxUserDepositAmount, _performanceFee, _vaultManager, _vaultSocialInfo)
}

// CreateVault is a paid mutator transaction binding the contract method 0x8f3c1ae6.
//
// Solidity: function createVault(string _name, string _symbol, uint256 _readinessTime, uint256 _duration, uint256 _minStartTradingAmount, uint256 _maxVaultCapAmount, uint256 _minUserDepositAmount, uint256 _maxUserDepositAmount, uint256 _performanceFee, address _vaultManager, (string,string,string,string) _vaultSocialInfo) returns()
func (_Contract *ContractSession) CreateVault(_name string, _symbol string, _readinessTime *big.Int, _duration *big.Int, _minStartTradingAmount *big.Int, _maxVaultCapAmount *big.Int, _minUserDepositAmount *big.Int, _maxUserDepositAmount *big.Int, _performanceFee *big.Int, _vaultManager common.Address, _vaultSocialInfo ICopyTradingVaultVaultSocialInfo) (*types.Transaction, error) {
	return _Contract.Contract.CreateVault(&_Contract.TransactOpts, _name, _symbol, _readinessTime, _duration, _minStartTradingAmount, _maxVaultCapAmount, _minUserDepositAmount, _maxUserDepositAmount, _performanceFee, _vaultManager, _vaultSocialInfo)
}

// CreateVault is a paid mutator transaction binding the contract method 0x8f3c1ae6.
//
// Solidity: function createVault(string _name, string _symbol, uint256 _readinessTime, uint256 _duration, uint256 _minStartTradingAmount, uint256 _maxVaultCapAmount, uint256 _minUserDepositAmount, uint256 _maxUserDepositAmount, uint256 _performanceFee, address _vaultManager, (string,string,string,string) _vaultSocialInfo) returns()
func (_Contract *ContractTransactorSession) CreateVault(_name string, _symbol string, _readinessTime *big.Int, _duration *big.Int, _minStartTradingAmount *big.Int, _maxVaultCapAmount *big.Int, _minUserDepositAmount *big.Int, _maxUserDepositAmount *big.Int, _performanceFee *big.Int, _vaultManager common.Address, _vaultSocialInfo ICopyTradingVaultVaultSocialInfo) (*types.Transaction, error) {
	return _Contract.Contract.CreateVault(&_Contract.TransactOpts, _name, _symbol, _readinessTime, _duration, _minStartTradingAmount, _maxVaultCapAmount, _minUserDepositAmount, _maxUserDepositAmount, _performanceFee, _vaultManager, _vaultSocialInfo)
}

// ExpireCancelOpen is a paid mutator transaction binding the contract method 0xdbbe2fa9.
//
// Solidity: function expireCancelOpen(address _vault, uint256 _orderId) returns()
func (_Contract *ContractTransactor) ExpireCancelOpen(opts *bind.TransactOpts, _vault common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "expireCancelOpen", _vault, _orderId)
}

// ExpireCancelOpen is a paid mutator transaction binding the contract method 0xdbbe2fa9.
//
// Solidity: function expireCancelOpen(address _vault, uint256 _orderId) returns()
func (_Contract *ContractSession) ExpireCancelOpen(_vault common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExpireCancelOpen(&_Contract.TransactOpts, _vault, _orderId)
}

// ExpireCancelOpen is a paid mutator transaction binding the contract method 0xdbbe2fa9.
//
// Solidity: function expireCancelOpen(address _vault, uint256 _orderId) returns()
func (_Contract *ContractTransactorSession) ExpireCancelOpen(_vault common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExpireCancelOpen(&_Contract.TransactOpts, _vault, _orderId)
}

// ExpireCloseVault is a paid mutator transaction binding the contract method 0xc6ac8b7f.
//
// Solidity: function expireCloseVault(address _vault) returns()
func (_Contract *ContractTransactor) ExpireCloseVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "expireCloseVault", _vault)
}

// ExpireCloseVault is a paid mutator transaction binding the contract method 0xc6ac8b7f.
//
// Solidity: function expireCloseVault(address _vault) returns()
func (_Contract *ContractSession) ExpireCloseVault(_vault common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ExpireCloseVault(&_Contract.TransactOpts, _vault)
}

// ExpireCloseVault is a paid mutator transaction binding the contract method 0xc6ac8b7f.
//
// Solidity: function expireCloseVault(address _vault) returns()
func (_Contract *ContractTransactorSession) ExpireCloseVault(_vault common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ExpireCloseVault(&_Contract.TransactOpts, _vault)
}

// ExpireRequestClose is a paid mutator transaction binding the contract method 0x107b39aa.
//
// Solidity: function expireRequestClose(address _vault, uint256 _orderId, uint256 _closeMargin) payable returns()
func (_Contract *ContractTransactor) ExpireRequestClose(opts *bind.TransactOpts, _vault common.Address, _orderId *big.Int, _closeMargin *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "expireRequestClose", _vault, _orderId, _closeMargin)
}

// ExpireRequestClose is a paid mutator transaction binding the contract method 0x107b39aa.
//
// Solidity: function expireRequestClose(address _vault, uint256 _orderId, uint256 _closeMargin) payable returns()
func (_Contract *ContractSession) ExpireRequestClose(_vault common.Address, _orderId *big.Int, _closeMargin *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExpireRequestClose(&_Contract.TransactOpts, _vault, _orderId, _closeMargin)
}

// ExpireRequestClose is a paid mutator transaction binding the contract method 0x107b39aa.
//
// Solidity: function expireRequestClose(address _vault, uint256 _orderId, uint256 _closeMargin) payable returns()
func (_Contract *ContractTransactorSession) ExpireRequestClose(_vault common.Address, _orderId *big.Int, _closeMargin *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExpireRequestClose(&_Contract.TransactOpts, _vault, _orderId, _closeMargin)
}

// ExtractVaultFee is a paid mutator transaction binding the contract method 0xb06cf454.
//
// Solidity: function extractVaultFee(address _token, address _receiver, uint256 _amount) returns()
func (_Contract *ContractTransactor) ExtractVaultFee(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "extractVaultFee", _token, _receiver, _amount)
}

// ExtractVaultFee is a paid mutator transaction binding the contract method 0xb06cf454.
//
// Solidity: function extractVaultFee(address _token, address _receiver, uint256 _amount) returns()
func (_Contract *ContractSession) ExtractVaultFee(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExtractVaultFee(&_Contract.TransactOpts, _token, _receiver, _amount)
}

// ExtractVaultFee is a paid mutator transaction binding the contract method 0xb06cf454.
//
// Solidity: function extractVaultFee(address _token, address _receiver, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) ExtractVaultFee(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExtractVaultFee(&_Contract.TransactOpts, _token, _receiver, _amount)
}

// GetPrice is a paid mutator transaction binding the contract method 0xaeb9fdbd.
//
// Solidity: function getPrice(bytes32 _feedId, bool _priceRollover, bytes[] _updateData) returns(uint256)
func (_Contract *ContractTransactor) GetPrice(opts *bind.TransactOpts, _feedId [32]byte, _priceRollover bool, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getPrice", _feedId, _priceRollover, _updateData)
}

// GetPrice is a paid mutator transaction binding the contract method 0xaeb9fdbd.
//
// Solidity: function getPrice(bytes32 _feedId, bool _priceRollover, bytes[] _updateData) returns(uint256)
func (_Contract *ContractSession) GetPrice(_feedId [32]byte, _priceRollover bool, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.GetPrice(&_Contract.TransactOpts, _feedId, _priceRollover, _updateData)
}

// GetPrice is a paid mutator transaction binding the contract method 0xaeb9fdbd.
//
// Solidity: function getPrice(bytes32 _feedId, bool _priceRollover, bytes[] _updateData) returns(uint256)
func (_Contract *ContractTransactorSession) GetPrice(_feedId [32]byte, _priceRollover bool, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.GetPrice(&_Contract.TransactOpts, _feedId, _priceRollover, _updateData)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _usdtAddress, address _manager, address _vaultAddress, address _tradingAddress, address _oracle) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _usdtAddress common.Address, _manager common.Address, _vaultAddress common.Address, _tradingAddress common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _usdtAddress, _manager, _vaultAddress, _tradingAddress, _oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _usdtAddress, address _manager, address _vaultAddress, address _tradingAddress, address _oracle) returns()
func (_Contract *ContractSession) Initialize(_usdtAddress common.Address, _manager common.Address, _vaultAddress common.Address, _tradingAddress common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _usdtAddress, _manager, _vaultAddress, _tradingAddress, _oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _usdtAddress, address _manager, address _vaultAddress, address _tradingAddress, address _oracle) returns()
func (_Contract *ContractTransactorSession) Initialize(_usdtAddress common.Address, _manager common.Address, _vaultAddress common.Address, _tradingAddress common.Address, _oracle common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _usdtAddress, _manager, _vaultAddress, _tradingAddress, _oracle)
}

// ManagerWhiteList is a paid mutator transaction binding the contract method 0x428a316e.
//
// Solidity: function managerWhiteList(address _user, bool _tag) returns()
func (_Contract *ContractTransactor) ManagerWhiteList(opts *bind.TransactOpts, _user common.Address, _tag bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "managerWhiteList", _user, _tag)
}

// ManagerWhiteList is a paid mutator transaction binding the contract method 0x428a316e.
//
// Solidity: function managerWhiteList(address _user, bool _tag) returns()
func (_Contract *ContractSession) ManagerWhiteList(_user common.Address, _tag bool) (*types.Transaction, error) {
	return _Contract.Contract.ManagerWhiteList(&_Contract.TransactOpts, _user, _tag)
}

// ManagerWhiteList is a paid mutator transaction binding the contract method 0x428a316e.
//
// Solidity: function managerWhiteList(address _user, bool _tag) returns()
func (_Contract *ContractTransactorSession) ManagerWhiteList(_user common.Address, _tag bool) (*types.Transaction, error) {
	return _Contract.Contract.ManagerWhiteList(&_Contract.TransactOpts, _user, _tag)
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

// ResetVaultInfo is a paid mutator transaction binding the contract method 0x12a99ca5.
//
// Solidity: function resetVaultInfo(string _twitter, string _photo, address _user) returns()
func (_Contract *ContractTransactor) ResetVaultInfo(opts *bind.TransactOpts, _twitter string, _photo string, _user common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "resetVaultInfo", _twitter, _photo, _user)
}

// ResetVaultInfo is a paid mutator transaction binding the contract method 0x12a99ca5.
//
// Solidity: function resetVaultInfo(string _twitter, string _photo, address _user) returns()
func (_Contract *ContractSession) ResetVaultInfo(_twitter string, _photo string, _user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ResetVaultInfo(&_Contract.TransactOpts, _twitter, _photo, _user)
}

// ResetVaultInfo is a paid mutator transaction binding the contract method 0x12a99ca5.
//
// Solidity: function resetVaultInfo(string _twitter, string _photo, address _user) returns()
func (_Contract *ContractTransactorSession) ResetVaultInfo(_twitter string, _photo string, _user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ResetVaultInfo(&_Contract.TransactOpts, _twitter, _photo, _user)
}

// ResetVaultStrategy is a paid mutator transaction binding the contract method 0x134c7327.
//
// Solidity: function resetVaultStrategy(address _user, string _strategy) returns()
func (_Contract *ContractTransactor) ResetVaultStrategy(opts *bind.TransactOpts, _user common.Address, _strategy string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "resetVaultStrategy", _user, _strategy)
}

// ResetVaultStrategy is a paid mutator transaction binding the contract method 0x134c7327.
//
// Solidity: function resetVaultStrategy(address _user, string _strategy) returns()
func (_Contract *ContractSession) ResetVaultStrategy(_user common.Address, _strategy string) (*types.Transaction, error) {
	return _Contract.Contract.ResetVaultStrategy(&_Contract.TransactOpts, _user, _strategy)
}

// ResetVaultStrategy is a paid mutator transaction binding the contract method 0x134c7327.
//
// Solidity: function resetVaultStrategy(address _user, string _strategy) returns()
func (_Contract *ContractTransactorSession) ResetVaultStrategy(_user common.Address, _strategy string) (*types.Transaction, error) {
	return _Contract.Contract.ResetVaultStrategy(&_Contract.TransactOpts, _user, _strategy)
}

// SetAssetToken is a paid mutator transaction binding the contract method 0x26bde5aa.
//
// Solidity: function setAssetToken(address _assetToken) returns()
func (_Contract *ContractTransactor) SetAssetToken(opts *bind.TransactOpts, _assetToken common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAssetToken", _assetToken)
}

// SetAssetToken is a paid mutator transaction binding the contract method 0x26bde5aa.
//
// Solidity: function setAssetToken(address _assetToken) returns()
func (_Contract *ContractSession) SetAssetToken(_assetToken common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAssetToken(&_Contract.TransactOpts, _assetToken)
}

// SetAssetToken is a paid mutator transaction binding the contract method 0x26bde5aa.
//
// Solidity: function setAssetToken(address _assetToken) returns()
func (_Contract *ContractTransactorSession) SetAssetToken(_assetToken common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAssetToken(&_Contract.TransactOpts, _assetToken)
}

// SetDepositAmountPeruser is a paid mutator transaction binding the contract method 0xf82735b8.
//
// Solidity: function setDepositAmountPeruser(uint256 _minDepositAmountPeruser, uint256 _maxDepositAmountPeruser) returns()
func (_Contract *ContractTransactor) SetDepositAmountPeruser(opts *bind.TransactOpts, _minDepositAmountPeruser *big.Int, _maxDepositAmountPeruser *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDepositAmountPeruser", _minDepositAmountPeruser, _maxDepositAmountPeruser)
}

// SetDepositAmountPeruser is a paid mutator transaction binding the contract method 0xf82735b8.
//
// Solidity: function setDepositAmountPeruser(uint256 _minDepositAmountPeruser, uint256 _maxDepositAmountPeruser) returns()
func (_Contract *ContractSession) SetDepositAmountPeruser(_minDepositAmountPeruser *big.Int, _maxDepositAmountPeruser *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDepositAmountPeruser(&_Contract.TransactOpts, _minDepositAmountPeruser, _maxDepositAmountPeruser)
}

// SetDepositAmountPeruser is a paid mutator transaction binding the contract method 0xf82735b8.
//
// Solidity: function setDepositAmountPeruser(uint256 _minDepositAmountPeruser, uint256 _maxDepositAmountPeruser) returns()
func (_Contract *ContractTransactorSession) SetDepositAmountPeruser(_minDepositAmountPeruser *big.Int, _maxDepositAmountPeruser *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDepositAmountPeruser(&_Contract.TransactOpts, _minDepositAmountPeruser, _maxDepositAmountPeruser)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _tag) returns()
func (_Contract *ContractTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _tag bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setManager", _manager, _tag)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _tag) returns()
func (_Contract *ContractSession) SetManager(_manager common.Address, _tag bool) (*types.Transaction, error) {
	return _Contract.Contract.SetManager(&_Contract.TransactOpts, _manager, _tag)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _tag) returns()
func (_Contract *ContractTransactorSession) SetManager(_manager common.Address, _tag bool) (*types.Transaction, error) {
	return _Contract.Contract.SetManager(&_Contract.TransactOpts, _manager, _tag)
}

// SetMaxPerformanceFee is a paid mutator transaction binding the contract method 0xdab7b06e.
//
// Solidity: function setMaxPerformanceFee(uint256 _performanceFee) returns()
func (_Contract *ContractTransactor) SetMaxPerformanceFee(opts *bind.TransactOpts, _performanceFee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMaxPerformanceFee", _performanceFee)
}

// SetMaxPerformanceFee is a paid mutator transaction binding the contract method 0xdab7b06e.
//
// Solidity: function setMaxPerformanceFee(uint256 _performanceFee) returns()
func (_Contract *ContractSession) SetMaxPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaxPerformanceFee(&_Contract.TransactOpts, _performanceFee)
}

// SetMaxPerformanceFee is a paid mutator transaction binding the contract method 0xdab7b06e.
//
// Solidity: function setMaxPerformanceFee(uint256 _performanceFee) returns()
func (_Contract *ContractTransactorSession) SetMaxPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaxPerformanceFee(&_Contract.TransactOpts, _performanceFee)
}

// SetMaxPositionSize is a paid mutator transaction binding the contract method 0x26a9c121.
//
// Solidity: function setMaxPositionSize(uint256 _positionSize) returns()
func (_Contract *ContractTransactor) SetMaxPositionSize(opts *bind.TransactOpts, _positionSize *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMaxPositionSize", _positionSize)
}

// SetMaxPositionSize is a paid mutator transaction binding the contract method 0x26a9c121.
//
// Solidity: function setMaxPositionSize(uint256 _positionSize) returns()
func (_Contract *ContractSession) SetMaxPositionSize(_positionSize *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaxPositionSize(&_Contract.TransactOpts, _positionSize)
}

// SetMaxPositionSize is a paid mutator transaction binding the contract method 0x26a9c121.
//
// Solidity: function setMaxPositionSize(uint256 _positionSize) returns()
func (_Contract *ContractTransactorSession) SetMaxPositionSize(_positionSize *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaxPositionSize(&_Contract.TransactOpts, _positionSize)
}

// SetMinManagerCommitAmount is a paid mutator transaction binding the contract method 0x67a8d5c0.
//
// Solidity: function setMinManagerCommitAmount(uint256 _managerCommitAmount) returns()
func (_Contract *ContractTransactor) SetMinManagerCommitAmount(opts *bind.TransactOpts, _managerCommitAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinManagerCommitAmount", _managerCommitAmount)
}

// SetMinManagerCommitAmount is a paid mutator transaction binding the contract method 0x67a8d5c0.
//
// Solidity: function setMinManagerCommitAmount(uint256 _managerCommitAmount) returns()
func (_Contract *ContractSession) SetMinManagerCommitAmount(_managerCommitAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinManagerCommitAmount(&_Contract.TransactOpts, _managerCommitAmount)
}

// SetMinManagerCommitAmount is a paid mutator transaction binding the contract method 0x67a8d5c0.
//
// Solidity: function setMinManagerCommitAmount(uint256 _managerCommitAmount) returns()
func (_Contract *ContractTransactorSession) SetMinManagerCommitAmount(_managerCommitAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinManagerCommitAmount(&_Contract.TransactOpts, _managerCommitAmount)
}

// SetMinVaultDuration is a paid mutator transaction binding the contract method 0x4d96ebd1.
//
// Solidity: function setMinVaultDuration(uint256 _minDuration) returns()
func (_Contract *ContractTransactor) SetMinVaultDuration(opts *bind.TransactOpts, _minDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinVaultDuration", _minDuration)
}

// SetMinVaultDuration is a paid mutator transaction binding the contract method 0x4d96ebd1.
//
// Solidity: function setMinVaultDuration(uint256 _minDuration) returns()
func (_Contract *ContractSession) SetMinVaultDuration(_minDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinVaultDuration(&_Contract.TransactOpts, _minDuration)
}

// SetMinVaultDuration is a paid mutator transaction binding the contract method 0x4d96ebd1.
//
// Solidity: function setMinVaultDuration(uint256 _minDuration) returns()
func (_Contract *ContractTransactorSession) SetMinVaultDuration(_minDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinVaultDuration(&_Contract.TransactOpts, _minDuration)
}

// SetReadinessTime is a paid mutator transaction binding the contract method 0x4789ed91.
//
// Solidity: function setReadinessTime(uint256 _minReadinessTime, uint256 _maxReadinessTime) returns()
func (_Contract *ContractTransactor) SetReadinessTime(opts *bind.TransactOpts, _minReadinessTime *big.Int, _maxReadinessTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setReadinessTime", _minReadinessTime, _maxReadinessTime)
}

// SetReadinessTime is a paid mutator transaction binding the contract method 0x4789ed91.
//
// Solidity: function setReadinessTime(uint256 _minReadinessTime, uint256 _maxReadinessTime) returns()
func (_Contract *ContractSession) SetReadinessTime(_minReadinessTime *big.Int, _maxReadinessTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetReadinessTime(&_Contract.TransactOpts, _minReadinessTime, _maxReadinessTime)
}

// SetReadinessTime is a paid mutator transaction binding the contract method 0x4789ed91.
//
// Solidity: function setReadinessTime(uint256 _minReadinessTime, uint256 _maxReadinessTime) returns()
func (_Contract *ContractTransactorSession) SetReadinessTime(_minReadinessTime *big.Int, _maxReadinessTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetReadinessTime(&_Contract.TransactOpts, _minReadinessTime, _maxReadinessTime)
}

// SetTradingAddress is a paid mutator transaction binding the contract method 0x071eb3c9.
//
// Solidity: function setTradingAddress(address _tradingAddress) returns()
func (_Contract *ContractTransactor) SetTradingAddress(opts *bind.TransactOpts, _tradingAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTradingAddress", _tradingAddress)
}

// SetTradingAddress is a paid mutator transaction binding the contract method 0x071eb3c9.
//
// Solidity: function setTradingAddress(address _tradingAddress) returns()
func (_Contract *ContractSession) SetTradingAddress(_tradingAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetTradingAddress(&_Contract.TransactOpts, _tradingAddress)
}

// SetTradingAddress is a paid mutator transaction binding the contract method 0x071eb3c9.
//
// Solidity: function setTradingAddress(address _tradingAddress) returns()
func (_Contract *ContractTransactorSession) SetTradingAddress(_tradingAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetTradingAddress(&_Contract.TransactOpts, _tradingAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Contract *ContractTransactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Contract *ContractSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultAddress(&_Contract.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Contract *ContractTransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultAddress(&_Contract.TransactOpts, _vaultAddress)
}

// SetVaultCap is a paid mutator transaction binding the contract method 0xf0f5907d.
//
// Solidity: function setVaultCap(uint256 _vaultCap) returns()
func (_Contract *ContractTransactor) SetVaultCap(opts *bind.TransactOpts, _vaultCap *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setVaultCap", _vaultCap)
}

// SetVaultCap is a paid mutator transaction binding the contract method 0xf0f5907d.
//
// Solidity: function setVaultCap(uint256 _vaultCap) returns()
func (_Contract *ContractSession) SetVaultCap(_vaultCap *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultCap(&_Contract.TransactOpts, _vaultCap)
}

// SetVaultCap is a paid mutator transaction binding the contract method 0xf0f5907d.
//
// Solidity: function setVaultCap(uint256 _vaultCap) returns()
func (_Contract *ContractTransactorSession) SetVaultCap(_vaultCap *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultCap(&_Contract.TransactOpts, _vaultCap)
}

// SetVaultFee is a paid mutator transaction binding the contract method 0x8c3d7819.
//
// Solidity: function setVaultFee(uint256 _vaultFee) returns()
func (_Contract *ContractTransactor) SetVaultFee(opts *bind.TransactOpts, _vaultFee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setVaultFee", _vaultFee)
}

// SetVaultFee is a paid mutator transaction binding the contract method 0x8c3d7819.
//
// Solidity: function setVaultFee(uint256 _vaultFee) returns()
func (_Contract *ContractSession) SetVaultFee(_vaultFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultFee(&_Contract.TransactOpts, _vaultFee)
}

// SetVaultFee is a paid mutator transaction binding the contract method 0x8c3d7819.
//
// Solidity: function setVaultFee(uint256 _vaultFee) returns()
func (_Contract *ContractTransactorSession) SetVaultFee(_vaultFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultFee(&_Contract.TransactOpts, _vaultFee)
}

// SetVaultduration is a paid mutator transaction binding the contract method 0x025bbfcf.
//
// Solidity: function setVaultduration(uint256 _minTime, uint256 _maxTime) returns()
func (_Contract *ContractTransactor) SetVaultduration(opts *bind.TransactOpts, _minTime *big.Int, _maxTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setVaultduration", _minTime, _maxTime)
}

// SetVaultduration is a paid mutator transaction binding the contract method 0x025bbfcf.
//
// Solidity: function setVaultduration(uint256 _minTime, uint256 _maxTime) returns()
func (_Contract *ContractSession) SetVaultduration(_minTime *big.Int, _maxTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultduration(&_Contract.TransactOpts, _minTime, _maxTime)
}

// SetVaultduration is a paid mutator transaction binding the contract method 0x025bbfcf.
//
// Solidity: function setVaultduration(uint256 _minTime, uint256 _maxTime) returns()
func (_Contract *ContractTransactorSession) SetVaultduration(_minTime *big.Int, _maxTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetVaultduration(&_Contract.TransactOpts, _minTime, _maxTime)
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

// ContractCreateVaultIterator is returned from FilterCreateVault and is used to iterate over the raw logs and unpacked data for CreateVault events raised by the Contract contract.
type ContractCreateVaultIterator struct {
	Event *ContractCreateVault // Event containing the contract specifics and raw log

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
func (it *ContractCreateVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreateVault)
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
		it.Event = new(ContractCreateVault)
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
func (it *ContractCreateVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreateVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreateVault represents a CreateVault event raised by the Contract contract.
type ContractCreateVault struct {
	Caller       common.Address
	VaultManager common.Address
	Vault        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateVault is a free log retrieval operation binding the contract event 0xb099da66a8504770c8deb56a35848ec236bde5e64dd9223d77c277b1ae09d433.
//
// Solidity: event CreateVault(address indexed caller, address indexed vaultManager, address indexed vault)
func (_Contract *ContractFilterer) FilterCreateVault(opts *bind.FilterOpts, caller []common.Address, vaultManager []common.Address, vault []common.Address) (*ContractCreateVaultIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var vaultManagerRule []interface{}
	for _, vaultManagerItem := range vaultManager {
		vaultManagerRule = append(vaultManagerRule, vaultManagerItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreateVault", callerRule, vaultManagerRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreateVaultIterator{contract: _Contract.contract, event: "CreateVault", logs: logs, sub: sub}, nil
}

// WatchCreateVault is a free log subscription operation binding the contract event 0xb099da66a8504770c8deb56a35848ec236bde5e64dd9223d77c277b1ae09d433.
//
// Solidity: event CreateVault(address indexed caller, address indexed vaultManager, address indexed vault)
func (_Contract *ContractFilterer) WatchCreateVault(opts *bind.WatchOpts, sink chan<- *ContractCreateVault, caller []common.Address, vaultManager []common.Address, vault []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var vaultManagerRule []interface{}
	for _, vaultManagerItem := range vaultManager {
		vaultManagerRule = append(vaultManagerRule, vaultManagerItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreateVault", callerRule, vaultManagerRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreateVault)
				if err := _Contract.contract.UnpackLog(event, "CreateVault", log); err != nil {
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

// ParseCreateVault is a log parse operation binding the contract event 0xb099da66a8504770c8deb56a35848ec236bde5e64dd9223d77c277b1ae09d433.
//
// Solidity: event CreateVault(address indexed caller, address indexed vaultManager, address indexed vault)
func (_Contract *ContractFilterer) ParseCreateVault(log types.Log) (*ContractCreateVault, error) {
	event := new(ContractCreateVault)
	if err := _Contract.contract.UnpackLog(event, "CreateVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExtractVaultFeeIterator is returned from FilterExtractVaultFee and is used to iterate over the raw logs and unpacked data for ExtractVaultFee events raised by the Contract contract.
type ContractExtractVaultFeeIterator struct {
	Event *ContractExtractVaultFee // Event containing the contract specifics and raw log

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
func (it *ContractExtractVaultFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExtractVaultFee)
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
		it.Event = new(ContractExtractVaultFee)
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
func (it *ContractExtractVaultFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExtractVaultFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExtractVaultFee represents a ExtractVaultFee event raised by the Contract contract.
type ContractExtractVaultFee struct {
	Sender   common.Address
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExtractVaultFee is a free log retrieval operation binding the contract event 0xa4940cc5a423ac8ee5cc0ad2c81c2019b1a7268272f414fb642c3abe7981a257.
//
// Solidity: event ExtractVaultFee(address sender, address token, address receiver, uint256 amount)
func (_Contract *ContractFilterer) FilterExtractVaultFee(opts *bind.FilterOpts) (*ContractExtractVaultFeeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExtractVaultFee")
	if err != nil {
		return nil, err
	}
	return &ContractExtractVaultFeeIterator{contract: _Contract.contract, event: "ExtractVaultFee", logs: logs, sub: sub}, nil
}

// WatchExtractVaultFee is a free log subscription operation binding the contract event 0xa4940cc5a423ac8ee5cc0ad2c81c2019b1a7268272f414fb642c3abe7981a257.
//
// Solidity: event ExtractVaultFee(address sender, address token, address receiver, uint256 amount)
func (_Contract *ContractFilterer) WatchExtractVaultFee(opts *bind.WatchOpts, sink chan<- *ContractExtractVaultFee) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExtractVaultFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExtractVaultFee)
				if err := _Contract.contract.UnpackLog(event, "ExtractVaultFee", log); err != nil {
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

// ParseExtractVaultFee is a log parse operation binding the contract event 0xa4940cc5a423ac8ee5cc0ad2c81c2019b1a7268272f414fb642c3abe7981a257.
//
// Solidity: event ExtractVaultFee(address sender, address token, address receiver, uint256 amount)
func (_Contract *ContractFilterer) ParseExtractVaultFee(log types.Log) (*ContractExtractVaultFee, error) {
	event := new(ContractExtractVaultFee)
	if err := _Contract.contract.UnpackLog(event, "ExtractVaultFee", log); err != nil {
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

// ContractNewAssetTokenIterator is returned from FilterNewAssetToken and is used to iterate over the raw logs and unpacked data for NewAssetToken events raised by the Contract contract.
type ContractNewAssetTokenIterator struct {
	Event *ContractNewAssetToken // Event containing the contract specifics and raw log

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
func (it *ContractNewAssetTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewAssetToken)
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
		it.Event = new(ContractNewAssetToken)
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
func (it *ContractNewAssetTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewAssetTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewAssetToken represents a NewAssetToken event raised by the Contract contract.
type ContractNewAssetToken struct {
	AssetToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewAssetToken is a free log retrieval operation binding the contract event 0x7219e96f865ac8a125ed70b79ab8d773cf98f57c775c654ac8d631dcc691e04b.
//
// Solidity: event NewAssetToken(address assetToken)
func (_Contract *ContractFilterer) FilterNewAssetToken(opts *bind.FilterOpts) (*ContractNewAssetTokenIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewAssetToken")
	if err != nil {
		return nil, err
	}
	return &ContractNewAssetTokenIterator{contract: _Contract.contract, event: "NewAssetToken", logs: logs, sub: sub}, nil
}

// WatchNewAssetToken is a free log subscription operation binding the contract event 0x7219e96f865ac8a125ed70b79ab8d773cf98f57c775c654ac8d631dcc691e04b.
//
// Solidity: event NewAssetToken(address assetToken)
func (_Contract *ContractFilterer) WatchNewAssetToken(opts *bind.WatchOpts, sink chan<- *ContractNewAssetToken) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewAssetToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewAssetToken)
				if err := _Contract.contract.UnpackLog(event, "NewAssetToken", log); err != nil {
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

// ParseNewAssetToken is a log parse operation binding the contract event 0x7219e96f865ac8a125ed70b79ab8d773cf98f57c775c654ac8d631dcc691e04b.
//
// Solidity: event NewAssetToken(address assetToken)
func (_Contract *ContractFilterer) ParseNewAssetToken(log types.Log) (*ContractNewAssetToken, error) {
	event := new(ContractNewAssetToken)
	if err := _Contract.contract.UnpackLog(event, "NewAssetToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewDepositAmountPeruserIterator is returned from FilterNewDepositAmountPeruser and is used to iterate over the raw logs and unpacked data for NewDepositAmountPeruser events raised by the Contract contract.
type ContractNewDepositAmountPeruserIterator struct {
	Event *ContractNewDepositAmountPeruser // Event containing the contract specifics and raw log

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
func (it *ContractNewDepositAmountPeruserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewDepositAmountPeruser)
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
		it.Event = new(ContractNewDepositAmountPeruser)
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
func (it *ContractNewDepositAmountPeruserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewDepositAmountPeruserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewDepositAmountPeruser represents a NewDepositAmountPeruser event raised by the Contract contract.
type ContractNewDepositAmountPeruser struct {
	MinDepositAmountPeruser *big.Int
	MaxDepositAmountPeruser *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNewDepositAmountPeruser is a free log retrieval operation binding the contract event 0xdb902bec16a5ce2f3edfc6cc579e56717f8e74722c5b20c2a4245438eed2c82d.
//
// Solidity: event NewDepositAmountPeruser(uint256 minDepositAmountPeruser, uint256 maxDepositAmountPeruser)
func (_Contract *ContractFilterer) FilterNewDepositAmountPeruser(opts *bind.FilterOpts) (*ContractNewDepositAmountPeruserIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewDepositAmountPeruser")
	if err != nil {
		return nil, err
	}
	return &ContractNewDepositAmountPeruserIterator{contract: _Contract.contract, event: "NewDepositAmountPeruser", logs: logs, sub: sub}, nil
}

// WatchNewDepositAmountPeruser is a free log subscription operation binding the contract event 0xdb902bec16a5ce2f3edfc6cc579e56717f8e74722c5b20c2a4245438eed2c82d.
//
// Solidity: event NewDepositAmountPeruser(uint256 minDepositAmountPeruser, uint256 maxDepositAmountPeruser)
func (_Contract *ContractFilterer) WatchNewDepositAmountPeruser(opts *bind.WatchOpts, sink chan<- *ContractNewDepositAmountPeruser) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewDepositAmountPeruser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewDepositAmountPeruser)
				if err := _Contract.contract.UnpackLog(event, "NewDepositAmountPeruser", log); err != nil {
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

// ParseNewDepositAmountPeruser is a log parse operation binding the contract event 0xdb902bec16a5ce2f3edfc6cc579e56717f8e74722c5b20c2a4245438eed2c82d.
//
// Solidity: event NewDepositAmountPeruser(uint256 minDepositAmountPeruser, uint256 maxDepositAmountPeruser)
func (_Contract *ContractFilterer) ParseNewDepositAmountPeruser(log types.Log) (*ContractNewDepositAmountPeruser, error) {
	event := new(ContractNewDepositAmountPeruser)
	if err := _Contract.contract.UnpackLog(event, "NewDepositAmountPeruser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewManagerIterator is returned from FilterNewManager and is used to iterate over the raw logs and unpacked data for NewManager events raised by the Contract contract.
type ContractNewManagerIterator struct {
	Event *ContractNewManager // Event containing the contract specifics and raw log

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
func (it *ContractNewManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewManager)
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
		it.Event = new(ContractNewManager)
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
func (it *ContractNewManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewManager represents a NewManager event raised by the Contract contract.
type ContractNewManager struct {
	Manager common.Address
	Tag     bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewManager is a free log retrieval operation binding the contract event 0x7da4ff26fbcc127d39f19783fd0463c814326353b279895452eacf342712256d.
//
// Solidity: event NewManager(address manager, bool tag)
func (_Contract *ContractFilterer) FilterNewManager(opts *bind.FilterOpts) (*ContractNewManagerIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewManager")
	if err != nil {
		return nil, err
	}
	return &ContractNewManagerIterator{contract: _Contract.contract, event: "NewManager", logs: logs, sub: sub}, nil
}

// WatchNewManager is a free log subscription operation binding the contract event 0x7da4ff26fbcc127d39f19783fd0463c814326353b279895452eacf342712256d.
//
// Solidity: event NewManager(address manager, bool tag)
func (_Contract *ContractFilterer) WatchNewManager(opts *bind.WatchOpts, sink chan<- *ContractNewManager) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewManager)
				if err := _Contract.contract.UnpackLog(event, "NewManager", log); err != nil {
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

// ParseNewManager is a log parse operation binding the contract event 0x7da4ff26fbcc127d39f19783fd0463c814326353b279895452eacf342712256d.
//
// Solidity: event NewManager(address manager, bool tag)
func (_Contract *ContractFilterer) ParseNewManager(log types.Log) (*ContractNewManager, error) {
	event := new(ContractNewManager)
	if err := _Contract.contract.UnpackLog(event, "NewManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewMaxPerformanceFeeIterator is returned from FilterNewMaxPerformanceFee and is used to iterate over the raw logs and unpacked data for NewMaxPerformanceFee events raised by the Contract contract.
type ContractNewMaxPerformanceFeeIterator struct {
	Event *ContractNewMaxPerformanceFee // Event containing the contract specifics and raw log

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
func (it *ContractNewMaxPerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewMaxPerformanceFee)
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
		it.Event = new(ContractNewMaxPerformanceFee)
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
func (it *ContractNewMaxPerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewMaxPerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewMaxPerformanceFee represents a NewMaxPerformanceFee event raised by the Contract contract.
type ContractNewMaxPerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewMaxPerformanceFee is a free log retrieval operation binding the contract event 0xce9d617af89e166c13d22e2990fd195c902a13adbc73dae1fa74fc353dc018da.
//
// Solidity: event NewMaxPerformanceFee(uint256 performanceFee)
func (_Contract *ContractFilterer) FilterNewMaxPerformanceFee(opts *bind.FilterOpts) (*ContractNewMaxPerformanceFeeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewMaxPerformanceFee")
	if err != nil {
		return nil, err
	}
	return &ContractNewMaxPerformanceFeeIterator{contract: _Contract.contract, event: "NewMaxPerformanceFee", logs: logs, sub: sub}, nil
}

// WatchNewMaxPerformanceFee is a free log subscription operation binding the contract event 0xce9d617af89e166c13d22e2990fd195c902a13adbc73dae1fa74fc353dc018da.
//
// Solidity: event NewMaxPerformanceFee(uint256 performanceFee)
func (_Contract *ContractFilterer) WatchNewMaxPerformanceFee(opts *bind.WatchOpts, sink chan<- *ContractNewMaxPerformanceFee) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewMaxPerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewMaxPerformanceFee)
				if err := _Contract.contract.UnpackLog(event, "NewMaxPerformanceFee", log); err != nil {
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

// ParseNewMaxPerformanceFee is a log parse operation binding the contract event 0xce9d617af89e166c13d22e2990fd195c902a13adbc73dae1fa74fc353dc018da.
//
// Solidity: event NewMaxPerformanceFee(uint256 performanceFee)
func (_Contract *ContractFilterer) ParseNewMaxPerformanceFee(log types.Log) (*ContractNewMaxPerformanceFee, error) {
	event := new(ContractNewMaxPerformanceFee)
	if err := _Contract.contract.UnpackLog(event, "NewMaxPerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewMaxPositionSizeIterator is returned from FilterNewMaxPositionSize and is used to iterate over the raw logs and unpacked data for NewMaxPositionSize events raised by the Contract contract.
type ContractNewMaxPositionSizeIterator struct {
	Event *ContractNewMaxPositionSize // Event containing the contract specifics and raw log

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
func (it *ContractNewMaxPositionSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewMaxPositionSize)
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
		it.Event = new(ContractNewMaxPositionSize)
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
func (it *ContractNewMaxPositionSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewMaxPositionSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewMaxPositionSize represents a NewMaxPositionSize event raised by the Contract contract.
type ContractNewMaxPositionSize struct {
	PositionSize *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewMaxPositionSize is a free log retrieval operation binding the contract event 0x442ac02443afe2784125c5c2c3fb915976861f513fdbbf68488ebc46c56c6c71.
//
// Solidity: event NewMaxPositionSize(uint256 positionSize)
func (_Contract *ContractFilterer) FilterNewMaxPositionSize(opts *bind.FilterOpts) (*ContractNewMaxPositionSizeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewMaxPositionSize")
	if err != nil {
		return nil, err
	}
	return &ContractNewMaxPositionSizeIterator{contract: _Contract.contract, event: "NewMaxPositionSize", logs: logs, sub: sub}, nil
}

// WatchNewMaxPositionSize is a free log subscription operation binding the contract event 0x442ac02443afe2784125c5c2c3fb915976861f513fdbbf68488ebc46c56c6c71.
//
// Solidity: event NewMaxPositionSize(uint256 positionSize)
func (_Contract *ContractFilterer) WatchNewMaxPositionSize(opts *bind.WatchOpts, sink chan<- *ContractNewMaxPositionSize) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewMaxPositionSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewMaxPositionSize)
				if err := _Contract.contract.UnpackLog(event, "NewMaxPositionSize", log); err != nil {
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

// ParseNewMaxPositionSize is a log parse operation binding the contract event 0x442ac02443afe2784125c5c2c3fb915976861f513fdbbf68488ebc46c56c6c71.
//
// Solidity: event NewMaxPositionSize(uint256 positionSize)
func (_Contract *ContractFilterer) ParseNewMaxPositionSize(log types.Log) (*ContractNewMaxPositionSize, error) {
	event := new(ContractNewMaxPositionSize)
	if err := _Contract.contract.UnpackLog(event, "NewMaxPositionSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewMinManagerCommitAmountIterator is returned from FilterNewMinManagerCommitAmount and is used to iterate over the raw logs and unpacked data for NewMinManagerCommitAmount events raised by the Contract contract.
type ContractNewMinManagerCommitAmountIterator struct {
	Event *ContractNewMinManagerCommitAmount // Event containing the contract specifics and raw log

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
func (it *ContractNewMinManagerCommitAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewMinManagerCommitAmount)
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
		it.Event = new(ContractNewMinManagerCommitAmount)
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
func (it *ContractNewMinManagerCommitAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewMinManagerCommitAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewMinManagerCommitAmount represents a NewMinManagerCommitAmount event raised by the Contract contract.
type ContractNewMinManagerCommitAmount struct {
	ManagerCommitAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewMinManagerCommitAmount is a free log retrieval operation binding the contract event 0xa936f5b95f3e55c7e83816f215d3e5773989e1afbaa1034e213d3ab6fe43ba60.
//
// Solidity: event NewMinManagerCommitAmount(uint256 managerCommitAmount)
func (_Contract *ContractFilterer) FilterNewMinManagerCommitAmount(opts *bind.FilterOpts) (*ContractNewMinManagerCommitAmountIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewMinManagerCommitAmount")
	if err != nil {
		return nil, err
	}
	return &ContractNewMinManagerCommitAmountIterator{contract: _Contract.contract, event: "NewMinManagerCommitAmount", logs: logs, sub: sub}, nil
}

// WatchNewMinManagerCommitAmount is a free log subscription operation binding the contract event 0xa936f5b95f3e55c7e83816f215d3e5773989e1afbaa1034e213d3ab6fe43ba60.
//
// Solidity: event NewMinManagerCommitAmount(uint256 managerCommitAmount)
func (_Contract *ContractFilterer) WatchNewMinManagerCommitAmount(opts *bind.WatchOpts, sink chan<- *ContractNewMinManagerCommitAmount) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewMinManagerCommitAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewMinManagerCommitAmount)
				if err := _Contract.contract.UnpackLog(event, "NewMinManagerCommitAmount", log); err != nil {
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

// ParseNewMinManagerCommitAmount is a log parse operation binding the contract event 0xa936f5b95f3e55c7e83816f215d3e5773989e1afbaa1034e213d3ab6fe43ba60.
//
// Solidity: event NewMinManagerCommitAmount(uint256 managerCommitAmount)
func (_Contract *ContractFilterer) ParseNewMinManagerCommitAmount(log types.Log) (*ContractNewMinManagerCommitAmount, error) {
	event := new(ContractNewMinManagerCommitAmount)
	if err := _Contract.contract.UnpackLog(event, "NewMinManagerCommitAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewReadinessTimeIterator is returned from FilterNewReadinessTime and is used to iterate over the raw logs and unpacked data for NewReadinessTime events raised by the Contract contract.
type ContractNewReadinessTimeIterator struct {
	Event *ContractNewReadinessTime // Event containing the contract specifics and raw log

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
func (it *ContractNewReadinessTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewReadinessTime)
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
		it.Event = new(ContractNewReadinessTime)
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
func (it *ContractNewReadinessTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewReadinessTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewReadinessTime represents a NewReadinessTime event raised by the Contract contract.
type ContractNewReadinessTime struct {
	MinReadinessTime *big.Int
	MaxReadinessTime *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewReadinessTime is a free log retrieval operation binding the contract event 0xc668aa8cc9e4d0a776615f9cce4896cb0d8fd126bfc2848acaa8583f516d3bea.
//
// Solidity: event NewReadinessTime(uint256 _minReadinessTime, uint256 _maxReadinessTime)
func (_Contract *ContractFilterer) FilterNewReadinessTime(opts *bind.FilterOpts) (*ContractNewReadinessTimeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewReadinessTime")
	if err != nil {
		return nil, err
	}
	return &ContractNewReadinessTimeIterator{contract: _Contract.contract, event: "NewReadinessTime", logs: logs, sub: sub}, nil
}

// WatchNewReadinessTime is a free log subscription operation binding the contract event 0xc668aa8cc9e4d0a776615f9cce4896cb0d8fd126bfc2848acaa8583f516d3bea.
//
// Solidity: event NewReadinessTime(uint256 _minReadinessTime, uint256 _maxReadinessTime)
func (_Contract *ContractFilterer) WatchNewReadinessTime(opts *bind.WatchOpts, sink chan<- *ContractNewReadinessTime) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewReadinessTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewReadinessTime)
				if err := _Contract.contract.UnpackLog(event, "NewReadinessTime", log); err != nil {
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

// ParseNewReadinessTime is a log parse operation binding the contract event 0xc668aa8cc9e4d0a776615f9cce4896cb0d8fd126bfc2848acaa8583f516d3bea.
//
// Solidity: event NewReadinessTime(uint256 _minReadinessTime, uint256 _maxReadinessTime)
func (_Contract *ContractFilterer) ParseNewReadinessTime(log types.Log) (*ContractNewReadinessTime, error) {
	event := new(ContractNewReadinessTime)
	if err := _Contract.contract.UnpackLog(event, "NewReadinessTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewTradingAddressIterator is returned from FilterNewTradingAddress and is used to iterate over the raw logs and unpacked data for NewTradingAddress events raised by the Contract contract.
type ContractNewTradingAddressIterator struct {
	Event *ContractNewTradingAddress // Event containing the contract specifics and raw log

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
func (it *ContractNewTradingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewTradingAddress)
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
		it.Event = new(ContractNewTradingAddress)
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
func (it *ContractNewTradingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewTradingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewTradingAddress represents a NewTradingAddress event raised by the Contract contract.
type ContractNewTradingAddress struct {
	TradingAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewTradingAddress is a free log retrieval operation binding the contract event 0xa91cdc15c51c24b03a04a853ad3ac853139634e4f3f272fe22f33d4e7f997a1a.
//
// Solidity: event NewTradingAddress(address tradingAddress)
func (_Contract *ContractFilterer) FilterNewTradingAddress(opts *bind.FilterOpts) (*ContractNewTradingAddressIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewTradingAddress")
	if err != nil {
		return nil, err
	}
	return &ContractNewTradingAddressIterator{contract: _Contract.contract, event: "NewTradingAddress", logs: logs, sub: sub}, nil
}

// WatchNewTradingAddress is a free log subscription operation binding the contract event 0xa91cdc15c51c24b03a04a853ad3ac853139634e4f3f272fe22f33d4e7f997a1a.
//
// Solidity: event NewTradingAddress(address tradingAddress)
func (_Contract *ContractFilterer) WatchNewTradingAddress(opts *bind.WatchOpts, sink chan<- *ContractNewTradingAddress) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewTradingAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewTradingAddress)
				if err := _Contract.contract.UnpackLog(event, "NewTradingAddress", log); err != nil {
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

// ParseNewTradingAddress is a log parse operation binding the contract event 0xa91cdc15c51c24b03a04a853ad3ac853139634e4f3f272fe22f33d4e7f997a1a.
//
// Solidity: event NewTradingAddress(address tradingAddress)
func (_Contract *ContractFilterer) ParseNewTradingAddress(log types.Log) (*ContractNewTradingAddress, error) {
	event := new(ContractNewTradingAddress)
	if err := _Contract.contract.UnpackLog(event, "NewTradingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultAddressIterator is returned from FilterNewVaultAddress and is used to iterate over the raw logs and unpacked data for NewVaultAddress events raised by the Contract contract.
type ContractNewVaultAddressIterator struct {
	Event *ContractNewVaultAddress // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultAddress)
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
		it.Event = new(ContractNewVaultAddress)
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
func (it *ContractNewVaultAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultAddress represents a NewVaultAddress event raised by the Contract contract.
type ContractNewVaultAddress struct {
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewVaultAddress is a free log retrieval operation binding the contract event 0x66211754288f8000fcfb529f1dcacece425f68e712d8b06359f5c2df22d59021.
//
// Solidity: event NewVaultAddress(address vaultAddress)
func (_Contract *ContractFilterer) FilterNewVaultAddress(opts *bind.FilterOpts) (*ContractNewVaultAddressIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultAddress")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultAddressIterator{contract: _Contract.contract, event: "NewVaultAddress", logs: logs, sub: sub}, nil
}

// WatchNewVaultAddress is a free log subscription operation binding the contract event 0x66211754288f8000fcfb529f1dcacece425f68e712d8b06359f5c2df22d59021.
//
// Solidity: event NewVaultAddress(address vaultAddress)
func (_Contract *ContractFilterer) WatchNewVaultAddress(opts *bind.WatchOpts, sink chan<- *ContractNewVaultAddress) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultAddress)
				if err := _Contract.contract.UnpackLog(event, "NewVaultAddress", log); err != nil {
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

// ParseNewVaultAddress is a log parse operation binding the contract event 0x66211754288f8000fcfb529f1dcacece425f68e712d8b06359f5c2df22d59021.
//
// Solidity: event NewVaultAddress(address vaultAddress)
func (_Contract *ContractFilterer) ParseNewVaultAddress(log types.Log) (*ContractNewVaultAddress, error) {
	event := new(ContractNewVaultAddress)
	if err := _Contract.contract.UnpackLog(event, "NewVaultAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultCapIterator is returned from FilterNewVaultCap and is used to iterate over the raw logs and unpacked data for NewVaultCap events raised by the Contract contract.
type ContractNewVaultCapIterator struct {
	Event *ContractNewVaultCap // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultCap)
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
		it.Event = new(ContractNewVaultCap)
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
func (it *ContractNewVaultCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultCap represents a NewVaultCap event raised by the Contract contract.
type ContractNewVaultCap struct {
	VaultCap *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewVaultCap is a free log retrieval operation binding the contract event 0x7fd9455ad268ee5756bed6477a7c817fcf02d680e5f94dfeef619fa4acb545de.
//
// Solidity: event NewVaultCap(uint256 vaultCap)
func (_Contract *ContractFilterer) FilterNewVaultCap(opts *bind.FilterOpts) (*ContractNewVaultCapIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultCap")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultCapIterator{contract: _Contract.contract, event: "NewVaultCap", logs: logs, sub: sub}, nil
}

// WatchNewVaultCap is a free log subscription operation binding the contract event 0x7fd9455ad268ee5756bed6477a7c817fcf02d680e5f94dfeef619fa4acb545de.
//
// Solidity: event NewVaultCap(uint256 vaultCap)
func (_Contract *ContractFilterer) WatchNewVaultCap(opts *bind.WatchOpts, sink chan<- *ContractNewVaultCap) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultCap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultCap)
				if err := _Contract.contract.UnpackLog(event, "NewVaultCap", log); err != nil {
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

// ParseNewVaultCap is a log parse operation binding the contract event 0x7fd9455ad268ee5756bed6477a7c817fcf02d680e5f94dfeef619fa4acb545de.
//
// Solidity: event NewVaultCap(uint256 vaultCap)
func (_Contract *ContractFilterer) ParseNewVaultCap(log types.Log) (*ContractNewVaultCap, error) {
	event := new(ContractNewVaultCap)
	if err := _Contract.contract.UnpackLog(event, "NewVaultCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultDurationIterator is returned from FilterNewVaultDuration and is used to iterate over the raw logs and unpacked data for NewVaultDuration events raised by the Contract contract.
type ContractNewVaultDurationIterator struct {
	Event *ContractNewVaultDuration // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultDuration)
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
		it.Event = new(ContractNewVaultDuration)
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
func (it *ContractNewVaultDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultDuration represents a NewVaultDuration event raised by the Contract contract.
type ContractNewVaultDuration struct {
	MinDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewVaultDuration is a free log retrieval operation binding the contract event 0xa86b1dd44e90bc3ae824363d1c7a86ac47f67dca02c54d04677dc62dc58672c0.
//
// Solidity: event NewVaultDuration(uint256 minDuration)
func (_Contract *ContractFilterer) FilterNewVaultDuration(opts *bind.FilterOpts) (*ContractNewVaultDurationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultDuration")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultDurationIterator{contract: _Contract.contract, event: "NewVaultDuration", logs: logs, sub: sub}, nil
}

// WatchNewVaultDuration is a free log subscription operation binding the contract event 0xa86b1dd44e90bc3ae824363d1c7a86ac47f67dca02c54d04677dc62dc58672c0.
//
// Solidity: event NewVaultDuration(uint256 minDuration)
func (_Contract *ContractFilterer) WatchNewVaultDuration(opts *bind.WatchOpts, sink chan<- *ContractNewVaultDuration) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultDuration)
				if err := _Contract.contract.UnpackLog(event, "NewVaultDuration", log); err != nil {
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

// ParseNewVaultDuration is a log parse operation binding the contract event 0xa86b1dd44e90bc3ae824363d1c7a86ac47f67dca02c54d04677dc62dc58672c0.
//
// Solidity: event NewVaultDuration(uint256 minDuration)
func (_Contract *ContractFilterer) ParseNewVaultDuration(log types.Log) (*ContractNewVaultDuration, error) {
	event := new(ContractNewVaultDuration)
	if err := _Contract.contract.UnpackLog(event, "NewVaultDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultFeeIterator is returned from FilterNewVaultFee and is used to iterate over the raw logs and unpacked data for NewVaultFee events raised by the Contract contract.
type ContractNewVaultFeeIterator struct {
	Event *ContractNewVaultFee // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultFee)
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
		it.Event = new(ContractNewVaultFee)
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
func (it *ContractNewVaultFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultFee represents a NewVaultFee event raised by the Contract contract.
type ContractNewVaultFee struct {
	VaultFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewVaultFee is a free log retrieval operation binding the contract event 0xd8bfeb223a6c7ae3cf0dd2eaa53b29a92ef784e00119c2a5c8a6117dcfa32d14.
//
// Solidity: event NewVaultFee(uint256 vaultFee)
func (_Contract *ContractFilterer) FilterNewVaultFee(opts *bind.FilterOpts) (*ContractNewVaultFeeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultFee")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultFeeIterator{contract: _Contract.contract, event: "NewVaultFee", logs: logs, sub: sub}, nil
}

// WatchNewVaultFee is a free log subscription operation binding the contract event 0xd8bfeb223a6c7ae3cf0dd2eaa53b29a92ef784e00119c2a5c8a6117dcfa32d14.
//
// Solidity: event NewVaultFee(uint256 vaultFee)
func (_Contract *ContractFilterer) WatchNewVaultFee(opts *bind.WatchOpts, sink chan<- *ContractNewVaultFee) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultFee)
				if err := _Contract.contract.UnpackLog(event, "NewVaultFee", log); err != nil {
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

// ParseNewVaultFee is a log parse operation binding the contract event 0xd8bfeb223a6c7ae3cf0dd2eaa53b29a92ef784e00119c2a5c8a6117dcfa32d14.
//
// Solidity: event NewVaultFee(uint256 vaultFee)
func (_Contract *ContractFilterer) ParseNewVaultFee(log types.Log) (*ContractNewVaultFee, error) {
	event := new(ContractNewVaultFee)
	if err := _Contract.contract.UnpackLog(event, "NewVaultFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultInfoIterator is returned from FilterNewVaultInfo and is used to iterate over the raw logs and unpacked data for NewVaultInfo events raised by the Contract contract.
type ContractNewVaultInfoIterator struct {
	Event *ContractNewVaultInfo // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultInfo)
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
		it.Event = new(ContractNewVaultInfo)
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
func (it *ContractNewVaultInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultInfo represents a NewVaultInfo event raised by the Contract contract.
type ContractNewVaultInfo struct {
	User    common.Address
	Caller  common.Address
	Twitter string
	Photo   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewVaultInfo is a free log retrieval operation binding the contract event 0x5dcbc88187bf8bfbe6192ab1f66acae24a46966971d7a4c8e0589c64ca458ce4.
//
// Solidity: event NewVaultInfo(address _user, address caller, string _twitter, string _photo)
func (_Contract *ContractFilterer) FilterNewVaultInfo(opts *bind.FilterOpts) (*ContractNewVaultInfoIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultInfo")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultInfoIterator{contract: _Contract.contract, event: "NewVaultInfo", logs: logs, sub: sub}, nil
}

// WatchNewVaultInfo is a free log subscription operation binding the contract event 0x5dcbc88187bf8bfbe6192ab1f66acae24a46966971d7a4c8e0589c64ca458ce4.
//
// Solidity: event NewVaultInfo(address _user, address caller, string _twitter, string _photo)
func (_Contract *ContractFilterer) WatchNewVaultInfo(opts *bind.WatchOpts, sink chan<- *ContractNewVaultInfo) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultInfo)
				if err := _Contract.contract.UnpackLog(event, "NewVaultInfo", log); err != nil {
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

// ParseNewVaultInfo is a log parse operation binding the contract event 0x5dcbc88187bf8bfbe6192ab1f66acae24a46966971d7a4c8e0589c64ca458ce4.
//
// Solidity: event NewVaultInfo(address _user, address caller, string _twitter, string _photo)
func (_Contract *ContractFilterer) ParseNewVaultInfo(log types.Log) (*ContractNewVaultInfo, error) {
	event := new(ContractNewVaultInfo)
	if err := _Contract.contract.UnpackLog(event, "NewVaultInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultStrategyIterator is returned from FilterNewVaultStrategy and is used to iterate over the raw logs and unpacked data for NewVaultStrategy events raised by the Contract contract.
type ContractNewVaultStrategyIterator struct {
	Event *ContractNewVaultStrategy // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultStrategy)
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
		it.Event = new(ContractNewVaultStrategy)
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
func (it *ContractNewVaultStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultStrategy represents a NewVaultStrategy event raised by the Contract contract.
type ContractNewVaultStrategy struct {
	User     common.Address
	Caller   common.Address
	Strategy string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewVaultStrategy is a free log retrieval operation binding the contract event 0xc50e878ff113c96cf1ce57da1428552b9997613af442885a42bf6e03dbb4785c.
//
// Solidity: event NewVaultStrategy(address _user, address caller, string _strategy)
func (_Contract *ContractFilterer) FilterNewVaultStrategy(opts *bind.FilterOpts) (*ContractNewVaultStrategyIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultStrategy")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultStrategyIterator{contract: _Contract.contract, event: "NewVaultStrategy", logs: logs, sub: sub}, nil
}

// WatchNewVaultStrategy is a free log subscription operation binding the contract event 0xc50e878ff113c96cf1ce57da1428552b9997613af442885a42bf6e03dbb4785c.
//
// Solidity: event NewVaultStrategy(address _user, address caller, string _strategy)
func (_Contract *ContractFilterer) WatchNewVaultStrategy(opts *bind.WatchOpts, sink chan<- *ContractNewVaultStrategy) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultStrategy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultStrategy)
				if err := _Contract.contract.UnpackLog(event, "NewVaultStrategy", log); err != nil {
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

// ParseNewVaultStrategy is a log parse operation binding the contract event 0xc50e878ff113c96cf1ce57da1428552b9997613af442885a42bf6e03dbb4785c.
//
// Solidity: event NewVaultStrategy(address _user, address caller, string _strategy)
func (_Contract *ContractFilterer) ParseNewVaultStrategy(log types.Log) (*ContractNewVaultStrategy, error) {
	event := new(ContractNewVaultStrategy)
	if err := _Contract.contract.UnpackLog(event, "NewVaultStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewVaultdurationIterator is returned from FilterNewVaultduration and is used to iterate over the raw logs and unpacked data for NewVaultduration events raised by the Contract contract.
type ContractNewVaultdurationIterator struct {
	Event *ContractNewVaultduration // Event containing the contract specifics and raw log

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
func (it *ContractNewVaultdurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewVaultduration)
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
		it.Event = new(ContractNewVaultduration)
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
func (it *ContractNewVaultdurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewVaultdurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewVaultduration represents a NewVaultduration event raised by the Contract contract.
type ContractNewVaultduration struct {
	MinTime *big.Int
	MaxTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewVaultduration is a free log retrieval operation binding the contract event 0xcfa6bee896794b44dea1e0b6143de968f9499849b8991b6ce29a9bb1b47a00b1.
//
// Solidity: event NewVaultduration(uint256 minTime, uint256 maxTime)
func (_Contract *ContractFilterer) FilterNewVaultduration(opts *bind.FilterOpts) (*ContractNewVaultdurationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewVaultduration")
	if err != nil {
		return nil, err
	}
	return &ContractNewVaultdurationIterator{contract: _Contract.contract, event: "NewVaultduration", logs: logs, sub: sub}, nil
}

// WatchNewVaultduration is a free log subscription operation binding the contract event 0xcfa6bee896794b44dea1e0b6143de968f9499849b8991b6ce29a9bb1b47a00b1.
//
// Solidity: event NewVaultduration(uint256 minTime, uint256 maxTime)
func (_Contract *ContractFilterer) WatchNewVaultduration(opts *bind.WatchOpts, sink chan<- *ContractNewVaultduration) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewVaultduration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewVaultduration)
				if err := _Contract.contract.UnpackLog(event, "NewVaultduration", log); err != nil {
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

// ParseNewVaultduration is a log parse operation binding the contract event 0xcfa6bee896794b44dea1e0b6143de968f9499849b8991b6ce29a9bb1b47a00b1.
//
// Solidity: event NewVaultduration(uint256 minTime, uint256 maxTime)
func (_Contract *ContractFilterer) ParseNewVaultduration(log types.Log) (*ContractNewVaultduration, error) {
	event := new(ContractNewVaultduration)
	if err := _Contract.contract.UnpackLog(event, "NewVaultduration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewWhiteListIterator is returned from FilterNewWhiteList and is used to iterate over the raw logs and unpacked data for NewWhiteList events raised by the Contract contract.
type ContractNewWhiteListIterator struct {
	Event *ContractNewWhiteList // Event containing the contract specifics and raw log

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
func (it *ContractNewWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewWhiteList)
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
		it.Event = new(ContractNewWhiteList)
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
func (it *ContractNewWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewWhiteList represents a NewWhiteList event raised by the Contract contract.
type ContractNewWhiteList struct {
	User common.Address
	Tag  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewWhiteList is a free log retrieval operation binding the contract event 0x8fecd5f8948f7f132a7781596491325ccecb323d81a24bcd7b946b1301fd554d.
//
// Solidity: event NewWhiteList(address user, bool tag)
func (_Contract *ContractFilterer) FilterNewWhiteList(opts *bind.FilterOpts) (*ContractNewWhiteListIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewWhiteList")
	if err != nil {
		return nil, err
	}
	return &ContractNewWhiteListIterator{contract: _Contract.contract, event: "NewWhiteList", logs: logs, sub: sub}, nil
}

// WatchNewWhiteList is a free log subscription operation binding the contract event 0x8fecd5f8948f7f132a7781596491325ccecb323d81a24bcd7b946b1301fd554d.
//
// Solidity: event NewWhiteList(address user, bool tag)
func (_Contract *ContractFilterer) WatchNewWhiteList(opts *bind.WatchOpts, sink chan<- *ContractNewWhiteList) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewWhiteList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewWhiteList)
				if err := _Contract.contract.UnpackLog(event, "NewWhiteList", log); err != nil {
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

// ParseNewWhiteList is a log parse operation binding the contract event 0x8fecd5f8948f7f132a7781596491325ccecb323d81a24bcd7b946b1301fd554d.
//
// Solidity: event NewWhiteList(address user, bool tag)
func (_Contract *ContractFilterer) ParseNewWhiteList(log types.Log) (*ContractNewWhiteList, error) {
	event := new(ContractNewWhiteList)
	if err := _Contract.contract.UnpackLog(event, "NewWhiteList", log); err != nil {
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
