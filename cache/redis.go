package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	_ "github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
)

var (
	ErrCacheSetFail = errors.New("cache: set fail")
)

type RedisCache struct {
	client redis.Cmdable
}

func NewRedisCache(client redis.Cmdable) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	result, err := r.client.Set(ctx, key, val, expiration).Result()
	if err != nil {
		return err
	}

	if result != "OK" {
		return fmt.Errorf("%w, %s", ErrCacheSetFail, key)
	}

	return nil
}

func (r *RedisCache) Get(ctx context.Context, key string) (any, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Delete(key string) error {
	// TODO implement me
	panic("implement me")
}
