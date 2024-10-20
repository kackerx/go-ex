package dp

import (
	"fmt"
	"testing"
)

func Test674(t *testing.T) {
	s := []int{
		1, 3, 5, 7,
	}
	fmt.Println(findLengthOfLCIS(s))
}

func findLengthOfLCIS(nums []int) int {
	res := 1

	for i := 0; i < len(nums); i++ {
		maxLen := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[j-1] {
				maxLen++
			} else {
				res = max(res, maxLen)
				i = j
				break
			}
		}
	}

	return res
}
