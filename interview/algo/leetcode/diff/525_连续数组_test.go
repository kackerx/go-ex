package diff

import (
	"fmt"
	"testing"
)

func Test525(t *testing.T) {
	nums := []int{0, 1, 0, 0, 1, 1, 0}
	result := findMaxLength(nums)
	fmt.Println(result) // 输出 6
}

func findMaxLength(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	valToIndex := make(map[int]int)
	res := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			preSum[i+1] = preSum[i] - 1
		} else {
			preSum[i+1] = preSum[i] + 1
		}
	}

	for i := 0; i < len(preSum); i++ {
		v, ok := valToIndex[preSum[i]]
		if ok {
			res = max(res, i-v)
		} else {
			valToIndex[preSum[i]] = i
		}
	}

	return res
}

type User struct {
	name string
}

func Test523(t *testing.T) {
	s := []int{23, 2, 4, 6, 7}
	fmt.Println(checkSubarraySum(s, 6))
}

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0

	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	valToIndex := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		val := preSum[i] % k
		if _, ok := valToIndex[val]; !ok {
			valToIndex[val] = i
		}
	}

	for i := 1; i < len(preSum); i++ {
		val := preSum[i] % k
		if idx, ok := valToIndex[val]; ok && i-idx >= 2 {
			return true
		}
	}

	return false
}
