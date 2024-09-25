package main

import "math/rand"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	cur := this.head
	res := cur.Val
	i := 1
	for cur != nil {
		if rand.Intn(i) == 0 {
			res = cur.Val
		}
		cur = cur.Next
		i++
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// leetcode submit region end(Prohibit modification and deletion)
