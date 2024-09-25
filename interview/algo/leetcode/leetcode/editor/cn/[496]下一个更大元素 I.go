package main

// leetcode submit region begin(Prohibit modification and deletion)
type stack []int

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len((*s))-1]
	return
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(e int) {
	*s = append(*s, e)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := nextGreater(nums2)

	m := make(map[int]int)
	for i, v := range nums2 {
		m[v] = greater[i]
	}

	res := make([]int, len(nums1))

	for i, v := range nums1 {
		res[i] = m[v]
	}

	return res
}

func nextGreater(nums []int) []int {
	res := make([]int, len(nums))
	s := &stack{}

	for i := len(nums) - 1; i >= 0; i-- {
		for !s.isEmpty() && s.peek() <= nums[i] {
			s.pop()
		}

		if s.isEmpty() {
			res[i] = -1
		} else {
			res[i] = s.peek()
		}

		s.push(nums[i])
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
