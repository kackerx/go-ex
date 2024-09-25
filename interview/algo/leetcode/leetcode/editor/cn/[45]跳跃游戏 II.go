package main

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	n := len(nums)
	end, farthest, jumps := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, nums[i]+i)
		if end == i { // 条件 end == i 的作用是判断当前跳跃范围是否结束。当 i 达到 end 时，表示我们已经在当前跳跃范围内走到了最远的位置，接下来需要进行一次新的跳跃。
			jumps++ // 只有end这个最远距离需要进一步跳跃时增加跳跃次数
			end = farthest
		}
	}
	return jumps
}

// leetcode submit region end(Prohibit modification and deletion)
