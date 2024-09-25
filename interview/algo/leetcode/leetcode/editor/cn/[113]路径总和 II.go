package main

import "slices"

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

	res := make([][]int, 0)
	traverse(root, targetSum, []int{}, &res)

	return res
}

func traverse(no *TreeNode, target int, track []int, res *[][]int) {
	if no == nil {
		return
	}

	track = append(track, no.Val)

	if no.Left == no.Right && target == no.Val {
		*res = append(*res, slices.Clone(track))
		return
	}

	traverse(no.Left, target-no.Val, track, res)
	traverse(no.Right, target-no.Val, track, res)
	track = track[:len(track)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
