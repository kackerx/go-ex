package repository

import (
	"context"

	"go-bobo/infra/persistence/dal"
)

type UserRepositoryImpl struct {
	dal *dal.UserDal
}

func (u *UserRepositoryImpl) GetUserList(ctx context.Context) {
	// TODO implement me
	panic("implement me")
}

func NewUserRepositoryImpl(dal *dal.UserDal) *UserRepositoryImpl {
	return &UserRepositoryImpl{dal: dal}
}
