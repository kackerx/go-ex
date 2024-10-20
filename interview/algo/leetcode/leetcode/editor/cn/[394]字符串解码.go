package main

import (
	"unicode"
)

// leetcode submit region begin(Prohibit modification and deletion)

func decodeString(s string) string {
	res, _ := dfs(s, 0)
	return res
}

func dfs(s string, i int) (string, int) {
	var (
		res   string
		multi int
	)

	for i < len(s) {
		switch {
		case unicode.IsDigit(rune(s[i])):
			multi = multi*10 + int(s[i]-'0')
		case unicode.IsLetter(rune(s[i])):
			res += string(s[i])
		case s[i] == '[':
			tmp, newIndex := dfs(s, i+1)
			i = newIndex
			for multi > 0 {
				res += tmp
				multi--
			}
		case s[i] == ']':
			return res, i
		}
		i++
	}

	return res, i
}

// leetcode submit region end(Prohibit modification and deletion)
