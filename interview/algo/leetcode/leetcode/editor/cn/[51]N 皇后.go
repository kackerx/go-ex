package main

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var res [][]string
	board := newBoard(n)

	backTrack(board, 0, &res)

	return res
}

func backTrack(board [][]byte, row int, res *[][]string) {
	if row == len(board) {
		var item []string
		for _, bytes := range board {
			item = append(item, string(bytes))
		}
		*res = append(*res, item)
	}

	n := len(board)
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}

		board[row][col] = 'Q'
		backTrack(board, row+1, res)
		board[row][col] = '.'
	}
}

func isValid(board [][]byte, row int, col int) bool {
	n := len(board)
	// 检测列上是否有Q
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检测右上方斜列是否有
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检测左上方斜列是否有
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func newBoard(n int) [][]byte {
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			res[i][j] = '.'
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
