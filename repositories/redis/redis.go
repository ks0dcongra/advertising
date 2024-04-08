package redis

import (
	"context"
	"advertising/configs"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepositoryInterface interface {
	Set(keyName string, value []byte, expiration time.Duration) error
	Get(keyName string) ([]byte, error)
}

type RedisRepositoryImpl struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisRepository() RedisRepositoryInterface {
	return &RedisRepositoryImpl{
		Client: configs.RedisConn,
		Ctx:    context.Background(),
	}
}

func (r *RedisRepositoryImpl) Set(keyName string, value []byte, expiration time.Duration) error {
	err := r.Client.Set(r.Ctx, keyName, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepositoryImpl) Get(keyName string) ([]byte, error) {
	val, err := r.Client.Get(r.Ctx, keyName).Bytes()
	if err != nil {
		return nil, err
	}
	return val, err
}
