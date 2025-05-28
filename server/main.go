package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/handlers"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func main() {
	ctx := context.Background()
	myncerCtx := core.MustGetMyncerCtx(ctx)

	http.HandleFunc(
		"/api/v1/users/create",
		ServerHandler(
			handlers.NewCreateUserHandler(),
			myncerCtx,
		),
	)
	http.HandleFunc(
		"/api/v1/users/login",
		ServerHandler(
			handlers.NewLoginUserHandler(),
			myncerCtx,
		),
	)
	http.HandleFunc(
		"/api/v1/users/list",
		ServerHandler(
			handlers.NewListUsersHandler(),
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

		reqContainer := h.GetRequestContainer(ctx)
		// Unmarshal request body here for usage in process request.
		if reqContainer != nil {
			if err := json.NewDecoder(r.Body).Decode(reqContainer); err != nil {
				core.Errorf(core.WrappedError(err, "failed to decode request body into container"))
				return
			}
		}

		prr := h.ProcessRequest(ctx, reqContainer, r, w)
		// Http status code writing.
		// Order matters: status code should be written before any response writer writes.
		// StatusOK is written by default if we're writing JSON to response writer.
		if prr.StatusCode != http.StatusOK && prr.StatusCode != 0 {
			w.WriteHeader(prr.StatusCode)
		}
		// Handler error logging.
		if prr.Err != nil {
			core.Errorf(prr.Err)
		}
		// Http message writing.
		if len(prr.MsgForHttp) > 0 {
			// Write the message for http alongside status code.
			if _, err := w.Write([]byte(prr.MsgForHttp)); err != nil {
				core.Errorf("failed to write failure message to writer")
			}
		}
	}
}
