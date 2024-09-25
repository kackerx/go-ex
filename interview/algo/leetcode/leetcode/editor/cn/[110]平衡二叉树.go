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
	var flag bool
	maxDepth(root, &flag)
	return !flag
}

func maxDepth(no *TreeNode, flag *bool) int {
	if no == nil {
		return 0
	}

	if *flag {
		return -1
	}

	left := maxDepth(no.Left, flag)
	right := maxDepth(no.Right, flag)

	if math.Abs(float64(left-right)) > 1 {
		*flag = true
	}

	return max(left, right) + 1
}

// leetcode submit region end(Prohibit modification and deletion)
