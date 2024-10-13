package leetcode

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type sp interface {
	Out(key string, val any)                  // 写入kv, 读该key的goroutine挂起的话则唤醒, 此方法不会阻塞立即返回
	Rd(key string, timeout time.Duration) any // 读取key, key不存在阻塞, 直到存在或者超时
}

type Map struct {
	data map[string]*entry
	mu   sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	val     any
	isExist bool
}

func (m *Map) Out(key string, val any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	e, ok := m.data[key]
	if !ok {
		m.data[key] = &entry{
			ch:      nil,
			val:     val,
			isExist: true,
		}
		return
	}

	e.val = val
	if !e.isExist {
		if e.ch != nil {
			close(e.ch)
			e.ch = nil
		}
		e.isExist = true
	}
}

func (m *Map) Rd(key string, timeout time.Duration) any {
	m.mu.RLock()
	e, ok := m.data[key]
	if ok && e.isExist {
		m.mu.RUnlock()
		return e.val
	}

	m.mu.RUnlock()
	m.mu.Lock() // double-check
	e, ok = m.data[key]
	if !ok {
		e = &entry{
			ch:      make(chan struct{}),
			val:     nil,
			isExist: false,
		}
		m.data[key] = e
	}
	m.mu.Unlock()
	if ok && e.isExist {
		return e.val
	}

	select {
	case <-e.ch:
		return e.val
	case <-time.After(timeout):
		fmt.Println("timeout")
		return nil
	}
}

func TestMap(t *testing.T) {
	m := &Map{
		data: make(map[string]*entry),
	}

	m.Out("a", 1)
	m.Out("b", 2)

	go func() {
		time.Sleep(time.Second * 8)
		m.Out("c", 3)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(m.Rd("c", time.Second*5))
		}()
	}

	select {}
}
