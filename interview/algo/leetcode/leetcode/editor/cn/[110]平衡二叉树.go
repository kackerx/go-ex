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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// leetcode submit region end(Prohibit modification and deletion)
