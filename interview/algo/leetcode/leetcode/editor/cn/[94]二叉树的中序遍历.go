package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)

	traverse(root, &res)

	return res
}

func traverse(no *TreeNode, res *[]int) {
	if no == nil {
		return
	}

	traverse(no.Left, res)
	*res = append(*res, no.Val)
	traverse(no.Right, res)
}
//leetcode submit region end(Prohibit modification and deletion)
