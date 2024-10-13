package main

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1) // m X n的有dp[m][n]条路径到达
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[1][i] = 1
	}

	for i := 0; i <= m; i++ {
		dp[i][1] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}


// leetcode submit region end(Prohibit modification and deletion)
