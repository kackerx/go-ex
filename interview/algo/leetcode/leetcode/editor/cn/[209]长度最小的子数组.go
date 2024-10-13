package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	var (
		l, r, sum int
		minLen    = math.MaxInt
	)

	for r < len(nums) {
		v := nums[r]
		r++
		sum += v

		for sum >= target {
			minLen = min(minLen, r-l)

			v = nums[l]
			l++
			sum -= v
		}

	}

	if minLen == math.MaxInt {
		return 0
	} else {
		return minLen
	}
}

// leetcode submit region end(Prohibit modification and deletion)
