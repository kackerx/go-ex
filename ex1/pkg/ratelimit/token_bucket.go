package ratelimit

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type TokenBucketLimiter struct {
	tokens  chan struct{}
	closeCh chan struct{}
}

func NewTokenBucketLimiter(cap int, interval time.Duration) *TokenBucketLimiter {
	closeCh := make(chan struct{})
	ticker := time.NewTicker(interval)
	ch := make(chan struct{}, cap)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-closeCh:
				return
			case <-ticker.C:
				select {
				case ch <- struct{}{}:
				default:
					// 没人取令牌
				}
			}
		}
	}()

	return &TokenBucketLimiter{tokens: ch, closeCh: closeCh}
}

func (t *TokenBucketLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		select {
		case <-t.closeCh:
			// 关闭故障检测后, 可以直接访问也可以禁止访问, 看抉择
			resp, err = handler(ctx, req)
		case <-ctx.Done():
			err = ctx.Err()
			return
		case <-t.tokens:
			resp, err = handler(ctx, req)
			// default:
			// 	err = errors.New("触发限流")
			// 	return
		}

		return
	}
}

func (t *TokenBucketLimiter) Close() error {
	close(t.closeCh)
	return nil
}
