package main

// leetcode submit region begin(Prohibit modification and deletion)
func closedIsland(grid [][]int) int {
	var count int
	m, n := len(grid), len(grid[0])
	// 求封闭的岛屿, 上下左右的陆地优先淹没
	for j := 0; j < n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m-1, j)
	}

	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])

	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid[i][j] == 1 {
		return
	}

	grid[i][j] = 1
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// leetcode submit region end(Prohibit modification and deletion)
