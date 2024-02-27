/*
Package cache has all cache-related implementations.
*/
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"literank.com/rest-books/infrastructure/config"
)

const (
	defaultTimeout = time.Second * 10
	defaultTTL     = time.Hour * 1
)

// RedisCache implements cache with redis
type RedisCache struct {
	c redis.UniversalClient
}

// NewRedisCache constructs a new RedisCache
func NewRedisCache(c *config.CacheConfig) *RedisCache {
	timeout := defaultTimeout
	if c.Timeout > 0 {
		timeout = time.Second * time.Duration(c.Timeout)
	}
	r := redis.NewClient(&redis.Options{
		Addr:         c.Address,
		Password:     c.Password,
		DB:           c.DB,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	})
	return &RedisCache{
		c: r,
	}
}

// Save sets key and value into the cache
func (r *RedisCache) Save(ctx context.Context, key, value string) error {
	if _, err := r.c.Set(ctx, key, value, defaultTTL).Result(); err != nil {
		return err
	}
	return nil
}

// Load reads the value by the key
func (r *RedisCache) Load(ctx context.Context, key string) (string, error) {
	value, err := r.c.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return value, nil
}
