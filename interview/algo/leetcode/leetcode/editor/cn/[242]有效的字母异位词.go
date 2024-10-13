package main

import "golang.org/x/exp/slices"

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	s1 := []rune(s)
	s2 := []rune(t)

	slices.Sort(s1)
	slices.Sort(s2)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s2[i] != s1[i] {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
