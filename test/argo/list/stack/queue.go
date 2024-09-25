package main

import "fmt"

type Queue struct {
	data []int
	size int
}

func (q *Queue) EnQueue(e int) {
	q.size++
	q.data = append(q.data, e)
}

func (q *Queue) Dequeue() int {
	e := q.data[0]
	q.data = q.data[1:]

	q.size--

	return e
}

func (q *Queue) Top() int {
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Show() {
	for _, item := range q.data {
		fmt.Println(item)
	}
}
