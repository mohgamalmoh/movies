package config

import (
	"github.com/go-redis/redis/v8"
)

func RedisConnection(cfg RedisDatabase) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Username: cfg.User,
		Password: cfg.Password,
		DB:       cfg.Name,
	})
	return db
}
