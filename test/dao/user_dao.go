package sqlmock

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
}

func NewUserDao(db *gorm.DB) UserRepository {
	return &UserDao{db: db}
}

type UserDao struct {
	db *gorm.DB
}

func (u *UserDao) Create(ctx context.Context, user *User) error {
	return u.db.Create(user).Error
}
