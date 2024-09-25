package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	var left, right, sum int
	minLen := math.MaxInt

	for right < len(nums) {
		val := nums[right]
		right++
		sum += val

		for sum >= target {
			minLen = min(minLen, right-left)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	} else {
		return minLen
	}
}
//leetcode submit region end(Prohibit modification and deletion)
