package greed

import (
	"fmt"
	"sort"
	"testing"
)

func Test452(t *testing.T) {
	s := [][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}
	fmt.Println(findMinArrowShots(s))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	count := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		intv := points[i]

		if intv[0] > end {
			count++
			end = intv[1]
		} else {
			end = min(end, points[i-1][1])
		}
	}

	return count
}
