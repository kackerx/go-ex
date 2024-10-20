package main

// leetcode submit region begin(Prohibit modification and deletion)
func possibleBipartition(n int, dislikes [][]int) bool {
	return isBipartite(buildGraph(dislikes, n))
}

func buildGraph(dislinks [][]int, n int) [][]int {
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, edge := range dislinks {
		v := edge[0]
		w := edge[1]

		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}

	return graph
}

func isBipartite(graph [][]int) bool {
	var (
		n       = len(graph)
		color   = make([]bool, n)
		visited = make([]bool, n)
		ok      bool // 默认false = 是二分图
	)

	var traverse func(graph [][]int, v int)
	traverse = func(graph [][]int, v int) {
		if ok {
			return
		}

		if visited[v] {
			return
		}

		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] { // 没访问过的节点, 染色, dfs
				color[w] = !color[v]
				traverse(graph, w) // dfs
			} else {
				if color[w] == color[v] {
					ok = true // 不是二分图了
					return
				}
			}
		}
	}

	for i := 0; i < n; i++ { // 图节点可能不是全连通, 所以对每个节点都当做起点全部遍历
		if !visited[i] {
			traverse(graph, i)
		}
	}

	return !ok
}

// leetcode submit region end(Prohibit modification and deletion)
