package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

const (
	cJwtCookieName = "jwt"
	cJwtSub        = "sub"
)

func GenerateJWTToken(jwtSecret string, userID string) (string, error) {
	claims := jwt.MapClaims{
		cJwtSub: userID,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
}

func MaybeGetUserFromRequest(
	ctx context.Context,
	myncerCtx *core.MyncerCtx, /*const*/
	r *http.Request, /*const*/
) (*myncer_pb.User /*@nullable*/, error) {
	userId, err := extractUserIdFromJWTCookie(myncerCtx.Config.JwtSecret, r)
	if err != nil {
		core.Printf(r.URL.Path)
		return nil, core.WrappedError(err, "failed to extract user id from jwt cookie")
	}
	user, err := myncerCtx.DB.UserStore.GetUserById(ctx, userId)
	if err != nil {
		return nil, core.WrappedError(err, fmt.Sprintf("failed to get user from id %s in jwt", userId))
	}
	return user, nil
}

func extractUserIdFromJWTCookie(
	jwtSecret string,
	r *http.Request, /*const*/
) (string /*userId*/, error) {
	cookie, err := r.Cookie(cJwtCookieName)
	if err != nil {
		return "", core.WrappedError(err, "failed to get jwt auth cookie")
	}

	token, err := jwt.Parse(
		cookie.Value,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, core.NewError("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return "", core.WrappedError(err, "failed to parse jwt token")
	}
	if !token.Valid {
		return "", core.NewError("jwt token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", core.NewError("invalid claims")
	}
	if claims[cJwtSub] == nil {
		return "", core.NewError("sub not found in jwt claims")
	}
	userId, ok := claims[cJwtSub].(string)
	if !ok {
		return "", core.NewError("expected sub in jwt token to be of string type")
	}

	return userId, nil
}

func isHttpOnly(mode myncer_pb.ServerMode) bool {
	switch mode {
	case myncer_pb.ServerMode_DEV:
		return false
	default:
		return true
	}
}
