package dp

import (
	"fmt"
	"math"
	"testing"
)

func TestFib(t *testing.T) {
	fmt.Println(coinchange([]int{3, 4, 5}, 17853))
}

func coinchange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[amout], 需要下标amount对应的硬币数, 所以必须长度是amount+1
	for i := range dp {
		dp[i] = amount + 1 // 表示无法到达的状态
	}

	dp[0] = 0 // base case
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 { // 当前金额减去当前选择硬币后大于零的话
				dp[i] = min(dp[i], dp[i-coin]+1) // 金额i和金额i去掉这个硬币所需的硬币个数+1, 穷解所有的coin后取最小的那个数
			}

		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res := math.MaxInt
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		if sub == -1 {
			continue
		}

		res = min(res, sub+1)
	}

	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}

func Sol(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -666
	}

	return coinChangeMemo(coins, amount, &memo)
}

func coinChangeMemo(coins []int, amount int, memo *[]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if (*memo)[amount] != -666 {
		return (*memo)[amount]
	}

	res := math.MaxInt
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		if sub == -1 {
			continue
		}

		res = min(res, sub+1)
	}

	if res == math.MaxInt {
		(*memo)[amount] = -1
	} else {
		(*memo)[amount] = res
	}

	return (*memo)[amount]
}
