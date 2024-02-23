package dto

import (
	"time"

	"gorm.io/gorm"
)

const (
	Pending = iota
	Done
)

type UserDiceBuyRecord struct {
	ID     uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserID string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Wallet string `json:"wallet" gorm:"column:wallet;type:varchar(96);null;index"`

	BlockNumber uint64 `json:"block_number" gorm:"column:block_number;type:bigint;not null;"`
	BlockHash   string `json:"block_hash" gorm:"column:block_hash;primaryKey;type:varchar(96);not null;"`
	TxIndex     uint64 `json:"tx_index" gorm:"column:tx_index;type:bigint;not null;"`
	TxHash      string `json:"tx_hash" gorm:"column:tx_hash;primaryKey;type:varchar(96);not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type StakingRecord struct {
	ID            uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserID        string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Wallet        string `json:"wallet" gorm:"column:wallet;type:varchar(96);null;index"`
	StakingAmount uint64 `json:"staking_amount" gorm:"column:staking_amount;type:varchar(100);not null;"`
	Symbol        string `json:"symbol" gorm:"column:symbol;type:varchar(100);not null;"`
	Address       string `json:"address" gorm:"column:address;type:varchar(100);"`

	BlockNumber uint64 `json:"block_number" gorm:"column:block_number;type:bigint;not null;"`
	BlockHash   string `json:"block_hash" gorm:"column:block_hash;primaryKey;type:varchar(96);not null;"`
	TxIndex     uint64 `json:"tx_index" gorm:"column:tx_index;type:bigint;not null;"`
	TxHash      string `json:"tx_hash" gorm:"column:tx_hash;primaryKey;type:varchar(96);not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UnStakingRecord struct {
	ID            uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserID        string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Wallet        string `json:"wallet" gorm:"column:wallet;type:varchar(96);null;index"`
	StakingAmount string `json:"staking_amount" gorm:"column:staking_amount;type:varchar(100);not null;"`
	Symbol        string `json:"symbol" gorm:"column:symbol;type:varchar(100);not null;"`
	Address       string `json:"address" gorm:"column:address;type:varchar(100);"`

	BlockNumber uint64 `json:"block_number" gorm:"column:block_number;type:bigint;not null;"`
	BlockHash   string `json:"block_hash" gorm:"column:block_hash;primaryKey;type:varchar(96);not null;"`
	TxIndex     uint64 `json:"tx_index" gorm:"column:tx_index;type:bigint;not null;"`
	TxHash      string `json:"tx_hash" gorm:"column:tx_hash;primaryKey;type:varchar(96);not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type User struct {
	UUID       string `json:"uuid" gorm:"column:uuid;type:char(32);primaryKey;"`
	Wallet     string `json:"wallet" gorm:"column:wallet;type:varchar(96);null;uniqueIndex"`
	Username   string `json:"username" gorm:"column:username;type:varchar(100);null;uniqueIndex"`
	ProfileUrl string `json:"profile_url" gorm:"column:profile_url;type:varchar(200);not null;"`

	StakingAmountETH  string `json:"staking_amount_eth" gorm:"column:staking_amount_eth;type:varchar(100);not null;"`
	StakingAmountUSDB string `json:"staking_amount_usdb" gorm:"column:staking_amount_usdb;type:varchar(100);not null;"`
	MintDice          bool   `json:"is_add" gorm:"column:is_add;type:tinyint(1);not null"`
	DiceSpeed         uint64 `json:"dice_speed" gorm:"column:dice_speed;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
