package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)

type MinHeap struct {
	data []int
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(e any) {
	h.data = append(h.data, e.(int))
}

func (h *MinHeap) Pop() any {
	ret := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return ret
}

func (h *MinHeap) Peek() int {
	return h.data[0]
}

func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i := 0; i < len(nums); i++ {
		if i < k {
			heap.Push(minHeap, nums[i])
		} else {
			if nums[i] >= minHeap.Peek() {
				heap.Pop(minHeap)
				heap.Push(minHeap, nums[i])
			}
		}
	}

	return minHeap.Peek()
}

// leetcode submit region end(Prohibit modification and deletion)
