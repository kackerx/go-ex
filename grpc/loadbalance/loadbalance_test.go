package loadbalance

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"testing"
	"time"

	_ "github.com/kackerx/go-ex/grpc/loadbalance/wrr"
	"github.com/kackerx/go-ex/grpc/server"
	"github.com/kackerx/go-ex/grpc/user"
	"github.com/stretchr/testify/suite"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type EtcdTestSuite struct {
	suite.Suite
	client *etcdv3.Client
}

func (s *EtcdTestSuite) SetupSuite() {
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 3 * time.Second,
	})
	s.Require().NoError(err)
	s.client = client
}

func (s *EtcdTestSuite) TestWrrClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Clean up old test entries
	em, _ := endpoints.NewManager(s.client, "service/user")
	for _, port := range []int{9999, 9998, 9997, 9996} {
		key := fmt.Sprintf("service/user/localhost:%d", port)
		em.DeleteEndpoint(ctx, key)
	}
	time.Sleep(100 * time.Millisecond)

	// Start servers
	go s.startServer(9999)
	go s.startServer(9998)
	go s.startServer(9997)
	go s.startServer(9996)
	time.Sleep(500 * time.Millisecond)

	bd, err := resolver.NewBuilder(s.client)
	s.Require().NoError(err)

	cc, err := grpc.NewClient(
		"service/user",
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"wrr":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(bd),
		grpc.WithBlock(),
	)
	s.Require().NoError(err)
	defer cc.Close()

	// Wait for connection ready
	ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel2()
	for cc.GetState() != connectivity.Ready {
		if !cc.WaitForStateChange(ctx2, cc.GetState()) {
			break
		}
	}

	client := user.NewUserClient(cc)

	// Make requests and verify load balancing across ports
	portsSeen := make(map[string]bool)
	for i := 0; i < 12; i++ {
		resp, err := client.Register(ctx, &user.RegisterReq{Name: "heheheh"})
		s.Require().NoError(err)

		// Extract port from response message
		var port string
		fmt.Sscanf(resp.Message, "Register success: heheheh, port: %s", &port)
		portsSeen[port] = true

		slog.Info("client request", "iter", i, "port", port)
		time.Sleep(50 * time.Millisecond)
	}

	// Verify WRR is working (should see multiple ports)
	if len(portsSeen) < 2 {
		s.T().Errorf("WRR load balancer not working: only saw %d port(s): %v", len(portsSeen), portsSeen)
	}
	slog.Info("WRR test completed", "ports_seen", portsSeen)
}

func (s *EtcdTestSuite) startServer(port int) {
	em, err := endpoints.NewManager(s.client, "service/user")
	if err != nil {
		slog.Error("endpoint manager failed", "error", err, "port", port)
		return
	}

	addr := fmt.Sprintf("localhost:%d", port)
	key := "service/user/" + addr

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := em.AddEndpoint(ctx, key, endpoints.Endpoint{Addr: addr}); err != nil {
		slog.Error("add endpoint failed", "error", err, "addr", addr)
		return
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("listen failed", "error", err, "addr", addr)
		return
	}

	server1 := grpc.NewServer()
	user.RegisterUserServer(server1, &server.UserServer{Port: port})

	slog.Info("server starting", "addr", addr)
	if err := server1.Serve(l); err != nil {
		slog.Error("server error", "error", err, "addr", addr)
	}
}

func TestEtcd(t *testing.T) {
	suite.Run(t, new(EtcdTestSuite))
}
