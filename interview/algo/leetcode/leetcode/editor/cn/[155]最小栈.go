package main

// leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	data     *stack
	minStack *stack
}

func Constructor() MinStack {
	return MinStack{
		data:     &stack{},
		minStack: &stack{},
	}
}

func (this *MinStack) Push(val int) {
	if this.minStack.IsEmpty() {
		this.minStack.Push(val)
	} else if !this.minStack.IsEmpty() && val <= this.minStack.Peek() {
		this.minStack.Push(val)
	}
	this.data.Push(val)
}

func (this *MinStack) Pop() {
	pop := this.data.Pop()
	if !this.minStack.IsEmpty() && this.minStack.Peek() >= pop {
		this.minStack.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.data.Peek()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Peek()
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

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// leetcode submit region end(Prohibit modification and deletion)
