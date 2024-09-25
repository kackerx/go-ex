package queue

import (
	"go-bobo/interview/algo/struct/heap"
)

type PriorityQueue struct {
	maxHeap *heap.MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap.NewMaxHeap(100)}
}

func (p *PriorityQueue) Enqueue(e int) {
	p.maxHeap.Add(int(e))
}

func (p *PriorityQueue) Dequeue() int {
	return p.maxHeap.ExtractMax()
}

func (p *PriorityQueue) GetFront() int {
	return p.maxHeap.FindMax()
}

func (p *PriorityQueue) Len() int {
	return p.maxHeap.Size()
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.maxHeap.IsEmpty()
}
