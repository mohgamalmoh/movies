package config

import (
	"github.com/go-redis/redis/v8"
	"strconv"
)

const (
	redisHost     = "redis"
	redisPort     = 6379
	redisUser     = "default"
	redisPassword = "pass"
	redisDBName   = 0
)

func RedisConnection() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + strconv.FormatUint(redisPort, 10), // Redis server address
		Username: redisUser,
		Password: redisPassword, // No password set
		DB:       redisDBName,   // Use default DB
	})
	return db
}
