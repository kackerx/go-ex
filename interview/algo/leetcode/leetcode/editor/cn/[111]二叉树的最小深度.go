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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil { // 避免节点的一边为空时返回了0
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	return min(left, right) + 1
}

// leetcode submit region end(Prohibit modification and deletion)
