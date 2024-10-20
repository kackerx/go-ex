package diff

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test238(t *testing.T) {
	s := []int{2, 3, -2, 4}
	fmt.Println(productExceptSelf(s))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// 构建前缀积
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	// 构建后缀积, 从右-->左的前缀积
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	fmt.Println(prefix, suffix)

	return res
}

func Test152(t *testing.T) {
	fmt.Println(maxProduct([]int{3, -1, 4}))
}

func maxProduct(nums []int) int {
	n := len(nums)
	pre := make([]int, n)

	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i]
	}

	fmt.Println(pre)
	res := slices.Max(pre)
	if res <= 0 {
		res = slices.Max(nums)
	}
	return res
}
