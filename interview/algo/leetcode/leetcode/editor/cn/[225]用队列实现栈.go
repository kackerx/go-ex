package main

// leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	s1, s2 []int
}

func Constructor() MyStack {
	return MyStack{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.s1 = append(this.s1, x)
	for len(this.s2) != 0 {
		val := this.s2[0]
		this.s2 = this.s2[1:len(this.s2)]
		this.s1 = append(this.s1, val)
	}

	this.s2, this.s1 = this.s1, this.s2
}

func (this *MyStack) Pop() int {
	val := this.s2[0]
	this.s2 = this.s2[1:len(this.s2)]
	return val
}

func (this *MyStack) Top() int {
	return this.s2[0]
}

func (this *MyStack) Empty() bool {
	return len(this.s2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// leetcode submit region end(Prohibit modification and deletion)
