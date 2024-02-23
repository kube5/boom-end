package service

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/game/handler/dao"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type BaseService interface {
	TransactionExecute(executor func(tx *gorm.DB) error) error
}

type Base struct {
	db *gorm.DB
}

func NewBaseService(db *db.DB) BaseService {
	return &Base{
		db: db.DB,
	}
}

func (b *Base) TransactionExecute(executor func(tx *gorm.DB) error) error {
	return b.db.Transaction(func(tx *gorm.DB) error {
		return executor(tx)
	})
}

func isEventProcessed(recordDao dao.GameMonitorRecordDao, chainId uint64, log *types.Log) (bool, error) {
	record, err := recordDao.Query(chainId, log.TxHash.String(), uint64(log.Index))
	if err != nil {
		return false, err
	}

	return record != nil, nil
}

func genRecordByEvent(chainId uint64, log *types.Log) *dto.GameMonitorRecord {
	return &dto.GameMonitorRecord{
		ChainID:     chainId,
		Contract:    log.Address.String(),
		BlockHash:   log.BlockHash.String(),
		BlockNumber: log.BlockNumber,
		TxIndex:     uint64(log.TxIndex),
		TxHash:      log.TxHash.String(),
		Index:       uint64(log.Index),
		Removed:     log.Removed,
	}
}
