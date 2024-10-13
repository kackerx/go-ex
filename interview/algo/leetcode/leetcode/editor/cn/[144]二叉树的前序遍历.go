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
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var helper func(no *TreeNode)
	helper = func(no *TreeNode) {
		if no == nil {
			return
		}

		res = append(res, no.Val)
		helper(no.Left)
		helper(no.Right)
	}

	helper(root)

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
