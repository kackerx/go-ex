package main

func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
