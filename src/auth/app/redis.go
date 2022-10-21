package app

import (
	"auth/config"

	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.HostPort,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DbNumber,
	})
	return rdb
}
