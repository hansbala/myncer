package grpc_handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewLogoutUserHandler() core.GrpcHandler[
	*myncer_pb.LogoutUserRequest,
	*myncer_pb.LogoutUserResponse,
] {
	return &logoutUserHandlerImpl{}
}

type logoutUserHandlerImpl struct{}

func (l *logoutUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.LogoutUserRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to logout")
	}
	return nil
}

func (l *logoutUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.LogoutUserRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.LogoutUserResponse] {
	logoutAuthCookie := auth.GetLogoutAuthCookie(core.ToMyncerCtx(ctx).Config.ServerMode)
	return core.NewGrpcHandlerResponse_WithCookies(
		&myncer_pb.LogoutUserResponse{Id: userInfo.GetId()},
		[]*http.Cookie{logoutAuthCookie},
	)
}
