package datastruct

import (
	"container/list"
)

type Node struct {
	key, val int
}

type LRUCache struct {
	data      *list.List
	m         map[int]*list.Element
	capaction int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		data:      list.New(),
		m:         make(map[int]*list.Element, capacity),
		capaction: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if !ok {
		return -1
	}

	this.data.MoveToFront(e)
	return e.Value.(*Node).val
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	if ok {
		e.Value.(*Node).val = value
		this.data.MoveToFront(e)
		return
	}

	if this.capaction == this.data.Len() {
		back := this.data.Back()
		this.data.Remove(back)
		delete(this.m, back.Value.(*Node).key)
	}

	node := &Node{key, value}
	e = this.data.PushFront(node)
	this.m[key] = e
}
