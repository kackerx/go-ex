package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		intv := points[i]

		if intv[0] > end {
			count++
			end = intv[1]
		}
	}

	return count
}

// leetcode submit region end(Prohibit modification and deletion)
