package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrKeyNotFound = errors.New("cache: key未找到")
)

type BuildInMapCache struct {
	data  map[string]*item
	close chan struct{}

	mu sync.RWMutex
}

func NewBuildInMapCache(internal time.Duration) *BuildInMapCache {
	res := &BuildInMapCache{
		data:  make(map[string]*item, 100),
		close: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(internal)
		for {
			select {
			case t := <-ticker.C:
				i := 0
				for key, val := range res.data {
					if i > 10000 {
						break
					}

					if !val.deadline.IsZero() && !val.deadline.Before(t) {
						delete(res.data, key)
					}
					i++
				}
			case <-res.close:
				return
			}
		}
	}()

	return res
}

func (b *BuildInMapCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	dl := time.Time{}
	if expiration != 0 {
		dl = time.Now().Add(expiration)
	}
	b.data[key] = &item{
		val:      val,
		deadline: dl,
	}

	return nil
}

func (b *BuildInMapCache) Get(key string) (any, error) {
	b.mu.RLock()
	if v, ok := b.data[key]; ok {
		b.mu.RUnlock()
		if v.deadlineBefore() {
			b.mu.Lock()
			defer b.mu.Unlock()
			v, ok = b.data[key]
			if !ok {
				return nil, fmt.Errorf("%w: key: %s", ErrKeyNotFound, key)
			}

			if v.deadlineBefore() {
				delete(b.data, key)
				return nil, fmt.Errorf("%w: key: %s", ErrKeyNotFound, key)
			}
		} else {
			return v, nil
		}
	}

	return nil, fmt.Errorf("%w, key: %s", ErrKeyNotFound, key)
}

func (b *BuildInMapCache) Delete(key string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.data, key)

	return nil
}

func (b *BuildInMapCache) delete(key string) {

}

func (b *BuildInMapCache) Close() error {
	// 用sync.Once或者select避免多次调用的panic
	close(b.close)
	return nil
}
