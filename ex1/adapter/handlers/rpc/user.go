package rpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	userPB "go-bobo/api/user"
	user "go-bobo/app/executor/user"
)

type UserHandler struct {
	app *user.UserApplication
}

func NewUserHandler(app *user.UserApplication) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (u UserHandler) ListUser(ctx context.Context, req *userPB.ListUserReq) (*userPB.ListUserResp, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) GetUserByMobile(ctx context.Context, req *userPB.GetUserByMobileReq) (*userPB.UserItem, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) GetUserByID(ctx context.Context, req *userPB.GetUserByIDReq) (*userPB.UserItem, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) AddUser(ctx context.Context, req *userPB.AddUserReq) (*userPB.AddUserResp, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) UpdateUser(ctx context.Context, req *userPB.UpdateUserReq) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) CheckPassWord(ctx context.Context, req *userPB.CheckPassWoreReq) (*userPB.CheckPassWordResp, error) {
	// TODO implement me
	panic("implement me")
}
