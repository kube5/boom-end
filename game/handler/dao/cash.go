package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type CashDao interface {
	Transaction(tx *gorm.DB) CashDao
	Create(f *dto.UserCashHistory) error
	QueryCashHistoryByUserId(userID string) ([]*dto.UserCashHistory, error)
	QueryCashHistory(id string) (*dto.UserCashHistory, error)
	QueryTotalCash(userID string) (uint64, error)
	Delete(f *dto.UserCashHistory) error
}

type Cash struct {
	db *db.DB
}

func (d *Cash) QueryCashHistoryByUserId(userId string) ([]*dto.UserCashHistory, error) {
	var scoreHistorys []*dto.UserCashHistory
	if err := d.db.Model(&dto.UserCashHistory{}).
		Where("user_id =?", userId).
		Order("created_at desc").
		Find(&scoreHistorys).Error; err != nil {
		return scoreHistorys, err
	}
	return scoreHistorys, nil
}

func (d *Cash) QueryCashHistory(id string) (*dto.UserCashHistory, error) {
	sh := &dto.UserCashHistory{}
	res := d.db.Where(&dto.UserCashHistory{
		UUID: id,
	}).Find(sh)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return sh, nil
}

func (d *Cash) QueryTotalCash(userID string) (uint64, error) {
	var sum uint64
	res := d.db.Model(&dto.UserCashHistory{}).Where(&dto.UserCashHistory{
		UserID: userID,
	}).Select("COALESCE(SUM(cash), 0)").Scan(&sum)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, res.Error
	}
	if res.RowsAffected == 0 {
		return 0, nil
	}
	return sum, nil
}

func NewCashDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (CashDao, error) {
	dao := &Cash{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *Cash) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.UserCashHistory{})
}

func (d *Cash) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *Cash) Create(history *dto.UserCashHistory) error {
	return d.db.Create(history).Error
}

func (d *Cash) Transaction(tx *gorm.DB) CashDao {
	return &Cash{
		db: &db.DB{DB: tx},
	}
}

func (d *Cash) Delete(f *dto.UserCashHistory) error {
	return nil
}
