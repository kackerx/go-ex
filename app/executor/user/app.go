package user

import (
	"context"

	"go-bobo/domain/aggregate/service"
	"go-bobo/infra/persistence/repository"
)

type UserApplication struct {
	userDomainService *service.UserDomainService
}

func NewUserApplication(userRepo *repository.UserRepositoryImpl) *UserApplication {
	return &UserApplication{
		userDomainService: service.NewUserDomainService(userRepo),
	}
}

func (u *UserApplication) ListUser(ctx context.Context) {
	u.userDomainService.GetUserList(ctx)
}
