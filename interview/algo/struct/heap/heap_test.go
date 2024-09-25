package heap

import (
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	h := NewMaxHeap(100)
	n := 1000

	a := make([]int, 0)
	for i := 0; i < 20; i++ {
		a = append(a, i)
	}

	h.Heapify(a)

	for i := 0; i < n; i++ {
		fmt.Println(h.ExtractMax())
	}
}

func TestInit(t *testing.T) {
	pq := &pp{}

	heap.Init(pq)
	heap.Push(pq, &Node{
		val:  2,
		next: nil,
	})
	heap.Push(pq, &Node{
		val:  0,
		next: nil,
	})
	heap.Push(pq, &Node{
		val:  1,
		next: nil,
	})

	fmt.Println(heap.Pop(pq))
	fmt.Println(heap.Pop(pq))
}
