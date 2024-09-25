package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, min(p.Val, q.Val), max(p.Val, q.Val))
}

func find(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val2 {
		return find(root.Left, val1, val2)
	}

	if root.Val < val1 {
		return find(root.Right, val1, val2)
	}

	return root
}

// leetcode submit region end(Prohibit modification and deletion)
