package main

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	var (
		res []string
	)

	if digits == "" {
		return res
	}

	mapping := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	backtrack1(digits, 0, mapping, "", &res)

	return res
}

func backtrack1(digits string, idx int, mapping []string, track string, res *[]string) {
	if len(track) == len(digits) {
		*res = append(*res, track)
		return
	}

	digit := digits[idx] - '0'
	for _, c := range mapping[digit] {
		track = track + string(c)
		backtrack1(digits, idx+1, mapping, track, res)
		track = track[:len(track)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
