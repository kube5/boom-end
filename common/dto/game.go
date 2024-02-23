package dto

import (
	"time"

	"gorm.io/gorm"
)

type GameContract struct {
	ChainID   uint64 `json:"chain_id" gorm:"column:chain_id;primaryKey;type:bigint;not null;"`
	Contract  string `json:"contract" gorm:"column:contract;primaryKey;type:varchar(96);not null;"`
	NextStart uint64 `json:"next_start" gorm:"column:next_start;type:bigint;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type GameMonitorRecord struct {
	ChainID     uint64 `json:"chain_id" gorm:"column:chain_id;primaryKey;type:bigint;not null;"`
	Contract    string `json:"contract" gorm:"column:contract;type:varchar(96);not null;"`
	BlockNumber uint64 `json:"block_number" gorm:"column:block_number;type:bigint;not null;"`
	BlockHash   string `json:"block_hash" gorm:"column:block_hash;primaryKey;type:varchar(96);not null;"`
	TxIndex     uint64 `json:"tx_index" gorm:"column:tx_index;type:bigint;not null;"`
	TxHash      string `json:"tx_hash" gorm:"column:tx_hash;primaryKey;type:varchar(96);not null;"`
	Index       uint64 `json:"index" gorm:"column:index;primaryKey;type:bigint;not null;"`
	Removed     bool   `json:"removed" gorm:"column:removed;primaryKey;type:boolean;not null;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserScoreHistory struct {
	UUID   string `json:"uuid" gorm:"column:uuid;type:char(32);primaryKey;not null"`
	UserID string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Score  int64  `json:"score" gorm:"column:score;type:bigint;not null"`
	Type   uint64 `json:"type" gorm:"column:type;type:tinyint;not null"`
	Desc   string `json:"desc" gorm:"column:desc;type:varchar(100);not null"`
	DescId string `json:"desc_id" gorm:"column:desc_id;type:varchar(100);not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserCashHistory struct {
	UUID                   string `json:"uuid" gorm:"column:uuid;type:char(32);primaryKey;not null"`
	UserScoreUsedHistoryID string `json:"user_score_used_history_id" gorm:"column:user_score_used_history_id;type:char(32);"`
	UserID                 string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Cash                   int64  `json:"cash" gorm:"column:cash;type:bigint;not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserScoreUsedHistory struct {
	UUID   string `json:"uuid" gorm:"column:uuid;type:char(32);primaryKey;not null"`
	UserID string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Score  int64  `json:"score" gorm:"column:score;type:bigint;not null"`
	Type   uint64 `json:"type" gorm:"column:type;type:tinyint;not null"`
	Desc   string `json:"desc" gorm:"column:desc;type:varchar(100);not null"`
	DescId string `json:"desc_id" gorm:"column:desc_id;type:varchar(100);not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
