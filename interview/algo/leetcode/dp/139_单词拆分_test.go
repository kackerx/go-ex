package dp

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func wordBreakk(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1) // dp[i] --> s的前i位字符即s[0:i]可以通过字典中的单词进行拆分
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && slices.Contains(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func Test139_trck(t *testing.T) {
	s := "aaab"
	wordDict := []string{"a", "aa", "ab"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	var (
		found    bool
		track    []string
		traverse func(i int)
		memo     = make(map[string]bool)
	)

	traverse = func(i int) {
		if found {
			return
		}

		if i == len(s) {
			fmt.Println(track)
			found = true
			return
		}

		suffix := s[i:]
		if memo[suffix] {
			return
		}

		for _, word := range wordDict {
			length := len(word)

			if i+length <= len(s) && s[i:i+length] == word {
				track = append(track, word)
				traverse(i + length)
				track = track[:len(track)-1]
			}
		}

		memo[suffix] = true
	}

	traverse(0)
	return found
}
