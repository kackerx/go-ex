package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	sort.Ints(nums)

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		left := nums[lo]
		right := nums[hi]
		sum := left + right

		if sum == target {
			return []int{}
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
