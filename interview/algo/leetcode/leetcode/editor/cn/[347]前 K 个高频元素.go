package main

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)

type pair struct {
	num, cnt int
}

type pq []pair

func (h *pq) Len() int {
	return len(*h)
}

func (h *pq) Less(i, j int) bool {
	return (*h)[i].cnt < (*h)[j].cnt
}

func (h *pq) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *pq) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *pq) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *pq) Peek() int {
	return (*h)[0].num
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	q := &pq{}
	heap.Init(q)

	for key, val := range m {
		if q.Len() < k {
			heap.Push(q, pair{
				num: key,
				cnt: val,
			})
			continue
		}

		if m[q.Peek()] < val {
			heap.Pop(q)
			heap.Push(q, pair{
				num: key,
				cnt: val,
			})
		}
	}

	res := make([]int, 0)

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(q).(pair).num)
	}

	return res
}


// leetcode submit region end(Prohibit modification and deletion)
