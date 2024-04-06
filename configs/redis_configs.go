package configs

import (
	"context"
	"errors"
	"os"
	"time"

	"strconv"

	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client

type redisConfig struct {
	Host            string
	Port            string
	Username        string
	Password        string
	MaxConnTimeout  int
	MaxWriteTimeout int
	MaxReadTimeout  int
	MaxPoolSize     int
	MinIdleCons     int
	RetryInterval   int64
}

func getRedisConfig() (*redisConfig, error) {
	var err error
	var host string
	var port string
	var username string
	var password string
	var maxConnTimeout int
	var maxWriteTimeout int
	var maxReadTimeout int
	var maxPoolSize int
	var minIdleCons int
	var redisRetryInterval int64

	host = os.Getenv("REDIS_HOST")
	if len(host) == 0 {
		return nil, errors.New("cannot get current redis host from env")
	}

	port = os.Getenv("REDIS_PORT")
	if len(port) == 0 {
		return nil, errors.New("cannot get current redis port from env")
	}

	username = os.Getenv("REDIS_USERNAME")
	if len(username) == 0 {
		return nil, errors.New("cannot get current redis username from env")
	}

	password = os.Getenv("REDIS_PASSWORD")
	if len(password) == 0 {
		return nil, errors.New("cannot get current redis password from env")
	}

	if len(os.Getenv("REDIS_RETRYINTERVAL")) == 0 {
		return nil, errors.New("cannot get current redis retry interval from env")
	}
	redisRetryInterval, err = strconv.ParseInt(os.Getenv("REDIS_RETRYINTERVAL"), 10, 64)
	if err != nil {
		return nil, err
	}

	if len(os.Getenv("REDIS_MAXCONNTIMEOUT")) == 0 {
		return nil, errors.New("cannot get current max conn timeout from env")
	}
	maxConnTimeout, err = strconv.Atoi(os.Getenv("REDIS_MAXCONNTIMEOUT"))
	if err != nil {
		return nil, err
	}

	if len(os.Getenv("REDIS_MAXWRITETIMEOUT")) == 0 {
		return nil, errors.New("cannot get current max write timeout from env")
	}
	maxWriteTimeout, err = strconv.Atoi(os.Getenv("REDIS_MAXWRITETIMEOUT"))
	if err != nil {
		return nil, err
	}

	if len(os.Getenv("REDIS_MAXREADTIMEOUT")) == 0 {
		return nil, errors.New("cannot get current max read timeout from env")
	}
	maxReadTimeout, err = strconv.Atoi(os.Getenv("REDIS_MAXREADTIMEOUT"))
	if err != nil {
		return nil, err
	}

	if len(os.Getenv("REDIS_MAXPOOLSIZE")) == 0 {
		return nil, errors.New("cannot get current redis max pool size from env")
	}
	maxPoolSize, err = strconv.Atoi(os.Getenv("REDIS_MAXPOOLSIZE"))
	if err != nil {
		return nil, err
	}

	if len(os.Getenv("REDIS_MINIDLECONS")) == 0 {
		return nil, errors.New("cannot get current redis minidlecons from env")
	}
	minIdleCons, err = strconv.Atoi(os.Getenv("REDIS_MINIDLECONS"))
	if err != nil {
		return nil, err
	}

	return &redisConfig{
		Host:            host,
		Port:            port,
		Username:        username,
		Password:        password,
		MaxConnTimeout:  maxConnTimeout,
		MaxWriteTimeout: maxWriteTimeout,
		MaxReadTimeout:  maxReadTimeout,
		MaxPoolSize:     maxPoolSize,
		MinIdleCons:     minIdleCons,
		RetryInterval:   redisRetryInterval,
	}, nil
}


func RedisSetup() error {	
	
	RedisConn = redis.NewClient(&redis.Options{
		Addr:         Cfg.Redis.Host + ":" + Cfg.Redis.Port,
		Username:     Cfg.Redis.Username,
		Password:     Cfg.Redis.Password,
		DB:           0,
		DialTimeout:  time.Duration(Cfg.Redis.MaxConnTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(Cfg.Redis.MaxWriteTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(Cfg.Redis.MaxReadTimeout) * time.Millisecond,
		PoolSize:     Cfg.Redis.MaxPoolSize,
		MinIdleConns: Cfg.Redis.MinIdleCons,
	})
	_, err := RedisConn.Do(context.Background(), "ping").Result()
	if err != nil {
		return err
	}

	return nil
}