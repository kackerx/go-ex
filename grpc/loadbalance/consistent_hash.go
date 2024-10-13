package loadbalance

import (
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

type ConsistentHashBalancer struct {
	length int
	conns  []balancer.SubConn
}

func (r *ConsistentHashBalancer) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	idx := rand.Intn(r.length)
	return balancer.PickResult{
		SubConn:  r.conns[idx],
		Done:     nil,
		Metadata: nil,
	}, nil
}

type ConsistentHashBuilder struct {
}

func (r *ConsistentHashBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	// TODO implement me
	panic("implement me")
}
