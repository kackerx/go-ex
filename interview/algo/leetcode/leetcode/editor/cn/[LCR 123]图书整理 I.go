package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBookList(head *ListNode) []int {
	res := make([]int, 0)
	if head == nil {
		return res
	}

	cur := head.Val
	list := reverseBookList(head.Next)
	list = append(list, cur)

	return list
}

// leetcode submit region end(Prohibit modification and deletion)
