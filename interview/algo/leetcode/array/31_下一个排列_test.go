package array

import (
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func Test31(t *testing.T) {
	s := []int{5, 4, 7, 6, 3, 2, 1}
	nextPermutation(s)

	fmt.Println(s)
}

func nextPermutation(nums []int) {
	k := len(nums) - 1
	for k-1 >= 0 && nums[k-1] >= nums[k] {
		k--
	}

	if k <= 0 {
		slices.Reverse(nums)
		return
	}

	t := len(nums) - 1
	for nums[t] <= nums[k-1] {
		t--
	}

	nums[k-1], nums[t] = nums[t], nums[k-1]

	sort.Ints(nums[k:])
}
