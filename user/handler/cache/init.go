package cache

import (
	"github.com/Mu-Exchange/Mu-End/common/redis"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/user/pkg/base"
)

func init() {
	base.RegisterComponents(NewRedis, NewUserCache, NewTgUserCache)
}

func NewRedis(cfg repo.CommonComponents) (*redis.Cache, error) {
	return redis.New(cfg.Cfg.Redis)
}
