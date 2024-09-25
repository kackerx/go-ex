package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1 := dummy
	for i := 0; i < n+1; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	p2 := dummy
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
