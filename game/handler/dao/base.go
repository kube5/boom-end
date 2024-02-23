package dao

import (
	"github.com/Mu-Exchange/Mu-End/common/db"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"go.uber.org/fx"
)

func NewBaseDao(lc fx.Lifecycle, sd fx.Shutdowner, cfg repo.CommonComponents) (*db.DB, error) {
	return db.NewDB(cfg.Cfg.DB, cfg.Logger)
}
