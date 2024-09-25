package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	var (
		res   [][]int
		track []int
		sum   int
	)

	backkkk(candidates, track, &res, 0, target, sum)
	return res
}

func backkkk(candidates []int, track []int, res *[][]int, index int, target int, sum int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backkkk(candidates, track, res, i, target, sum)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
