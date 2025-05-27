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
) error {
	// No user permissions required to create user.
	return nil
}

func (c *createUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) error {
	restReq, ok := (reqBody).(*api.CreateUserRequest)
	if !ok {
		return core.NewError("could not cast to create user request")
	}
	if err := c.validateRequest(restReq); err != nil {
		return core.WrappedError(err, "failed to validate create user request")
	}
	protoUser := createProtoUser(restReq)
	if err := core.ToMyncerCtx(ctx).DB.UserStore.CreateUser(ctx, protoUser); err != nil {
		return core.WrappedError(err, "failed to create user")
	}
	restResp := api.NewCreateUserResponse()
	restResp.SetId(protoUser.GetId())
	return WriteJSONOk(resp, restResp)
}

func (c *createUserHandlerImpl) validateRequest(req *api.CreateUserRequest /*const*/) error {
	if len(req.GetEmail()) == 0 {
		return core.NewError("email is required")
	}
	if len(req.GetFirstName()) == 0 {
		return core.NewError("first name is required")
	}
	if len(req.GetLastName()) == 0 {
		return core.NewError("last name is required")
	}
	return nil
}

func createProtoUser(req *api.CreateUserRequest /*const*/) *myncer_pb.User {
	return &myncer_pb.User{
		Id:        uuid.New().String(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
	}
}
