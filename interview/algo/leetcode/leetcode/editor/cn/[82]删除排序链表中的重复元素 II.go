package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	p, q := dummy, head
	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			for q.Next != nil && q.Val == q.Next.Val {
				q = q.Next
			}

			q = q.Next
			if q == nil {
				p.Next = nil
			}
		} else {
			p.Next = q
			p = p.Next
			q = q.Next
		}
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
