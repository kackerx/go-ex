package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)

	for i := 0; res < len(g) && i < len(s); i++ {
		if s[i] >= g[res] {
			res++
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
