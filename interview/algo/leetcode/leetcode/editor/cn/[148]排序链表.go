package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMin(head)
	left := head
	right := mid.Next
	mid.Next = nil

	leftList := sortList(left)
	rightList := sortList(right)

	return merge(leftList, rightList)
}

func getMin(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	p, p1, p2 := dummy, left, right

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
