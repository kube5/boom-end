package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/Mu-Exchange/Mu-End/common/config"
)

type Cache struct {
	*redis.Client
}

func NewRedisClient(cfg *config.Redis) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	}
	if cfg.TLS {
		opt.TLSConfig = &tls.Config{}
	}
	rds := redis.NewClient(opt)
	if _, err := rds.Ping(context.Background()).Result(); err != nil {
		fmt.Println("redis ping err:", err)
		return nil, err
	}

	return rds, nil
}

func New(cfg *config.Redis) (*Cache, error) {
	rds, err := NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}
	cache := &Cache{
		rds,
	}
	return cache, nil
}
