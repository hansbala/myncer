package auth

import (
	"context"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type userContextKey struct{}

func ContextWithUser(ctx context.Context, user *myncer_pb.User /*const*/) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

func UserFromContext(ctx context.Context) *myncer_pb.User /*@nullable*/ {
	user, ok := ctx.Value(userContextKey{}).(*myncer_pb.User)
	if ok {
		return user
	}
	return nil
}
