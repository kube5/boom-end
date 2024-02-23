// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumVault.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CancelOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumVault.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CloseOnTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"closeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"brokerage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"vaultStatus\",\"type\":\"bool\"}],\"name\":\"CloseVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"closeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"brokerage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"vaultStatus\",\"type\":\"bool\"}],\"name\":\"CloseVaultFactroy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"followers\",\"type\":\"uint256\"}],\"name\":\"Follower\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"readinessTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userMinDepositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxCapValue\",\"type\":\"uint256\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"twitter\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"photo\",\"type\":\"string\"}],\"name\":\"NewSocialInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"strategy\",\"type\":\"string\"}],\"name\":\"NewStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumVault.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"OpenOnTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"closeMargin\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"RequestClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"RequestOpenLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"RequestOpenMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isAdd\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"UpdateMargin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOpen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeVaultFactroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalvalue\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalvalue\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"followers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllOrderList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTotalVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"getTradingTotalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_assetToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trading\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"readinessTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStartTradingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVaultCapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUserDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUserDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultManagementFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minManagerCommitAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDurationTime\",\"type\":\"uint256\"}],\"internalType\":\"structICopyTradingVault.VaultInfo\",\"name\":\"_vaultInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"vaultName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultTwitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultPhoto\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultStrategy\",\"type\":\"string\"}],\"internalType\":\"structICopyTradingVault.VaultSocialInfo\",\"name\":\"_socialInfo\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCapValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxOpenTradingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTradingNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDurationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStartValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_wasExecuted\",\"type\":\"bool\"},{\"internalType\":\"enumCloseType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"onTradeClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_wasExecuted\",\"type\":\"bool\"}],\"name\":\"onTradeOpen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderPairId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderStatusList\",\"outputs\":[{\"internalType\":\"enumVault.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readinessTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_closeMargin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestClose\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTradeType\",\"name\":\"_t\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"_base\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestOpenLimit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBase\",\"name\":\"_base\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestOpenMarket\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isAdd\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"}],\"name\":\"updateMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"name\":\"updateOpenRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_twitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_photo\",\"type\":\"string\"}],\"name\":\"updateSocialInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_strategy\",\"type\":\"string\"}],\"name\":\"updateStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"}],\"name\":\"updateTPAndSL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDepositInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userMaxDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userMinDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManagementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultSocialInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"vaultName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultTwitter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultPhoto\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultStrategy\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Contract *ContractCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Contract *ContractSession) Asset() (common.Address, error) {
	return _Contract.Contract.Asset(&_Contract.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Contract *ContractCallerSession) Asset() (common.Address, error) {
	return _Contract.Contract.Asset(&_Contract.CallOpts)
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

// ConvertToAssets is a free data retrieval call binding the contract method 0x181e7b3b.
//
// Solidity: function convertToAssets(uint256 shares, uint256 _totalvalue) view returns(uint256)
func (_Contract *ContractCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int, _totalvalue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "convertToAssets", shares, _totalvalue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x181e7b3b.
//
// Solidity: function convertToAssets(uint256 shares, uint256 _totalvalue) view returns(uint256)
func (_Contract *ContractSession) ConvertToAssets(shares *big.Int, _totalvalue *big.Int) (*big.Int, error) {
	return _Contract.Contract.ConvertToAssets(&_Contract.CallOpts, shares, _totalvalue)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x181e7b3b.
//
// Solidity: function convertToAssets(uint256 shares, uint256 _totalvalue) view returns(uint256)
func (_Contract *ContractCallerSession) ConvertToAssets(shares *big.Int, _totalvalue *big.Int) (*big.Int, error) {
	return _Contract.Contract.ConvertToAssets(&_Contract.CallOpts, shares, _totalvalue)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xdb2088f4.
//
// Solidity: function convertToShares(uint256 assets, uint256 _totalvalue) view returns(uint256)
func (_Contract *ContractCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int, _totalvalue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "convertToShares", assets, _totalvalue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xdb2088f4.
//
// Solidity: function convertToShares(uint256 assets, uint256 _totalvalue) view returns(uint256)
func (_Contract *ContractSession) ConvertToShares(assets *big.Int, _totalvalue *big.Int) (*big.Int, error) {
	return _Contract.Contract.ConvertToShares(&_Contract.CallOpts, assets, _totalvalue)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xdb2088f4.
//
// Solidity: function convertToShares(uint256 assets, uint256 _totalvalue) view returns(uint256)
func (_Contract *ContractCallerSession) ConvertToShares(assets *big.Int, _totalvalue *big.Int) (*big.Int, error) {
	return _Contract.Contract.ConvertToShares(&_Contract.CallOpts, assets, _totalvalue)
}

// CurrentDepositAmount is a free data retrieval call binding the contract method 0xc6834d1c.
//
// Solidity: function currentDepositAmount() view returns(uint256)
func (_Contract *ContractCaller) CurrentDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDepositAmount is a free data retrieval call binding the contract method 0xc6834d1c.
//
// Solidity: function currentDepositAmount() view returns(uint256)
func (_Contract *ContractSession) CurrentDepositAmount() (*big.Int, error) {
	return _Contract.Contract.CurrentDepositAmount(&_Contract.CallOpts)
}

// CurrentDepositAmount is a free data retrieval call binding the contract method 0xc6834d1c.
//
// Solidity: function currentDepositAmount() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentDepositAmount() (*big.Int, error) {
	return _Contract.Contract.CurrentDepositAmount(&_Contract.CallOpts)
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

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Contract *ContractCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Contract *ContractSession) Duration() (*big.Int, error) {
	return _Contract.Contract.Duration(&_Contract.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Contract *ContractCallerSession) Duration() (*big.Int, error) {
	return _Contract.Contract.Duration(&_Contract.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_Contract *ContractCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_Contract *ContractSession) FactoryAddress() (common.Address, error) {
	return _Contract.Contract.FactoryAddress(&_Contract.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_Contract *ContractCallerSession) FactoryAddress() (common.Address, error) {
	return _Contract.Contract.FactoryAddress(&_Contract.CallOpts)
}

// Followers is a free data retrieval call binding the contract method 0xac3e3d8d.
//
// Solidity: function followers() view returns(uint256)
func (_Contract *ContractCaller) Followers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "followers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Followers is a free data retrieval call binding the contract method 0xac3e3d8d.
//
// Solidity: function followers() view returns(uint256)
func (_Contract *ContractSession) Followers() (*big.Int, error) {
	return _Contract.Contract.Followers(&_Contract.CallOpts)
}

// Followers is a free data retrieval call binding the contract method 0xac3e3d8d.
//
// Solidity: function followers() view returns(uint256)
func (_Contract *ContractCallerSession) Followers() (*big.Int, error) {
	return _Contract.Contract.Followers(&_Contract.CallOpts)
}

// GetAllOrderList is a free data retrieval call binding the contract method 0x23fffefa.
//
// Solidity: function getAllOrderList() view returns(uint256[])
func (_Contract *ContractCaller) GetAllOrderList(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllOrderList")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllOrderList is a free data retrieval call binding the contract method 0x23fffefa.
//
// Solidity: function getAllOrderList() view returns(uint256[])
func (_Contract *ContractSession) GetAllOrderList() ([]*big.Int, error) {
	return _Contract.Contract.GetAllOrderList(&_Contract.CallOpts)
}

// GetAllOrderList is a free data retrieval call binding the contract method 0x23fffefa.
//
// Solidity: function getAllOrderList() view returns(uint256[])
func (_Contract *ContractCallerSession) GetAllOrderList() ([]*big.Int, error) {
	return _Contract.Contract.GetAllOrderList(&_Contract.CallOpts)
}

// GetCurrentTotalVault is a free data retrieval call binding the contract method 0x736691b0.
//
// Solidity: function getCurrentTotalVault() view returns(uint256)
func (_Contract *ContractCaller) GetCurrentTotalVault(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCurrentTotalVault")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTotalVault is a free data retrieval call binding the contract method 0x736691b0.
//
// Solidity: function getCurrentTotalVault() view returns(uint256)
func (_Contract *ContractSession) GetCurrentTotalVault() (*big.Int, error) {
	return _Contract.Contract.GetCurrentTotalVault(&_Contract.CallOpts)
}

// GetCurrentTotalVault is a free data retrieval call binding the contract method 0x736691b0.
//
// Solidity: function getCurrentTotalVault() view returns(uint256)
func (_Contract *ContractCallerSession) GetCurrentTotalVault() (*big.Int, error) {
	return _Contract.Contract.GetCurrentTotalVault(&_Contract.CallOpts)
}

// GetVaultState is a free data retrieval call binding the contract method 0x4a8c110a.
//
// Solidity: function getVaultState() view returns(bool)
func (_Contract *ContractCaller) GetVaultState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVaultState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetVaultState is a free data retrieval call binding the contract method 0x4a8c110a.
//
// Solidity: function getVaultState() view returns(bool)
func (_Contract *ContractSession) GetVaultState() (bool, error) {
	return _Contract.Contract.GetVaultState(&_Contract.CallOpts)
}

// GetVaultState is a free data retrieval call binding the contract method 0x4a8c110a.
//
// Solidity: function getVaultState() view returns(bool)
func (_Contract *ContractCallerSession) GetVaultState() (bool, error) {
	return _Contract.Contract.GetVaultState(&_Contract.CallOpts)
}

// MaxCapValue is a free data retrieval call binding the contract method 0x8b207675.
//
// Solidity: function maxCapValue() view returns(uint256)
func (_Contract *ContractCaller) MaxCapValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxCapValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCapValue is a free data retrieval call binding the contract method 0x8b207675.
//
// Solidity: function maxCapValue() view returns(uint256)
func (_Contract *ContractSession) MaxCapValue() (*big.Int, error) {
	return _Contract.Contract.MaxCapValue(&_Contract.CallOpts)
}

// MaxCapValue is a free data retrieval call binding the contract method 0x8b207675.
//
// Solidity: function maxCapValue() view returns(uint256)
func (_Contract *ContractCallerSession) MaxCapValue() (*big.Int, error) {
	return _Contract.Contract.MaxCapValue(&_Contract.CallOpts)
}

// MaxOpenTradingAmount is a free data retrieval call binding the contract method 0x06420da1.
//
// Solidity: function maxOpenTradingAmount() view returns(uint256)
func (_Contract *ContractCaller) MaxOpenTradingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxOpenTradingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOpenTradingAmount is a free data retrieval call binding the contract method 0x06420da1.
//
// Solidity: function maxOpenTradingAmount() view returns(uint256)
func (_Contract *ContractSession) MaxOpenTradingAmount() (*big.Int, error) {
	return _Contract.Contract.MaxOpenTradingAmount(&_Contract.CallOpts)
}

// MaxOpenTradingAmount is a free data retrieval call binding the contract method 0x06420da1.
//
// Solidity: function maxOpenTradingAmount() view returns(uint256)
func (_Contract *ContractCallerSession) MaxOpenTradingAmount() (*big.Int, error) {
	return _Contract.Contract.MaxOpenTradingAmount(&_Contract.CallOpts)
}

// MaxTradingNumber is a free data retrieval call binding the contract method 0x10cd277b.
//
// Solidity: function maxTradingNumber() view returns(uint256)
func (_Contract *ContractCaller) MaxTradingNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxTradingNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTradingNumber is a free data retrieval call binding the contract method 0x10cd277b.
//
// Solidity: function maxTradingNumber() view returns(uint256)
func (_Contract *ContractSession) MaxTradingNumber() (*big.Int, error) {
	return _Contract.Contract.MaxTradingNumber(&_Contract.CallOpts)
}

// MaxTradingNumber is a free data retrieval call binding the contract method 0x10cd277b.
//
// Solidity: function maxTradingNumber() view returns(uint256)
func (_Contract *ContractCallerSession) MaxTradingNumber() (*big.Int, error) {
	return _Contract.Contract.MaxTradingNumber(&_Contract.CallOpts)
}

// MinDurationTime is a free data retrieval call binding the contract method 0x53e77c5a.
//
// Solidity: function minDurationTime() view returns(uint256)
func (_Contract *ContractCaller) MinDurationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDurationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDurationTime is a free data retrieval call binding the contract method 0x53e77c5a.
//
// Solidity: function minDurationTime() view returns(uint256)
func (_Contract *ContractSession) MinDurationTime() (*big.Int, error) {
	return _Contract.Contract.MinDurationTime(&_Contract.CallOpts)
}

// MinDurationTime is a free data retrieval call binding the contract method 0x53e77c5a.
//
// Solidity: function minDurationTime() view returns(uint256)
func (_Contract *ContractCallerSession) MinDurationTime() (*big.Int, error) {
	return _Contract.Contract.MinDurationTime(&_Contract.CallOpts)
}

// MinStartValue is a free data retrieval call binding the contract method 0x49ad71eb.
//
// Solidity: function minStartValue() view returns(uint256)
func (_Contract *ContractCaller) MinStartValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStartValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStartValue is a free data retrieval call binding the contract method 0x49ad71eb.
//
// Solidity: function minStartValue() view returns(uint256)
func (_Contract *ContractSession) MinStartValue() (*big.Int, error) {
	return _Contract.Contract.MinStartValue(&_Contract.CallOpts)
}

// MinStartValue is a free data retrieval call binding the contract method 0x49ad71eb.
//
// Solidity: function minStartValue() view returns(uint256)
func (_Contract *ContractCallerSession) MinStartValue() (*big.Int, error) {
	return _Contract.Contract.MinStartValue(&_Contract.CallOpts)
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

// OrderIndex is a free data retrieval call binding the contract method 0x864d1260.
//
// Solidity: function orderIndex(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) OrderIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "orderIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIndex is a free data retrieval call binding the contract method 0x864d1260.
//
// Solidity: function orderIndex(uint256 ) view returns(uint256)
func (_Contract *ContractSession) OrderIndex(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OrderIndex(&_Contract.CallOpts, arg0)
}

// OrderIndex is a free data retrieval call binding the contract method 0x864d1260.
//
// Solidity: function orderIndex(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) OrderIndex(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OrderIndex(&_Contract.CallOpts, arg0)
}

// OrderList is a free data retrieval call binding the contract method 0x76f75e7f.
//
// Solidity: function orderList(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) OrderList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "orderList", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderList is a free data retrieval call binding the contract method 0x76f75e7f.
//
// Solidity: function orderList(uint256 ) view returns(uint256)
func (_Contract *ContractSession) OrderList(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OrderList(&_Contract.CallOpts, arg0)
}

// OrderList is a free data retrieval call binding the contract method 0x76f75e7f.
//
// Solidity: function orderList(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) OrderList(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OrderList(&_Contract.CallOpts, arg0)
}

// OrderPairId is a free data retrieval call binding the contract method 0xd9347e42.
//
// Solidity: function orderPairId(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) OrderPairId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "orderPairId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderPairId is a free data retrieval call binding the contract method 0xd9347e42.
//
// Solidity: function orderPairId(uint256 ) view returns(uint256)
func (_Contract *ContractSession) OrderPairId(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OrderPairId(&_Contract.CallOpts, arg0)
}

// OrderPairId is a free data retrieval call binding the contract method 0xd9347e42.
//
// Solidity: function orderPairId(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) OrderPairId(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.OrderPairId(&_Contract.CallOpts, arg0)
}

// OrderStatusList is a free data retrieval call binding the contract method 0x8ada6738.
//
// Solidity: function orderStatusList(uint256 ) view returns(uint8)
func (_Contract *ContractCaller) OrderStatusList(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "orderStatusList", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OrderStatusList is a free data retrieval call binding the contract method 0x8ada6738.
//
// Solidity: function orderStatusList(uint256 ) view returns(uint8)
func (_Contract *ContractSession) OrderStatusList(arg0 *big.Int) (uint8, error) {
	return _Contract.Contract.OrderStatusList(&_Contract.CallOpts, arg0)
}

// OrderStatusList is a free data retrieval call binding the contract method 0x8ada6738.
//
// Solidity: function orderStatusList(uint256 ) view returns(uint8)
func (_Contract *ContractCallerSession) OrderStatusList(arg0 *big.Int) (uint8, error) {
	return _Contract.Contract.OrderStatusList(&_Contract.CallOpts, arg0)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Contract *ContractCaller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Contract *ContractSession) PerformanceFee() (*big.Int, error) {
	return _Contract.Contract.PerformanceFee(&_Contract.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_Contract *ContractCallerSession) PerformanceFee() (*big.Int, error) {
	return _Contract.Contract.PerformanceFee(&_Contract.CallOpts)
}

// ReadinessTime is a free data retrieval call binding the contract method 0x404b12e3.
//
// Solidity: function readinessTime() view returns(uint256)
func (_Contract *ContractCaller) ReadinessTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "readinessTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadinessTime is a free data retrieval call binding the contract method 0x404b12e3.
//
// Solidity: function readinessTime() view returns(uint256)
func (_Contract *ContractSession) ReadinessTime() (*big.Int, error) {
	return _Contract.Contract.ReadinessTime(&_Contract.CallOpts)
}

// ReadinessTime is a free data retrieval call binding the contract method 0x404b12e3.
//
// Solidity: function readinessTime() view returns(uint256)
func (_Contract *ContractCallerSession) ReadinessTime() (*big.Int, error) {
	return _Contract.Contract.ReadinessTime(&_Contract.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Contract *ContractCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Contract *ContractSession) StartTime() (*big.Int, error) {
	return _Contract.Contract.StartTime(&_Contract.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Contract *ContractCallerSession) StartTime() (*big.Int, error) {
	return _Contract.Contract.StartTime(&_Contract.CallOpts)
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

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Contract *ContractCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Contract *ContractSession) TotalAssets() (*big.Int, error) {
	return _Contract.Contract.TotalAssets(&_Contract.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Contract *ContractCallerSession) TotalAssets() (*big.Int, error) {
	return _Contract.Contract.TotalAssets(&_Contract.CallOpts)
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

// UserDepositInfo is a free data retrieval call binding the contract method 0xd4e47031.
//
// Solidity: function userDepositInfo(address ) view returns(uint256)
func (_Contract *ContractCaller) UserDepositInfo(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userDepositInfo", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDepositInfo is a free data retrieval call binding the contract method 0xd4e47031.
//
// Solidity: function userDepositInfo(address ) view returns(uint256)
func (_Contract *ContractSession) UserDepositInfo(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserDepositInfo(&_Contract.CallOpts, arg0)
}

// UserDepositInfo is a free data retrieval call binding the contract method 0xd4e47031.
//
// Solidity: function userDepositInfo(address ) view returns(uint256)
func (_Contract *ContractCallerSession) UserDepositInfo(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserDepositInfo(&_Contract.CallOpts, arg0)
}

// UserMaxDepositAmount is a free data retrieval call binding the contract method 0xfcbdbd6d.
//
// Solidity: function userMaxDepositAmount() view returns(uint256)
func (_Contract *ContractCaller) UserMaxDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userMaxDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxDepositAmount is a free data retrieval call binding the contract method 0xfcbdbd6d.
//
// Solidity: function userMaxDepositAmount() view returns(uint256)
func (_Contract *ContractSession) UserMaxDepositAmount() (*big.Int, error) {
	return _Contract.Contract.UserMaxDepositAmount(&_Contract.CallOpts)
}

// UserMaxDepositAmount is a free data retrieval call binding the contract method 0xfcbdbd6d.
//
// Solidity: function userMaxDepositAmount() view returns(uint256)
func (_Contract *ContractCallerSession) UserMaxDepositAmount() (*big.Int, error) {
	return _Contract.Contract.UserMaxDepositAmount(&_Contract.CallOpts)
}

// UserMinDepositAmount is a free data retrieval call binding the contract method 0x4b5a33dd.
//
// Solidity: function userMinDepositAmount() view returns(uint256)
func (_Contract *ContractCaller) UserMinDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userMinDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMinDepositAmount is a free data retrieval call binding the contract method 0x4b5a33dd.
//
// Solidity: function userMinDepositAmount() view returns(uint256)
func (_Contract *ContractSession) UserMinDepositAmount() (*big.Int, error) {
	return _Contract.Contract.UserMinDepositAmount(&_Contract.CallOpts)
}

// UserMinDepositAmount is a free data retrieval call binding the contract method 0x4b5a33dd.
//
// Solidity: function userMinDepositAmount() view returns(uint256)
func (_Contract *ContractCallerSession) UserMinDepositAmount() (*big.Int, error) {
	return _Contract.Contract.UserMinDepositAmount(&_Contract.CallOpts)
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

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Contract *ContractCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Contract *ContractSession) VaultManager() (common.Address, error) {
	return _Contract.Contract.VaultManager(&_Contract.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Contract *ContractCallerSession) VaultManager() (common.Address, error) {
	return _Contract.Contract.VaultManager(&_Contract.CallOpts)
}

// VaultSocialInfo is a free data retrieval call binding the contract method 0xdf696f43.
//
// Solidity: function vaultSocialInfo() view returns(string vaultName, string vaultTwitter, string vaultPhoto, string vaultStrategy)
func (_Contract *ContractCaller) VaultSocialInfo(opts *bind.CallOpts) (struct {
	VaultName     string
	VaultTwitter  string
	VaultPhoto    string
	VaultStrategy string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultSocialInfo")

	outstruct := new(struct {
		VaultName     string
		VaultTwitter  string
		VaultPhoto    string
		VaultStrategy string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VaultName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VaultTwitter = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VaultPhoto = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.VaultStrategy = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// VaultSocialInfo is a free data retrieval call binding the contract method 0xdf696f43.
//
// Solidity: function vaultSocialInfo() view returns(string vaultName, string vaultTwitter, string vaultPhoto, string vaultStrategy)
func (_Contract *ContractSession) VaultSocialInfo() (struct {
	VaultName     string
	VaultTwitter  string
	VaultPhoto    string
	VaultStrategy string
}, error) {
	return _Contract.Contract.VaultSocialInfo(&_Contract.CallOpts)
}

// VaultSocialInfo is a free data retrieval call binding the contract method 0xdf696f43.
//
// Solidity: function vaultSocialInfo() view returns(string vaultName, string vaultTwitter, string vaultPhoto, string vaultStrategy)
func (_Contract *ContractCallerSession) VaultSocialInfo() (struct {
	VaultName     string
	VaultTwitter  string
	VaultPhoto    string
	VaultStrategy string
}, error) {
	return _Contract.Contract.VaultSocialInfo(&_Contract.CallOpts)
}

// VaultStatus is a free data retrieval call binding the contract method 0x24a8166e.
//
// Solidity: function vaultStatus() view returns(bool)
func (_Contract *ContractCaller) VaultStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "vaultStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultStatus is a free data retrieval call binding the contract method 0x24a8166e.
//
// Solidity: function vaultStatus() view returns(bool)
func (_Contract *ContractSession) VaultStatus() (bool, error) {
	return _Contract.Contract.VaultStatus(&_Contract.CallOpts)
}

// VaultStatus is a free data retrieval call binding the contract method 0x24a8166e.
//
// Solidity: function vaultStatus() view returns(bool)
func (_Contract *ContractCallerSession) VaultStatus() (bool, error) {
	return _Contract.Contract.VaultStatus(&_Contract.CallOpts)
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

// CloseVault is a paid mutator transaction binding the contract method 0x18976fa2.
//
// Solidity: function closeVault() returns()
func (_Contract *ContractTransactor) CloseVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "closeVault")
}

// CloseVault is a paid mutator transaction binding the contract method 0x18976fa2.
//
// Solidity: function closeVault() returns()
func (_Contract *ContractSession) CloseVault() (*types.Transaction, error) {
	return _Contract.Contract.CloseVault(&_Contract.TransactOpts)
}

// CloseVault is a paid mutator transaction binding the contract method 0x18976fa2.
//
// Solidity: function closeVault() returns()
func (_Contract *ContractTransactorSession) CloseVault() (*types.Transaction, error) {
	return _Contract.Contract.CloseVault(&_Contract.TransactOpts)
}

// CloseVaultFactroy is a paid mutator transaction binding the contract method 0xa63c3356.
//
// Solidity: function closeVaultFactroy() returns()
func (_Contract *ContractTransactor) CloseVaultFactroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "closeVaultFactroy")
}

// CloseVaultFactroy is a paid mutator transaction binding the contract method 0xa63c3356.
//
// Solidity: function closeVaultFactroy() returns()
func (_Contract *ContractSession) CloseVaultFactroy() (*types.Transaction, error) {
	return _Contract.Contract.CloseVaultFactroy(&_Contract.TransactOpts)
}

// CloseVaultFactroy is a paid mutator transaction binding the contract method 0xa63c3356.
//
// Solidity: function closeVaultFactroy() returns()
func (_Contract *ContractTransactorSession) CloseVaultFactroy() (*types.Transaction, error) {
	return _Contract.Contract.CloseVaultFactroy(&_Contract.TransactOpts)
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

// Deposit is a paid mutator transaction binding the contract method 0xffab8573.
//
// Solidity: function deposit(uint256 assets, bytes[] updateData) returns(uint256)
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", assets, updateData)
}

// Deposit is a paid mutator transaction binding the contract method 0xffab8573.
//
// Solidity: function deposit(uint256 assets, bytes[] updateData) returns(uint256)
func (_Contract *ContractSession) Deposit(assets *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, assets, updateData)
}

// Deposit is a paid mutator transaction binding the contract method 0xffab8573.
//
// Solidity: function deposit(uint256 assets, bytes[] updateData) returns(uint256)
func (_Contract *ContractTransactorSession) Deposit(assets *big.Int, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, assets, updateData)
}

// GetTradingTotalValue is a paid mutator transaction binding the contract method 0x2596d331.
//
// Solidity: function getTradingTotalValue(bytes[] updateData) returns(uint256)
func (_Contract *ContractTransactor) GetTradingTotalValue(opts *bind.TransactOpts, updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getTradingTotalValue", updateData)
}

// GetTradingTotalValue is a paid mutator transaction binding the contract method 0x2596d331.
//
// Solidity: function getTradingTotalValue(bytes[] updateData) returns(uint256)
func (_Contract *ContractSession) GetTradingTotalValue(updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.GetTradingTotalValue(&_Contract.TransactOpts, updateData)
}

// GetTradingTotalValue is a paid mutator transaction binding the contract method 0x2596d331.
//
// Solidity: function getTradingTotalValue(bytes[] updateData) returns(uint256)
func (_Contract *ContractTransactorSession) GetTradingTotalValue(updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.GetTradingTotalValue(&_Contract.TransactOpts, updateData)
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

// Initialize is a paid mutator transaction binding the contract method 0x11d9dfa6.
//
// Solidity: function initialize(string _name, string _symbol, address _assetToken, address _manager, address _factory, address _trading, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _vaultInfo, (string,string,string,string) _socialInfo) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _assetToken common.Address, _manager common.Address, _factory common.Address, _trading common.Address, _vaultInfo ICopyTradingVaultVaultInfo, _socialInfo ICopyTradingVaultVaultSocialInfo) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _name, _symbol, _assetToken, _manager, _factory, _trading, _vaultInfo, _socialInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x11d9dfa6.
//
// Solidity: function initialize(string _name, string _symbol, address _assetToken, address _manager, address _factory, address _trading, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _vaultInfo, (string,string,string,string) _socialInfo) returns()
func (_Contract *ContractSession) Initialize(_name string, _symbol string, _assetToken common.Address, _manager common.Address, _factory common.Address, _trading common.Address, _vaultInfo ICopyTradingVaultVaultInfo, _socialInfo ICopyTradingVaultVaultSocialInfo) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _assetToken, _manager, _factory, _trading, _vaultInfo, _socialInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x11d9dfa6.
//
// Solidity: function initialize(string _name, string _symbol, address _assetToken, address _manager, address _factory, address _trading, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _vaultInfo, (string,string,string,string) _socialInfo) returns()
func (_Contract *ContractTransactorSession) Initialize(_name string, _symbol string, _assetToken common.Address, _manager common.Address, _factory common.Address, _trading common.Address, _vaultInfo ICopyTradingVaultVaultInfo, _socialInfo ICopyTradingVaultVaultSocialInfo) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _assetToken, _manager, _factory, _trading, _vaultInfo, _socialInfo)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares) returns(uint256)
func (_Contract *ContractTransactor) Mint(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mint", shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares) returns(uint256)
func (_Contract *ContractSession) Mint(shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares) returns(uint256)
func (_Contract *ContractTransactorSession) Mint(shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, shares)
}

// OnTradeClose is a paid mutator transaction binding the contract method 0x657409e2.
//
// Solidity: function onTradeClose(uint256 _orderId, uint256 , bool _wasExecuted, uint8 ) returns()
func (_Contract *ContractTransactor) OnTradeClose(opts *bind.TransactOpts, _orderId *big.Int, arg1 *big.Int, _wasExecuted bool, arg3 uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "onTradeClose", _orderId, arg1, _wasExecuted, arg3)
}

// OnTradeClose is a paid mutator transaction binding the contract method 0x657409e2.
//
// Solidity: function onTradeClose(uint256 _orderId, uint256 , bool _wasExecuted, uint8 ) returns()
func (_Contract *ContractSession) OnTradeClose(_orderId *big.Int, arg1 *big.Int, _wasExecuted bool, arg3 uint8) (*types.Transaction, error) {
	return _Contract.Contract.OnTradeClose(&_Contract.TransactOpts, _orderId, arg1, _wasExecuted, arg3)
}

// OnTradeClose is a paid mutator transaction binding the contract method 0x657409e2.
//
// Solidity: function onTradeClose(uint256 _orderId, uint256 , bool _wasExecuted, uint8 ) returns()
func (_Contract *ContractTransactorSession) OnTradeClose(_orderId *big.Int, arg1 *big.Int, _wasExecuted bool, arg3 uint8) (*types.Transaction, error) {
	return _Contract.Contract.OnTradeClose(&_Contract.TransactOpts, _orderId, arg1, _wasExecuted, arg3)
}

// OnTradeOpen is a paid mutator transaction binding the contract method 0xb2a11f85.
//
// Solidity: function onTradeOpen(uint256 _orderId, bool _wasExecuted) returns()
func (_Contract *ContractTransactor) OnTradeOpen(opts *bind.TransactOpts, _orderId *big.Int, _wasExecuted bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "onTradeOpen", _orderId, _wasExecuted)
}

// OnTradeOpen is a paid mutator transaction binding the contract method 0xb2a11f85.
//
// Solidity: function onTradeOpen(uint256 _orderId, bool _wasExecuted) returns()
func (_Contract *ContractSession) OnTradeOpen(_orderId *big.Int, _wasExecuted bool) (*types.Transaction, error) {
	return _Contract.Contract.OnTradeOpen(&_Contract.TransactOpts, _orderId, _wasExecuted)
}

// OnTradeOpen is a paid mutator transaction binding the contract method 0xb2a11f85.
//
// Solidity: function onTradeOpen(uint256 _orderId, bool _wasExecuted) returns()
func (_Contract *ContractTransactorSession) OnTradeOpen(_orderId *big.Int, _wasExecuted bool) (*types.Transaction, error) {
	return _Contract.Contract.OnTradeOpen(&_Contract.TransactOpts, _orderId, _wasExecuted)
}

// RequestClose is a paid mutator transaction binding the contract method 0xc04398ee.
//
// Solidity: function requestClose(uint256 _orderId, uint256 _closeMargin, address ) payable returns()
func (_Contract *ContractTransactor) RequestClose(opts *bind.TransactOpts, _orderId *big.Int, _closeMargin *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestClose", _orderId, _closeMargin, arg2)
}

// RequestClose is a paid mutator transaction binding the contract method 0xc04398ee.
//
// Solidity: function requestClose(uint256 _orderId, uint256 _closeMargin, address ) payable returns()
func (_Contract *ContractSession) RequestClose(_orderId *big.Int, _closeMargin *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestClose(&_Contract.TransactOpts, _orderId, _closeMargin, arg2)
}

// RequestClose is a paid mutator transaction binding the contract method 0xc04398ee.
//
// Solidity: function requestClose(uint256 _orderId, uint256 _closeMargin, address ) payable returns()
func (_Contract *ContractTransactorSession) RequestClose(_orderId *big.Int, _closeMargin *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestClose(&_Contract.TransactOpts, _orderId, _closeMargin, arg2)
}

// RequestOpenLimit is a paid mutator transaction binding the contract method 0xf750704f.
//
// Solidity: function requestOpenLimit(uint8 _t, (address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 limit, address ) payable returns()
func (_Contract *ContractTransactor) RequestOpenLimit(opts *bind.TransactOpts, _t uint8, _base TradeBase, updateData [][]byte, limit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestOpenLimit", _t, _base, updateData, limit, arg4)
}

// RequestOpenLimit is a paid mutator transaction binding the contract method 0xf750704f.
//
// Solidity: function requestOpenLimit(uint8 _t, (address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 limit, address ) payable returns()
func (_Contract *ContractSession) RequestOpenLimit(_t uint8, _base TradeBase, updateData [][]byte, limit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenLimit(&_Contract.TransactOpts, _t, _base, updateData, limit, arg4)
}

// RequestOpenLimit is a paid mutator transaction binding the contract method 0xf750704f.
//
// Solidity: function requestOpenLimit(uint8 _t, (address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 limit, address ) payable returns()
func (_Contract *ContractTransactorSession) RequestOpenLimit(_t uint8, _base TradeBase, updateData [][]byte, limit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenLimit(&_Contract.TransactOpts, _t, _base, updateData, limit, arg4)
}

// RequestOpenMarket is a paid mutator transaction binding the contract method 0x8f92dab8.
//
// Solidity: function requestOpenMarket((address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 slippage, address ) payable returns()
func (_Contract *ContractTransactor) RequestOpenMarket(opts *bind.TransactOpts, _base TradeBase, updateData [][]byte, slippage *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestOpenMarket", _base, updateData, slippage, arg3)
}

// RequestOpenMarket is a paid mutator transaction binding the contract method 0x8f92dab8.
//
// Solidity: function requestOpenMarket((address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 slippage, address ) payable returns()
func (_Contract *ContractSession) RequestOpenMarket(_base TradeBase, updateData [][]byte, slippage *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenMarket(&_Contract.TransactOpts, _base, updateData, slippage, arg3)
}

// RequestOpenMarket is a paid mutator transaction binding the contract method 0x8f92dab8.
//
// Solidity: function requestOpenMarket((address,uint256,uint256,bool,uint256,uint256,uint256) _base, bytes[] updateData, uint256 slippage, address ) payable returns()
func (_Contract *ContractTransactorSession) RequestOpenMarket(_base TradeBase, updateData [][]byte, slippage *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RequestOpenMarket(&_Contract.TransactOpts, _base, updateData, slippage, arg3)
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

// UpdateMargin is a paid mutator transaction binding the contract method 0xfd77f87f.
//
// Solidity: function updateMargin(uint256 _orderId, uint256 _amount, bool _isAdd, bytes[] _updateData) returns()
func (_Contract *ContractTransactor) UpdateMargin(opts *bind.TransactOpts, _orderId *big.Int, _amount *big.Int, _isAdd bool, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMargin", _orderId, _amount, _isAdd, _updateData)
}

// UpdateMargin is a paid mutator transaction binding the contract method 0xfd77f87f.
//
// Solidity: function updateMargin(uint256 _orderId, uint256 _amount, bool _isAdd, bytes[] _updateData) returns()
func (_Contract *ContractSession) UpdateMargin(_orderId *big.Int, _amount *big.Int, _isAdd bool, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMargin(&_Contract.TransactOpts, _orderId, _amount, _isAdd, _updateData)
}

// UpdateMargin is a paid mutator transaction binding the contract method 0xfd77f87f.
//
// Solidity: function updateMargin(uint256 _orderId, uint256 _amount, bool _isAdd, bytes[] _updateData) returns()
func (_Contract *ContractTransactorSession) UpdateMargin(_orderId *big.Int, _amount *big.Int, _isAdd bool, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMargin(&_Contract.TransactOpts, _orderId, _amount, _isAdd, _updateData)
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

// UpdateSocialInfo is a paid mutator transaction binding the contract method 0x268526b1.
//
// Solidity: function updateSocialInfo(string _twitter, string _photo) returns()
func (_Contract *ContractTransactor) UpdateSocialInfo(opts *bind.TransactOpts, _twitter string, _photo string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSocialInfo", _twitter, _photo)
}

// UpdateSocialInfo is a paid mutator transaction binding the contract method 0x268526b1.
//
// Solidity: function updateSocialInfo(string _twitter, string _photo) returns()
func (_Contract *ContractSession) UpdateSocialInfo(_twitter string, _photo string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSocialInfo(&_Contract.TransactOpts, _twitter, _photo)
}

// UpdateSocialInfo is a paid mutator transaction binding the contract method 0x268526b1.
//
// Solidity: function updateSocialInfo(string _twitter, string _photo) returns()
func (_Contract *ContractTransactorSession) UpdateSocialInfo(_twitter string, _photo string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSocialInfo(&_Contract.TransactOpts, _twitter, _photo)
}

// UpdateStrategy is a paid mutator transaction binding the contract method 0x73031b75.
//
// Solidity: function updateStrategy(string _strategy) returns()
func (_Contract *ContractTransactor) UpdateStrategy(opts *bind.TransactOpts, _strategy string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateStrategy", _strategy)
}

// UpdateStrategy is a paid mutator transaction binding the contract method 0x73031b75.
//
// Solidity: function updateStrategy(string _strategy) returns()
func (_Contract *ContractSession) UpdateStrategy(_strategy string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStrategy(&_Contract.TransactOpts, _strategy)
}

// UpdateStrategy is a paid mutator transaction binding the contract method 0x73031b75.
//
// Solidity: function updateStrategy(string _strategy) returns()
func (_Contract *ContractTransactorSession) UpdateStrategy(_strategy string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStrategy(&_Contract.TransactOpts, _strategy)
}

// UpdateTPAndSL is a paid mutator transaction binding the contract method 0x04c7ec71.
//
// Solidity: function updateTPAndSL(uint256 orderId, uint256 tp, uint256 sl, bytes[] _updateData) returns()
func (_Contract *ContractTransactor) UpdateTPAndSL(opts *bind.TransactOpts, orderId *big.Int, tp *big.Int, sl *big.Int, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateTPAndSL", orderId, tp, sl, _updateData)
}

// UpdateTPAndSL is a paid mutator transaction binding the contract method 0x04c7ec71.
//
// Solidity: function updateTPAndSL(uint256 orderId, uint256 tp, uint256 sl, bytes[] _updateData) returns()
func (_Contract *ContractSession) UpdateTPAndSL(orderId *big.Int, tp *big.Int, sl *big.Int, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTPAndSL(&_Contract.TransactOpts, orderId, tp, sl, _updateData)
}

// UpdateTPAndSL is a paid mutator transaction binding the contract method 0x04c7ec71.
//
// Solidity: function updateTPAndSL(uint256 orderId, uint256 tp, uint256 sl, bytes[] _updateData) returns()
func (_Contract *ContractTransactorSession) UpdateTPAndSL(orderId *big.Int, tp *big.Int, sl *big.Int, _updateData [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTPAndSL(&_Contract.TransactOpts, orderId, tp, sl, _updateData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
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
	Index   *big.Int
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelOpen is a free log retrieval operation binding the contract event 0x70aae6781f454c2184c7f647b5bc336fc425da19d7081dab4ed4e3c92c09d98b.
//
// Solidity: event CancelOpen(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) FilterCancelOpen(opts *bind.FilterOpts, orderId []*big.Int, index []*big.Int) (*ContractCancelOpenIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CancelOpen", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &ContractCancelOpenIterator{contract: _Contract.contract, event: "CancelOpen", logs: logs, sub: sub}, nil
}

// WatchCancelOpen is a free log subscription operation binding the contract event 0x70aae6781f454c2184c7f647b5bc336fc425da19d7081dab4ed4e3c92c09d98b.
//
// Solidity: event CancelOpen(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) WatchCancelOpen(opts *bind.WatchOpts, sink chan<- *ContractCancelOpen, orderId []*big.Int, index []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CancelOpen", orderIdRule, indexRule)
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

// ParseCancelOpen is a log parse operation binding the contract event 0x70aae6781f454c2184c7f647b5bc336fc425da19d7081dab4ed4e3c92c09d98b.
//
// Solidity: event CancelOpen(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) ParseCancelOpen(log types.Log) (*ContractCancelOpen, error) {
	event := new(ContractCancelOpen)
	if err := _Contract.contract.UnpackLog(event, "CancelOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCloseOnTradeIterator is returned from FilterCloseOnTrade and is used to iterate over the raw logs and unpacked data for CloseOnTrade events raised by the Contract contract.
type ContractCloseOnTradeIterator struct {
	Event *ContractCloseOnTrade // Event containing the contract specifics and raw log

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
func (it *ContractCloseOnTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCloseOnTrade)
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
		it.Event = new(ContractCloseOnTrade)
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
func (it *ContractCloseOnTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCloseOnTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCloseOnTrade represents a CloseOnTrade event raised by the Contract contract.
type ContractCloseOnTrade struct {
	OrderId *big.Int
	Index   *big.Int
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCloseOnTrade is a free log retrieval operation binding the contract event 0x643a81b584fafd2fabb5e05c23f194090fdb7be856ac728d1c61ad5bc525026e.
//
// Solidity: event CloseOnTrade(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) FilterCloseOnTrade(opts *bind.FilterOpts, orderId []*big.Int, index []*big.Int) (*ContractCloseOnTradeIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CloseOnTrade", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &ContractCloseOnTradeIterator{contract: _Contract.contract, event: "CloseOnTrade", logs: logs, sub: sub}, nil
}

// WatchCloseOnTrade is a free log subscription operation binding the contract event 0x643a81b584fafd2fabb5e05c23f194090fdb7be856ac728d1c61ad5bc525026e.
//
// Solidity: event CloseOnTrade(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) WatchCloseOnTrade(opts *bind.WatchOpts, sink chan<- *ContractCloseOnTrade, orderId []*big.Int, index []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CloseOnTrade", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCloseOnTrade)
				if err := _Contract.contract.UnpackLog(event, "CloseOnTrade", log); err != nil {
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

// ParseCloseOnTrade is a log parse operation binding the contract event 0x643a81b584fafd2fabb5e05c23f194090fdb7be856ac728d1c61ad5bc525026e.
//
// Solidity: event CloseOnTrade(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) ParseCloseOnTrade(log types.Log) (*ContractCloseOnTrade, error) {
	event := new(ContractCloseOnTrade)
	if err := _Contract.contract.UnpackLog(event, "CloseOnTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCloseVaultIterator is returned from FilterCloseVault and is used to iterate over the raw logs and unpacked data for CloseVault events raised by the Contract contract.
type ContractCloseVaultIterator struct {
	Event *ContractCloseVault // Event containing the contract specifics and raw log

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
func (it *ContractCloseVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCloseVault)
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
		it.Event = new(ContractCloseVault)
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
func (it *ContractCloseVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCloseVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCloseVault represents a CloseVault event raised by the Contract contract.
type ContractCloseVault struct {
	CloseAddress common.Address
	VaultAddress common.Address
	CloseTime    *big.Int
	VaultFee     *big.Int
	Brokerage    *big.Int
	VaultStatus  bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCloseVault is a free log retrieval operation binding the contract event 0xc2e1998855d430f0d090f38d8b2b0c89d224c8061aa97c16f0723817b55ba600.
//
// Solidity: event CloseVault(address indexed closeAddress, address indexed vaultAddress, uint256 closeTime, uint256 vaultFee, uint256 brokerage, bool vaultStatus)
func (_Contract *ContractFilterer) FilterCloseVault(opts *bind.FilterOpts, closeAddress []common.Address, vaultAddress []common.Address) (*ContractCloseVaultIterator, error) {

	var closeAddressRule []interface{}
	for _, closeAddressItem := range closeAddress {
		closeAddressRule = append(closeAddressRule, closeAddressItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CloseVault", closeAddressRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractCloseVaultIterator{contract: _Contract.contract, event: "CloseVault", logs: logs, sub: sub}, nil
}

// WatchCloseVault is a free log subscription operation binding the contract event 0xc2e1998855d430f0d090f38d8b2b0c89d224c8061aa97c16f0723817b55ba600.
//
// Solidity: event CloseVault(address indexed closeAddress, address indexed vaultAddress, uint256 closeTime, uint256 vaultFee, uint256 brokerage, bool vaultStatus)
func (_Contract *ContractFilterer) WatchCloseVault(opts *bind.WatchOpts, sink chan<- *ContractCloseVault, closeAddress []common.Address, vaultAddress []common.Address) (event.Subscription, error) {

	var closeAddressRule []interface{}
	for _, closeAddressItem := range closeAddress {
		closeAddressRule = append(closeAddressRule, closeAddressItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CloseVault", closeAddressRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCloseVault)
				if err := _Contract.contract.UnpackLog(event, "CloseVault", log); err != nil {
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

// ParseCloseVault is a log parse operation binding the contract event 0xc2e1998855d430f0d090f38d8b2b0c89d224c8061aa97c16f0723817b55ba600.
//
// Solidity: event CloseVault(address indexed closeAddress, address indexed vaultAddress, uint256 closeTime, uint256 vaultFee, uint256 brokerage, bool vaultStatus)
func (_Contract *ContractFilterer) ParseCloseVault(log types.Log) (*ContractCloseVault, error) {
	event := new(ContractCloseVault)
	if err := _Contract.contract.UnpackLog(event, "CloseVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCloseVaultFactroyIterator is returned from FilterCloseVaultFactroy and is used to iterate over the raw logs and unpacked data for CloseVaultFactroy events raised by the Contract contract.
type ContractCloseVaultFactroyIterator struct {
	Event *ContractCloseVaultFactroy // Event containing the contract specifics and raw log

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
func (it *ContractCloseVaultFactroyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCloseVaultFactroy)
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
		it.Event = new(ContractCloseVaultFactroy)
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
func (it *ContractCloseVaultFactroyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCloseVaultFactroyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCloseVaultFactroy represents a CloseVaultFactroy event raised by the Contract contract.
type ContractCloseVaultFactroy struct {
	CloseAddress common.Address
	VaultAddress common.Address
	CloseTime    *big.Int
	VaultFee     *big.Int
	Brokerage    *big.Int
	VaultStatus  bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCloseVaultFactroy is a free log retrieval operation binding the contract event 0xf0abc1167bb421bdd727447d809e12b44506dbe42b00a6d715120f872298c8ab.
//
// Solidity: event CloseVaultFactroy(address indexed closeAddress, address indexed vaultAddress, uint256 closeTime, uint256 vaultFee, uint256 brokerage, bool vaultStatus)
func (_Contract *ContractFilterer) FilterCloseVaultFactroy(opts *bind.FilterOpts, closeAddress []common.Address, vaultAddress []common.Address) (*ContractCloseVaultFactroyIterator, error) {

	var closeAddressRule []interface{}
	for _, closeAddressItem := range closeAddress {
		closeAddressRule = append(closeAddressRule, closeAddressItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CloseVaultFactroy", closeAddressRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractCloseVaultFactroyIterator{contract: _Contract.contract, event: "CloseVaultFactroy", logs: logs, sub: sub}, nil
}

// WatchCloseVaultFactroy is a free log subscription operation binding the contract event 0xf0abc1167bb421bdd727447d809e12b44506dbe42b00a6d715120f872298c8ab.
//
// Solidity: event CloseVaultFactroy(address indexed closeAddress, address indexed vaultAddress, uint256 closeTime, uint256 vaultFee, uint256 brokerage, bool vaultStatus)
func (_Contract *ContractFilterer) WatchCloseVaultFactroy(opts *bind.WatchOpts, sink chan<- *ContractCloseVaultFactroy, closeAddress []common.Address, vaultAddress []common.Address) (event.Subscription, error) {

	var closeAddressRule []interface{}
	for _, closeAddressItem := range closeAddress {
		closeAddressRule = append(closeAddressRule, closeAddressItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CloseVaultFactroy", closeAddressRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCloseVaultFactroy)
				if err := _Contract.contract.UnpackLog(event, "CloseVaultFactroy", log); err != nil {
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

// ParseCloseVaultFactroy is a log parse operation binding the contract event 0xf0abc1167bb421bdd727447d809e12b44506dbe42b00a6d715120f872298c8ab.
//
// Solidity: event CloseVaultFactroy(address indexed closeAddress, address indexed vaultAddress, uint256 closeTime, uint256 vaultFee, uint256 brokerage, bool vaultStatus)
func (_Contract *ContractFilterer) ParseCloseVaultFactroy(log types.Log) (*ContractCloseVaultFactroy, error) {
	event := new(ContractCloseVaultFactroy)
	if err := _Contract.contract.UnpackLog(event, "CloseVaultFactroy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Contract contract.
type ContractDepositIterator struct {
	Event *ContractDeposit // Event containing the contract specifics and raw log

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
func (it *ContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeposit)
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
		it.Event = new(ContractDeposit)
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
func (it *ContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeposit represents a Deposit event raised by the Contract contract.
type ContractDeposit struct {
	Sender common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 indexed assets, uint256 indexed shares)
func (_Contract *ContractFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, assets []*big.Int, shares []*big.Int) (*ContractDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetsRule []interface{}
	for _, assetsItem := range assets {
		assetsRule = append(assetsRule, assetsItem)
	}
	var sharesRule []interface{}
	for _, sharesItem := range shares {
		sharesRule = append(sharesRule, sharesItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Deposit", senderRule, assetsRule, sharesRule)
	if err != nil {
		return nil, err
	}
	return &ContractDepositIterator{contract: _Contract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 indexed assets, uint256 indexed shares)
func (_Contract *ContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractDeposit, sender []common.Address, assets []*big.Int, shares []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetsRule []interface{}
	for _, assetsItem := range assets {
		assetsRule = append(assetsRule, assetsItem)
	}
	var sharesRule []interface{}
	for _, sharesItem := range shares {
		sharesRule = append(sharesRule, sharesItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Deposit", senderRule, assetsRule, sharesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeposit)
				if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 indexed assets, uint256 indexed shares)
func (_Contract *ContractFilterer) ParseDeposit(log types.Log) (*ContractDeposit, error) {
	event := new(ContractDeposit)
	if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFollowerIterator is returned from FilterFollower and is used to iterate over the raw logs and unpacked data for Follower events raised by the Contract contract.
type ContractFollowerIterator struct {
	Event *ContractFollower // Event containing the contract specifics and raw log

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
func (it *ContractFollowerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFollower)
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
		it.Event = new(ContractFollower)
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
func (it *ContractFollowerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFollowerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFollower represents a Follower event raised by the Contract contract.
type ContractFollower struct {
	Followers *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollower is a free log retrieval operation binding the contract event 0x27a4ec1ae6c7e8bb681a587eabc128c2cda3f6052567885233f83dbf3f49e904.
//
// Solidity: event Follower(uint256 followers)
func (_Contract *ContractFilterer) FilterFollower(opts *bind.FilterOpts) (*ContractFollowerIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Follower")
	if err != nil {
		return nil, err
	}
	return &ContractFollowerIterator{contract: _Contract.contract, event: "Follower", logs: logs, sub: sub}, nil
}

// WatchFollower is a free log subscription operation binding the contract event 0x27a4ec1ae6c7e8bb681a587eabc128c2cda3f6052567885233f83dbf3f49e904.
//
// Solidity: event Follower(uint256 followers)
func (_Contract *ContractFilterer) WatchFollower(opts *bind.WatchOpts, sink chan<- *ContractFollower) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Follower")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFollower)
				if err := _Contract.contract.UnpackLog(event, "Follower", log); err != nil {
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

// ParseFollower is a log parse operation binding the contract event 0x27a4ec1ae6c7e8bb681a587eabc128c2cda3f6052567885233f83dbf3f49e904.
//
// Solidity: event Follower(uint256 followers)
func (_Contract *ContractFilterer) ParseFollower(log types.Log) (*ContractFollower, error) {
	event := new(ContractFollower)
	if err := _Contract.contract.UnpackLog(event, "Follower", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Contract contract.
type ContractInitializeIterator struct {
	Event *ContractInitialize // Event containing the contract specifics and raw log

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
func (it *ContractInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialize)
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
		it.Event = new(ContractInitialize)
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
func (it *ContractInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialize represents a Initialize event raised by the Contract contract.
type ContractInitialize struct {
	PerformanceFee       *big.Int
	ReadinessTime        *big.Int
	Duration             *big.Int
	UserMinDepositAmount *big.Int
	MaxCapValue          *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x485315dc1f29f580769d5cbd5ad09f9f58438ccc423fc55786693af76dfd0607.
//
// Solidity: event Initialize(uint256 performanceFee, uint256 readinessTime, uint256 duration, uint256 userMinDepositAmount, uint256 maxCapValue)
func (_Contract *ContractFilterer) FilterInitialize(opts *bind.FilterOpts) (*ContractInitializeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &ContractInitializeIterator{contract: _Contract.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x485315dc1f29f580769d5cbd5ad09f9f58438ccc423fc55786693af76dfd0607.
//
// Solidity: event Initialize(uint256 performanceFee, uint256 readinessTime, uint256 duration, uint256 userMinDepositAmount, uint256 maxCapValue)
func (_Contract *ContractFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *ContractInitialize) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialize)
				if err := _Contract.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x485315dc1f29f580769d5cbd5ad09f9f58438ccc423fc55786693af76dfd0607.
//
// Solidity: event Initialize(uint256 performanceFee, uint256 readinessTime, uint256 duration, uint256 userMinDepositAmount, uint256 maxCapValue)
func (_Contract *ContractFilterer) ParseInitialize(log types.Log) (*ContractInitialize, error) {
	event := new(ContractInitialize)
	if err := _Contract.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ContractNewSocialInfoIterator is returned from FilterNewSocialInfo and is used to iterate over the raw logs and unpacked data for NewSocialInfo events raised by the Contract contract.
type ContractNewSocialInfoIterator struct {
	Event *ContractNewSocialInfo // Event containing the contract specifics and raw log

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
func (it *ContractNewSocialInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewSocialInfo)
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
		it.Event = new(ContractNewSocialInfo)
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
func (it *ContractNewSocialInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewSocialInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewSocialInfo represents a NewSocialInfo event raised by the Contract contract.
type ContractNewSocialInfo struct {
	Caller  common.Address
	Twitter string
	Photo   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSocialInfo is a free log retrieval operation binding the contract event 0x8a05696680470bcb82e9ac8f3389a293c1b1b4ebdca13d487bee501aad7e256b.
//
// Solidity: event NewSocialInfo(address caller, string twitter, string photo)
func (_Contract *ContractFilterer) FilterNewSocialInfo(opts *bind.FilterOpts) (*ContractNewSocialInfoIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewSocialInfo")
	if err != nil {
		return nil, err
	}
	return &ContractNewSocialInfoIterator{contract: _Contract.contract, event: "NewSocialInfo", logs: logs, sub: sub}, nil
}

// WatchNewSocialInfo is a free log subscription operation binding the contract event 0x8a05696680470bcb82e9ac8f3389a293c1b1b4ebdca13d487bee501aad7e256b.
//
// Solidity: event NewSocialInfo(address caller, string twitter, string photo)
func (_Contract *ContractFilterer) WatchNewSocialInfo(opts *bind.WatchOpts, sink chan<- *ContractNewSocialInfo) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewSocialInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewSocialInfo)
				if err := _Contract.contract.UnpackLog(event, "NewSocialInfo", log); err != nil {
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

// ParseNewSocialInfo is a log parse operation binding the contract event 0x8a05696680470bcb82e9ac8f3389a293c1b1b4ebdca13d487bee501aad7e256b.
//
// Solidity: event NewSocialInfo(address caller, string twitter, string photo)
func (_Contract *ContractFilterer) ParseNewSocialInfo(log types.Log) (*ContractNewSocialInfo, error) {
	event := new(ContractNewSocialInfo)
	if err := _Contract.contract.UnpackLog(event, "NewSocialInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewStrategyIterator is returned from FilterNewStrategy and is used to iterate over the raw logs and unpacked data for NewStrategy events raised by the Contract contract.
type ContractNewStrategyIterator struct {
	Event *ContractNewStrategy // Event containing the contract specifics and raw log

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
func (it *ContractNewStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewStrategy)
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
		it.Event = new(ContractNewStrategy)
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
func (it *ContractNewStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewStrategy represents a NewStrategy event raised by the Contract contract.
type ContractNewStrategy struct {
	Caller   common.Address
	Strategy string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewStrategy is a free log retrieval operation binding the contract event 0x54d670c4fd6f2a963fe334d441e020b3c97415d0f0bbae5f4812a34a304ccfc5.
//
// Solidity: event NewStrategy(address caller, string strategy)
func (_Contract *ContractFilterer) FilterNewStrategy(opts *bind.FilterOpts) (*ContractNewStrategyIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewStrategy")
	if err != nil {
		return nil, err
	}
	return &ContractNewStrategyIterator{contract: _Contract.contract, event: "NewStrategy", logs: logs, sub: sub}, nil
}

// WatchNewStrategy is a free log subscription operation binding the contract event 0x54d670c4fd6f2a963fe334d441e020b3c97415d0f0bbae5f4812a34a304ccfc5.
//
// Solidity: event NewStrategy(address caller, string strategy)
func (_Contract *ContractFilterer) WatchNewStrategy(opts *bind.WatchOpts, sink chan<- *ContractNewStrategy) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewStrategy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewStrategy)
				if err := _Contract.contract.UnpackLog(event, "NewStrategy", log); err != nil {
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

// ParseNewStrategy is a log parse operation binding the contract event 0x54d670c4fd6f2a963fe334d441e020b3c97415d0f0bbae5f4812a34a304ccfc5.
//
// Solidity: event NewStrategy(address caller, string strategy)
func (_Contract *ContractFilterer) ParseNewStrategy(log types.Log) (*ContractNewStrategy, error) {
	event := new(ContractNewStrategy)
	if err := _Contract.contract.UnpackLog(event, "NewStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOpenOnTradeIterator is returned from FilterOpenOnTrade and is used to iterate over the raw logs and unpacked data for OpenOnTrade events raised by the Contract contract.
type ContractOpenOnTradeIterator struct {
	Event *ContractOpenOnTrade // Event containing the contract specifics and raw log

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
func (it *ContractOpenOnTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOpenOnTrade)
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
		it.Event = new(ContractOpenOnTrade)
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
func (it *ContractOpenOnTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOpenOnTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOpenOnTrade represents a OpenOnTrade event raised by the Contract contract.
type ContractOpenOnTrade struct {
	OrderId *big.Int
	Index   *big.Int
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenOnTrade is a free log retrieval operation binding the contract event 0xdbd1e106195ad17edfed5cb4bfc95758cc3219e116a69b5619f69e7241f94953.
//
// Solidity: event OpenOnTrade(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) FilterOpenOnTrade(opts *bind.FilterOpts, orderId []*big.Int, index []*big.Int) (*ContractOpenOnTradeIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OpenOnTrade", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOpenOnTradeIterator{contract: _Contract.contract, event: "OpenOnTrade", logs: logs, sub: sub}, nil
}

// WatchOpenOnTrade is a free log subscription operation binding the contract event 0xdbd1e106195ad17edfed5cb4bfc95758cc3219e116a69b5619f69e7241f94953.
//
// Solidity: event OpenOnTrade(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) WatchOpenOnTrade(opts *bind.WatchOpts, sink chan<- *ContractOpenOnTrade, orderId []*big.Int, index []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OpenOnTrade", orderIdRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOpenOnTrade)
				if err := _Contract.contract.UnpackLog(event, "OpenOnTrade", log); err != nil {
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

// ParseOpenOnTrade is a log parse operation binding the contract event 0xdbd1e106195ad17edfed5cb4bfc95758cc3219e116a69b5619f69e7241f94953.
//
// Solidity: event OpenOnTrade(uint256 indexed orderId, uint256 indexed index, uint8 status)
func (_Contract *ContractFilterer) ParseOpenOnTrade(log types.Log) (*ContractOpenOnTrade, error) {
	event := new(ContractOpenOnTrade)
	if err := _Contract.contract.UnpackLog(event, "OpenOnTrade", log); err != nil {
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
	OrderId      *big.Int
	CloseMargin  *big.Int
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestClose is a free log retrieval operation binding the contract event 0x65e03d81870a70279568cada208e70be82cdaf16444955f5b3f78479db986339.
//
// Solidity: event RequestClose(uint256 indexed orderId, uint256 indexed closeMargin, address indexed vaultAddress)
func (_Contract *ContractFilterer) FilterRequestClose(opts *bind.FilterOpts, orderId []*big.Int, closeMargin []*big.Int, vaultAddress []common.Address) (*ContractRequestCloseIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var closeMarginRule []interface{}
	for _, closeMarginItem := range closeMargin {
		closeMarginRule = append(closeMarginRule, closeMarginItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RequestClose", orderIdRule, closeMarginRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRequestCloseIterator{contract: _Contract.contract, event: "RequestClose", logs: logs, sub: sub}, nil
}

// WatchRequestClose is a free log subscription operation binding the contract event 0x65e03d81870a70279568cada208e70be82cdaf16444955f5b3f78479db986339.
//
// Solidity: event RequestClose(uint256 indexed orderId, uint256 indexed closeMargin, address indexed vaultAddress)
func (_Contract *ContractFilterer) WatchRequestClose(opts *bind.WatchOpts, sink chan<- *ContractRequestClose, orderId []*big.Int, closeMargin []*big.Int, vaultAddress []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var closeMarginRule []interface{}
	for _, closeMarginItem := range closeMargin {
		closeMarginRule = append(closeMarginRule, closeMarginItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RequestClose", orderIdRule, closeMarginRule, vaultAddressRule)
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

// ParseRequestClose is a log parse operation binding the contract event 0x65e03d81870a70279568cada208e70be82cdaf16444955f5b3f78479db986339.
//
// Solidity: event RequestClose(uint256 indexed orderId, uint256 indexed closeMargin, address indexed vaultAddress)
func (_Contract *ContractFilterer) ParseRequestClose(log types.Log) (*ContractRequestClose, error) {
	event := new(ContractRequestClose)
	if err := _Contract.contract.UnpackLog(event, "RequestClose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRequestOpenLimitIterator is returned from FilterRequestOpenLimit and is used to iterate over the raw logs and unpacked data for RequestOpenLimit events raised by the Contract contract.
type ContractRequestOpenLimitIterator struct {
	Event *ContractRequestOpenLimit // Event containing the contract specifics and raw log

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
func (it *ContractRequestOpenLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRequestOpenLimit)
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
		it.Event = new(ContractRequestOpenLimit)
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
func (it *ContractRequestOpenLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRequestOpenLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRequestOpenLimit represents a RequestOpenLimit event raised by the Contract contract.
type ContractRequestOpenLimit struct {
	OrderId      *big.Int
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestOpenLimit is a free log retrieval operation binding the contract event 0xf8d8295a2e10ef0d796fb30db26609d5c2e5282e1ed6627eaf42b62998f86d17.
//
// Solidity: event RequestOpenLimit(uint256 indexed orderId, address indexed vaultAddress)
func (_Contract *ContractFilterer) FilterRequestOpenLimit(opts *bind.FilterOpts, orderId []*big.Int, vaultAddress []common.Address) (*ContractRequestOpenLimitIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RequestOpenLimit", orderIdRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRequestOpenLimitIterator{contract: _Contract.contract, event: "RequestOpenLimit", logs: logs, sub: sub}, nil
}

// WatchRequestOpenLimit is a free log subscription operation binding the contract event 0xf8d8295a2e10ef0d796fb30db26609d5c2e5282e1ed6627eaf42b62998f86d17.
//
// Solidity: event RequestOpenLimit(uint256 indexed orderId, address indexed vaultAddress)
func (_Contract *ContractFilterer) WatchRequestOpenLimit(opts *bind.WatchOpts, sink chan<- *ContractRequestOpenLimit, orderId []*big.Int, vaultAddress []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RequestOpenLimit", orderIdRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRequestOpenLimit)
				if err := _Contract.contract.UnpackLog(event, "RequestOpenLimit", log); err != nil {
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

// ParseRequestOpenLimit is a log parse operation binding the contract event 0xf8d8295a2e10ef0d796fb30db26609d5c2e5282e1ed6627eaf42b62998f86d17.
//
// Solidity: event RequestOpenLimit(uint256 indexed orderId, address indexed vaultAddress)
func (_Contract *ContractFilterer) ParseRequestOpenLimit(log types.Log) (*ContractRequestOpenLimit, error) {
	event := new(ContractRequestOpenLimit)
	if err := _Contract.contract.UnpackLog(event, "RequestOpenLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRequestOpenMarketIterator is returned from FilterRequestOpenMarket and is used to iterate over the raw logs and unpacked data for RequestOpenMarket events raised by the Contract contract.
type ContractRequestOpenMarketIterator struct {
	Event *ContractRequestOpenMarket // Event containing the contract specifics and raw log

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
func (it *ContractRequestOpenMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRequestOpenMarket)
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
		it.Event = new(ContractRequestOpenMarket)
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
func (it *ContractRequestOpenMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRequestOpenMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRequestOpenMarket represents a RequestOpenMarket event raised by the Contract contract.
type ContractRequestOpenMarket struct {
	OrderId      *big.Int
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestOpenMarket is a free log retrieval operation binding the contract event 0x52896e33b03258815819a1a383b293e1615f7fb7165dc36f4ed542bb3063a83a.
//
// Solidity: event RequestOpenMarket(uint256 indexed orderId, address indexed vaultAddress)
func (_Contract *ContractFilterer) FilterRequestOpenMarket(opts *bind.FilterOpts, orderId []*big.Int, vaultAddress []common.Address) (*ContractRequestOpenMarketIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RequestOpenMarket", orderIdRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRequestOpenMarketIterator{contract: _Contract.contract, event: "RequestOpenMarket", logs: logs, sub: sub}, nil
}

// WatchRequestOpenMarket is a free log subscription operation binding the contract event 0x52896e33b03258815819a1a383b293e1615f7fb7165dc36f4ed542bb3063a83a.
//
// Solidity: event RequestOpenMarket(uint256 indexed orderId, address indexed vaultAddress)
func (_Contract *ContractFilterer) WatchRequestOpenMarket(opts *bind.WatchOpts, sink chan<- *ContractRequestOpenMarket, orderId []*big.Int, vaultAddress []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RequestOpenMarket", orderIdRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRequestOpenMarket)
				if err := _Contract.contract.UnpackLog(event, "RequestOpenMarket", log); err != nil {
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

// ParseRequestOpenMarket is a log parse operation binding the contract event 0x52896e33b03258815819a1a383b293e1615f7fb7165dc36f4ed542bb3063a83a.
//
// Solidity: event RequestOpenMarket(uint256 indexed orderId, address indexed vaultAddress)
func (_Contract *ContractFilterer) ParseRequestOpenMarket(log types.Log) (*ContractRequestOpenMarket, error) {
	event := new(ContractRequestOpenMarket)
	if err := _Contract.contract.UnpackLog(event, "RequestOpenMarket", log); err != nil {
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
	OrderId      *big.Int
	Amount       *big.Int
	IsAdd        bool
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateMargin is a free log retrieval operation binding the contract event 0xc421704e75a679bacbd7abc6d53b71f1d57d2fba8d4da363849fc14836b245c1.
//
// Solidity: event UpdateMargin(uint256 indexed orderId, uint256 indexed amount, bool _isAdd, address indexed vaultAddress)
func (_Contract *ContractFilterer) FilterUpdateMargin(opts *bind.FilterOpts, orderId []*big.Int, amount []*big.Int, vaultAddress []common.Address) (*ContractUpdateMarginIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateMargin", orderIdRule, amountRule, vaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateMarginIterator{contract: _Contract.contract, event: "UpdateMargin", logs: logs, sub: sub}, nil
}

// WatchUpdateMargin is a free log subscription operation binding the contract event 0xc421704e75a679bacbd7abc6d53b71f1d57d2fba8d4da363849fc14836b245c1.
//
// Solidity: event UpdateMargin(uint256 indexed orderId, uint256 indexed amount, bool _isAdd, address indexed vaultAddress)
func (_Contract *ContractFilterer) WatchUpdateMargin(opts *bind.WatchOpts, sink chan<- *ContractUpdateMargin, orderId []*big.Int, amount []*big.Int, vaultAddress []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	var vaultAddressRule []interface{}
	for _, vaultAddressItem := range vaultAddress {
		vaultAddressRule = append(vaultAddressRule, vaultAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateMargin", orderIdRule, amountRule, vaultAddressRule)
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

// ParseUpdateMargin is a log parse operation binding the contract event 0xc421704e75a679bacbd7abc6d53b71f1d57d2fba8d4da363849fc14836b245c1.
//
// Solidity: event UpdateMargin(uint256 indexed orderId, uint256 indexed amount, bool _isAdd, address indexed vaultAddress)
func (_Contract *ContractFilterer) ParseUpdateMargin(log types.Log) (*ContractUpdateMargin, error) {
	event := new(ContractUpdateMargin)
	if err := _Contract.contract.UnpackLog(event, "UpdateMargin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	Sender common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 indexed assets, uint256 indexed shares)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, assets []*big.Int, shares []*big.Int) (*ContractWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetsRule []interface{}
	for _, assetsItem := range assets {
		assetsRule = append(assetsRule, assetsItem)
	}
	var sharesRule []interface{}
	for _, sharesItem := range shares {
		sharesRule = append(sharesRule, sharesItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw", senderRule, assetsRule, sharesRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 indexed assets, uint256 indexed shares)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw, sender []common.Address, assets []*big.Int, shares []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetsRule []interface{}
	for _, assetsItem := range assets {
		assetsRule = append(assetsRule, assetsItem)
	}
	var sharesRule []interface{}
	for _, sharesItem := range shares {
		sharesRule = append(sharesRule, sharesItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw", senderRule, assetsRule, sharesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 indexed assets, uint256 indexed shares)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
