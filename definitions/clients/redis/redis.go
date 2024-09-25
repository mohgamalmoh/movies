package redis

import "time"

type RedisClient interface {
	SetCache(key, value string, timeout time.Duration) error
	GetCache(key string) (string, error)
}
