package main
//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	var slow, fast, count int

	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			slow++
			nums[slow] = nums[fast]
		}

		count++
		fast++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			count = 0
		}
	}

	return slow + 1
}
//leetcode submit region end(Prohibit modification and deletion)
