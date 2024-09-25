package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(matrix [][]int) int {
	// dp := make([][]int, len(matrix)) // 从第一行开始落到dp[i][j]的最小路径和位dp(matrix, i, j)
	memo := make([][]int, len(matrix))
	for i := range memo {
		memo[i] = make([]int, len(matrix))
		for j := range memo[i] {
			memo[i][j] = 666666
		}
	}

	var dp func(matrix [][]int, i, j int) int
	dp = func(matrix [][]int, i, j int) int {
		// 1、索引合法性检查
		if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
			return 99999
		}

		if i == 0 {
			return matrix[0][j]
		}

		if memo[i][j] != 666666 {
			return memo[i][j]
		}

		memo[i][j] = matrix[i][j] + min(
			dp(matrix, i-1, j),
			dp(matrix, i-1, j-1),
			dp(matrix, i-1, j+1),
		)

		return memo[i][j]
	}

	res := math.MaxInt
	// 终点可能在 matrix[n-1] 的任意一列
	for j := 0; j < len(matrix); j++ {
		colSum := dp(matrix, len(matrix)-1, j)
		if colSum < res {
			res = colSum
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
