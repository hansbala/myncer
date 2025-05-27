package main

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/handlers"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func main() {
	ctx := context.Background()
	myncerCtx := core.MustGetMyncerCtx(ctx)

	http.HandleFunc(
		"/users/create",
		ServerHandler(
			handlers.NewCreateUserHandler(),
			myncerCtx,
		),
	)
	core.Printf("Myncer listening on port 8080")
	if err := http.ListenAndServe(":8080", nil /*handler*/); err != nil {
		core.Errorf("failed: ", err)
	}
}

func ServerHandler(h core.Handler, myncerCtx *core.MyncerCtx /*const*/) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create custom ctx passing myncer ctx down with it.
		ctx := core.WithMyncerCtx(r.Context(), myncerCtx)
		// TODO: Get the user with JWT auth here.
		user := &myncer_pb.User{Id: "some-id", FirstName: "devuser"}

		if err := h.CheckUserPermissions(ctx, user); err != nil {
			core.Errorf(core.WrappedError(err, "check user perms failed"))
		}

		if err := h.ProcessRequest(ctx, r, w); err != nil {
			core.Errorf(core.WrappedError(err, "process request failed"))
		}
	}
}
