package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewCurrentUserHandler() core.GrpcHandler[
	*myncer_pb.CurrentUserRequest,
	*myncer_pb.CurrentUserResponse,
] {
	return &currentUserHandlerImpl{}
}

type currentUserHandlerImpl struct{}

func (l *currentUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CurrentUserRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user not authenticated")
	}
	return nil
}

func (l *currentUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.CurrentUserRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.CurrentUserResponse] {
	return core.NewGrpcHandlerResponse_OK(
		&myncer_pb.CurrentUserResponse{User: getPublicUser(userInfo)},
	)
}

func getPublicUser(userInfo *myncer_pb.User /*const*/) *myncer_pb.PublicUser {
	return &myncer_pb.PublicUser{
		Id:        userInfo.GetId(),
		FirstName: userInfo.GetFirstName(),
		LastName:  userInfo.GetLastName(),
		Email:     userInfo.GetEmail(),
	}
}
