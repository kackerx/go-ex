package main

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var l, r, res int
	win := make(map[byte]int, 0)

	for r < len(s) {
		c := s[r]
		r++

		win[c]++

		for win[c] > 1 {
			d := s[l]
			win[d]--
			l++
		}

		res = max(res, r-l)

	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
