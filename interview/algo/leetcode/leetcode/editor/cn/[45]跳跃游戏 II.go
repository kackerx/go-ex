package main
//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var end, jumps, farthest int

	for i := 0; i < len(nums); i++ {
		farthest = max(farthest, i+nums[i])

		if end == i {
			jumps++
			end = farthest

			if end >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}

//leetcode submit region end(Prohibit modification and deletion)
