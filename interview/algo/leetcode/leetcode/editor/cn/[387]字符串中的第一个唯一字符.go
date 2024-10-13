package main

// leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	ht := [26]int{}

	for i := 0; i < len(s); i++ {
		ht[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if ht[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
