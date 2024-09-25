package ratelimit

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type LeakyBucketLimiter struct {
	producer *time.Ticker
	closeCh  chan struct{}
}

func NewLeakyBucketLimiter(interval time.Duration) *LeakyBucketLimiter {
	ticker := time.NewTicker(interval)
	closeCh := make(chan struct{})

	return &LeakyBucketLimiter{producer: ticker, closeCh: closeCh}
}

func (l *LeakyBucketLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		select {
		case <-l.producer.C:
			resp, err = handler(ctx, req)
		case <-ctx.Done():
			err = ctx.Err()
			return
		case <-l.closeCh:
			l.producer.Stop()
			return
		}

		return
	}
}

func (l *LeakyBucketLimiter) Close() {
	close(l.closeCh)
}
