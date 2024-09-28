package cache

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-bobo/cache/mocks"
)

func TestRedisLock(t *testing.T) {
	ctrl := gomock.NewController(t)

	testCases := []struct {
		name       string
		key        string
		expairtion time.Duration
		mock       func(ctrl *gomock.Controller) redis.Cmdable

		wantLock *Lock
		wantErr  error
	}{
		{
			name:       "lock",
			key:        "keyyy",
			expairtion: time.Second * 5,
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				cmd.EXPECT().SetNX(
					context.Background(),
					"keyyy",
					gomock.Any(),
					time.Second*5,
				).Return(redis.NewBoolResult(false, context.DeadlineExceeded))

				return cmd
			},
			wantLock: &Lock{
				key: "keyyy",
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cli := NewClient(tc.mock(ctrl))
			lock, err := cli.TryLock(context.Background(), tc.key, tc.expairtion)
			// assert.Equal(t, tc.wantErr, err)
			// require.Equal(t, tc.wantErr, err)
			// require.NoError(t, err)
			// require.Error(t, err)
			require.ErrorIs(t, err, context.DeadlineExceeded, "返回预期的超时错误")
			if err != nil {
				return
			}

			assert.Equal(t, tc.wantLock.key, lock.key)
			// assert.NotEmpty(t, lock.val)
		})
	}
}
