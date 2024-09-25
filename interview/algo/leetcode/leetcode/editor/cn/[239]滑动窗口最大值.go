package main

import (
	"container/list"
)

//leetcode submit region begin(Prohibit modification and deletion)
type MonotonicQueue struct {
	maxq *list.List
}

func (m *MonotonicQueue) Push(e int) {
	for !m.IsEmpty() && e > m.maxq.Back().Value.(int) {
		m.maxq.Remove(m.maxq.Back())
	}
	m.maxq.PushBack(e)
}

func (m *MonotonicQueue) Pop(e int) {
	if !m.IsEmpty() && m.maxq.Front().Value.(int) == e {
		m.maxq.Remove(m.maxq.Front())
	}
}

func (m *MonotonicQueue) Max() int {
	if !m.IsEmpty() {
		return m.maxq.Front().Value.(int)
	}

	return -1
}

func (m *MonotonicQueue) IsEmpty() bool {
	return m.maxq.Len() == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	win := &MonotonicQueue{list.New()}
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			win.Push(nums[i])
			continue
		}

		win.Push(nums[i])
		res = append(res, win.Max())
		win.Pop(nums[i - k + 1])
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
