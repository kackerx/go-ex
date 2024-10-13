package main

// leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// (sum + target) must be even to partition
	if (sum+target)%2 == 1 || sum < abs(target) {
		return 0
	}
	newTarget := (sum + target) / 2

	return kn(nums, newTarget)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func kn(nums []int, sum int) int {
	n := len(nums)
	// dp[j] --> 装满重量j, 有dp[j]种方法
	dp := make([]int, sum+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[sum]
}



// leetcode submit region end(Prohibit modification and deletion)
