package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}

		fast++
	}

	return slow
}

// leetcode submit region end(Prohibit modification and deletion)
