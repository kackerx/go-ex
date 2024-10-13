package dp

import (
	"fmt"
	"testing"
)

func Test746(t *testing.T) {
	s := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(s))
}

func minCostClimbingStairs(cost []int) int {
	// minCost := math.MaxInt

	dp := make([]int, len(cost)+1) // 楼顶下标是len(cost), 跳到第i个台阶最低花费为dp[i]
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]) // 要么跳2步要么跳1步, 那么从这两种情况来的花费的最小值就是本次dp[i]的最小花费
	}

	return dp[len(cost)]
}
