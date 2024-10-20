package main

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	num := 1
	up := 0
	down := n - 1
	left := 0
	right := n - 1

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for {
		for i := left; i <= right; i++ {
			res[up][i] = num
			num++
		}
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res[i][right] = num
			num++
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res[down][i] = num
			num++
		}
		down--
		if down < up {
			break
		}

		for i := down; i >= up; i-- {
			res[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
