package string

import (
	"fmt"
	"testing"
)

func TestReve(t *testing.T) {
	s := "我爱中国"

	fmt.Println(reverse(s))
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
