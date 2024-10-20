package main

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	valToIndex := make(map[int]int)
	res := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			preSum[i+1] = preSum[i] - 1
		} else {
			preSum[i+1] = preSum[i] + 1
		}
	}

	for i := 0; i < len(preSum); i++ {
		v, ok := valToIndex[preSum[i]]
		if ok {
			res = max(res, i-v)
		} else {
			valToIndex[preSum[i]] = i
		}
	}

	return res
}
// leetcode submit region end(Prohibit modification and deletion)
