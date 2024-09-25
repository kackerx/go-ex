package queue

import "fmt"

type LoopQueue[T any] struct {
	data                []T
	front, tail, length int
}

func NewLoopQueue[T any](cap int) *LoopQueue[T] {
	return &LoopQueue[T]{data: make([]T, cap, cap*2)}
}

func (q *LoopQueue[T]) Enqueue(e T) {
	if q.IsFull() {
		panic("队列满了")
	}

	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.length++
}

func (q *LoopQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("队列为空")
	}

	res := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.length--

	return res
}

func (q *LoopQueue[T]) GetFront() T {
	if q.IsEmpty() {
		panic("队列未空")
	}
	return q.data[q.front]
}

func (q *LoopQueue[T]) Len() int { return q.length }

func (q *LoopQueue[T]) Cap() int { return len(q.data) - 1 }

func (q *LoopQueue[T]) IsEmpty() bool { return q.front == q.tail }

func (q *LoopQueue[T]) IsFull() bool { return q.front == (q.tail+1)%len(q.data) }

func (q *LoopQueue[T]) String() string {
	var res string

	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		res += fmt.Sprintf("%v", q.data[i])
		if (i+1)%len(q.data) != q.tail {
			res += "-"
		}
	}

	return res
}
