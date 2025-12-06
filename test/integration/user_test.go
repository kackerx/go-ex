package integration

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("kackerx")
	defer func() {
		fmt.Println("end")
	}()

	panic("hehe")
}
