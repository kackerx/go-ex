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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 利用定义，计算左右子树的最大深度
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	// 整棵树的最大深度等于左右子树的最大深度取最大值，
	// 然后再加上根节点自己
	res := max(leftMax, rightMax) + 1

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)
