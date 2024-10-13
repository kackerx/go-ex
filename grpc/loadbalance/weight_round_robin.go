package loadbalance

import (
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

type WeightRoundRobin struct {
	conns []*conn
	mu    sync.Mutex
}

func (w *WeightRoundRobin) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	if len(w.conns) == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	var totalWeight uint32
	var res *conn
	w.mu.Lock()
	defer w.mu.Unlock()
	for _, conn := range w.conns {
		totalWeight += conn.efficientWeight
		// 选择时增加每个节点的当前权重, 并且挑选最大值
		conn.currnetWeight += conn.efficientWeight
		if res == nil || res.currnetWeight < conn.currnetWeight {
			res = conn
		}
	}

	// 选择后减少选中节点的权重
	res.currnetWeight -= totalWeight
	return balancer.PickResult{
		SubConn: res.c,
		Done: func(info balancer.DoneInfo) {
			w.mu.Lock()
			defer w.mu.Unlock()
			if info.Err == nil {
				res.efficientWeight++
			} else {
				res.efficientWeight--
			}
		},
		Metadata: nil,
	}, nil
}

type WeightRoundRobinBuilder struct {
}

func (w *WeightRoundRobinBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	// TODO implement me
	panic("implement me")
}

type conn struct {
	c               balancer.SubConn
	weight          uint32
	currnetWeight   uint32
	efficientWeight uint32
}
