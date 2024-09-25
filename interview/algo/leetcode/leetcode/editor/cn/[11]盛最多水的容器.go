package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxVal := 0

	for l < r {
		maxVal = max(maxVal, (r-l)*min(height[l], height[r]))

		// 移动高度较低的一边, 因为移动高的一边的话矩形的高度不会变还是较低那边, 但是宽度窄了
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxVal
}

// leetcode submit region end(Prohibit modification and deletion)
