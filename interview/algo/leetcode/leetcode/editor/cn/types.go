package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i].Val < (*m)[j].Val
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() any {
	ret := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return ret
}

type stack []*ListNode

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Push(e *ListNode) {
	*s = append(*s, e)
}

func (s *stack) Pop() *ListNode {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
