package main

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	bs := []byte(s)

	bs = clearSpace(bs)
	reve(bs, 0, len(bs)-1)

	for i := 0; i < len(bs); {
		for j := i; j < len(bs); j++ {
			if j == len(bs)-1 {
				reve(bs, i, j)
				i = j + 1
			} else if bs[j] == ' ' {
				reve(bs, i, j-1)
				i = j + 1
			}
		}
	}

	return string(bs)
}

func reve(bs []byte, start, end int) {
	for start < end {
		bs[start], bs[end] = bs[end], bs[start]
		start++
		end--
	}
}

func clearSpace(bs []byte) []byte {
	res := make([]byte, 0)
	for i := 0; i < len(bs); i++ {
		if i == 0 && bs[i] == ' ' {
			continue
		}

		if bs[i] == ' ' && bs[i-1] == ' ' {
			continue
		}

		res = append(res, bs[i])
	}

	if res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
