package main

import (
	"golang.org/x/exp/slices"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0] // 以i结尾的最大和子数组

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	// 因为dp数组的定义不是当前数组最大子数组和所以不能直接返回
	return slices.Max(dp)
}

// leetcode submit region end(Prohibit modification and deletion)
