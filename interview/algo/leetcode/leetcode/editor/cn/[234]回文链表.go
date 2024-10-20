package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	reveNode := reve(slow)

	p := reveNode
	q := head
	for p != nil {
		if p.Val != q.Val {
			return false
		}

		p = p.Next
		q = q.Next
	}

	return true
}

func reve(node *ListNode) *ListNode {
	cur := node
	var pre, next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// leetcode submit region end(Prohibit modification and deletion)
