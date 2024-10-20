package main

import (
	"fmt"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		code := encode(str)
		m[code] = append(m[code], str)
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func encode(s string) string {
	count := [26]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	sb := strings.Builder{}
	for _, v := range count {
		sb.WriteString(fmt.Sprintf(`%d#`, v))
	}

	return sb.String()
}

// leetcode submit region end(Prohibit modification and deletion)
