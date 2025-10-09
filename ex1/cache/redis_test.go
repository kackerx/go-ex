package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	"go-bobo/cache/mocks"
)

func TestRedisCache_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name string

		mock func() redis.Cmdable

		key        string
		value      string
		expiration time.Duration

		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "set value success",
			mock: func() redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)

				status := redis.NewStatusCmd(context.Background())
				status.SetVal("OK")

				cmd.EXPECT().
					Set(context.Background(), "key1", "val1", time.Second).
					Return(status)

				return cmd
			},
			key:        "key1",
			value:      "val1",
			expiration: time.Second,
			wantErr:    nil,
		},
		{
			name: "timeout",
			mock: func() redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)

				status := redis.NewStatusCmd(context.Background())
				status.SetErr(context.DeadlineExceeded)

				cmd.EXPECT().
					Set(context.Background(), "key2", "val2", time.Second).
					Return(status)

				return cmd
			},
			key:        "key2",
			value:      "val2",
			expiration: time.Second,
			wantErr:    context.DeadlineExceeded,
		},
		{
			name: "unexpected msg",
			mock: func() redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)

				status := redis.NewStatusCmd(context.Background())
				status.SetVal("no ok")

				cmd.EXPECT().
					Set(context.Background(), "key2", "val2", time.Second).
					Return(status)

				return cmd
			},
			key:        "key2",
			value:      "val2",
			expiration: time.Second,
			wantErr:    fmt.Errorf("%w, %s", ErrCacheSetFail, "key2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRedisCache(tt.mock())
			err := c.Set(context.Background(), tt.key, tt.value, tt.expiration)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
