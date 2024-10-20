package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}

	fmt.Println(dp)
	return slices.Max(dp)
}

// leetcode submit region end(Prohibit modification and deletion)
