package rpc_handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewCreateUserHandler() core.GrpcHandler[
	*myncer_pb.CreateUserRequest,
	*myncer_pb.CreateUserResponse,
] {
	return &createUserImpl{}
}

type createUserImpl struct{}

func (c *createUserImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CreateUserRequest, /*const*/
) error {
	// No user permissions required to create user.
	return nil
}

func (c *createUserImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CreateUserRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.CreateUserResponse] {
	if err := c.validateRequest(reqBody); err != nil {
		return core.NewGrpcHandlerResponse_BadRequest[*myncer_pb.CreateUserResponse](
			core.WrappedError(err, "failed to validate create user request"),
		)
	}
	protoUser, err := c.createProtoUser(reqBody)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.CreateUserResponse](
			core.WrappedError(err, "failed to construct proto user"),
		)
	}
	if err := core.ToMyncerCtx(ctx).DB.UserStore.CreateUser(ctx, protoUser); err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.CreateUserResponse](
			core.WrappedError(err, "failed to create user"),
		)
	}
	return core.NewGrpcHandlerResponse_OK(&myncer_pb.CreateUserResponse{Id: protoUser.GetId()})
}

func (c *createUserImpl) validateRequest(req *myncer_pb.CreateUserRequest /*const*/) error {
	return ValidateUserFields(
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}

func (c *createUserImpl) createProtoUser(
	req *myncer_pb.CreateUserRequest, /*const*/
) (*myncer_pb.User, error) {
	return GetProtoUser(
		uuid.New().String(),
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}
