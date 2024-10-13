package main

// leetcode submit region begin(Prohibit modification and deletion)
func removeCoveredIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return b[1] - a[1]
		}
		return a[0] - b[0]
	})

	var res int
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		ints := intervals[i]

		// 覆盖区间
		if left <= ints[0] && right >= ints[1] {
			res++
		}

		// 相交区间, 合并
		if right >= ints[0] && right <= ints[1] {
			right = ints[1]
		}

		// 不相交, 更新起点, 终点
		if right < ints[0] {
			left = ints[0]
			right = ints[1]
		}
	}

	return len(intervals) - res
}

// leetcode submit region end(Prohibit modification and deletion)
