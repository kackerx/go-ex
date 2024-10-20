package main

// leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var found bool

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, word, i, j, 0, &found)
			if found {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j, p int, found *bool) {
	if p == len(word) {
		*found = true
		return
	}

	m, n := len(board), len(board[0])

	if *found {
		return
	}

	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[p] {
		return
	}

	tmp := board[i][j]
	board[i][j] = '#'

	dfs(board, word, i-1, j, p+1, found)
	dfs(board, word, i+1, j, p+1, found)
	dfs(board, word, i, j-1, p+1, found)
	dfs(board, word, i, j+1, p+1, found)

	board[i][j] = tmp
}

// leetcode submit region end(Prohibit modification and deletion)
