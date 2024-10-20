package main
//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	var (
		track string
		res   []string
	)

	backtrack(track, n, n, &res)

	return res
}

func backtrack(track string, left, right int, res *[]string) {
	if right < left {
		return
	}

	if left < 0 || right < 0 {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, track)
		return
	}

	track += "("
	backtrack(track, left-1, right, res)
	track = track[:len(track)-1]

	track += ")"
	backtrack(track, left, right-1, res)
	track = track[:len(track)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
