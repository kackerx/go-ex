package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func reconstructQueue(people [][]int) [][]int {
	var res [][]int

	sort.Slice(people, func(i, j int) bool { // 我们首先按照身高高到低排序, 保证当前元素之前的一定比自己高
		if people[i][0] == people[j][0] { // 然后身高相同的话按照升序排序
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	for i := 0; i < len(people); i++ {
		// 当前元素的k是前面有几个, 所以必须等于i, 不为i就在i为止插入
		insertIndex := people[i][1]

		res = append(res[:insertIndex],
			append(
				[][]int{people[i]},
				res[insertIndex:]...,
			)...)
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
