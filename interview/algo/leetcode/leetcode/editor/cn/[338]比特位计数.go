package main

import "math/bits"

// leetcode submit region begin(Prohibit modification and deletion)
func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		res[i] = bits.OnesCount(uint(i))
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
