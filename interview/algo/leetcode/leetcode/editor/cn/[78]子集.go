package main
//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	var (
		res   [][]int
		track []int
	)

	backtrack78(nums, 0, track, &res)

	return res
}

func backtrack78(nums []int, idx int, track []int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := idx; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack78(nums, i+1, track, res)
		track = track[:len(track)-1]
	}

}
//leetcode submit region end(Prohibit modification and deletion)
