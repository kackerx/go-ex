package main


// leetcode submit region begin(Prohibit modification and deletion)
type stack []int

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len((*s))-1]
	return
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(e int) {
	*s = append(*s, e)
}
type MinStack struct {
	data     *stack
	minStack *stack
}

func Constructor() MinStack {
	return MinStack{&stack{}, &stack{}}
}

func (m *MinStack) Push(val int) {
	if m.minStack.isEmpty() || m.minStack.peek() >= val {
		m.minStack.push(val)
	}

	m.data.push(val)
}

func (m *MinStack) Pop() {
	if m.Top() == m.GetMin() {
		m.minStack.pop()
	}
	m.data.pop()
}

func (m *MinStack) Top() int {
	return m.data.peek()
}

func (m *MinStack) GetMin() int {
	return m.minStack.peek()
}

// ---

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// leetcode submit region end(Prohibit modification and deletion)
