package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewEditUserHandler() core.Handler {
	return &editUserHandlerImpl{}
}

type editUserHandlerImpl struct{}

var _ core.Handler = (*editUserHandlerImpl)(nil)

func (eh *editUserHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return &api.EditUserRequest{}
}

func (eh *editUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required")
	}
	return nil
}

func (eh *editUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	restRequest, ok := reqBody.(api.EditUserRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("expected edit user request"),
		)
	}

	if err := eh.validateRequest(&restRequest); err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("failed to validate edit user request"),
		)
	}

	user, err := eh.getUpdatedUser(userInfo, &restRequest)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to get updated user"),
		)
	}

	if err := core.ToMyncerCtx(ctx).DB.UserStore.EditUser(ctx, user); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.NewError("failed to edit user"),
		)
	}

	if err := WriteJSONOk(resp, rest_helpers.UserProtoToRest(user)); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write user to response"),
		)
	}

	return core.NewProcessRequestResponse_OK()
}

func (eh *editUserHandlerImpl) validateRequest(req *api.EditUserRequest /*const*/) error {
	return validateUserFields(
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}

func (eh *editUserHandlerImpl) getUpdatedUser(
	user *myncer_pb.User, /*const*/
	restReq *api.EditUserRequest, /*const*/
) (*myncer_pb.User, error) {
	return getProtoUser(
		user.GetId(),
		restReq.GetFirstName(),
		restReq.GetLastName(),
		restReq.GetEmail(),
		restReq.GetPassword(),
	)
}
