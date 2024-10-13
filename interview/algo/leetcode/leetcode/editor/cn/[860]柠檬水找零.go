package main

// leetcode submit region begin(Prohibit modification and deletion)
func lemonadeChange(bills []int) bool {
	var ten, five int
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five < 0 {
				return false
			}
			five--
			ten++
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
