package main

import "math/bits"

// leetcode submit region begin(Prohibit modification and deletion)
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

// leetcode submit region end(Prohibit modification and deletion)
