package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var (
		res   [][]int
		track []int
		used  = make([]bool, len(nums))
	)
	slices.Sort(nums)
	backkk(nums, track, used, &res)
	return res
}

func backkk(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
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
		backkk(nums, track, used, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
