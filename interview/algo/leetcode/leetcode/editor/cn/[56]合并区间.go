package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		lastMerged := merged[len(merged)-1]

		if cur[0] <= lastMerged[1] {
			lastMerged[1] = max(lastMerged[1], cur[1])
		} else {
			merged = append(merged, cur)
		}
	}

	return merged
}

// leetcode submit region end(Prohibit modification and deletion)
