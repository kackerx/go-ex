package main

import (
	"fmt"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
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

func simplifyPath(path string) string {
	s := &Stack[string]{}
	pathList := strings.Split(path, "/")
	slices.Reverse(pathList)
	for _, item := range pathList {
		if item == "." || item == "" {
			continue
		}

		s.Push(item)
	}

	s2 := &Stack[string]{}
	for !s.IsEmpty() {
		v := s.Pop()
		if v == ".." {
			if s2.IsEmpty() {
				continue
			}

		} else {
			s2.Push(v)
		}
	}

	var res string
	for !s2.IsEmpty() {
		res = fmt.Sprintf("/%s", s2.Pop()) + res
	}

	if res == "" {
		return "/"
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
