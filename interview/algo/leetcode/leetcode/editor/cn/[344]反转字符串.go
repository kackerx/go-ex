package main
//leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte)  {
	left, right := 0, len(s)-1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
//leetcode submit region end(Prohibit modification and deletion)
