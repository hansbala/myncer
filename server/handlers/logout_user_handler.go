package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/auth"
)

func NewLogoutUserHandler() core.Handler {
	return &logoutUserHandlerImpl{}
}

type logoutUserHandlerImpl struct{}

var _ core.Handler = (*logoutUserHandlerImpl)(nil)

func (lu *logoutUserHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil
}

func (lu *logoutUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to logout")
	}
	return nil
}

func (lu *logoutUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	auth.ClearAuthCookie(resp, core.ToMyncerCtx(ctx).Config.ServerMode)
	if _, err := resp.Write([]byte("Logged out successfully")); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write logout success message"),
		)
	}
	return core.NewProcessRequestResponse_OK()
}
