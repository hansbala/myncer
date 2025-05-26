package services

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

type UserService interface {
	CreateUser(
		ctx context.Context, 
		req *myncer_pb.CreateUserRequest, /*const*/
	) (*myncer_pb.CreateUserResponse, error)
}

func NewUserService(database *core.Database) UserService {
	return &userServiceImpl{
		db: database,
	}
}

type userServiceImpl struct {
	db *core.Database /*const*/
}

var _ UserService = (*userServiceImpl)(nil)
	
func (us *userServiceImpl) CreateUser(
	ctx context.Context,
	req *myncer_pb.CreateUserRequest, /*const*/
) (*myncer_pb.CreateUserResponse, error) {
	return &myncer_pb.CreateUserResponse{GeneratedId: "some-random-gid"}, nil
}
