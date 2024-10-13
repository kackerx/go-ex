package main

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		intv := intervals[i]

		if intv[0] >= end {
			count++
			end = intv[1]
		}

	}

	return len(intervals) - count // count是最多有几个区间不重叠, 那么总区间减去count
}

// leetcode submit region end(Prohibit modification and deletion)
