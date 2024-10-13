package leetcode

import "context"

type Task func()

type TaskPool struct {
	tasks chan Task
	close chan struct{}
}

func (t *TaskPool) Go(ctx context.Context, task Task) error {
	select {
	case t.tasks <- task:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *TaskPool) Close() {
	close(t.close)
}

func NewTaskPool(cap, workCount int) *TaskPool {
	res := &TaskPool{tasks: make(chan Task, cap), close: make(chan struct{})}

	for i := 0; i < workCount; i++ {
		go func() {
			for {
				select {
				case fn := <-res.tasks:
					defer func() {
						if r := recover(); r != nil {
							return
						}
					}()
					fn()
				case <-res.close:
					return
				}
			}
		}()
	}

	return res
}
