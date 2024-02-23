package dao

import (
	"errors"

	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type GameMonitorRecordDao interface {
	Transaction(tx *gorm.DB) GameMonitorRecordDao
	Query(chainId uint64, txHash string, index uint64) (*dto.GameMonitorRecord, error)
	Create(contract *dto.GameMonitorRecord) error
}

type GameMonitorRecord struct {
	db *db.DB
}

func NewGameMonitorRecordDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (GameMonitorRecordDao, error) {
	dao := &GameMonitorRecord{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *GameMonitorRecord) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.GameMonitorRecord{})
}

func (d *GameMonitorRecord) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *GameMonitorRecord) Transaction(tx *gorm.DB) GameMonitorRecordDao {
	return &GameMonitorRecord{
		db: &db.DB{DB: tx},
	}
}

func (d *GameMonitorRecord) Query(chainId uint64, txHash string, index uint64) (*dto.GameMonitorRecord, error) {
	record := &dto.GameMonitorRecord{}
	res := d.db.Where(&dto.GameMonitorRecord{
		ChainID: chainId,
		TxHash:  txHash,
	}).Where("game_monitor_records.index = ?", index).Find(record)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return record, nil
}

func (d *GameMonitorRecord) Create(monitor *dto.GameMonitorRecord) error {
	return d.db.Create(monitor).Error
}
