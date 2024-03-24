package dto

import (
	"time"

	"gorm.io/gorm"
)

type TUserScoreHistory struct {
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

type TUserCashHistory struct {
	UUID                   string `json:"uuid" gorm:"column:uuid;type:char(32);primaryKey;not null"`
	UserScoreUsedHistoryID string `json:"user_score_used_history_id" gorm:"column:user_score_used_history_id;type:char(32);"`
	UserID                 string `json:"user_id" gorm:"column:user_id;type:char(32);"`
	Cash                   int64  `json:"cash" gorm:"column:cash;type:bigint;not null"`
	Dice1                  int64  `json:"dice1" gorm:"column:dice1;type:bigint;not null"`
	Dice2                  int64  `json:"dice2" gorm:"column:dice2;type:bigint;not null"`
	Level                  int64  `json:"level" gorm:"column:level;type:bigint;not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type TUserScoreUsedHistory struct {
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
