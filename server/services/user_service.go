package services

import (
	"context"

	"connectrpc.com/connect"

	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/grpc_handlers"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	myncer_pb_connect "github.com/hansbala/myncer/proto/myncer/myncer_pbconnect"
)

func NewUserService() *UserService {
	return &UserService{
		createUserHandler: grpc_handlers.NewCreateUserHandler(),
	}
}

type UserService struct{
	createUserHandler core.GrpcHandler[*myncer_pb.CreateUserRequest, *myncer_pb.CreateUserResponse]
}

var _ myncer_pb_connect.UserServiceHandler = (*UserService)(nil)

func (u *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[myncer_pb.CreateUserRequest], /*const*/
) (*connect.Response[myncer_pb.CreateUserResponse], error) {
	return OrchestrateHandler(ctx, u.createUserHandler, req.Msg)
}

func (u *UserService) LoginUser(
	ctx context.Context, 
	req *connect.Request[myncer_pb.LoginUserRequest], /*const*/
) (*connect.Response[myncer_pb.LoginUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}

func (u *UserService) LogoutUser(
	ctx context.Context, 
	req *connect.Request[myncer_pb.LogoutUserRequest], /*const*/
) (*connect.Response[myncer_pb.LogoutUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}

func (u *UserService) EditUser(
	ctx context.Context,
	req *connect.Request[myncer_pb.EditUserRequest], /*const*/
) (*connect.Response[myncer_pb.EditUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}

func (u *UserService) GetCurrentUser(
	ctx context.Context, 
	req *connect.Request[myncer_pb.CurrentUserRequest], /*const*/
) (*connect.Response[myncer_pb.CurrentUserResponse], error) {
	return nil, core.NewError("not implemented yet")
}
