package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 */

var maxDiameter = 0

func diameterOfBinaryTree(root *TreeNode) int {
	// 记录最大直径的长度
	maxDepth(root)
	return maxDiameter
}

func maxDepth(no *TreeNode) int {
	if no == nil {
		return 0
	}
	leftdepth := maxDepth(no.Left)
	rightDepth := maxDepth(no.Right)

	maxDiameter = max(maxDiameter, leftdepth+rightDepth)

	return max(leftdepth, rightDepth) + 1
}

// leetcode submit region end(Prohibit modification and deletion)
