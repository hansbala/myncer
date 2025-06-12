package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/handlers"
)

var (
	cHandlersMap = map[string]core.Handler{
		// User handlers.
		"/api/v1/users/create": handlers.NewCreateUserHandler(),
		"/api/v1/users/login":  handlers.NewLoginUserHandler(),
		"/api/v1/users/me":     handlers.NewCurrentUserHandler(),
		"/api/v1/users/edit":   handlers.NewEditUserHandler(),
	}
)

func main() {
	ctx := context.Background()
	myncerCtx := core.MustGetMyncerCtx(ctx)

	for pattern, handler := range cHandlersMap {
		http.Handle(pattern, WithCors(ServerHandler(handler, myncerCtx)))
	}
	core.Printf("Myncer listening on port 8080")
	if err := http.ListenAndServe(":8080", nil /*handler*/); err != nil {
		core.Errorf("failed: ", err)
	}
}

func WithCors(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Allow frontend origin
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Credentials", "true") // for cookies
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			h.ServeHTTP(w, r)
		},
	)
}

func ServerHandler(h core.Handler, myncerCtx *core.MyncerCtx /*const*/) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create custom ctx passing myncer ctx down with it.
		ctx := core.WithMyncerCtx(r.Context(), myncerCtx)

		// Get user based on JWT auth.
		user, err := auth.MaybeGetUserFromRequest(ctx, r)
		if err != nil {
			// Not a fatal case. Expected for unathenticated endpoints.
			// Logging error for now but if it gets too much, we can remove.
			core.Warning(core.WrappedError(err, "failed to get user from request"))
		}

		// Unmarshal request body here for usage in process request.
		// Do this before checking perms since handlers often need to check perms based on the body.
		reqContainer := h.GetRequestContainer(ctx)
		if reqContainer != nil {
			if err := json.NewDecoder(r.Body).Decode(reqContainer); err != nil {
				core.Errorf(core.WrappedError(err, "failed to decode request body into container"))
				return
			}
		}

		// Check perms.
		if err := h.CheckUserPermissions(ctx, user, reqContainer); err != nil {
			core.Errorf(core.WrappedError(err, "check user perms failed"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		prr := h.ProcessRequest(ctx, user, reqContainer, r, w)
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
