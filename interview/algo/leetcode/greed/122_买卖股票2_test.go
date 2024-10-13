package greed

import (
	"fmt"
	"testing"
)

func Test122(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var res int

	for i := 1; i < len(prices); i++ {
		res += max(prices[i]-prices[i-1], 0)
	}

	return res
}
