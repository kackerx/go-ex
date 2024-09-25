package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()

	l.Add(0, 0)
	l.Add(1, 1)
	l.Add(2, 2)
	fmt.Println(l)

	l.AddFirst(5)
	fmt.Println(l)
	fmt.Println(l.Len())

	fmt.Println(l.Get(2), l.Get(3))

	l.Set(2, 100)
	fmt.Println(l)

	fmt.Println("====")

	e := l.Remove(2)

	fmt.Println(l, e)
}

func TestName(t *testing.T) {
	reverse(&Node{data: 0, next: NewNode(1, NewNode(2, NewNode(3, nil)))})
}

func traverse(nums []int, index int) {
	if index == len(nums) {
		return
	}

	traverse(nums, index+1)
	fmt.Println(nums[index])
}

func reverse(head *Node) {
	if head == nil {
		return
	}

	fmt.Println(head.data)
	reverse(head.next)
}

func reverseList(no *Node) *Node {
	if no.next == nil {
		return no
	}

	newHead := reverseList(no.next)
	no.next.next = no
	no.next = nil

	return newHead
}

func TestReverseList(t *testing.T) {
	no := reverseList(&Node{data: 0, next: NewNode(1, NewNode(2, NewNode(3, nil)))})
	reverse(no)
}
