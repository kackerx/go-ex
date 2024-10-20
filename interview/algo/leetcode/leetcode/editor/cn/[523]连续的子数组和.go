package main

// leetcode submit region begin(Prohibit modification and deletion)
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0

	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	valToIndex := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		val := preSum[i] % k
		if _, ok := valToIndex[val]; !ok {
			valToIndex[val] = i
		}
	}

	for i := 1; i < len(preSum); i++ {
		val := preSum[i] % k
		if idx, ok := valToIndex[val]; ok && i-idx >= 2 {
			return true
		}
	}

	return false
}


// leetcode submit region end(Prohibit modification and deletion)
