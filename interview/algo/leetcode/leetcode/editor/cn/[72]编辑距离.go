package main

// leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		} else if j == -1 {
			return i + 1
		}

		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			return min(
				dp(i, j-1)+1,   // 插入
				dp(i-1, j)+1,   // 删除
				dp(i-1, j-1)+1, // 替换
			)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

// leetcode submit region end(Prohibit modification and deletion)
