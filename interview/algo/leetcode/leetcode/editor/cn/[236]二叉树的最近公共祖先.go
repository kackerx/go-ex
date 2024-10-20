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
	return find(root, p.Val, q.Val)
}

func find(node *TreeNode, val1, val2 int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == val1 || node.Val == val2 {
		return node
	}

	left := find(node.Left, val1, val2)
	right := find(node.Right, val1, val2)

	if left != nil && right != nil {
		return node
	}

	if left != nil {
		return left
	} else {
		return right
	}
}

// leetcode submit region end(Prohibit modification and deletion)
