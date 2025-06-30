package rpc_handlers

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewLoginUserHandler() core.GrpcHandler[
	*myncer_pb.LoginUserRequest,
	*myncer_pb.LoginUserResponse,
] {
	return &loginHandlerImpl{}
}

type loginHandlerImpl struct{}

func (l *loginHandlerImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.LoginUserRequest, /*const*/
) error {
	// Logging in requires no user permissions.
	return nil
}

func (l *loginHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.LoginUserRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.LoginUserResponse] {
	myncerCtx := core.ToMyncerCtx(ctx)
	user, err := myncerCtx.DB.UserStore.GetUserByEmail(ctx, reqBody.GetEmail())
	if err != nil {
		// Not found is unathorized to prevent leaking if user exists.
		if errors.Is(err, core.CUserNotFoundError) {
			return core.NewGrpcHandlerResponse_Unauthorized[*myncer_pb.LoginUserResponse](err)
		}
		// Fallback to Internal server error.
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.LoginUserResponse](
			core.WrappedError(err, "failed to get user by email"),
		)
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.GetHashedPassword()),
		[]byte(reqBody.GetPassword()),
	); err != nil {
		return core.NewGrpcHandlerResponse_Unauthorized[*myncer_pb.LoginUserResponse](
			core.WrappedError(err, "invalid password"),
		)
	}

	// Sets the JWT auth cookie.
	jwtToken, err := auth.GenerateJWTToken(myncerCtx.Config.JwtSecret, user.GetId())
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.LoginUserResponse](
			core.WrappedError(err, "failed to generate jwt token"),
		)
	}

	return core.NewGrpcHandlerResponse_WithCookies(
		&myncer_pb.LoginUserResponse{Id: user.GetId()},
		[]*http.Cookie{auth.GetAuthCookie(jwtToken, myncerCtx.Config.ServerMode)},
	)
}
