package service

import (
	"context"

	"go-bobo/domain/aggregate/repository"
)

type UserDomainService struct {
	userRepo repository.UserRepository
}

func NewUserDomainService(userRepo repository.UserRepository) *UserDomainService {
	return &UserDomainService{
		userRepo: userRepo,
	}
}

func (u *UserDomainService) GetUserList(ctx context.Context) {
	u.userRepo.GetUserList(ctx)
}
