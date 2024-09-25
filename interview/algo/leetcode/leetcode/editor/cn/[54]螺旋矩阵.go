package main

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	var res []int
	m := len(matrix)
	n := len(matrix[0])

	upper := 0
	lower := m - 1
	left := 0
	right := n - 1

	for len(res) < m*n {
		if upper <= lower {
			for i := left; i <= right; i++ {
				res = append(res, matrix[upper][i])
			}
			upper++
		}

		if left <= right {
			for i := upper; i <= lower; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

	}
}

// leetcode submit region end(Prohibit modification and deletion)
