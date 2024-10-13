package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func Test76(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	win, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	var (
		l, r, valid, start int
		length             = math.MaxInt
	)

	for r < len(s) {
		c := s[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for valid == len(need) {
			if r-l < length {
				start = l
				length = r - l
			}

			d := s[l]
			l++
			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}

				win[d]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	} else {
		return s[start : start+length]
	}
}

func Test567(t *testing.T) {
	fmt.Println(checkInclusion("ad", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	win, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	var (
		l, r, valid int
	)

	for r < len(s2) {
		c := s2[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if v == win[c] {
				valid++
			}
		}

		for r-l >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[l]
			l++

			if v, ok := need[d]; ok {
				if v == win[c] {
					valid--
				}
				win[d]--
			}
		}
	}

	return false
}

func Test3(t *testing.T) {

}

func lengthOfLongestSubstring(s string) int {
	maxRes := math.MinInt
	win := make(map[byte]int)

	var l, r int
	for r < len(s) {
		c := s[r]
		r++
		win[c]++

		for win[c] > 1 {
			d := s[l]
			l++
			win[d]--
		}

		maxRes = max(maxRes, r-l)
	}

	return maxRes
}

func Test438(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	win, need := make(map[byte]int), make(map[byte]int)
	res := make([]int, 0)

	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	var l, r, valid int

	for r < len(s) {
		c := s[r]
		r++
		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for r-l >= len(p) {
			if valid == len(need) {
				res = append(res, l)
			}
			d := s[l]
			l++
			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}
				win[d]--
			}
		}
	}

	return res
}
