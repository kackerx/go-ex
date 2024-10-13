package main

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left < 0 || left >= len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func rightBound(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left-1 < 0 || left-1 >= len(nums) {
		return -1
	}

	if nums[left-1] == target {
		return left - 1
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
