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

func find(no *TreeNode, val1, val2 int) *TreeNode {
	if no == nil {
		return nil
	}

	// 前序 (先遇到的一定是公共祖先
	if no.Val == val1 || no.Val == val2 {
		return no
	}

	left := find(no.Left, val1, val2)
	right := find(no.Right, val1, val2)

	// 后续位置 (左子树和右子树都不为nil那就是当前节点
	if left != nil && right != nil {
		return no
	}

	if left != nil {
		return left
	} else {
		return right
	}
}

// leetcode submit region end(Prohibit modification and deletion)
