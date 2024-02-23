package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserDao interface {
	Transaction(tx *gorm.DB) UserDao
	FindByUserID(userID string) (*dto.User, error)
	FindByAddress(from string) (*dto.User, error)
	FindByUserName(userName string) (*dto.User, error)
	Create(user *dto.User) error
	UpdateAddress(user *dto.User) error
	UpdateUser(user *dto.User) (*dto.User, error)
	BatchCreate(users []*dto.User) error
	QueryBySocialName(socialName, name string) (*dto.User, error)
}

type User struct {
	db *db.DB
}

func (d *User) QueryBySocialName(socialName, name string) (*dto.User, error) {
	user := &dto.User{}
	res := d.db.Where(socialName+" = ?", name).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func NewUserDao(lc fx.Lifecycle, sd fx.Shutdowner, db *db.DB) (UserDao, error) {
	dao := &User{
		db: db,
	}
	basic.RegisterHook(lc, sd, dao)
	return dao, nil
}

func (d *User) Start(sd fx.Shutdowner) error {
	return d.db.AutoMigrate(&dto.User{})
}

func (d *User) Stop(sd fx.Shutdowner) error {
	return nil
}

func (d *User) Transaction(tx *gorm.DB) UserDao {
	return &User{
		db: &db.DB{DB: tx},
	}
}

func (d *User) UpdateAddress(user *dto.User) error {
	return d.db.Model(&dto.User{UUID: user.UUID}).Updates(dto.User{Wallet: user.Wallet}).Error
}

func (d *User) UpdateUser(user *dto.User) (*dto.User, error) {
	if err := d.db.Model(&dto.User{UUID: user.UUID}).Updates(user).Error; err != nil {
		return nil, err
	}

	return d.FindByUserID(user.UUID)
}

func (d *User) FindByUserID(userID string) (*dto.User, error) {
	user := &dto.User{}
	res := d.db.Where(&dto.User{
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

func (d *User) FindByUserName(username string) (*dto.User, error) {
	user := &dto.User{}
	res := d.db.Where(&dto.User{
		Username: username,
	}).Find(user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}

func (d *User) FindByAddress(addr string) (*dto.User, error) {
	user := &dto.User{}
	res := d.db.Where(&dto.User{
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

func (d *User) Create(user *dto.User) error {
	var omitColumns []string
	if user.Wallet == "" {
		omitColumns = append(omitColumns, "wallet")
	}
	if user.Username == "" {
		omitColumns = append(omitColumns, "username")
	}
	omitColumns = append(omitColumns, "facebook", "instagram", "github", "telegram", "mirror", "linkedin")
	if len(omitColumns) > 0 {
		return d.db.Omit(omitColumns...).Create(user).Error
	}
	return d.db.Create(user).Error
}

func (d *User) BatchCreate(users []*dto.User) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return errors.Wrapf(err, "register user[%s] failed", user.UUID)
			}
		}
		return nil
	})
}
