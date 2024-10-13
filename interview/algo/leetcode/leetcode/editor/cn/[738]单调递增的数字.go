package main

import "strconv"

// leetcode submit region begin(Prohibit modification and deletion)
func monotoneIncreasingDigits(n int) int {
	// 332, 从最高位也就是个位和前一位比较, 小于就置为9然后前一位--
	bs := []byte(strconv.Itoa(n))

	flag := len(bs)
	for i := len(bs) - 1; i > 0; i-- {
		if bs[i] < bs[i-1] {
			bs[i-1]--
			flag = i
		}
	}

	for i := flag; i < len(bs); i++ {
		bs[i] = '9'
	}

	res, _ := strconv.Atoi(string(bs))
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
