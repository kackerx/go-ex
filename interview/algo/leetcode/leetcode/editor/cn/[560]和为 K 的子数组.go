package main

// leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	var (
		res    int
		curSum int
		count  = map[int]int{0: 1}
	)

	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		if v, ok := count[curSum-k]; ok {
			res += v
		}

		count[curSum]++
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
