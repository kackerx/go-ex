package dp

import (
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func Test494(t *testing.T) {
	// s := []int{100}
	// fmt.Println(findTargetSumWays(s, -200))
}

func TestBack(t *testing.T) {
	s := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(findTargetSumWaysBack(s, 3))
}

func findTargetSumWaysBack(nums []int, target int) int {
	sum := target
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return 0
	}
	sum = sum / 2

	sort.Ints(nums)
	var (
		res   [][]int
		track []int
	)
	backtrack(nums, sum, 0, 0, track, &res)

	return len(res)
}

func backtrack(nums []int, target, sum, idx int, track []int, res *[][]int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target {
		return
	}

	for i := idx; i < len(nums); i++ {
		// if i > idx && nums[i] == nums[i-1] {
		// 	continue
		// }

		track = append(track, nums[i])
		sum += nums[i]
		backtrack(nums, target, sum, i+1, track, res)
		sum -= nums[i]
		track = track[:len(track)-1]
	}
}
