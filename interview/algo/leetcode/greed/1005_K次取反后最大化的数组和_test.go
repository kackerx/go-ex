package greed

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func Test1001(t *testing.T) {
	s := []int{2, -3, -1, 5, -4}
	fmt.Println(largestSumAfterKNegations(s, 2))
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
