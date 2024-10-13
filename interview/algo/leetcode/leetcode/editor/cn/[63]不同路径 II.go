package main

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ { // 遍历顺序左到右上到下, 可以由图得知
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 { // 目标点是障碍物则dp[i][j] = 0
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[m-1][n-1]
}

// leetcode submit region end(Prohibit modification and deletion)
