package main

import (
	"context"
	"fmt"
)

type Task func()

type TaskPool struct {
	tasks   chan Task
	cap     int
	closeCh chan struct{}
}

func NewTaskPool(workerCnt, cap int) *TaskPool {
	pool := &TaskPool{tasks: make(chan Task, cap), closeCh: make(chan struct{})}

	for i := 0; i < workerCnt; i++ {
		go func() {
			for {
				select {
				case <-pool.closeCh:
					fmt.Println("closed")
					return
				case f, ok := <-pool.tasks:
					if !ok {
						return
					}
					f()
				}
			}
		}()
	}

	return pool
}

func (p *TaskPool) Go(ctx context.Context, t Task) (err error) {
	select {
	case p.tasks <- t:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (p *TaskPool) Close() {
	close(p.tasks)
	close(p.closeCh)
}
