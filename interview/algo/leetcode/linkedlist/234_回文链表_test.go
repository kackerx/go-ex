package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	sb := strings.Builder{}
	cur := l
	for cur != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", cur.Val))
		cur = cur.Next
		if cur == nil {
			sb.WriteString("N")
		}
	}

	return sb.String()
}

func Test234(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(isPalindrome(h))
}

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
