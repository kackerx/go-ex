package main

// leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	l := 0
	r := x + 1
	for l < r { // r = x+1, 因此范围是[l, r)
		mid := l + (r-l+1)/2 // l = mid时会出现的坑, 范围不再缩小, 因此
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}

// leetcode submit region end(Prohibit modification and deletion)
