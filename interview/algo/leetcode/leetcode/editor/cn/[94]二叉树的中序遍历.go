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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(no *TreeNode, res *[]int) {
	if no == nil {
		return
	}

	inorder(no.Left, res)
	*res = append(*res, no.Val)
	inorder(no.Right, res)

}

// leetcode submit region end(Prohibit modification and deletion)
