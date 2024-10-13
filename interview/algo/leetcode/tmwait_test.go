package leetcode

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDD(t *testing.T) {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}
	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

const (
	a = iota
	b = iota
)

const (
	name = "mm"
	c    = iota
	hehe = "mm"
	d    = iota
)

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-time.After(timeout):
		return true
	case <-ch:
		return false
	}
}

// D B D B B x B D C B

func TestHehe(t *testing.T) {
	target := 345
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	size := 2
	resCh := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < len(data); i += size {
		end := i + size
		if end > len(data) {
			end = len(data)
		}

		go find(ctx, data, i, end, target, resCh)
	}

	select {
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
		cancel()
	case i := <-resCh:
		fmt.Println("found", i)
		cancel()
	}
}

func find(ctx context.Context, nums []int, start, end, target int, resCh chan int) {
	for i := start; i < end; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("task cancel")
			return
		default:
			time.Sleep(time.Second * 1)
			if nums[i] == target {
				resCh <- i
				return
			}
		}
	}
}
