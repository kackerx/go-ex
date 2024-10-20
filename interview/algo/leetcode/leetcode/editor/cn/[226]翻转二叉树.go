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
func invertTree(root *TreeNode) *TreeNode {
	return reverse(root)
}

func reverse(no *TreeNode) *TreeNode {
	if no == nil {
		return no
	}

	// no.Left = reverse(no.Right)
	// no.Right = reverse(no.Left)
	no.Left, no.Right = reverse(no.Right), reverse(no.Left)

	return no
}

// leetcode submit region end(Prohibit modification and deletion)
