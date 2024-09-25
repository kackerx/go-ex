package main
//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	s := &set{}
	for _, num := range nums {
		s.add(num)
	}

	var maxLen int

	for num := range *s {
		if s.has(num - 1) {
			continue
		}

		curNum := num
		curCnt := 1
		for {
			if s.has(curNum + 1) {
				curNum++
				curCnt++
			} else {
				break
			}
		}
		maxLen = max(maxLen, curCnt)
	}

	return maxLen
}

type set map[int]struct{}

func (s *set) has(key int) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *set) add(key int) {
	(*s)[key] = struct{}{}
}
//leetcode submit region end(Prohibit modification and deletion)
