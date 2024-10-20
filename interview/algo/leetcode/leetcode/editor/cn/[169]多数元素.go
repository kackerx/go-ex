package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	var res int
	maxVal := math.MinInt
	for k, v := range m {
		if v > maxVal {
			maxVal = v
			res = k
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
