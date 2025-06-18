package services

import (
	"context"

	"connectrpc.com/connect"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	myncer_pb_connect "github.com/hansbala/myncer/proto/myncer/myncer_pbconnect"
)

func NewUserService() *UserService {
	return &UserService{}
}

type UserService struct{}

var _ myncer_pb_connect.UserServiceHandler = (*UserService)(nil)

func (s *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[myncer_pb.CreateUserRequest], /*const*/
) (*connect.Response[myncer_pb.CreateUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}


func (s *UserService) LoginUser(
	context.Context, 
	*connect.Request[myncer_pb.LoginUserRequest],
) (*connect.Response[myncer_pb.LoginUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}

func (s *UserService) LogoutUser(
	context.Context, 
	*connect.Request[myncer_pb.LogoutUserRequest],
) (*connect.Response[myncer_pb.LogoutUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}

func (s *UserService) EditUser(
	context.Context,
	*connect.Request[myncer_pb.EditUserRequest],
) (*connect.Response[myncer_pb.EditUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}

func (s *UserService) GetCurrentUser(
	context.Context, 
	*connect.Request[myncer_pb.CurrentUserRequest],
) (*connect.Response[myncer_pb.CurrentUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}
