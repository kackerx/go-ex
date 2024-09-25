package main

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	win, need := make(map[byte]int, 0), make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	var l, r, valid int

	for r < len(s2) {
		c := s2[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		if r-l >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[l]
			l++

			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}
				win[d]--
			}
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)
