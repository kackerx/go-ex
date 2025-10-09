package ratelimit

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
)

type FixWindowLimiter struct {
	// 窗口起始时间
	timestamp int64
	// 时间窗口大小
	interval int64
	// 窗口内允许通过的最大请求数量
	rate int64
	// 当前窗口请求数量, 需重置
	cnt int64
}

func (t *FixWindowLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		current := time.Now().UnixNano()
		ts := atomic.LoadInt64(&t.timestamp)
		cnt := atomic.LoadInt64(&t.cnt)

		if ts+t.interval < current {
			if atomic.CompareAndSwapInt64(&t.timestamp, ts, current) {
				atomic.StoreInt64(&t.cnt, 0)
			}
		}

		cnt = atomic.AddInt64(&t.cnt, 1)
		defer atomic.AddInt64(&t.cnt, -1)

		if cnt > t.rate {
			err = errors.New("瓶颈")
			return
		}

		return handler(ctx, req)
	}
}
