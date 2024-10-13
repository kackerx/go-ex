package main

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backtrack216(k, n, 0, 1, []int{}, &res)

	return res
}

func backtrack216(k, n, sum, index int, track []int, res *[][]int) {
	if len(track) == k && sum == n {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum >= n {
		return
	}

	for i := index; i < 10; i++ {
		track = append(track, i)
		sum += i
		backtrack216(k, n, sum, i+1, track, res)
		sum -= i
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
