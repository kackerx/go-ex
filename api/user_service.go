package main

import "context"

var _ CRUDService[User, UserCreateReq, UserUpdateReq] = (*UserService)(nil)

func NewUserRepo() UserRepository {
	return nil
}

type UserRepository interface {
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type UserService struct {
	userRepo UserRepository
}

func (s *UserService) Create(ctx context.Context, req *UserCreateReq) (*User, error) {
	panic("unimplemented")
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (s *UserService) Get(ctx context.Context, id string) (*User, error) {
	panic("unimplemented")
}

func (s *UserService) List(ctx context.Context, req ListRequest) (*ListResponse[User], error) {
	panic("unimplemented")
}

func (s *UserService) Update(ctx context.Context, req *UserUpdateReq) (*User, error) {
	panic("unimplemented")
}
