package main

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i] // n ^ n = 0, n ^ 0 = n
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
