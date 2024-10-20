package diff

import (
	"fmt"
	"slices"
	"testing"
)

func Test560(t *testing.T) {
	// s := []int{1, 1, 1}
	s := []int{1, -1, 0}
	// s := []int{1, 2, 3}
	fmt.Println(subarraySum(s, 0))
}

func subarraySum(nums []int, k int) int {
	var (
		res   [][]int
		track []int
	)

	backtrack(nums, 0, 0, k, track, &res)

	fmt.Println(res)
	return len(res)
}

func backtrack(nums []int, idx, sum, k int, track []int, res *[][]int) {
	// if len(track) == len(nums) {
	// *res = append(*res, slices.Clone(track))
	fmt.Println("sumk", sum, "---", k, "---", track)
	// 	return
	// }

	if sum == k && len(track) != 0 {
		*res = append(*res, slices.Clone(track))
		// return
	}
	// if sum > k {
	// 	return
	// }

	for i := idx; i < len(nums); i++ {
		if i > idx && i-idx >= 1 && len(track) != 0 {
			continue
		}
		track = append(track, nums[i])
		sum += nums[i]

		backtrack(nums, i+1, sum, k, track, res)

		track = track[:len(track)-1]
		sum -= nums[i]
	}
}
