package main

// leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1) // dp[i] --> s的前i位字符即s[0:i]可以通过字典中的单词进行拆分
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && slices.Contains(wordDict, s[j:i]) { // [0:j)可以拆分, [j:i)又是dict的单词, 所以[0:i)也可以
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)
