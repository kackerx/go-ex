package main

// leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(num uint32) int {
	var sum int
	for num != 0 {
		num = num & (num - 1) // n&(n-1) 消除最右侧的1
		sum++
	}

	return sum
}

// leetcode submit region end(Prohibit modification and deletion)
