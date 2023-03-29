package database

import (
	"go-fiber-practice/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	return rdb
}
