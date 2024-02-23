package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ScoreDao interface {
	Transaction(tx *gorm.DB) ScoreDao
	Create(f *dto.UserScoreHistory) error
	QueryScoreHistoryByUserId(userID string) ([]*dto.UserScoreHistory, error)
	QueryScoreHistoryByUserIdAndType(userID string, scoreType int64) ([]*dto.UserScoreHistory, error)
	QueryScoreHistory(id string) (*dto.UserScoreHistory, error)
	QueryTotalScore(userID string) (uint64, error)
	QueryTotalScoreByType(userID string, scoreType uint64) (uint64, error)
	Delete(f *dto.UserScoreHistory) error
}

type Score struct {
	db *db.DB
}

func (d *Score) QueryScoreHistoryByUserId(userId string) ([]*dto.UserScoreHistory, error) {
	var scoreHistorys []*dto.UserScoreHistory
	if err := d.db.Model(&dto.UserScoreHistory{}).
		Where("user_id =?", userId).
		Order("created_at desc").
		Find(&scoreHistorys).Error; err != nil {
		return scoreHistorys, err
	}
	return scoreHistorys, nil
}

func (d *Score) QueryScoreHistoryByUserIdAndType(userId string, scoreType int64) ([]*dto.UserScoreHistory, error) {
	var scoreHistorys []*dto.UserScoreHistory
	if err := d.db.Where(&dto.UserScoreHistory{
		UserID: userId,
		Type:   uint64(scoreType),
	}).Order("created_at desc").
		Find(&scoreHistorys).Error; err != nil {
		return scoreHistorys, err
	}
	return scoreHistorys, nil
}

func (d *Score) QueryScoreHistory(id string) (*dto.UserScoreHistory, error) {
	sh := &dto.UserScoreHistory{}
	res := d.db.Where(&dto.UserScoreHistory{
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

func (d *Score) QueryTotalScore(userID string) (uint64, error) {
	var sum uint64
	res := d.db.Model(&dto.UserScoreHistory{}).Where(&dto.UserScoreHistory{
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

func (d *Score) QueryTotalScoreByType(userID string, scoreType uint64) (uint64, error) {
	var sum uint64
	res := d.db.Model(&dto.UserScoreHistory{}).Where(&dto.UserScoreHistory{
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

func NewScoreDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (ScoreDao, error) {
	dao := &Score{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *Score) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.UserScoreHistory{})
}

func (d *Score) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *Score) Create(history *dto.UserScoreHistory) error {
	return d.db.Create(history).Error
}

func (d *Score) Transaction(tx *gorm.DB) ScoreDao {
	return &Score{
		db: &db.DB{DB: tx},
	}
}

func (d *Score) Delete(f *dto.UserScoreHistory) error {
	return nil
}
