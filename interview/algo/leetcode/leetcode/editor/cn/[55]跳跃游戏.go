package main

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i { // [3, 2, 1, 0, 4] farthest = 3时, i=3, 只能跳0步永远到不了
			return false
		}
	}

	return farthest >= len(nums)-1
}

// leetcode submit region end(Prohibit modification and deletion)
