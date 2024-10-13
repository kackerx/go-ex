package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func Test90(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	var (
		res   [][]int
		track []int
	)

	sort.Ints(nums)
	backtrack90(nums, 0, track, &res)
	return res
}

func backtrack90(nums []int, idx int, track []int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtrack90(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func Test78(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var (
		res   [][]int
		track []int
	)

	backtrack78(nums, 0, track, &res)

	return res
}

func backtrack78(nums []int, idx int, track []int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := idx; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack78(nums, i+1, track, res)
		track = track[:len(track)-1]
	}

}

func Test47(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack47(nums, used, []int{}, &res)

	return res
}

func backtrack47(nums []int, used []bool, track []int, res *[][]int) {
	if len(nums) == len(track) {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack47(nums, used, track, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}

func Test46(t *testing.T) {
	fmt.Println(permute([]int{1, 1, 2}))
}

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack46(nums, used, []int{}, &res)
	return res
}

func backtrack46(nums []int, used []bool, track []int, res *[][]int) {
	if len(track) == len(nums) {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack46(nums, used, track, res)
		used[i] = false
		track = track[:len(track)-1]
	}
}

func Test39(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrack39(candidates, target, 0, 0, []int{}, &res)
	return res
}

func backtrack39(candidates []int, target, sum, idx int, track []int, res *[][]int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target {
		return
	}

	for i := idx; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack39(candidates, target, sum, i, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}

func Test77(t *testing.T) {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var res [][]int
	backtrack77(n, k, 1, []int{}, &res)
	return res
}

func backtrack77(n, k, idx int, track []int, res *[][]int) {
	if len(track) == k {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := idx; i < n+1; i++ {
		track = append(track, i)
		backtrack77(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func Test40(t *testing.T) {
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1}, 4))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	backtrack40(candidates, target, 0, 0, []int{}, &res)

	return res
}

func backtrack40(candidates []int, target, sum, index int, track []int, res *[][]int) {
	if sum == target {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack40(candidates, target, sum, i+1, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}

func Test216(t *testing.T) {
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backtrack216(k, n, 0, 1, []int{}, &res)

	return res
}

func backtrack216(k, n, sum, index int, track []int, res *[][]int) {
	if len(track) == k && sum == n {
		*res = append(*res, slices.Clone(track))
		return
	}

	if sum >= n {
		return
	}

	for i := index; i < 10; i++ {
		track = append(track, i)
		sum += i
		backtrack216(k, n, sum, i+1, track, res)
		sum -= i
		track = track[:len(track)-1]
	}
}
