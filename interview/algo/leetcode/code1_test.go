package leetcode

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	// a := []string{"a", "b", "c", "d", "a", "b", "c", "a"}
	// a := []string{"a", "a", "b", "a", "a", "b", "a", "a", "a"}
	// a := []string{"a", "b", "c", "a", "b", "y"}

	fmt.Println(kmp("abxabcabcaby", "abcaby"))
}

func kmp(s string, pat string) int {
	next := buildNext([]byte(pat))

	j := 0
	for i := 0; i < len(s); i++ {
		if j > 0 && s[i] != pat[j] {
			j = next[j-1] // 遇到不匹配的位置, j跳到上一个最大前缀匹配
		}

		if s[i] == pat[j] { // 如果等于了j++
			j++
		}

		if j == len(pat) { // 此时走完了pat, 就证明找到了
			return i - j + 1
		}
	}

	return -1
}

// 构造next数组
// j i
// a b c d a b c a
// 0 1 2 3 4 5 6 7
// 0 0 0 0 1 2 3 1
// ---
// a a b a a b a a a
// 0 1 2 3 4 5 6 7 8
func buildNext(pat []byte) []int {
	// i为后缀末尾, j为前缀末尾, 只有长度在2的时候才有才后缀, 所以当aa时, j = 0, i = 1, 数组初始化0为0, 因为第一个元素没有前后缀
	// 如果pat[i] != pat[j], j = next[j-1]继续判断pat[i]和pat[j], next[i] = j, i++
	// 如果pat[i] == pat[j], next[i] = ++j, i++, j++
	next := make([]int, len(pat))
	j := 0
	for i := 1; i < len(pat); i++ {
		for j > 0 && pat[i] != pat[j] {
			j = next[j-1]
		}

		if pat[i] == pat[j] {
			j++
		}

		next[i] = j
	}

	return next
}

func TestFoo(t *testing.T) {
	// a := 1

}
