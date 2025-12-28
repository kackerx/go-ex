package main

import (
	"context"
	"fmt"

	"github.com/kackerx/go-ex/grpc/registry/user"
)

type UserServer struct {
	user.UnimplementedUserServer
}

func (u *UserServer) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	return &user.RegisterResp{
		Message: fmt.Sprintf("Register success: %s", req.Name),
	}, nil
}
