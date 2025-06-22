package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewTemplateHandler() core.GrpcHandler[
	*myncer_pb.CreateUserRequest,
	*myncer_pb.CreateUserResponse,
] {
	return &templateHandlerImpl{}
}

type templateHandlerImpl struct{}

func (l *templateHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CreateUserRequest, /*const*/
) error {
	return nil
}

func (l *templateHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CreateUserRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.CreateUserResponse] {
	return nil
}
