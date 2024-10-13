package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	userPB "go-bobo/api/user"
)

type UserServer struct {
}

func (s *UserServer) ListUser(ctx context.Context, req *userPB.ListUserReq) (*userPB.ListUserResp, error) {
	return &userPB.ListUserResp{
		Total: 1001,
		Data: []*userPB.UserItem{
			{
				Id:       0,
				Mobile:   "18838974677",
				Password: "",
				Nickname: "kacker",
				Gender:   0,
				Role:     0,
			},
		},
	}, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *userPB.GetUserByMobileReq) (*userPB.UserItem, error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserServer) GetUserByID(ctx context.Context, req *userPB.GetUserByIDReq) (*userPB.UserItem, error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserServer) AddUser(ctx context.Context, req *userPB.AddUserReq) (*userPB.AddUserResp, error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserServer) UpdateUser(ctx context.Context, req *userPB.UpdateUserReq) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

func (s *UserServer) CheckPassWord(ctx context.Context, req *userPB.CheckPassWoreReq) (*userPB.CheckPassWordResp, error) {
	// TODO implement me
	panic("implement me")
}
