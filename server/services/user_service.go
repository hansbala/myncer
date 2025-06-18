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
	req *connect.Request[proto.CreateUserRequest],
) (*connect.Response[proto.User], error) {
	core.Printf("Request headers: ", req.Header())
	res := connect.NewResponse(&myncer_pb.User{
		FirstName: "Hello from Hans!",
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

