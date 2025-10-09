package argo

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Name string
}

func TestFoo(t *testing.T) {
	s := []int{1}
	a := &User{"kacker"}
	fmt.Printf("%p, %p\n", s, a)
	Foo(s, a)
}

func Foo(s []int, h *User) {
}

func TestSol(t *testing.T) {
	// IsUnique("abc中国")
	fmt.Println(reverseStr("abcde"))
	fmt.Println(reverseStr("123456"))
}

func IsUnique(s string) bool {
	for _, v := range s {
		fmt.Println(v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	return false
}

func reverseStr(s string) string {
	left, right := 0, len(s)-1
	str := []rune(s)
	for left <= right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}

	return string(str)
}

func sol() {
	letter, number := make(chan int), make(chan int)

	go func() {
		for i := range number {
			time.Sleep(time.Second * 1)
			fmt.Println(i)
			i++
			fmt.Println(i)
			if i == 28 {
				close(letter)
				close(number)
				return
			}
			letter <- i
		}
	}()

	number <- 1
	for i := range letter {
		l := i + 63
		fmt.Printf("%c\n", l)
		l++
		fmt.Printf("%c\n", l)
		number <- l - 63
	}

	fmt.Println("main end")
}
