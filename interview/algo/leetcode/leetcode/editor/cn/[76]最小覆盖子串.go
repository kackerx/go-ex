package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)

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

// leetcode submit region end(Prohibit modification and deletion)
