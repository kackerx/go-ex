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

func (l *ListNode) String() string {
	sb := strings.Builder{}
	cur := l
	for cur != nil {
		sb.WriteString(fmt.Sprintf("%d->", cur.Val))
		cur = cur.Next
	}

	return sb.String()
}

type minHeap []*ListNode

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i].Val < (*m)[j].Val
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() any {
	ret := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return ret
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Push(e int) {
	*s = append(*s, e)
}

func (s *stack) Pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) String() string {
	sb := strings.Builder{}
	for _, v := range *s {
		sb.WriteString(fmt.Sprintf("%d->", v))
	}

	return sb.String()
}

type set map[int]struct{}

func (s *set) Has(v int) bool {
	if _, ok := (*s)[v]; ok {
		return true
	}

	return false
}

func (s *set) Add(v int) {
	(*s)[v] = struct{}{}
}

func (s *set) Rem(v int) {
	delete(*s, v)
}
