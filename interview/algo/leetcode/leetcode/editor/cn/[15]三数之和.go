package main

import "slices"

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var res [][]int

	slices.Sort(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		twoRes := twoSum(nums, i+1, -nums[i])
		for _, tuple := range twoRes {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}

		for i+1 < n && nums[i] == nums[i+1] {
			i++
		}

	}

	return res
}

func twoSum(nums []int, start, target int) [][]int {
	res := make([][]int, 0)
	l, r := start, len(nums)-1
	for l < r {
		left, right := nums[l], nums[r]
		sum := nums[l] + nums[r]
		if sum == target {
			res = append(res, []int{nums[l], nums[r]})
			for l < r && nums[l] == left {
				l++
			}
			for l < r && nums[r] == right {
				r--
			}
		} else if sum < target {
			for l < r && nums[l] == left {
				l++
			}
		} else {
			for l < r && nums[r] == right {
				r--
			}
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
