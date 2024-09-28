package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := &minHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &ListNode{}
	p := dummy
	for h.Len() != 0 {
		no := heap.Pop(h).(*ListNode)
		p.Next = &ListNode{Val: no.Val}
		p = p.Next
		if no.Next != nil {
			heap.Push(h, no.Next)
		}
	}

	return dummy.Next
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

// leetcode submit region end(Prohibit modification and deletion)
