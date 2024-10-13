package main

// leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1) // dp[i] --> 第i阶, dp[i]种方法

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
