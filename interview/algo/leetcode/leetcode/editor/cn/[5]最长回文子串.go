package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := palidrome(s, i, i)
		s2 := palidrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palidrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}

// leetcode submit region end(Prohibit modification and deletion)
