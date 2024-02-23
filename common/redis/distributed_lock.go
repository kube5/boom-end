package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Mu-Exchange/Mu-End/common/config"
)

type DistributedLock struct {
	redisClient *redis.Client
}

func NewDistributedLock(cfg *config.Redis) (*DistributedLock, error) {
	rds, err := NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}

	return &DistributedLock{
		redisClient: rds,
	}, nil
}

func (lock *DistributedLock) AcquireLock(key string, timeout time.Duration) bool {
	isSet, err := lock.redisClient.SetNX(context.Background(), key, "locked", timeout).Result()
	if err != nil {
		return false
	}

	return isSet
}

func (lock *DistributedLock) ReleaseLock(key string) bool {
	_, err := lock.redisClient.Del(context.Background(), key).Result()
	if err != nil {
		return false
	}

	return true
}
