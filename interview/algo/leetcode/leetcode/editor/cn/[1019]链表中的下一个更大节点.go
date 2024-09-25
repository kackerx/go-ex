package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type stack []int

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(e int) {
	*s = append(*s, e)
}

func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	s := &stack{}
	greater := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for !s.isEmpty() && s.peek() <= nums[i] {
			s.pop()
		}

		if s.isEmpty() {
			greater[i] = 0
		} else {
			greater[i] = s.peek()
		}

		s.push(nums[i])
	}

	return greater
}

// leetcode submit region end(Prohibit modification and deletion)
