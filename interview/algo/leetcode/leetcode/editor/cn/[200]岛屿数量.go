package main

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	var count int

	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return // 超出索引边界
	}

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// leetcode submit region end(Prohibit modification and deletion)
