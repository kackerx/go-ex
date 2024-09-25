package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	cur.Next = head
	for i := 0; i < n-k%n; i++ {
		cur = cur.Next
	}

	newHead := cur.Next
	cur.Next = nil

	return newHead
}

// leetcode submit region end(Prohibit modification and deletion)
