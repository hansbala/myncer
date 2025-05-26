package main

import (
	"context"

	"github.com/google/uuid"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func main() {
	ctx := context.Background()
	ctx = core.WithMyncerCtx(ctx, core.MustGetMyncerCtx(ctx))
	if err := core.ToMyncerCtx(ctx).DB.UserStore.CreateUser(
		&myncer_pb.User{
			Id:        uuid.NewString(),
			FirstName: "hans",
			LastName:  "bala",
			Email:     "hansbala@hansbala.com",
		},
	); err != nil {
		panic(core.WrappedError(err, "failed to create test user"))
	}
	for {
	}
}
