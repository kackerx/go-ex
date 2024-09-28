package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	stk := &stack{}
	p := head
	for p != nil {
		stk.Push(p)
		p = p.Next
	}

	p = head
	for p != nil {
		top := stk.Pop()
		next := p.Next
		if next == top || next == top.Next {
			top.Next = nil
			break
		}
		p.Next = top
		top.Next = next
		p = next
	}
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

// leetcode submit region end(Prohibit modification and deletion)
