package linked_list

import "fmt"

type Node struct {
	data int
	next *Node
}

func NewNode(data int, next *Node) *Node {
	return &Node{data: data, next: next}
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList { return &LinkedList{head: NewNode(0, nil)} }

func (l *LinkedList) Len() int { return l.length }

func (l *LinkedList) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList) Add(index, e int) {
	prev := l.head

	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = NewNode(e, prev.next)
	l.length++
}

func (l *LinkedList) AddFirst(e int) { l.Add(0, e) }

func (l *LinkedList) AddLast(e int) { l.Add(l.length, e) }

func (l *LinkedList) Get(index int) int {
	cur := l.head.next

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.data
}

func (l *LinkedList) GetFirst() int {
	return l.Get(0)
}

func (l *LinkedList) GetLast() int {
	return l.Get(l.length - 1)
}

func (l *LinkedList) Set(index, e int) {
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.data = e
}

func (l *LinkedList) Remove(index int) int {
	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	res := prev.next
	prev.next = prev.next.next

	l.length--

	return res.data
}

func (l *LinkedList) RemoveFirst() int {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() int {
	return l.Remove(l.length - 1)
}

func (l *LinkedList) String() string {
	var res string

	cur := l.head.next
	for cur != nil {
		res += fmt.Sprintf("%d", cur.data)
		if cur.next != nil {
			res += " --> "
		}
		cur = cur.next
	}

	return res
}
