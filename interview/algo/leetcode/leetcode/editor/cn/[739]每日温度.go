package main

// leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(temperatures []int) []int {
	stk := &stack{}
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for !stk.IsEmpty() && temperatures[stk.Peek()] <= temperatures[i] {
			stk.Pop()
		}

		if stk.IsEmpty() {
			res[i] = 0
		} else {
			res[i] = stk.Peek() - i
		}

		stk.Push(i)
	}

	return res
}

type stack []int

func (s *stack) Len() int {
	return len((*s))
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Push(e int) {
	(*s) = append((*s), e)
}

func (s *stack) Pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
