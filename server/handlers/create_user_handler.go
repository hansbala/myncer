package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewCreateUserHandler() core.Handler {
	return &createUserHandlerImpl{}
}

type createUserHandlerImpl struct{}

var _ core.Handler = (*createUserHandlerImpl)(nil)

func (c *createUserHandlerImpl) GetRequestContainer(ctx context.Context) any {
	return &api.CreateUserRequest{}
}

func (c *createUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	// No user permissions required to create user.
	return nil
}

func (c *createUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse {
	restReq, ok := (reqBody).(*api.CreateUserRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("could not cast to create user request"),
		)
	}
	if err := c.validateRequest(restReq); err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "failed to validate create user request"),
		)
	}

	protoUser, err := c.createProtoUser(restReq)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to construct proto user"),
		)
	}
	if err := core.ToMyncerCtx(ctx).DB.UserStore.CreateUser(ctx, protoUser); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to create user"),
		)
	}

	restResp := api.NewCreateUserResponse()
	restResp.SetId(protoUser.GetId())
	if err := WriteJSONOk(resp, restResp); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write response"),
		)
	}

	return core.NewProcessRequestResponse_OK()
}

func (c *createUserHandlerImpl) validateRequest(req *api.CreateUserRequest /*const*/) error {
	return validateUserFields(
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}

func (c *createUserHandlerImpl) createProtoUser(
	req *api.CreateUserRequest, /*const*/
) (*myncer_pb.User, error) {
	return getProtoUser(
		uuid.New().String(),
		req.GetFirstName(),
		req.GetLastName(),
		req.GetEmail(),
		req.GetPassword(),
	)
}

