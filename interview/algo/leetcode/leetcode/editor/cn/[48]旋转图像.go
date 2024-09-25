package main

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	n := len(matrix)
	// 先沿对角线镜像对称二维矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// swap(matrix[i][j], matrix[j][i]);
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		row := matrix[i]
		j, k := 0, n-1
		for j < k {
			row[j], row[k] = row[k], row[j]
			j++
			k--
		}
	}
}
func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for j > i {
		// swap(arr[i], arr[j]);
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// leetcode submit region end(Prohibit modification and deletion)
