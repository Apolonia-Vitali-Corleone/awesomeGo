// Package database pkg/database/redis.go
package database

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go1/config"
)

func InitRedis(cfg config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb, nil
}
