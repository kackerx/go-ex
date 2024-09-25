package queue

type Queue[T any] interface {
	Enqueue(e T)
	Dequeue() T
	GetFront() T
	Len() int
	IsEmpty() bool
}
