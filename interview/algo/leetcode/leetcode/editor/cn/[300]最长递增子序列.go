package main

import "golang.org/x/exp/slices"

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n) // dp[i] --> 到i为止递增序列长度为dp[i]

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}

// leetcode submit region end(Prohibit modification and deletion)
