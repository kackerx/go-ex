package dal

import (
	"gorm.io/gorm"

	"go-bobo/infra/persistence/dal/po"
)

type UserDal struct {
	db *gorm.DB
}

func NewUserDal(db *gorm.DB) *UserDal {
	return &UserDal{db: db}
}

func (u *UserDal) Migrate() (err error) {
	return u.db.AutoMigrate(&po.User{})
}
