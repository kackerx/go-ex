package stack

type Stock[T any] interface {
	Len() int
	IsEmpty() bool
	Push(e int)
	Pop() int
	Peek() int
}
