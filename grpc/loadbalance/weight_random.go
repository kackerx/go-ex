package loadbalance

import (
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

type WeightRandomBalancer struct {
	length      int
	totalWeight uint32
	conns       []*WeightConn
}

type WeightConn struct {
	c      balancer.SubConn
	weight uint32
}

func (r *WeightRandomBalancer) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	target := rand.Intn(int(r.totalWeight))
	for _, conn := range r.conns {
		target -= int(conn.weight)
		if target < 0 {
			return balancer.PickResult{
				SubConn:  conn.c,
				Done:     nil,
				Metadata: nil,
			}, nil
		}
	}

	return balancer.PickResult{}, nil
}

type WeightRandomBuilder struct {
}

func (r *WeightRandomBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	var (
		total uint32
		conns []*WeightConn
	)

	for sub, subInfo := range info.ReadySCs {
		weight := subInfo.Address.Attributes.Value("weight").(uint32)
		total += weight
		conns = append(conns, &WeightConn{
			c:      sub,
			weight: weight,
		})
	}

	return &WeightRandomBalancer{
		length:      len(conns),
		totalWeight: total,
		conns:       conns,
	}
}
