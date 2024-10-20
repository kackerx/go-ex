package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	lt, gt, i := -1, len(nums), 0

	for i < gt {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else {
			gt--
			nums[gt], nums[i] = nums[i], nums[gt]
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
