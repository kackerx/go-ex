package grpc

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"

	userPB "go-bobo/api/user"
	"go-bobo/grpc/loadbalance"
)

var (
	client userPB.UserClient
	ctx    = context.Background()
)

func setUp() {
	balancer.Register(base.NewBalancerBuilder("ROUND_ROBIN", &loadbalance.RoundRobinBuilder{}, base.Config{true}))
	conn, err := grpc.Dial(
		"localhost:8080,localhost:8081",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`
{"LoadBalancingPolicy": "ROUND_ROBIN"}
`),
	)
	if err != nil {
		panic(fmt.Sprintf("client: NewClient fail: %v", err))
	}

	client = userPB.NewUserClient(conn)
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestRoundBobinBalance(t *testing.T) {
	setUp()

	user, err := client.ListUser(ctx, &userPB.ListUserReq{
		Page:     0,
		PageSize: 0,
	})
	if err != nil {
		panic(fmt.Sprintf("client: ListUser faild: %v", err))
	}

	fmt.Printf("%v\n", user)
}

func TestFoo(t *testing.T) {
	s := []int{1, 2, 3}
	for i := range s {
		fmt.Println(s[i])
	}
}
