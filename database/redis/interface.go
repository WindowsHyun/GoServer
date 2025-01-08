package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	GetClient() *redis.Client
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) int64
	HSet(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) error
	HGet(ctx context.Context, key string, field string, dest interface{}) error
	HDel(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key string, field string) bool
	LPush(ctx context.Context, key string, values ...interface{}) error
	RPop(ctx context.Context, key string, dest interface{}) error
	BRPop(ctx context.Context, keys []string, timeout time.Duration, dest interface{}) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
	Transaction(ctx context.Context, fn func(tx *redis.Tx) error) error
}
