package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {

	h := &Heap{}
	heap.Init(h)

	for i := 0; i < len(nums); i++ {
		if i < k {
			heap.Push(h, nums[i])
		} else if nums[i] > h.Peek() {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	return heap.Pop(h).(int)
}

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func (h *Heap) Peek() int {
	return (*h)[0]
}

// leetcode submit region end(Prohibit modification and deletion)
