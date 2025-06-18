package services

import (
	"context"

	"connectrpc.com/connect"

	// "github.com/hansbala/myncer/core"
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
	// currUser := core.ToMyncerCtx(ctx).RequestUser
	// core.DebugPrintJson(currUser)
	res := connect.NewResponse(&myncer_pb.CreateUserResponse{
		Greeting: "Hello from Hans!",
	})
	return res, nil
}
