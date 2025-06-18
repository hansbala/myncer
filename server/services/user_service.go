package services

import (
	"context"

	"connectrpc.com/connect"

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
	res := connect.NewResponse(&myncer_pb.CreateUserResponse{
		Greeting: "Hello from Hans!",
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
