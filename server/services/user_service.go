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
		createUserHandler:  grpc_handlers.NewCreateUserHandler(),
		loginUserHandler:   grpc_handlers.NewLoginUserHandler(),
		currentUserHandler: grpc_handlers.NewCurrentUserHandler(),
		logoutUserHandler:  grpc_handlers.NewLogoutUserHandler(),
		editUserHandler:    grpc_handlers.NewEditUserHandler(),
	}
}

type UserService struct {
	createUserHandler  core.GrpcHandler[*myncer_pb.CreateUserRequest, *myncer_pb.CreateUserResponse]
	loginUserHandler   core.GrpcHandler[*myncer_pb.LoginUserRequest, *myncer_pb.LoginUserResponse]
	currentUserHandler core.GrpcHandler[*myncer_pb.CurrentUserRequest, *myncer_pb.CurrentUserResponse]
	logoutUserHandler  core.GrpcHandler[*myncer_pb.LogoutUserRequest, *myncer_pb.LogoutUserResponse]
	editUserHandler    core.GrpcHandler[*myncer_pb.EditUserRequest, *myncer_pb.EditUserResponse]
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
	return OrchestrateHandler(ctx, u.loginUserHandler, req.Msg)
}

func (u *UserService) LogoutUser(
	ctx context.Context,
	req *connect.Request[myncer_pb.LogoutUserRequest], /*const*/
) (*connect.Response[myncer_pb.LogoutUserResponse], error) {
	return OrchestrateHandler(ctx, u.logoutUserHandler, req.Msg)
}

func (u *UserService) EditUser(
	ctx context.Context,
	req *connect.Request[myncer_pb.EditUserRequest], /*const*/
) (*connect.Response[myncer_pb.EditUserResponse], error) {
	return OrchestrateHandler(ctx, u.editUserHandler, req.Msg)
}

func (u *UserService) GetCurrentUser(
	ctx context.Context,
	req *connect.Request[myncer_pb.CurrentUserRequest], /*const*/
) (*connect.Response[myncer_pb.CurrentUserResponse], error) {
	return OrchestrateHandler(ctx, u.currentUserHandler, req.Msg)
}
