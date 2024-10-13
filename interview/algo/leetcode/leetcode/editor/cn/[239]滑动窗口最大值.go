package main

import "container/list"

// leetcode submit region begin(Prohibit modification and deletion)

type MaxQueue struct {
	*list.List
}

func (m *MaxQueue) Push(e int) {
	for m.Len() > 0 && m.Back().Value.(int) < e {
		m.Remove(m.Back())
	}

	m.PushBack(e)
}

func (m *MaxQueue) Pop(e int) {
	if m.Front().Value.(int) == e {
		m.Remove(m.Front())
	}
}

func (m *MaxQueue) Max() int {
	return m.Front().Value.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	mq := &MaxQueue{list.New()}

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			mq.Push(nums[i])
		} else {
			mq.Push(nums[i])
			res = append(res, mq.Max())
			mq.Pop(nums[i-k+1])
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
