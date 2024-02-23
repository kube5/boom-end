package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

const DefaultTimeout = 5

type DB struct {
	*gorm.DB
}

func NewDB(config *config.DB, log *logrus.Logger) (*DB, error) {
	dialector, err := newMysqlDB(config, &config.Master)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(
			log,
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  If(config.EnableLog, logger.Info, logger.Warn),
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		return nil, err
	}

	if config.EnableSlave {
		slaveDialector, err := newMysqlDB(config, &config.Slave)
		if err != nil {
			return nil, err
		}
		err = gormDB.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{slaveDialector},
		}))
		if err != nil {
			return nil, err
		}
	}

	db := &DB{gormDB}
	return db, nil
}

func newMysqlDB(cfg *config.DB, dbInfo *config.DBInfo) (gorm.Dialector, error) {
	dsn := getDSN(*dbInfo)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime.Duration)

	return mysql.New(mysql.Config{
		Conn: db,
	}), nil
}

func getDSN(cfg config.DBInfo) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&&parseTime=true&loc=Local&timeout=%ds",
		cfg.User,
		cfg.Password,
		cfg.IP,
		cfg.Port,
		cfg.DBName,
		DefaultTimeout)
}

func If(isTrue bool, a, b logger.LogLevel) logger.LogLevel {
	if isTrue {
		return a
	}
	return b
}
