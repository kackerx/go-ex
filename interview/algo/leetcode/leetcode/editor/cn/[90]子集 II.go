package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	slices.Sort(nums)
	backtr(nums, track, 0, &res)
	return res
}

func backtr(nums []int, track []int, index int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := index; i < len(nums); i++ {
		// 剪枝操作, 值相同的相邻树枝, 遍历第一条, 查看递归🌲, 每个枝是一个循环记下index模拟
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtr(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
