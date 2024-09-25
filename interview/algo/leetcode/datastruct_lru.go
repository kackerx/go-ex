package leetcode

type Node struct {
	key, val   int
	next, prev *Node
}

func NewNode(key int, val int, next, prev *Node) *Node {
	return &Node{key: key, val: val, next: next, prev: prev}
}

type doubleList struct {
	head, tail *Node
	size       int
}

func NewDoubleList() *doubleList {
	head := NewNode(0, 0, nil, nil)
	tail := NewNode(0, 0, nil, nil)
	head.next = tail
	tail.prev = head
	return &doubleList{head: head, tail: tail}
}

func (d *doubleList) Remove(no *Node) {
	no.prev.next = no.next
	no.next.prev = no.prev
	d.size--
}

func (d *doubleList) RemoveLast() *Node {
	if d.size == 0 {
		return nil
	}

	delNode := d.tail.prev
	d.Remove(d.tail.prev)
	d.size--

	return delNode
}

func (d *doubleList) AddFirst(no *Node) {
	no.next = d.head.next
	no.prev = d.head
	d.head.next.prev = no
	d.head.next = no
	d.size++
}

type LRUCache struct {
	cache   *doubleList
	cap     int
	nodeMap map[int]*Node
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{cap: capacity, cache: NewDoubleList(), nodeMap: make(map[int]*Node)}
}

func (l *LRUCache) Get(key int) int {
	if no, ok := l.nodeMap[key]; !ok {
		l.cache.Remove(no)
		l.cache.AddFirst(no)
		return no.val
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	if no, ok := l.nodeMap[key]; ok {
		no.val = value
		l.cache.Remove(no)
		l.cache.AddFirst(no)
		return
	}

	if l.cap == l.cache.size {
		last := l.cache.RemoveLast()
		delete(l.nodeMap, last.key)
	}
	no := NewNode(key, value, nil, nil)
	l.cache.AddFirst(no)
	l.nodeMap[key] = no
}
