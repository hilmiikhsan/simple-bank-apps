package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/simple-bank-apps/config"
	log "github.com/sirupsen/logrus"
)

type RedisConfig interface {
	Connect(ctx context.Context) (*redis.Client, error)
}

func NewRedisConfig(cfg *config.Config) RedisConfig {
	return &redisConfig{
		cfg: cfg,
	}
}

type redisConfig struct {
	cfg *config.Config
}

func (r *redisConfig) Connect(ctx context.Context) (*redis.Client, error) {
	timeout := time.Duration(r.cfg.RedisServer.Timeout) * time.Second

	rdb := redis.NewClient(&redis.Options{
		Addr:        r.cfg.RedisServer.Addr,
		Password:    r.cfg.RedisServer.Password,
		DialTimeout: timeout,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("cannot connect to redis")
		return nil, err
	}

	fmt.Printf("success connect to redis %s", rdb)
	return rdb, nil
}
