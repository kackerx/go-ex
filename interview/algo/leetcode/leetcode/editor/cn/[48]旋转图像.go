package main

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	m := len(matrix)

	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		reverse(row)
	}
}

func reverse(row []int) {
	i, j := 0, len(row)-1
	for i < j {
		row[i], row[j] = row[j], row[i]
		i++
		j--
	}
}

// leetcode submit region end(Prohibit modification and deletion)
