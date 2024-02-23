package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ScoreUsedDao interface {
	Transaction(tx *gorm.DB) ScoreUsedDao
	Create(f *dto.UserScoreUsedHistory) error
	QueryScoreUsedHistoryByUserId(userID string) ([]*dto.UserScoreUsedHistory, error)
	QueryScoreUsedHistoryByUserIdAndType(userID string, scoreType int64) ([]*dto.UserScoreUsedHistory, error)
	QueryScoreUsedHistory(id string) (*dto.UserScoreUsedHistory, error)
	QueryTotalScoreUsed(userID string) (uint64, error)
	QueryTotalScoreUsedByType(userID string, scoreType uint64) (uint64, error)
	Delete(f *dto.UserScoreUsedHistory) error
}

type ScoreUsed struct {
	db *db.DB
}

func (d *ScoreUsed) QueryScoreUsedHistoryByUserId(userId string) ([]*dto.UserScoreUsedHistory, error) {
	var scoreHistorys []*dto.UserScoreUsedHistory
	if err := d.db.Model(&dto.UserScoreUsedHistory{}).
		Where("user_id =?", userId).
		Order("created_at desc").
		Find(&scoreHistorys).Error; err != nil {
		return scoreHistorys, err
	}
	return scoreHistorys, nil
}

func (d *ScoreUsed) QueryScoreUsedHistoryByUserIdAndType(userId string, scoreType int64) ([]*dto.UserScoreUsedHistory, error) {
	var scoreHistorys []*dto.UserScoreUsedHistory
	if err := d.db.Where(&dto.UserScoreUsedHistory{
		UserID: userId,
		Type:   uint64(scoreType),
	}).Order("created_at desc").
		Find(&scoreHistorys).Error; err != nil {
		return scoreHistorys, err
	}
	return scoreHistorys, nil
}

func (d *ScoreUsed) QueryScoreUsedHistory(id string) (*dto.UserScoreUsedHistory, error) {
	sh := &dto.UserScoreUsedHistory{}
	res := d.db.Where(&dto.UserScoreUsedHistory{
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

func (d *ScoreUsed) QueryTotalScoreUsed(userID string) (uint64, error) {
	var sum uint64
	res := d.db.Model(&dto.UserScoreUsedHistory{}).Where(&dto.UserScoreUsedHistory{
		UserID: userID,
	}).Select("COALESCE(SUM(score), 0)").Scan(&sum)
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

func (d *ScoreUsed) QueryTotalScoreUsedByType(userID string, scoreType uint64) (uint64, error) {
	var sum uint64
	res := d.db.Model(&dto.UserScoreUsedHistory{}).Where(&dto.UserScoreUsedHistory{
		UserID: userID,
		Type:   scoreType,
	}).Select("COALESCE(SUM(score), 0)").Scan(&sum)
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

func NewScoreUsedDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (ScoreUsedDao, error) {
	dao := &ScoreUsed{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *ScoreUsed) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.UserScoreUsedHistory{})
}

func (d *ScoreUsed) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *ScoreUsed) Create(history *dto.UserScoreUsedHistory) error {
	return d.db.Create(history).Error
}

func (d *ScoreUsed) Transaction(tx *gorm.DB) ScoreUsedDao {
	return &ScoreUsed{
		db: &db.DB{DB: tx},
	}
}

func (d *ScoreUsed) Delete(f *dto.UserScoreUsedHistory) error {
	return nil
}
