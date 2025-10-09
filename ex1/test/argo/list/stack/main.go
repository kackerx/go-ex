package main

import "fmt"

type Stack struct {
	mainQueue *Queue
	tempQueue *Queue
	top, size int
}

func (s *Stack) Push(e int) {
	s.tempQueue.EnQueue(e)

	for !s.mainQueue.IsEmpty() {
		s.tempQueue.EnQueue(s.mainQueue.Dequeue())
	}

	s.mainQueue, s.tempQueue = s.tempQueue, s.mainQueue
	s.top = e
}

func (s *Stack) Pop() int {
	return s.mainQueue.Dequeue()
}

func (s *Stack) Peek() {
	fmt.Println(s.top)
}

func main() {
	s := &Stack{
		mainQueue: new(Queue),
		tempQueue: new(Queue),
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	s.Push(4)
	s.Push(5)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
