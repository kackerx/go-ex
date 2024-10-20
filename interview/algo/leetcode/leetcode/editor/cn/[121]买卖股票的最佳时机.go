package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			// 根据状态转移方程可得：
			//   dp[i][0]
			// = max(dp[-1][0], dp[-1][1] + prices[i])
			// = max(0, -infinity + prices[i]) = 0
			// = max(dp[-1][0], dp[-1][1] + prices[i])
			// = max(0, -infinity + prices[i]) = 0

			dp[i][1] = -prices[i]
			// 根据状态转移方程可得：
			//   dp[i][1]
			// = max(dp[-1][1], dp[-1][0] - prices[i])
			// = max(-infinity, 0 - prices[i])
			// = -prices[i]
			// = max(dp[-1][1], dp[-1][0] - prices[i])
			// = max(-infinity, 0 - prices[i])
			// = -prices[i]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

// leetcode submit region end(Prohibit modification and deletion)
