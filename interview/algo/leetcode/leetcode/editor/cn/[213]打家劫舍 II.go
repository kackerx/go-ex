package main

// leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(subRob(nums[1:]), subRob(nums[:len(nums)-1])) // 既然首尾相连, 那就不要首或者不要尾然后取最大值
}

func subRob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
