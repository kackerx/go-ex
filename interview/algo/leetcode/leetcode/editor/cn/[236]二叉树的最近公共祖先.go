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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var helper func(node *TreeNode, val1, val2 int) *TreeNode
	helper = func(node *TreeNode, val1, val2 int) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == val1 || node.Val == val2 {
			return node
		}

		left := helper(node.Left, val1, val2)
		right := helper(node.Right, val1, val2)

		if left != nil && right != nil {
			return node
		}

		if left != nil {
			return left
		}
		return right
	}

	return helper(root, p.Val, q.Val)
}

// leetcode submit region end(Prohibit modification and deletion)
