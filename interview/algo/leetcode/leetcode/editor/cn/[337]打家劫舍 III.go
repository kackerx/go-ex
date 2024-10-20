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
func rob(root *TreeNode) int {
	dp := traverse(root)
	return max(dp[0], dp[1])
}

func traverse(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	dp := make([]int, 2)

	leftDP := traverse(node.Left)
	rightDP := traverse(node.Right)

	dp[0] = max(leftDP[0], leftDP[1]) + max(rightDP[0], rightDP[1])
	dp[1] = node.Val + leftDP[0] + rightDP[0]

	return dp
}

// leetcode submit region end(Prohibit modification and deletion)
