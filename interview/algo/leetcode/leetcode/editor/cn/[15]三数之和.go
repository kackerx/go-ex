package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		tuples := twoSum1(nums, i+1, -nums[i])
		for _, ints := range tuples {
			ints = append(ints, nums[i])
			res = append(res, ints)
		}

		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

	}

	return res
}

func twoSum1(nums []int, start, target int) [][]int {
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	var res [][]int

	lo := start
	hi := len(nums) - 1
	for lo < hi {
		left := nums[lo]
		right := nums[hi]

		sum := left + right
		if sum == target {
			res = append(res, []int{nums[lo], nums[hi]})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else {
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}

	return res
}


// leetcode submit region end(Prohibit modification and deletion)
