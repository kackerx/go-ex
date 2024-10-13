package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	var res int

	for i := 1; i < len(prices); i++ {
		res += max(prices[i]-prices[i-1], 0) // 当天价格减去前一天价格为正有收益我们就加入收益和
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
