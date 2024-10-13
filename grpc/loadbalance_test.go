package grpc

import (
	"net"
	"testing"

	"google.golang.org/grpc"

	userPB "go-bobo/api/user"
)

func TestServer(t *testing.T) {
	srv := grpc.NewServer()

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	userPB.RegisterUserServer(srv, &UserServer{})

	t.Log("8080 running")
	if err = srv.Serve(listen); err != nil {
		panic(err)
	}
}
