package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
// coinChange 使用硬币coins, 凑出amount所需要最少得硬币个数
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
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
		sub := coinChangeMemo(coins, amount-coin, memo)
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

// leetcode submit region end(Prohibit modification and deletion)
