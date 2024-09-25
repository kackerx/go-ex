package leetcode

import "container/list"

type LfuCache struct {
	keyToVal   map[int]int
	keyToFreq  map[int]int
	freqToKeys map[int]*list.List
	keyToNode  map[int]*list.Element
	minFreq    int
	cap        int
	size       int
}

func NewLfuCache(cap int) *LfuCache {
	return &LfuCache{
		make(map[int]int),
		make(map[int]int),
		make(map[int]*list.List),
		make(map[int]*list.Element),
		0,
		cap,
		0,
	}
}

func (l *LfuCache) Get(key int) int {
	val, ok := l.keyToVal[key]
	if !ok {
		return -1
	}

	l.increaseFreq(key)
	return val
}

func (l *LfuCache) Push(key, value int) {
	if _, ok := l.keyToVal[key]; ok {
		l.keyToVal[key] = value
		l.increaseFreq(key)
		return
	}

	// 不存在
	// 满了删除频率最小
	if l.cap <= l.size {
		l.removeMinFreqKey(key)
	}

	l.keyToVal[key] = value
	l.keyToFreq[key] = 1
	l.keyToNode[key] = l.freqToKeys[1].PushBack(key)
	l.minFreq = 1
	l.size++
}

func (l *LfuCache) removeMinFreqKey(key int) {
	minList := l.freqToKeys[l.minFreq]
	deleteNode := minList.Front()
	deleteKey := deleteNode.Value.(int)
	minList.Remove(deleteNode)
	if minList.Len() == 0 {
		delete(l.freqToKeys, l.minFreq)
	}

	delete(l.keyToVal, deleteKey)
	delete(l.keyToFreq, deleteKey)
	delete(l.keyToNode, deleteKey)
	l.size--
}

func (l *LfuCache) increaseFreq(key int) {
	// key的当前频率
	curFreq := l.keyToFreq[key]
	l.keyToFreq[key]++

	// 从当前频率队列中删除
	l.freqToKeys[curFreq].Remove(l.keyToNode[key])
	if l.freqToKeys[curFreq].Len() == 0 {
		delete(l.freqToKeys, key)
		if curFreq == l.minFreq {
			l.minFreq++
		}
	}

	// 去+1频率队列中新增
	if _, ok := l.freqToKeys[curFreq+1]; !ok {
		l.freqToKeys[curFreq+1] = list.New()
	}
	l.keyToNode[key] = l.freqToKeys[curFreq+1].PushBack(key)
}
