package main

// leetcode submit region begin(Prohibit modification and deletion)
type stock []int

func (s *stock) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stock) pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len((*s))-1]
	return
}

func (s *stock) isEmpty() bool {
	return len(*s) == 0
}

func (s *stock) push(e int) {
	*s = append(*s, e)
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	st := &stock{}

	for i := 0; i < len(temperatures); i++ {
		for !st.isEmpty() {
			if temperatures[i] > temperatures[st.peek()] {
				index := st.pop()
				res[index] = i - index
			} else {
				break
			}
		}

		st.push(i)
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
