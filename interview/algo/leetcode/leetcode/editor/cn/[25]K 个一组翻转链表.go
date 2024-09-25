package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}

		b = b.Next
	}

	newHead := ReveN(a, k) // 先翻转第一个, 递归翻转剩下的, a.Next
	a.Next = reverseKGroup(b, k)
	return newHead
}

// leetcode submit region end(Prohibit modification and deletion)
