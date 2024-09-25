package ratelimit

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type SlideWindowLimiter struct {
	queue    list.List
	interval int64 `json:"interval"`
	rate     int

	mu sync.Mutex
}

func (s *SlideWindowLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		now := time.Now().UnixNano()
		boundary := now - s.interval

		s.mu.Lock()
		ts := s.queue.Front()
		// 删除不在窗口的数据
		for ts.Value.(int64) < boundary {
			s.queue.Remove(ts)
			ts = s.queue.Front()
		}

		curLen := s.queue.Len()
		s.mu.Unlock()
		if curLen >= s.rate {
			err = errors.New("瓶颈")
			return
		}

		resp, err = handler(ctx, req)
		s.queue.PushBack(now)

		return
	}
}

func (s *SlideWindowLimiter) name() {

}

type Stack []int

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}
