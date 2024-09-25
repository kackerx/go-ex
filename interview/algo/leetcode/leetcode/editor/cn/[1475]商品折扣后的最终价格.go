package main

// leetcode submit region begin(Prohibit modification and deletion)
type stack []int

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) Pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len((*s))-1]
	return
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(e int) {
	*s = append(*s, e)
}

func finalPrices(prices []int) []int {
	st := stack{}
	res := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		for !st.IsEmpty() {
			if prices[i] <= prices[st.Peek()] {
				topIndex := st.Pop()
				res[topIndex] = prices[topIndex] - prices[i]
			} else {
				break
			}
		}
		st.Push(i)
	}

	for !st.IsEmpty() {
		topIndex := st.Pop()
		res[topIndex] = prices[topIndex]
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
