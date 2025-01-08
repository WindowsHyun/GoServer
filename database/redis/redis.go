package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *RedisRepository) GetClient() *redis.Client {
	return r.Client
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisRepository) Get(ctx context.Context, key string, dest interface{}) error {
	return r.Client.Get(ctx, key).Scan(dest)
}

func (r *RedisRepository) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}

func (r *RedisRepository) Exists(ctx context.Context, keys ...string) int64 {
	return r.Client.Exists(ctx, keys...).Val()
}

func (r *RedisRepository) HSet(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) error {
	err := r.Client.HSet(ctx, key, field, value).Err()
	if err != nil {
		return err
	}
	if expiration > 0 {
		return r.Expire(ctx, key, expiration)
	}
	return nil
}

func (r *RedisRepository) HGet(ctx context.Context, key string, field string, dest interface{}) error {
	return r.Client.HGet(ctx, key, field).Scan(dest)
}

func (r *RedisRepository) HDel(ctx context.Context, key string, fields ...string) error {
	return r.Client.HDel(ctx, key, fields...).Err()
}

func (r *RedisRepository) HExists(ctx context.Context, key string, field string) bool {
	return r.Client.HExists(ctx, key, field).Val()
}

func (r *RedisRepository) LPush(ctx context.Context, key string, values ...interface{}) error {
	return r.Client.LPush(ctx, key, values...).Err()
}

func (r *RedisRepository) RPop(ctx context.Context, key string, dest interface{}) error {
	val, err := r.Client.RPop(ctx, key).Result()
	if err != nil {
		return err
	}
	return scanValue(val, dest)
}

func (r *RedisRepository) BRPop(ctx context.Context, keys []string, timeout time.Duration, dest interface{}) error {
	val, err := r.Client.BRPop(ctx, timeout, keys...).Result()
	if err != nil {
		return err
	}
	return scanValue(val[1], dest)
}

func (r *RedisRepository) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.Client.Expire(ctx, key, expiration).Err()
}

func (r *RedisRepository) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

func (r *RedisRepository) Transaction(ctx context.Context, fn func(tx *redis.Tx) error) error {
	return r.Client.Watch(ctx, func(tx *redis.Tx) error {
		return fn(tx)
	})
}

func scanValue(val interface{}, dest interface{}) error {
	return redis.NewStringResult(val.(string), nil).Scan(dest)
}
