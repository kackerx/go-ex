package cache

import (
	"context"
	"time"
)

type WriteThrough struct {
	Cache
	StoreFunc func(ctx context.Context, key string, val any) error
}

func (w *WriteThrough) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	if err := w.StoreFunc(ctx, key, val); err != nil {
		return err
	}

	return w.Cache.Set(ctx, key, val, expiration)
}
