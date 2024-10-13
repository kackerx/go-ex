package main
//leetcode submit region begin(Prohibit modification and deletion)
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if len(nums) < 1 {
		return n
	}
	res := 1
	preDiff := 0
	for i := 0; i < n-1; i++ {
		curDiff := nums[i+1] - nums[i]
		if preDiff >= 0 && curDiff < 0 || preDiff <= 0 && curDiff > 0 {
			res++
			preDiff = curDiff
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
