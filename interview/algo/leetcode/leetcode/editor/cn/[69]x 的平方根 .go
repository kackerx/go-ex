package main

// leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	lo := 0
	hi := x

	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if mid*mid < x {
			lo = mid
		} else if mid*mid == x {
			return mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}

// leetcode submit region end(Prohibit modification and deletion)
