package leetcode

import "go-bobo/interview/algo/struct/stack"

type MyQueue[T comparable] struct {
	in  *stack.ArrayStock[T]
	out *stack.ArrayStock[T]
}

func NewMyQueue[T comparable]() *MyQueue[T] {
	return &MyQueue[T]{
		in:  stack.NewArrayStock[T](100),
		out: stack.NewArrayStock[T](100),
	}
}

func (q *MyQueue[T]) Enqueue(e T) {
	q.in.Push(e)
}

func (q *MyQueue[T]) Dequeue() T {
	if !q.out.IsEmpty() {
		return q.out.Pop()
	}

	for !q.in.IsEmpty() {
		q.out.Push(q.in.Pop())
	}

	return q.out.Pop()
}
