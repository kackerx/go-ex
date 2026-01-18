package main

import (
	"errors"
	"fmt"
	"time"
)

type Link struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}

type User struct {
	Name string
	Age  int
}

func main() {
	us := &User{
		Name: "kacker",
		Age:  0,
	}

	fmt.Println(us)
}

func foo(u *User) error {
	return errors.New(u.Name)
}
