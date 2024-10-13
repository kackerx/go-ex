package main

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack46(nums, used, []int{}, &res)
	return res
}

func backtrack46(nums []int, used []bool, track []int, res *[][]int) {
	if len(track) == len(nums) {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack46(nums, used, track, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
