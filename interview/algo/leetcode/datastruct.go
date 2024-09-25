package leetcode

import (
	"fmt"
	"strings"
)

type Stack[T comparable] []T

func (s *Stack[T]) Push(e T) {
	*s = append(*s, e)
}

func (s *Stack[T]) Pop() (e T) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("^")
	for i := s.Len() - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%v->", (*s)[i]))
	}

	return sb.String()
}

type ListNode struct {
	Val  int
	Next *ListNode
}
