package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)

type pq []int

func (h *pq) Len() int {
	return len((*h))
}

func (h *pq) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *pq) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *pq) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *pq) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *pq) Peek() int {
	return (*h)[0]
}

func findKthLargest(nums []int, k int) int {
	q := &pq{}
	heap.Init(q)

	for _, num := range nums {
		if q.Len() < k {
			heap.Push(q, num)
		} else if q.Peek() <= num { // 最小堆保持k个最大的值
			heap.Pop(q)
			heap.Push(q, num)
		}
	}

	return q.Peek()
}

// leetcode submit region end(Prohibit modification and deletion)
