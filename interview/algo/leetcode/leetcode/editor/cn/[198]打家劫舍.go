package main

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n) // 考虑偷当前下标i所偷到最大的钱

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
