package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	var (
		res   [][]int
		track []int
	)
	slices.Sort(candidates)
	backk(candidates, 0, track, &res, target, 0)
	return res
}

func backk(candidates []int, index int, track []int, res *[][]int, target, sum int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target { // 已经大于目标值直接return
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		sum += candidates[i]
		backk(candidates, i+1, track, res, target, sum)
		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
