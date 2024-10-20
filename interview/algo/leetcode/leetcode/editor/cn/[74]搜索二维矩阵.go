package main

// leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l := 0
	r := m*n - 1

	for l <= r {
		mid := l + (r-l)/2
		val := get(matrix, mid)
		if val == target {
			return true
		} else if val > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func get(matrix [][]int, index int) int {
	n := len(matrix[0])
	return matrix[index/n][index%n] // 整体是升序二分, 一维中的索引index/n是行数, %n是列数
}

// leetcode submit region end(Prohibit modification and deletion)
