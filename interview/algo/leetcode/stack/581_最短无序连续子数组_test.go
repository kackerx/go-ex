package stack

import (
	"fmt"
	"math"
	"testing"
)

func Test581(t *testing.T) {
	s := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(s))
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := math.MaxInt, math.MinInt
	incr := &stack{}
	for i := 0; i < n; i++ {
		for !incr.IsEmpty() && nums[incr.Peek()] > nums[i] {
			left = min(left, incr.Pop())
		}
		incr.Push(i)
	}

	decr := &stack{}
	for i := n - 1; i >= 0; i-- {
		for !decr.IsEmpty() && nums[decr.Peek()] < nums[i] {
			right = max(right, decr.Pop())
		}
		decr.Push(i)
	}

	fmt.Println(incr)
	fmt.Println(decr)

	if left == math.MaxInt && right == math.MinInt {
		return 0
	}
	return right - left + 1
}

type stack struct {
	data []int
}

func (s *stack) Push(e int) {
	s.data = append(s.data, e)
}

func (s *stack) Pop() int {
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret
}

func (s *stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}
