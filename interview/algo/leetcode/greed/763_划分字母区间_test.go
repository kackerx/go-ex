package greed

import (
	"fmt"
	"testing"
)

func Test763(t *testing.T) {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(len(s))

	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	ht := farthest(s)
	res := make([]int, 0)

	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		end = max(end, ht[s[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res
}

func farthest(s string) [26]int {
	ht := [26]int{}

	for i := 0; i < len(s); i++ {
		ht[s[i]-'a'] = i
	}

	return ht
}
