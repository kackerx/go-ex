package main
//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	back3(n, k, 1, 0, []int{}, &res)
	return res
}

func back3(n, k, idx, sum int, track []int, res *[][]int) {
	if len(track) == k && sum == n {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := idx; i <= 9; i++ {
		track = append(track, i)
		sum += i
		back3(n, k, i+1, sum, track, res)
		sum -= i
		track = track[:len(track)-1]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
