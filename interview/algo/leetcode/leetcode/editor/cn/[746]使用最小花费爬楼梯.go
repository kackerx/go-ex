package main

import (
	"fmt"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	// minCost := math.MaxInt

	dp := make([]int, len(cost)+1) // 跳到第i个台阶最低花费为dp[i]
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	fmt.Println(dp)

	return dp[len(cost)]
}

// leetcode submit region end(Prohibit modification and deletion)
