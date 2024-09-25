package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)
	backtrackk(nums, 0, []int{}, &res)

	return res
}

func backtrackk(nums []int, index int, track []int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := index; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrackk(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
