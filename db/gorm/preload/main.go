package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/slice"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name  string  `gorm:"type:varchar(32);not null;"`
	Books []*Book `gorm:"foreignKey:UserID"`
}

type Book struct {
	gorm.Model

	UserID uint  `gorm:""`
	User   *User `gorm:"foreignKey:UserID"`
}

var (
	db  *gorm.DB
	err error
)

func init() {
	// Standard PostgreSQL DSN format: "host=... user=... password=... dbname=... port=... sslmode=..."
	dsn := "host=127.0.0.1 user=postgres password=Wasd4044 dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// The original error is wrapped, so we print the detailed reason.
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	db.AutoMigrate(&User{}, &Book{})
}

func main() {
	a := []int{1}
	a = slice.InsertAt(a, 0, 0)
	fmt.Println(a)
}

func createBook() {
	book := &Book{
		UserID: 1,
		User: &User{
			Name: "John",
		},
	}
	err = db.Create(book).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("book created")
}
