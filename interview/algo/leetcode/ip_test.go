package leetcode

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	mu       sync.Mutex
	close    chan struct{}
}

func (b *Ban) Close() {
	close(b.close)
}

func NewBan() *Ban {
	res := &Ban{visitIPs: make(map[string]time.Time), close: make(chan struct{})}

	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case now := <-ticker.C:
				res.mu.Lock()
				for k, v := range res.visitIPs {
					if now.Sub(v).Minutes() >= 3 {
						delete(res.visitIPs, k)
					}
				}
				res.mu.Unlock()
			case <-res.close:
				ticker.Stop()
				return
			}

		}
	}()

	return res
}

func (o *Ban) visit(ip string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func TestName(t *testing.T) {
	var success int32
	ban := NewBan()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt32(&success, 1)
				}
			}(j)
		}

	}

	wg.Wait()
	fmt.Println("success:", success)
}
