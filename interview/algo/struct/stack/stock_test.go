package stack

import (
	"fmt"
	"testing"
)

func TestArrayStock(t *testing.T) {
	s := NewArrayStock[int](10)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println("peek", s.Peek())

	fmt.Println(s.Pop())
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
}

func TestLinkedListStack(t *testing.T) {
	s := NewLinkedListStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println("peek", s.Peek())

	fmt.Println(s.Pop())
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
}
