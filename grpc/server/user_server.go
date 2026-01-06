package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kackerx/go-ex/grpc/user"
)

type UserServer struct {
	Port int
	user.UnimplementedUserServer
}

func (u *UserServer) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	slog.Info("from server: ", "port", u.Port)
	return &user.RegisterResp{
		Message: fmt.Sprintf("Register success: %s, port: %d", req.Name, u.Port),
	}, nil
}
