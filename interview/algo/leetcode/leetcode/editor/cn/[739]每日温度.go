package main

// leetcode submit region begin(Prohibit modification and deletion)

type Stack struct {
	data []int
}

func (s *Stack) Pop() int {
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}

func (s *Stack) Push(e int) {
	s.data = append(s.data, e)
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	s := &Stack{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		for !s.IsEmpty() && temperatures[s.Peek()] <= temperatures[i] {
			s.Pop()
		}

		if !s.IsEmpty() {
			res[i] = s.Peek() - i
		}

		s.Push(i)
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
