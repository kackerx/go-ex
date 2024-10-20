package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	a, b := head, head

	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverseN(a, k)
	a.Next = reverseKGroup(b, k)

	return newHead
}

func reverseN(node *ListNode, n int) *ListNode {
	var pre, next *ListNode
	cur := node
	for i := 0; i < n; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	node.Next = cur

	return pre
}


// leetcode submit region end(Prohibit modification and deletion)
