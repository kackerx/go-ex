package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt

	var oneSideMax func(node *TreeNode) int
	oneSideMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max(0, oneSideMax(node.Left))
		rightMax := max(0, oneSideMax(node.Right))

		res = max(res, leftMax+rightMax+node.Val)

		return max(leftMax, rightMax) + node.Val
	}

	oneSideMax(root)
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
