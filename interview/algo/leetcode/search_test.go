package leetcode

import (
	"fmt"
	"testing"
)

var nums = []int{1, 2, 2, 3, 3, 3, 5}

func TestBs(t *testing.T) {
	// fmt.Println(bs(nums, 5))
	fmt.Println(leftBound1(nums, 10))
}

func bs(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// leftBound1 当target不存在时 --> 大于taget的最小值
func leftBound1(nums []int, target int) int {
	left := 0
	right := len(nums) // [left, right)

	for left < right { // 由于right不在数组合法索引, 所以范围是<不是<=
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 索引越界说明无目标元素
	// 如果不在这里返回-1, 那么返回的left要么是0要么就是len(nums)就是大于target的最小值索引, left-1 == floor
	if left < 0 || right >= len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

// floor 小于target的最大值
func floor(nums []int, target int) int {
	return leftBound1(nums, target) - 1
}

// rightBound 当target不存在时 --> 小于target的最大值
func rightBound(nums []int, target int) int {
	left := 0
	right := len(nums) // [left, right)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1 // 增大左区间, mid已经搜索过了可以不要
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 索引越界无元素
	// 此时如果不返回-1, left-1返回的是小于target的最大值, 因为for循环结束时left=right, left最后的赋值是mid+1, nums[mid]=target, 所以left-1 = mid
	if left-1 < 0 || left-1 >= len(nums) {
		return -1
	}

	if nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}
}
