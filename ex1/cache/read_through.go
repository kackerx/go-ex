package cache

import "context"

type ReadThroughCache struct {
	Cache
	LoadFunc func(ctx context.Context, key string) (any, error)
}

func (r *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	get, err := r.Cache.Get(ctx, key)
	if err != nil {
		return nil, nil
	}

	return get, err
}
