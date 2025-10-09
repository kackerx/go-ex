package cache

import (
	"context"
	_ "embed"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	ErrFailedToPreemptLock = errors.New("reids-lock: 抢锁失败")
	ErrUnLockNotFount      = errors.New("redis-locl: 解锁失败, 锁不存在")

	//go:embed lua/unlock.lua
	luaUnlock string
	//go:embed lua/refresh.lua
	luaRefresh string
)

// Client 对redis.Cmdable二次封装
type Client struct {
	client redis.Cmdable
}

func NewClient(client redis.Cmdable) *Client {
	return &Client{client: client}
}

func (c *Client) TryLock(ctx context.Context, key string, expiration time.Duration) (lock *Lock, err error) {
	val := uuid.New().String()
	ok, err := c.client.SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		return
	}

	if !ok {
		return nil, ErrFailedToPreemptLock
	}

	return &Lock{
		client:     c.client,
		key:        key,
		val:        val,
		expiration: expiration,
	}, nil
}

type Lock struct {
	client     redis.Cmdable
	key        string
	val        string
	expiration time.Duration
}

func (l *Lock) Refresh(ctx context.Context) (err error) {
	_, err = l.client.Eval(ctx, luaRefresh, []string{l.key}, l.val, l.expiration.Seconds()).Int64()
	return err
}

func (l *Lock) UnLock(ctx context.Context, key string) (err error) {
	cnt, err := l.client.Eval(ctx, luaUnlock, []string{key}, l.val).Int64()
	if err != nil {
		return
	}

	if cnt != 1 {
		return ErrUnLockNotFount
	}

	return
}
