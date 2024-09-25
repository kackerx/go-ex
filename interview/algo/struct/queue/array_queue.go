package queue

import (
	"fmt"

	"go-bobo/interview/algo/struct/array"
)

type ArrayQueue[T comparable] struct {
	arr *array.Array[T]
}

func NewArrayQueue[T comparable](cap int) *ArrayQueue[T] {
	return &ArrayQueue[T]{arr: array.NewArray[T](cap)}
}

func (q *ArrayQueue[T]) Enqueue(e T) { q.arr.AddLast(e) }

func (q *ArrayQueue[T]) Dequeue() T {
	return q.arr.RemoveFirst()
}

func (q *ArrayQueue[T]) GetFront() T {
	return q.arr.GetFirst()
}

func (q *ArrayQueue[T]) Len() int {
	return q.arr.Len()
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.arr.IsEmpty()
}

func (q *ArrayQueue[T]) String() string {
	var res string
	for i := 0; i < q.arr.Len(); i++ {
		item, err := q.arr.Get(i)
		if err != nil {
			panic(err)
		}

		res += fmt.Sprintf("%v", item)
		if i != q.arr.Len()-1 {
			res += "-"
		}
	}

	return res
}
