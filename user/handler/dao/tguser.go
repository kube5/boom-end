package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TgUserDao interface {
	Transaction(tx *gorm.DB) TgUserDao
	FindByTgUserID(userID string) (*dto.TgUser, error)
	FindByAddress(from string) (*dto.TgUser, error)
	FindByTgUserName(userName string) (*dto.TgUser, error)
	Create(user *dto.TgUser) error
	UpdateAddress(user *dto.TgUser) error
	UpdateTgUser(user *dto.TgUser) (*dto.TgUser, error)
	BatchCreate(users []*dto.TgUser) error
	QueryBySocialName(socialName, name string) (*dto.TgUser, error)
}

type TgUser struct {
	db *db.DB
}

func (d *TgUser) QueryBySocialName(socialName, name string) (*dto.TgUser, error) {
	user := &dto.TgUser{}
	res := d.db.Where(socialName+" = ?", name).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func NewTgUserDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (TgUserDao, error) {
	dao := &TgUser{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *TgUser) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.TgUser{})
}

func (d *TgUser) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *TgUser) Transaction(tx *gorm.DB) TgUserDao {
	return &TgUser{
		db: &db.DB{DB: tx},
	}
}

func (d *TgUser) UpdateAddress(user *dto.TgUser) error {
	return d.db.Model(&dto.TgUser{UUID: user.UUID}).Updates(dto.TgUser{Wallet: user.Wallet}).Error
}

func (d *TgUser) UpdateTgUser(user *dto.TgUser) (*dto.TgUser, error) {
	if err := d.db.Model(&dto.TgUser{UUID: user.UUID}).Updates(user).Error; err != nil {
		return nil, err
	}

	return d.FindByTgUserID(user.UUID)
}

func (d *TgUser) FindByTgUserID(userID string) (*dto.TgUser, error) {
	user := &dto.TgUser{}
	res := d.db.Where(&dto.TgUser{
		UUID: userID,
	}).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func (d *TgUser) FindByTgUserName(username string) (*dto.TgUser, error) {
	user := &dto.TgUser{}
	res := d.db.Where(&dto.TgUser{
		Tgname: username,
	}).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func (d *TgUser) FindByAddress(addr string) (*dto.TgUser, error) {
	user := &dto.TgUser{}
	res := d.db.Where(&dto.TgUser{
		Wallet: addr,
	}).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func (d *TgUser) Create(user *dto.TgUser) error {
	var omitColumns []string
	if user.Wallet == "" {
		omitColumns = append(omitColumns, "wallet")
	}
	if user.Tgname == "" {
		omitColumns = append(omitColumns, "username")
	}
	omitColumns = append(omitColumns, "facebook", "instagram", "github", "telegram", "mirror", "linkedin")
	if len(omitColumns) > 0 {
		return d.db.Omit(omitColumns...).Create(user).Error
	}
	return d.db.Create(user).Error
}

func (d *TgUser) BatchCreate(users []*dto.TgUser) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return errors.Wrapf(err, "register user[%s] failed", user.UUID)
			}
		}
		return nil
	})
}
