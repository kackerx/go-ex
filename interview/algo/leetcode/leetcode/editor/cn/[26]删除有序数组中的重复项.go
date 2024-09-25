package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// leetcode submit region end(Prohibit modification and deletion)
