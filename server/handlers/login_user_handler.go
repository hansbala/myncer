package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"golang.org/x/crypto/bcrypt"
)

func NewLoginUserHandler() core.Handler {
	return &loginUserHandlerImpl{}
}

type loginUserHandlerImpl struct{}

var _ core.Handler = (*loginUserHandlerImpl)(nil)

func (l *loginUserHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return &api.UserLoginRequest{}
}

func (l *loginUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
) error {
	// Logging in requires no user permissions.
	return nil
}

func (l *loginUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	restReq, ok := (reqBody).(*api.UserLoginRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("failed to get user login request"),
		)
	}

	myncerCtx := core.ToMyncerCtx(ctx)
	user, err := myncerCtx.DB.UserStore.GetUserByEmail(ctx, restReq.GetEmail())
	if err != nil {
		// Not found is unathorized to prevent leaking if user exists.
		if errors.Is(err, core.CUserNotFoundError) {
			return core.NewProcessRequestResponse_Unauthorized(err)
		}
		// Fallback to Internal server error.
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to get user by email"),
		)
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.GetHashedPassword()), 
		[]byte(restReq.GetPassword()),
	); err != nil {
		return core.NewProcessRequestResponse_Unauthorized(
			core.WrappedError(err, "invalid password"),
		)
	}

	// Sets the JWT auth cookie.
	jwtToken, err := auth.GenerateJWTToken(myncerCtx.Config.JwtSecret, user.GetId())
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to generate jwt token"),
		)
	}
	auth.SetAuthCookie(resp, jwtToken)
	
	if _, err := resp.Write([]byte("Success. Set auth cookie")); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write success message"),
		)
	}
	return core.NewProcessRequestResponse_OK()
}
