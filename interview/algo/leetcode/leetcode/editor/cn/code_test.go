package main

import "testing"

func Test167(t *testing.T) {
}

type Stack []int

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Len() int {
	return len(*s)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
