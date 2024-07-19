package repository

import "context"

type UserRepository interface {
	GetUserList(ctx context.Context)
}
