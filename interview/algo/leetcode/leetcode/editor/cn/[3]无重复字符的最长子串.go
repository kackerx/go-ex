package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if s == "" {return 0}
	maxRes := math.MinInt
	win := make(map[byte]int)

	var l, r int
	for r < len(s) {
		c := s[r]
		r++
		win[c]++

		for win[c] > 1 {
			d := s[l]
			l++
			win[d]--
		}

		maxRes = max(maxRes, r-l)
	}

	return maxRes
}

// leetcode submit region end(Prohibit modification and deletion)
