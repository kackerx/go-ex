package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	k := len(nums) - 1
	for k-1 >= 0 && nums[k-1] >= nums[k] {
		k--
	}

	if k <= 0 {
		slices.Reverse(nums)
		return
	}

	t := len(nums) - 1
	for nums[t] <= nums[k-1] {
		t--
	}

	nums[k-1], nums[t] = nums[t], nums[k-1]
	sort.Ints(nums[k:])
}

// leetcode submit region end(Prohibit modification and deletion)
