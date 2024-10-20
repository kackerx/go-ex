package main

// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		visited  = make([]bool, numCourses)
		onPath   = make([]bool, numCourses)
		hasCycle bool
	)

	var traverse func(graph [][]int, s int)
	traverse = func(graph [][]int, s int) {
		if onPath[s] {
			hasCycle = true
		}

		if visited[s] || hasCycle {
			return
		}

		visited[s] = true // 遍历过的置灰
		onPath[s] = true  // 当前节点所有目标节点遍历完取消选择, 调用堆栈
		for _, v := range graph[s] {
			traverse(graph, v)
		}

		onPath[s] = false
	}

	graph := buildGraph(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			traverse(graph, i)
		}
	}

	return !hasCycle
}

func buildGraph(n int, prerequisites [][]int) (graph [][]int) {
	graph = make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for _, edge := range prerequisites {
		from := edge[0]
		to := edge[1]

		graph[from] = append(graph[from], to) // from --> to: 边的方向是被依赖, from被to依赖, 修完from才能修to
	}

	return
}

// leetcode submit region end(Prohibit modification and deletion)
