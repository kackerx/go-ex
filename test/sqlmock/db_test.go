package sqlmock

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDB(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	testCases := []struct {
		name string
		user *User

		mock func(t *testing.T) *sql.DB

		ctx context.Context

		wantID  uint
		wantErr error
	}{
		{
			name: "插入成功",
			ctx:  context.Background(),
			user: &User{
				ID:    1,
				Name:  "test",
				Email: "test@test.com",
			},
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				require.NoError(err)
				result := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT INTO `users` .*").
					WillReturnResult(result)
				return mockDB
			},
			wantID:  uint(1),
			wantErr: nil,
		},
		{
			name: "邮箱冲突",
			ctx:  context.Background(),
			user: &User{
				Name:  "test",
				Email: "test@test.com",
			},
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				require.NoError(err)

				result := sqlmock.NewErrorResult(errors.New("邮箱冲突"))
				mock.ExpectExec("INSERT INTO `users` .*").
					WillReturnResult(result)

				return mockDB
			},
			wantErr: errors.New("邮箱冲突"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, err := gorm.Open(mysql.New(mysql.Config{
				Conn:                      tt.mock(t),
				SkipInitializeWithVersion: true, // 初始化时跳过版本检查, SELECT VERSION();
			}), &gorm.Config{
				DisableAutomaticPing:   true, // 禁用自动ping, 因为sqlmock不支持ping
				SkipDefaultTransaction: true, // 默认单条sql也会开启事务, BEGIN;someSQL;COMMIT;
			})
			require.NoError(err)

			userDao := NewUserDao(db)
			err = userDao.Create(tt.ctx, tt.user)

			assert.Equal(tt.wantErr, err)
			assert.Equal(tt.wantID, tt.user.ID)
		})
	}
}
