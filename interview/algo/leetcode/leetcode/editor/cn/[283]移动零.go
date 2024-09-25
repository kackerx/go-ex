package main

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

// leetcode submit region end(Prohibit modification and deletion)

type User struct {
	Name string
	Age  int
}
