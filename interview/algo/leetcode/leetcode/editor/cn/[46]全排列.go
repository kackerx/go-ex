package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, track, &res)

	return res
}

func backtrack(nums []int, track []int, res *[][]int) {
	if len(track) == len(nums) {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := 0; i < len(nums); i++ {
		if slices.Contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])
		backtrack(nums, track, res)
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
