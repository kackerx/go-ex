package stack

import "go-bobo/interview/algo/struct/linked_list"

type LinkedListStack struct {
	list *linked_list.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{list: linked_list.NewLinkedList()}
}

func (l *LinkedListStack) Len() int {
	return l.list.Len()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedListStack) Push(e int) {
	l.list.AddFirst(e)
}

func (l *LinkedListStack) Pop() int {
	return l.list.RemoveFirst()
}

func (l *LinkedListStack) Peek() int {
	return l.list.GetFirst()
}

func (l *LinkedListStack) String() string {
	return l.list.String()
}
