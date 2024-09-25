package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
func partition(head *ListNode, x int) *ListNode {
	l1 := &ListNode{}
	l2 := &ListNode{}

	p1 := l1
	p2 := l2

	cur := head
	for cur != nil {
		if cur.Val < x {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}

		next := cur.Next
		cur.Next = nil
		cur = next
	}

	// 合并两个有序链表
	p1.Next = l2.Next
	return l1.Next
}

// leetcode submit region end(Prohibit modification and deletion)
