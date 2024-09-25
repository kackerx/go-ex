package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewArrayQueue[int](10)

	q.Enqueue(1)
	q.Enqueue(2)

	fmt.Println(q)

	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q)
}

func TestLoopQueue(t *testing.T) {
	q := NewLoopQueue[int](10)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q)
}
