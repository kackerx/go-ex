package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// type Node struct {
// 	Val    int
// 	Next   *Node
// 	Random *Node
// }

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	cur := head
	for cur != nil {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	res := head.Next
	cur = head
	for cur != nil {
		newCur := cur.Next
		cur.Next = newCur.Next
		if newCur.Next != nil {
			newCur.Next = newCur.Next.Next
		}
		cur = cur.Next
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
