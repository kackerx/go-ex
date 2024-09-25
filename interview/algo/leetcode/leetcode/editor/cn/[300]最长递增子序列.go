package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)) // 含义: 以nums[i]该元素结尾的最长子序列长度
	for i := range dp {
		dp[i] = 1 // basecase: 当i = 0时包含自己, 所以初始为1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1) // i = 5; 比nums[5]小的同时i<5在5之前的dp数组选择最大的
			}
		}
	}

	return slices.Max(dp)
	// 数学归纳法: dp[0..i-1]都已知, 想办法求出dp[i]就结束嘞
}

// leetcode submit region end(Prohibit modification and deletion)
