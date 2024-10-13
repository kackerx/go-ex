package greed

import (
	"fmt"
	"testing"
)

func Test376(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if len(nums) < 1 { // 根据题意只有一个元素时摆动为1
		return n
	}
	res := 1     // 默认最后一个值有一个摆动
	preDiff := 0 // 默认下标0前面有平坡, preDeff是0
	for i := 0; i < n-1; i++ {
		curDiff := nums[i+1] - nums[i]
		if preDiff >= 0 && curDiff < 0 || preDiff <= 0 && curDiff > 0 { // 摆动正负结果++, 同时更新preDiff
			res++
			preDiff = curDiff
		}
	}

	return res
}
