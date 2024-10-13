package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack47(nums, used, []int{}, &res)

	return res
}

func backtrack47(nums []int, used []bool, track []int, res *[][]int) {
	if len(nums) == len(track) {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack47(nums, used, track, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
