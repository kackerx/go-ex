package main

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	bs := cleanSpace(s)
	rev(bs, 0, len(bs)-1)

	for i := 0; i < len(bs); {
		for j := i; j < len(bs); j++ {
			if j+1 == len(bs) || bs[j+1] == ' ' {
				rev(bs, i, j)
				i = j + 2
				break
			}
		}
	}

	return string(bs)

}

func cleanSpace(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			bs = append(bs, s[i])
		} else if len(bs) > 0 && bs[len(bs)-1] != ' ' {
			bs = append(bs, ' ')
		}
	}

	if bs[len(bs)-1] == ' ' {
		bs = bs[:len(bs)-1]
	}

	return bs
}

func rev(bs []byte, i, j int) {
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}

// leetcode submit region end(Prohibit modification and deletion)
