package main

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i += 2 * k {
		end := i + k - 1
		if len(bs)-i < k {
			end = len(bs) - 1
		} else if len(bs)-i < 2*k && len(bs)-i >= k {
			end = i + k - 1
		}
		reverseS(bs, i, end)
	}

	return string(bs)
}

func reverseS(bs []byte, start, end int) {
	for start < end {
		bs[start], bs[end] = bs[end], bs[start]
		start++
		end--
	}
}

// leetcode submit region end(Prohibit modification and deletion)
