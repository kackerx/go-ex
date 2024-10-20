package diff

import (
	"fmt"
	"testing"
)

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)

	preSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	return NumArray{preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

func Test303(t *testing.T) {
	s := []int{-2, 0, 3, -5, 2, -1}
	//      0, -2, -2, 1, -4, -2, -3
	na := Constructor(s)
	fmt.Println(na.preSum)

	fmt.Println(na.SumRange(0, 2))
}
