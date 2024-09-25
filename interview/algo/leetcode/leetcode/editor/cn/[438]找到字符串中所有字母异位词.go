package main

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	win, need := make(map[byte]int, 0), make(map[byte]int, 0)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	var res []int
	var l, r, valid int

	for r < len(s) {
		c := s[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for r-l >= len(p) {
			if valid == len(need) {
				res = append(res, l)
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

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
