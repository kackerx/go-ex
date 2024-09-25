package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return ReveN(head, right)
	}

	pre := head
	for i := 0; i < left-1-1; i++ {
		pre = pre.Next
	}

	pre.Next = ReveN(pre.Next, right-left+1)
	return head
}

func ReveN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, cur, next *ListNode
	cur = head
	for i := 0; i < n; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	head.Next = cur
	return pre
}

// leetcode submit region end(Prohibit modification and deletion)
