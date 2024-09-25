package stack

import (
	"fmt"

	"go-bobo/interview/algo/struct/array"
)

type ArrayStock[T comparable] struct {
	array *array.Array[T]
}

func NewArrayStock[T comparable](cap int) *ArrayStock[T] {
	return &ArrayStock[T]{array: array.NewArray[T](cap)}
}

func (a *ArrayStock[T]) Len() int {
	return a.array.Len()
}

func (a *ArrayStock[T]) IsEmpty() bool { return a.array.IsEmpty() }

func (a *ArrayStock[T]) Push(e T) { a.array.AddLast(e) }

func (a *ArrayStock[T]) Pop() T {
	return a.array.RemoveLast()
}

func (a *ArrayStock[T]) Peek() T {
	return a.array.GetLast()
}

func (a *ArrayStock[T]) String() string {
	var res string
	for i := 0; i < a.array.Len(); i++ {
		item, err := a.array.Get(i)
		if err != nil {
			panic(err)
		}

		res += fmt.Sprintf("%v", item)
		if i != a.array.Len()-1 {
			res += "-"
		}
	}

	return res
}
