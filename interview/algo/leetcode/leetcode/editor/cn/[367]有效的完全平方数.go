package main

// leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	lo := 2
	hi := num/2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			lo = mid + 1
		} else {
			hi = mid-1
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)
