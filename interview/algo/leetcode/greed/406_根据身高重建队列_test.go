package greed

import (
	"fmt"
	"sort"
	"testing"
)

func TestFoo(t *testing.T) {
	s := []int{1, 2, 3}
	s = append(s[:1], append([]int{1}, s[1:]...)...)
	fmt.Println(s)
}

func Test406(t *testing.T) {
	s := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	fmt.Println(reconstructQueue(s))
}

func reconstructQueue(people [][]int) [][]int {
	var res [][]int

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	for i := 0; i < len(people); i++ {
		if people[i][1] == i {
			res = append(res, people[i])
		} else {
			insertIndex := people[i][1]
			res = append(res[:insertIndex], append(
				[][]int{people[i]},
				res[insertIndex:]...,
			)...)
		}
	}

	return res
}
