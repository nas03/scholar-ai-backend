package utils

import (
	"context"
	"time"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/redis/go-redis/v9"
)

// IRedisCache defines the interface for Redis cache operations.
type IRedisCache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, data any) error
	SetEx(ctx context.Context, key string, data any, exp time.Duration) error
	Del(ctx context.Context, key string) error
}

// RedisCache implements IRedisCache using a Redis client.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new RedisCache instance.
func NewRedisCache() IRedisCache {
	return &RedisCache{
		client: global.Redis,
	}
}

// Get retrieves a value by key.
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Set stores a value without expiration.
func (r *RedisCache) Set(ctx context.Context, key string, data any) error {
	return r.client.Set(ctx, key, data, 0).Err()
}

// SetEx stores a value with expiration time.
func (r *RedisCache) SetEx(ctx context.Context, key string, data any, exp time.Duration) error {
	return r.client.SetEx(ctx, key, data, exp).Err()
}

// Del deletes a key.
func (r *RedisCache) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Keys returns all keys matching the pattern.
func (r *RedisCache) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}
