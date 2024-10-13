package main

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	s := &set{}

	for _, v := range nums1 {
		s.Add(v)
	}

	for _, v := range nums2 {
		if s.Has(v) {
			res = append(res, v)
			s.Rem(v)
		}
	}

	return res
}

type set map[int]struct{}

func (s *set) Has(v int) bool {
	if _, ok := (*s)[v]; ok {
		return true
	}

	return false
}

func (s *set) Add(v int) {
	(*s)[v] = struct{}{}
}

func (s *set) Rem(v int) {
	delete(*s, v)
}

// leetcode submit region end(Prohibit modification and deletion)
