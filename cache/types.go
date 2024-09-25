package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, val any, expiration time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
}

type item struct {
	val      any
	deadline time.Time
}

func (i *item) deadlineBefore() bool {
	return !i.deadline.IsZero() && i.deadline.Before(time.Now())
}
