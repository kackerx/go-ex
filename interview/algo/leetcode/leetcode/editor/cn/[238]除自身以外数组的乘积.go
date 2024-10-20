package main

// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// 构建前缀积
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	// 构建后缀积, 从右-->左的前缀积
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	res[0] = suffix[1]
	res[n-1] = prefix[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = prefix[i-1] * suffix[i+1]
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
