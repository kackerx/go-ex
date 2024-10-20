package main

// leetcode submit region begin(Prohibit modification and deletion)
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		visited   = make([]bool, numCourses)
		onPath    = make([]bool, numCourses)
		hasCycle  bool
		postOrder = make([]int, 0)
	)

	var traverse func(graph [][]int, s int)
	traverse = func(graph [][]int, s int) {
		if onPath[s] {
			hasCycle = true
		}

		if visited[s] || hasCycle {
			return
		}
		visited[s] = true
		onPath[s] = true

		for _, v := range graph[s] {
			traverse(graph, v)
		}

		onPath[s] = false
		postOrder = append(postOrder, s)
	}

	graph := buildGraph(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			traverse(graph, i)
		}
	}

	if hasCycle {
		return nil
	}

	return postOrder
}

func buildGraph(n int, prerequisites [][]int) (graph [][]int) {
	graph = make([][]int, n)
	// fi
	for _, prerequisite := range prerequisites {
		from := prerequisite[0]
		to := prerequisite[1]

		graph[from] = append(graph[from], to)
	}

	return
}

// leetcode submit region end(Prohibit modification and deletion)
