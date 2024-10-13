package greed

import (
	"fmt"
	"math"
	"testing"
)

func Test53(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxSum = max(maxSum, sum) // 每次都累加, 然后和最大和比较交换
		if sum < 0 {              // 和如果是负数, 和从0开始, 从下一个下标重新计算
			sum = 0
		}
	}

	return maxSum
}
