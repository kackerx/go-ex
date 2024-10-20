package linkedlist

import (
	"fmt"
	"testing"
)

var head = &ListNode{
	Val:  0,
	Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
}

func TestReverseN(t *testing.T) {
	fmt.Println(head)

	newHead := reverseN(head, 3)
	fmt.Println(newHead)
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

func TestReverseBetween(t *testing.T) {
	fmt.Println(reverseBetween(head, 2, 3))
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right-left+1)
	}

	pre := head
	for i := 1; i < left-1; i++ {
		pre = pre.Next
	}

	pre.Next = reverseN(pre.Next, right-left+1)

	return pre
}

func TestReverseKGroup(t *testing.T) {
	fmt.Println(reverseKGroup(head, 4))
}

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
