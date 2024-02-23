package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type GameContractDao interface {
	Transaction(tx *gorm.DB) GameContractDao
	QueryByChainIdAndContract(chainId uint64, contract string) (*dto.GameContract, error)
	Create(contract *dto.GameContract) error
	Update(contract *dto.GameContract) error
}

type GameContract struct {
	db *db.DB
}

func NewGameContractDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (GameContractDao, error) {
	dao := &GameContract{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *GameContract) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.GameContract{})
}

func (d *GameContract) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *GameContract) Transaction(tx *gorm.DB) GameContractDao {
	return &GameContract{
		db: &db.DB{DB: tx},
	}
}

func (d *GameContract) QueryByChainIdAndContract(chainId uint64, contract string) (*dto.GameContract, error) {
	monitor := &dto.GameContract{}
	res := d.db.Where(&dto.GameContract{
		ChainID:  chainId,
		Contract: contract,
	}).Find(monitor)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return monitor, nil
}

func (d *GameContract) Create(monitor *dto.GameContract) error {
	return d.db.Create(monitor).Error
}

func (d *GameContract) Update(monitor *dto.GameContract) error {
	return d.db.Model(&dto.GameContract{
		ChainID:  monitor.ChainID,
		Contract: monitor.Contract,
	}).Updates(monitor).Error
}
