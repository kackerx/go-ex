package hash

import (
	"fmt"
	"math"
	"testing"
)

func Test169(t *testing.T) {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	ht := [50000]int{}
	for _, num := range nums {
		ht[num]++
	}

	fmt.Println(ht)

	maxVal := math.MinInt
	var ret int
	for i, v := range ht {
		if v > maxVal {
			maxVal = v
			ret = i
		}
	}

	return ret
}
