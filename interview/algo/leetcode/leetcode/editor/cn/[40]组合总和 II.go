package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	backtrack40(candidates, target, 0, 0, []int{}, &res)

	return res
}

func backtrack40(candidates []int, target, sum, index int, track []int, res *[][]int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack40(candidates, target, sum, i+1, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
