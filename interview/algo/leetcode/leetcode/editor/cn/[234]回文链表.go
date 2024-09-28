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

	newHead := rev1(slow)
	p := head

	for newHead != nil {
		if p.Val != newHead.Val {
			return false
		}

		p = p.Next
		newHead = newHead.Next
	}

	return true
}

func rev1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
// leetcode submit region end(Prohibit modification and deletion)
