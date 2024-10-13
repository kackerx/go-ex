package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	win, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	var (
		l, r, valid, start int
		length             = math.MaxInt
	)

	for r < len(s) {
		c := s[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for valid == len(need) {
			if r-l < length {
				start = l
				length = r - l
			}

			d := s[l]
			l++
			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}

				win[d]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	} else {
		return s[start : start+length]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
