package main

// leetcode submit region begin(Prohibit modification and deletion)
func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		res   [][]int
		track []int
	)

	traverse(graph, 0, track, &res)

	return res
}

func traverse(graph [][]int, s int, track []int, res *[][]int) {
	track = append(track, s)

	n := len(graph)
	if s == n-1 {
		*res = append(*res, append([]int{}, track...))
	}

	for _, v := range graph[s] {
		traverse(graph, v, track, res)
	}

	track = track[:len(track)-1]
}

// leetcode submit region end(Prohibit modification and deletion)
