package array

import (
	"errors"
	"fmt"
)

type Array[T comparable] struct {
	data   []T
	length int
}

func (a *Array[T]) String() string {
	var res []T
	for i := 0; i < a.length; i++ {
		res = append(res, a.data[i])
	}
	return fmt.Sprintf("len: %d - cap: %d %v", a.length, a.Cap(), res)
}

func (a *Array[T]) Len() int {
	return a.length
}

func (a *Array[T]) Cap() int {
	return len(a.data)
}

func (a *Array[T]) IsFull() bool {
	return a.length == len(a.data)
}

func (a *Array[T]) IsLegal(index int) (err error) {
	if index < 0 || index >= a.length {
		return errors.New("索引不合法")
	}

	return
}

func (a *Array[T]) IsEmpty() bool {
	return a.length == 0
}

func (a *Array[T]) AddLast(e T) {
	a.Add(a.length, e)
}

func (a *Array[T]) AddFirst(e T) {
	a.Add(0, e)
}

func (a *Array[T]) Container(e T) bool {
	for i := 0; i < a.length; i++ {
		if a.data[i] == e {
			return true
		}
	}

	return false
}

func (a *Array[T]) Find() {
	// todo
}

func (a *Array[T]) Remove(index int) (e T) {
	e = a.data[index]
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}

	a.length--
	return
}

func (a *Array[T]) RemoveFirst() T {
	return a.Remove(0)
}

func (a *Array[T]) RemoveLast() T {
	return a.Remove(a.length - 1)
}

func (a *Array[T]) GetFirst() T {
	return a.data[0]
}

func (a *Array[T]) GetLast() T {
	return a.data[a.length-1]
}

func (a *Array[T]) Add(index int, e T) (err error) {
	if a.IsFull() {
		// todo
	}

	if index < 0 || index > a.length {
		return errors.New("索引位置无效")
	}

	// 每个元素后移一位
	for i := a.length - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	// 赋值维护
	a.data[index] = e
	a.length++

	return
}

func (a *Array[T]) Get(index int) (e T, err error) {
	if index < 0 || index >= a.length {
		return e, errors.New("")
	}

	return a.data[index], nil
}

func (a *Array[T]) Set(index int, e T) (err error) {
	if err = a.IsLegal(index); err != nil {
		return
	}

	a.data[index] = e

	return
}

func NewArray[T comparable](cap int) *Array[T] {
	return &Array[T]{data: make([]T, cap, cap*2), length: 0}
}
