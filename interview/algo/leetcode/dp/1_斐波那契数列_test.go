package dp

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(fib(8))
}

// 1. dp数组及下标含义
// 2. 递推公式
// 3. dp数组basecase
// 4. 遍历顺序
func fib(n int) int {
	dp := make([]int, n+1) // 第n项的fib值是dp[n]

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
