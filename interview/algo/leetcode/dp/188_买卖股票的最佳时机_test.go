package dp

import (
	"fmt"
	"testing"
)

func Test188(t *testing.T) {

}

/*
i: 天数, k: 交易数的上限, 0/1: 是否持有股票, 开始购买buy减少交易上限k
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i]), // 今天没有持有股票的状态的两种选择, 1. 前一天没有持有今天休息k还是k, 2. 前一天持有, 今天卖掉获取收益
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - price[i]),// 昨天持有股票今天休息, 或者昨天未持有股票且昨天交易上限是[k-1], 今天买入当天的交易上限为k

// base case
dp[-1][...][0] = 0 --> i=-1还没开始, 利润为0
dp[-1][...][1] = -infinity --> 还没开始不可能持有股票, 所以设置一个最小值

dp[...][0][0] = 0 --> k=0, 没权利交易, 利润为0
dp[...][0][1] = -infinity --> 没权利交易不可能持有股票, 最小值
*/

func Test122(t *testing.T) {

}

func maxProfit122(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func maxProfit(k int, prices []int) int {
	dp := make([][][]int, 0)

	return dp[1][0][0]
}

func Test121(t *testing.T) {
	s := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit121(s))
}

func maxProfit121(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			// 根据状态转移方程可得：
			//   dp[i][0]
			// = max(dp[-1][0], dp[-1][1] + prices[i])
			// = max(0, -infinity + prices[i]) = 0
			// = max(dp[-1][0], dp[-1][1] + prices[i])
			// = max(0, -infinity + prices[i]) = 0

			dp[i][1] = -prices[i]
			// 根据状态转移方程可得：
			//   dp[i][1]
			// = max(dp[-1][1], dp[-1][0] - prices[i])
			// = max(-infinity, 0 - prices[i])
			// = -prices[i]
			// = max(dp[-1][1], dp[-1][0] - prices[i])
			// = max(-infinity, 0 - prices[i])
			// = -prices[i]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}
