package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	maxVal := math.MinInt

	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxVal = max(maxVal, area)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVal
}

// leetcode submit region end(Prohibit modification and deletion)
