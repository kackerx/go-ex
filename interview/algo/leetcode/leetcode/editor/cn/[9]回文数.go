package main

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	tmp := x
	y := 0

	for tmp > 0 {
		last := tmp % 10
		tmp = tmp / 10
		y = y*10 + last
	}

	return x == y
}

// leetcode submit region end(Prohibit modification and deletion)
