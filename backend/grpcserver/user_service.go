package grpcserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/services"
)

func NewUserServiceServer(svc services.UserService) myncer_pb.UserServiceServer {
	return &userServiceServer{userService: svc}
}

type userServiceServer struct {
	myncer_pb.UnimplementedUserServiceServer
	userService services.UserService
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *myncer_pb.CreateUserRequest) (*myncer_pb.CreateUserResponse, error) {
	resp, err := s.userService.CreateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateUser failed: %v", err)
	}
	return resp, nil
}
