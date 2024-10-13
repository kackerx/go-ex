package dp

func climbStairs(n int) int {
	/*
	*** 1阶 --> 1
	*** 2阶 --> 2
	*** 3阶 --> 3 = 1+2
	*** 4阶 --> 5 = 2+3
	 */
	dp := make([]int, n+1) // dp[i] --> 第i阶, dp[i]种方法

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
