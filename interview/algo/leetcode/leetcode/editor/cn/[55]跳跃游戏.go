package main

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	var cover int
	for i := 0; i <= cover; i++ { // i最大走到覆盖范围内
		cover = max(cover, i+nums[i]) // 不断更新覆盖范围为当前下标能跳最大值
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)
