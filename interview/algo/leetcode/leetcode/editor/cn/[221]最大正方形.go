package main

// leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	height := len(matrix)
	width := len(matrix[0])

	dp := make([][]int, height+1) // dp[height+1][width+1] --> 第i行第j列为右下角正方形最大边长
	for i := range dp {
		dp[i] = make([]int, width+1)
	}

	side := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i][j], dp[i+1][j]) + 1
				side = max(side, dp[i+1][j+1])
			}
		}
	}

	return side * side
}

// leetcode submit region end(Prohibit modification and deletion)
