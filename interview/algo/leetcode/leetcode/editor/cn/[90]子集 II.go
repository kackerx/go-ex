package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	var (
		res   [][]int
		track []int
	)

	sort.Ints(nums)
	backtrack90(nums, 0, track, &res)
	return res
}

func backtrack90(nums []int, idx int, track []int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtrack90(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
