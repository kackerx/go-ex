package loadbalance

import (
	"fmt"
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

var nodeList = []string{
	"a",
	"b",
	"c",
	"d",
	"e",
}

type RoundRobinBalancer struct {
	conns []balancer.SubConn
	index int32
}

func (r *RoundRobinBalancer) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	idx := atomic.AddInt32(&r.index, 1)
	conn := r.conns[int(idx)%len(r.conns)]
	fmt.Printf("Get index: %d\n", idx)

	return balancer.PickResult{
		SubConn: conn,
		Done: func(info balancer.DoneInfo) {
			fmt.Printf("Get index: %d\n", idx)
		},
		Metadata: nil,
	}, nil
}

type RoundRobinBuilder struct {
}

func (r *RoundRobinBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	var conns []balancer.SubConn
	for conn := range info.ReadySCs {
		conns = append(conns, conn)
	}

	return &RoundRobinBalancer{
		conns: conns,
		index: -1,
	}
}
