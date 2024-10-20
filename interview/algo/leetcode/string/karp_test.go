package string

import (
	"fmt"
	"math"
	"testing"
)

func TestKarp(t *testing.T) {
	s := "8267"

	num := 0
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
		fmt.Println(num)
	}

	fmt.Println(math.Log10(120))
}
