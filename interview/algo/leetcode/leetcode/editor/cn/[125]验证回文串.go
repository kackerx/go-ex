package main

import (
	"strings"
	"unicode"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	sb := strings.Builder{}
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			sb.WriteRune(unicode.ToLower(v))
		}
	}

	s = sb.String()
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
