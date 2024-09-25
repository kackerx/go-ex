package leetcode

import "go-bobo/interview/algo/struct/queue"

type MyStack[T comparable] struct {
	tmp *queue.ArrayQueue[T]
	q   *queue.ArrayQueue[T]
}

func NewMyStack[T comparable]() *MyStack[T] {
	return &MyStack[T]{
		queue.NewArrayQueue[T](10),
		queue.NewArrayQueue[T](10),
	}
}

func (s *MyStack[T]) Push(e T) {
	s.tmp.Enqueue(e)     // 临时队列入队
	for !s.q.IsEmpty() { // 出队队列全部进入临时队列
		s.tmp.Enqueue(s.q.Dequeue())
	}

	s.tmp, s.q = s.q, s.tmp
}

func (s *MyStack[T]) Pop() T {
	return s.q.Dequeue()
}
