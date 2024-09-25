package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	patternList := strings.Split(pattern, "")
	wordList := strings.Split(s, " ")
	wordSet := map[string]bool{}
	patternToWord := map[string]string{}

	if len(wordList) != len(patternList) {return false}

	for i, c := range patternList {
		word := wordList[i]
		if w, ok := patternToWord[c]; !ok {
			if wordSet[word] {
				return false
			}
			patternToWord[c] = word
		} else {
			if w != word {
				return false
			}
		}

		wordSet[word] = true
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
