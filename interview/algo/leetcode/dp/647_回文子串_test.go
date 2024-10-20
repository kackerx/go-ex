package dp

import (
	"fmt"
	"testing"
)

func Test647(t *testing.T) {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n) // dp[i] --> 前i位的最大回文子串有dp[i]个
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	res := 0

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}

	return res
}

func palindrom(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return s[i+1 : j]
}
