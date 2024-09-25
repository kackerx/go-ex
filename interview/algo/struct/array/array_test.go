package array

import (
	"fmt"
	"math"
	"testing"
)

func Test76(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	var (
		need, win   = make(map[byte]int), make(map[byte]int)
		left, right int // [left, right), 初始[0, 0)窗口内无元素
		valid       int // 满足need条件的字符个数 valid == need.len
		start       int
		length      = math.MaxInt // 最小覆盖子串的, 起始索引和长度
	)
	for i := 0; i < len(t); i++ { // 初始化need
		need[t[i]]++
	}

	for right < len(s) {
		c := s[right] // 逐字符遍历
		right++       // 扩大窗口

		if v, ok := need[c]; ok { // 如果是条件字符, 那么进入win, 同时判断如果和要求字符数量一致则vaild++
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for valid == len(need) { // 收缩条件
			// 更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// 待移除字符
			d := s[left]
			left++
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

func TestName(t *testing.T) {
	arr := NewArray[int](10)
	arr.AddLast(1)
	arr.AddLast(1)
	arr.AddLast(1)

	if err := arr.Add(2, 4); err != nil {
		panic(err)
	}

	fmt.Println(arr)
	fmt.Println(arr.GetFirst(), arr.GetLast())

	arr.Remove(1)

	fmt.Println(arr)

	arr.RemoveFirst()

	fmt.Println(arr)

	arr.RemoveLast()

	fmt.Println(arr)
}

func TestFoo(t *testing.T) {
	s := "Hello, 世界"
	for index, char := range s {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
	for index, char := range []rune(s) {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}

func TestDiff(t *testing.T) {
	d := NewDifference(make([]int, 5))

	d.Increment(1-1, 2-1, 10)
	d.Increment(2-1, 3-1, 20)
	d.Increment(2-1, 5-1, 25)

	fmt.Println(d.Res())
}
