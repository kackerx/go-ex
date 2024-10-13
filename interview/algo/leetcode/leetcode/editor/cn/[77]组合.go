package main
//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	var res [][]int
	backtrack77(n, k, 1, []int{}, &res)
	return res
}

func backtrack77(n, k, idx int, track []int, res *[][]int) {
	if len(track) == k {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := idx; i < n+1; i++ {
		track = append(track, i)
		backtrack77(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
