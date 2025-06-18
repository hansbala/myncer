package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewCurrentUserHandler() core.Handler {
	return &currentUserHandlerImpl{}
}

type currentUserHandlerImpl struct{}

var _ core.Handler = (*currentUserHandlerImpl)(nil)

func (c *currentUserHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil // GET request so no body.
}

func (c *currentUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required")
	}
	return nil
}

func (c *currentUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	if err := WriteJSONOk(resp, rest_helpers.UserProtoToRest(userInfo)); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write user to response"),
		)
	}
	return core.NewProcessRequestResponse_OK()
}
