package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right-left+1)
	}

	pre := head
	for i := 1; i < left-1; i++ {
		pre = pre.Next
	}

	pre.Next = reverseN(pre.Next, right-left+1)

	return head
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
//leetcode submit region end(Prohibit modification and deletion)
