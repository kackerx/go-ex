package main

import (
	"golang.org/x/exp/slices"
)

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var totalSum int
	var nodeList [][]int

	var helper func(node *TreeNode, track []int)

	helper = func(node *TreeNode, track []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			track = append(track, node.Val)
			nodeList = append(nodeList, slices.Clone(track))
			return
		}

		track = append(track, node.Val)
		helper(node.Left, track)
		helper(node.Right, track)
		track = track[:len(track)-1]
	}

	helper(root, []int{})

	for _, ints := range nodeList {
		var sum int
		var carry = 1
		for i := len(ints) - 1; i >= 0; i-- {
			sum += carry * ints[i]
			carry *= 10
		}
		totalSum += sum
	}

	return totalSum
}

// leetcode submit region end(Prohibit modification and deletion)
