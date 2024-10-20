package main

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	up := 0
	down := m - 1
	left := 0
	right := n - 1

	res := make([]int, 0)

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		// 读完当前行后, 看是否退出循环
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
