package greed

import (
	"fmt"
	"testing"
)

func Test55(t *testing.T) {
	s := []int{2, 3, 1, 1, 4}
	// s := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(s))
}

func canJump(nums []int) bool {
	var cover int
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		fmt.Println(cover)
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}

func Test45(t *testing.T) {
	s := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(s))
}

func jump(nums []int) int {
	var end, jumps, farthest int

	for i := 0; i < len(nums); i++ {
		farthest = max(farthest, i+nums[i])

		if end == i {
			jumps++
			end = farthest

			if end >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}
