package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var res, rank int

	var traverse func(no *TreeNode)
	traverse = func(no *TreeNode) {
		if no == nil {
			return
		}

		traverse(no.Left)
		rank++
		if rank == k {
			res = no.Val
			return
		}
		traverse(no.Right)
	}

	traverse(root)
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
