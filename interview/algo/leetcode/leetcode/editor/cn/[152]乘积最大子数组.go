package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	maxVal := nums[0]
	iMax := nums[0]
	iMin := nums[0]

	for i := 1; i < len(nums); i++ {
		tmpMin, tmpMax := iMin, iMax
		iMin = min(nums[i], tmpMin*nums[i], tmpMax*nums[i])
		iMax = max(nums[i], tmpMin*nums[i], tmpMax*nums[i])

		maxVal = max(maxVal, iMax)
	}

	return maxVal
}

// leetcode submit region end(Prohibit modification and deletion)
