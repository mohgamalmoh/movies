package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedisClient(rdb *redis.Client) Redis {
	return Redis{rdb: rdb}
}

// Initialize a Redis client
var ctx = context.Background()

// Set a key-value pair in Redis
func (c Redis) SetCache(key string, value string, timeout time.Duration) error {
	err := c.rdb.Set(ctx, key, value, timeout).Err() // TTL of 10 minutes
	if err != nil {
		return err
	}
	return nil
}

// Get a value from Redis
func (c Redis) GetCache(key string) (val string, err error) {
	val, err = c.rdb.Get(ctx, key).Result()
	return
}
