package wrr

import (
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

const name = "wrr"

// newBuilder creates a new wrr balancer builder.
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(name, &wrrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type wrrPickerBuilder struct{}

func (*wrrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	var conns []balancer.SubConn
	for sc := range info.ReadySCs {
		conns = append(conns, sc)
	}

	return &wrrPicker{
		subConns: conns,
	}
}

type wrrPicker struct {
	subConns []balancer.SubConn
	next     atomic.Uint64
	mu       sync.Mutex
}

func (p *wrrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.subConns) == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	nextIndex := p.next.Add(1)
	sc := p.subConns[int(nextIndex%uint64(len(p.subConns)))]
	return balancer.PickResult{SubConn: sc}, nil
}