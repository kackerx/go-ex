package ratelimit

import (
	"context"
	_ "embed"
	"errors"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

//go:embed lua/fix_window.lua
var luaFixWindow string

type RedisFixWindowLimiter struct {
	client  redis.Cmdable
	service string
}

func (r *RedisFixWindowLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		limit, err := r.limit(ctx)
		if err != nil {
			return
		}

		if limit {
			err = errors.New("执行限流")
			return
		}

		return handler(ctx, req)
	}
}

func (r *RedisFixWindowLimiter) limit(ctx context.Context) (bool, error) {
	return r.client.Eval(ctx, luaFixWindow, []string{r.service}).Bool()
}
