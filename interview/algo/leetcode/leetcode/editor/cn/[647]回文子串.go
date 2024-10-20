package main

// leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n) // dp[i] --> 前i位的最大回文子串有dp[i]个
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	res := 0

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
