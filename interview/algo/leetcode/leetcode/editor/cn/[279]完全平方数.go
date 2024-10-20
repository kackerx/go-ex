package main

// leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	dp := make([]int, n+1) // 和为i的平方数的最小数量是dp[i]

	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
