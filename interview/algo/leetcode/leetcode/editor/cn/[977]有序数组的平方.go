package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r, k := 0, len(nums)-1, len(nums)-1
	for l <= r {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res[k] = nums[l] * nums[l]
			l++
		} else {
			res[k] = nums[r] * nums[r]
			r--
		}
		k--
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
