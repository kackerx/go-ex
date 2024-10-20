package main

import "golang.org/x/exp/slices"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var (
		res   [][]int
		track []int
	)

	traverse(root, targetSum, track, &res)
	return res
}

func traverse(node *TreeNode, sum int, track []int, res *[][]int) {
	if node == nil {
		return
	}

	track = append(track, node.Val)
	if node.Left == nil && node.Right == nil && sum == node.Val {
		*res = append(*res, slices.Clone(track))
		return
	}

	traverse(node.Left, sum-node.Val, track, res)
	traverse(node.Right, sum-node.Val, track, res)

	track = track[:len(track)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
