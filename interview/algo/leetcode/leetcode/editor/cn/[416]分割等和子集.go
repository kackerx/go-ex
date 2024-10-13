package main

// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	sum = sum / 2

	// dp[i][j] --> 取前i个数, 可以凑满重量为j的背包
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 { // 第i个数装不下了
				dp[i][j] = dp[i-1][j]
			} else {
				// 要么不包含当前数, 要么包含当前数(去掉当前数的重量)
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][sum]
}

// leetcode submit region end(Prohibit modification and deletion)
