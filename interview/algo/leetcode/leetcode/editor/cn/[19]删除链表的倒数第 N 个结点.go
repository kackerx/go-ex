package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	slow := dummy
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
