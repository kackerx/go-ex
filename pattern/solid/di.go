package main

type User struct {
	Name string
}

type UserService struct {
	userRepo Repository
}

func NewUserService(userRepo Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type Repository interface {
	GetUser(id int) *User
}

type userApplication struct {
	userService UserService
	userRepo    Repository
}

type userRepo struct {
}

func (u *userRepo) GetUser(id int) *User {
	return &User{
		Name: "John Doe",
	}
}

func main() {
	NewUserService(&userRepo{})
}
