package dto

import (
	"time"

	"gorm.io/gorm"
)

type Market struct {
	ID                uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID              string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	TotalUsers        uint64 `json:"total_users" gorm:"column:total_users;type:bigint;not null;"`
	TotalTrade        uint64 `json:"total_trade" gorm:"column:total_trade;type:bigint;not null;"`
	TradeVolume       uint64 `json:"trade_volume" gorm:"column:trade_volume;type:bigint;not null;"`
	TotalTradeFee     uint64 `json:"total_trade_fee" gorm:"column:total_trade_fee;type:bigint;not null;"`
	VaultLock         uint64 `json:"vault_lock" gorm:"column:vault_lock;type:bigint;not null;"`
	LiquidationAmount uint64 `json:"liquidation_amount" gorm:"column:liquidation_amount;type:bigint;not null;"`
	LiquidationCount  uint64 `json:"liquidation_count" gorm:"column:liquidation_count;type:bigint;not null;"`
	OpenInterestLong  uint64 `json:"open_interest_long" gorm:"column:open_interest_long;type:bigint;not null;"`
	OpenInterestShort uint64 `json:"open_interest_short" gorm:"column:open_interest_short;type:bigint;not null;"`
	Profit            uint64 `json:"profit" gorm:"column:profit;type:bigint;not null;"`
	Loss              uint64 `json:"loss" gorm:"column:loss;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Pair struct {
	ID                uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID              string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	TotalUsers        uint64 `json:"total_users" gorm:"column:total_users;type:bigint;not null;"`
	TotalTrade        uint64 `json:"total_trade" gorm:"column:total_trade;type:bigint;not null;"`
	TradeVolume       uint64 `json:"trade_volume" gorm:"column:trade_volume;type:bigint;not null;"`
	TotalTradeFee     uint64 `json:"total_trade_fee" gorm:"column:total_trade_fee;type:bigint;not null;"`
	LiquidationAmount uint64 `json:"liquidation_amount" gorm:"column:liquidation_amount;type:bigint;not null;"`
	LiquidationCount  uint64 `json:"liquidation_count" gorm:"column:liquidation_count;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

//type User struct {
//	ID                 uint64    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
//	UUID               string    `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
//	Deposited          uint64    `json:"deposited" gorm:"column:deposited;type:bigint;not null;"`
//	Withdrawed         uint64    `json:"withdrawed" gorm:"column:withdrawed;type:bigint;not null;"`
//	TotalTrade         uint64    `json:"total_trade" gorm:"column:total_trade;type:bigint;not null;"`
//	TradeVolume        uint64    `json:"trade_volume" gorm:"column:trade_volume;type:bigint;not null;"`
//	TradePrincipal     uint64    `json:"trade_principal" gorm:"column:trade_principal;type:bigint;not null;"`
//	CloseCount         uint64    `json:"close_count" gorm:"column:close_count;type:bigint;not null;"`
//	OpenPosition       uint64    `json:"open_position" gorm:"column:open_position;type:bigint;not null;"`
//	OpenPositionCount  uint64    `json:"open_position_count" gorm:"column:open_position_count;type:bigint;not null;"`
//	OpenPrincipal      uint64    `json:"open_principal" gorm:"column:open_principal;type:bigint;not null;"`
//	ProfitAmount       uint64    `json:"profit_amount" gorm:"column:profit_amount;type:bigint;not null;"`
//	GainTradeCount     uint64    `json:"gain_trade_count" gorm:"column:gain_trade_count;type:bigint;not null;"`
//	AmountGainRate     big.Float `json:"amount_gain_rate" gorm:"column:amount_gain_rate;type:decimal(10,2);not null;"`
//	TradeCountGainRate big.Float `json:"trade_count_gain_rate" gorm:"column:trade_count_gain_rate;type:decimal(10,2);not null;"`
//	LiquidationAmount  uint64    `json:"liquidation_amount" gorm:"column:liquidation_amount;type:bigint;not null;"`
//	LiquidationCount   uint64    `json:"liquidation_count" gorm:"column:liquidation_count;type:bigint;not null;"`
//
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
//}

type OpenTrade struct {
	ID          uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID        string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	UserId      string `json:"user_id" gorm:"column:user_id;type:char(32);index;"`
	Trader      string `json:"trader" gorm:"column:trader;type:char(64);index;"`
	PairIndex   uint64 `json:"pair_index" gorm:"column:pair_index;type:bigint;not null;"`
	Margin      uint64 `json:"margin" gorm:"column:margin;type:bigint;not null;"`
	LeftMargin  uint64 `json:"left_margin" gorm:"column:left_margin;type:bigint;not null;"`
	OpenPrice   uint64 `json:"open_price" gorm:"column:open_price;type:bigint;not null;"`
	Long        bool   `json:"long" gorm:"column:long;type:tinyint(1);not null"`
	Leverage    uint64 `json:"leverage" gorm:"column:leverage;type:bigint;not null;"`
	Tp          uint64 `json:"tp" gorm:"column:tp;type:bigint;not null;"`
	Sl          uint64 `json:"sl" gorm:"column:sl;type:bigint;not null;"`
	Status      uint64 `json:"status" gorm:"column:status;type:bigint;not null;"`
	OpenTime    uint64 `json:"open_time" gorm:"column:open_time;type:bigint;not null;"`
	CloseTime   uint64 `json:"close_time" gorm:"column:close_time;type:bigint;not null;"`
	OpenFee     uint64 `json:"open_fee" gorm:"column:open_fee;type:bigint;not null;"`
	CloseFee    uint64 `json:"close_fee" gorm:"column:close_fee;type:bigint;not null;"`
	RolloverFee uint64 `json:"rollover_fee" gorm:"column:rollover_fee;type:bigint;not null;"`
	FundingFee  uint64 `json:"funding_fee" gorm:"column:funding_fee;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type OpenRequest struct {
	ID         uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID       string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	UserId     string `json:"user_id" gorm:"column:user_id;type:char(32);index;"`
	Trader     string `json:"trader" gorm:"column:trader;type:char(64);index;"`
	PairIndex  uint64 `json:"pair_index" gorm:"column:pair_index;type:bigint;not null;"`
	Margin     uint64 `json:"margin" gorm:"column:margin;type:bigint;not null;"`
	Long       bool   `json:"long" gorm:"column:long;type:tinyint(1);not null"`
	Leverage   uint64 `json:"leverage" gorm:"column:leverage;type:bigint;not null;"`
	Tp         uint64 `json:"tp" gorm:"column:tp;type:bigint;not null;"`
	Sl         uint64 `json:"sl" gorm:"column:sl;type:bigint;not null;"`
	TradeType  uint64 `json:"trade_type" gorm:"column:trade_type;type:bigint;not null;"`
	Status     uint64 `json:"status" gorm:"column:status;type:bigint;not null;"`
	CallTarget string `json:"call_target" gorm:"column:call_target;type:char(512);index;"`
	MaxPrice   uint64 `json:"max_price" gorm:"column:max_price;type:bigint;not null;"`
	MinPrice   uint64 `json:"min_price" gorm:"column:min_price;type:bigint;not null;"`
	Time       uint64 `json:"time" gorm:"column:time;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CloseRequest struct {
	ID          uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID        string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	UserId      string `json:"user_id" gorm:"column:user_id;type:char(32);index;"`
	CloseMargin uint64 `json:"close_margin" gorm:"column:close_margin;type:bigint;not null;"`
	Index       uint64 `json:"index" gorm:"column::index;type:bigint;not null;"`
	RequestTime uint64 `json:"request_time" gorm:"column:request_time;type:bigint;not null;"`
	Status      uint64 `json:"status" gorm:"column:status;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CloseHistory struct {
	ID          uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID        string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	UserId      string `json:"user_id" gorm:"column:user_id;type:char(32);index;"`
	CloseMargin uint64 `json:"close_margin" gorm:"column:close_margin;type:bigint;not null;"`
	OrderId     uint64 `json:"order_id" gorm:"column:order_id;type:bigint;not null;"`
	PairIndex   uint64 `json:"pair_index" gorm:"column:pair_index;type:bigint;not null;"`
	CloseFee    uint64 `json:"close_fee" gorm:"column:close_fee;type:bigint;not null;"`
	Leverage    uint64 `json:"leverage" gorm:"column:leverage;type:bigint;not null;"`
	Long        bool   `json:"long" gorm:"column:long;type:tinyint(1);not null"`
	OpenPrice   uint64 `json:"open_price" gorm:"column:open_price;type:bigint;not null;"`
	RolloverFee uint64 `json:"rollover_fee" gorm:"column:rollover_fee;type:bigint;not null;"`
	FundingFee  uint64 `json:"funding_fee" gorm:"column:funding_fee;type:bigint;not null;"`
	OpenTime    uint64 `json:"open_time" gorm:"column:open_time;type:bigint;not null;"`
	CloseTime   uint64 `json:"close_time" gorm:"column:close_time;type:bigint;not null;"`
	CloseType   uint64 `json:"close_type" gorm:"column:close_type;type:bigint;not null;"`
	Profit      uint64 `json:"profit" gorm:"column:profit;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UpDateMargin struct {
	ID      uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UUID    string `json:"uuid" gorm:"column:uuid;type:char(32);index;"`
	UserId  string `json:"user_id" gorm:"column:user_id;type:char(32);index;"`
	OrderId uint64 `json:"order_id" gorm:"column:order_id;type:bigint;not null;"`
	Amount  uint64 `json:"amount" gorm:"column:amount;type:bigint;not null;"`
	IsAdd   bool   `json:"is_add" gorm:"column:is_add;type:tinyint(1);not null"`
	Time    uint64 `json:"time" gorm:"column:time;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

//type MonitorContract struct {
//	ChainId   uint64 `json:"chain_id" gorm:"column:chain_id;primaryKey;type:bigint;not null;"`
//	Contracts  string `json:"contract" gorm:"column:contract;primaryKey;type:varchar(96);not null;"`
//	NextBlock uint64 `json:"next_start" gorm:"column:next_start;type:bigint;not null;"`
//
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
//}
//
//type MonitorLog struct {
//	ChainId     uint64 `json:"chain_id" gorm:"column:chain_id;primaryKey;type:bigint;not null;"`
//	Contracts    string `json:"contract" gorm:"column:contract;type:varchar(96);not null;"`
//	BlockNumber uint64 `json:"block_number" gorm:"column:block_number;type:bigint;not null;"`
//	BlockHash   string `json:"block_hash" gorm:"column:block_hash;primaryKey;type:varchar(96);not null;"`
//	TxIndex     uint64 `json:"tx_index" gorm:"column:tx_index;type:bigint;not null;"`
//	TxHash      string `json:"tx_hash" gorm:"column:tx_hash;primaryKey;type:varchar(96);not null;"`
//	Index       uint64 `json:"index" gorm:"column:index;primaryKey;type:bigint;not null;"`
//	Removed     bool   `json:"removed" gorm:"column:removed;primaryKey;type:boolean;not null;"`
//
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
//}
