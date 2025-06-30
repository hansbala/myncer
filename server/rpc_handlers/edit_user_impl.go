package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewEditUserHandler() core.GrpcHandler[
	*myncer_pb.EditUserRequest,
	*myncer_pb.EditUserResponse,
] {
	return &editUserHandlerImpl{}
}

type editUserHandlerImpl struct{}

func (eu *editUserHandlerImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.EditUserRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to edit user details")
	}
	return nil
}

func (eu *editUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.EditUserRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.EditUserResponse] {
	if err := eu.validateRequest(reqBody); err != nil {
		return core.NewGrpcHandlerResponse_BadRequest[*myncer_pb.EditUserResponse](
			core.NewError("failed to validate edit user request"),
		)
	}

	user, err := eu.getUpdatedUser(userInfo, reqBody)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.EditUserResponse](
			core.WrappedError(err, "failed to get updated user"),
		)
	}

	if err := core.ToMyncerCtx(ctx).DB.UserStore.EditUser(ctx, user); err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.EditUserResponse](
			core.WrappedError(err, "failed to edit user in database"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(&myncer_pb.EditUserResponse{User: getPublicUser(user)})
}

func (eu *editUserHandlerImpl) validateRequest(req *myncer_pb.EditUserRequest /*const*/) error {
	return ValidateUserFields(
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}

func (eu *editUserHandlerImpl) getUpdatedUser(
	userInfo *myncer_pb.User, /*const*/
	req *myncer_pb.EditUserRequest, /*const*/
) (*myncer_pb.User, error) {
	return GetProtoUser(
		userInfo.GetId(),
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}
