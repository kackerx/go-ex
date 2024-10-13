package main

// leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
	// j x (i-j), j x dp[i-j]
	dp := make([]int, n+1) // dp[i] --> 对于数i, 所拆分的子数的最大之积

	dp[2] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j]) // 拆成两个数j和(i-j), 拆成固定的j和其余情况下最大值dp[i-j]
		}
	}

	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
