package main

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/kackerx/go-ex/grpc/server"
	"github.com/kackerx/go-ex/grpc/user"
	"github.com/stretchr/testify/suite"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
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

func (s *EtcdTestSuite) TestClient() {
	bd, err := resolver.NewBuilder(s.client)
	s.Require().NoError(err)

	cc, err := grpc.NewClient("etcd:///service/user", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithResolvers(bd))
	s.Require().NoError(err)

	client := user.NewUserClient(cc)
	resp, err := client.Register(context.Background(), &user.RegisterReq{
		Name:    "heheheh",
		Address: "localhost:9999",
	})
	s.Require().NoError(err)

	fmt.Println(resp.Message)
}

func (s *EtcdTestSuite) TestServer() {
	em, err := endpoints.NewManager(s.client, "service/user")
	s.Require().NoError(err)

	addr := "localhost:9999"
	key := "service/user/" + addr // 本机实例key
	err = em.AddEndpoint(context.Background(), key, endpoints.Endpoint{
		Addr: addr,
	})
	s.Require().NoError(err)

	l, err := net.Listen("tcp", ":9999")
	s.Require().NoError(err)
	defer l.Close()

	server1 := grpc.NewServer()
	user.RegisterUserServer(server1, &server.UserServer{})
	err = server1.Serve(l)

	s.T().Log(err)
}

func TestEtcd(t *testing.T) {
	suite.Run(t, new(EtcdTestSuite))
}
