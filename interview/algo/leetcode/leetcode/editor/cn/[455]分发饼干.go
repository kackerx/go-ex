package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	var res int

	for i := 0; res < len(g) && i < len(s); i++ {
		if s[i] >= g[res] { // res初始化为0, 最小的饼干满足最小的孩子, 后续饼干不满足res就不增加, g[res]也不会增加, 直到满足了该孩子才轮到下一个孩子
			res++
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
