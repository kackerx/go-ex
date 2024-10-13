package main
//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrack39(candidates, target, 0, 0, []int{}, &res)
	return res
}

func backtrack39(candidates []int, target, sum, idx int, track []int, res *[][]int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target {
		return
	}

	for i := idx; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack39(candidates, target, sum, i, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
