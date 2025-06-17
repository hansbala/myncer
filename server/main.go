package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/datasources"
	"github.com/hansbala/myncer/handlers"
	"github.com/hansbala/myncer/llm"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/sync_engine"
)

func main() {
	ctx := context.Background()
	spotifyClient := datasources.NewSpotifyClient()
	youtubeClient := datasources.NewYouTubeClient()
	myncerCtx := core.MustGetMyncerCtx(
		ctx,
		&core.DatasourceClients{
			SpotifyClient: spotifyClient,
			YoutubeClient: youtubeClient,
		},
		&core.LlmClients{
			GeminiLlmClient: llm.NewGeminiLlmClient(),
			OpenAILlmClient: llm.NewOpenAILlmClient(),
		},
	)

	for pattern, handler := range GetHandlersMap(myncerCtx) {
		http.Handle(pattern, WithCors(ServerHandler(handler, myncerCtx), myncerCtx))
	}
	core.Printf("Myncer listening on port 8080")
	if err := http.ListenAndServe(":8080", nil /*handler*/); err != nil {
		core.Errorf("failed: ", err)
	}
}

func GetHandlersMap(
	myncerCtx *core.MyncerCtx, /*const*/
) map[string]core.Handler {
	return map[string]core.Handler{
		// User handlers.
		"/api/v1/users/create": handlers.NewCreateUserHandler(),
		"/api/v1/users/login":  handlers.NewLoginUserHandler(),
		"/api/v1/users/logout": handlers.NewLogoutUserHandler(),
		"/api/v1/users/me":     handlers.NewCurrentUserHandler(),
		"/api/v1/users/edit":   handlers.NewEditUserHandler(),
		// Datasource handlers.
		"/api/v1/auth/{datasource}/exchange":              handlers.NewAuthExchangeHandler(),
		"/api/v1/datasources/list":                        handlers.NewListDatasourcesHandler(),
		"/api/v1/datasources/{datasource}/playlists/list": handlers.NewListDatasourcePlaylistsHandler(),
		// Syncs handlers.
		"/api/v1/syncs/create": handlers.NewCreateSyncHandler(),
		"/api/v1/syncs/list":   handlers.NewListSyncsHandler(),
		"/api/v1/syncs/run":    handlers.NewRunSyncHandler(sync_engine.NewSyncEngine()),
	}
}

func WithCors(h http.Handler, myncerCtx *core.MyncerCtx /*const*/) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowedOrigin := ""
			if myncerCtx.Config.ServerMode == myncer_pb.ServerMode_DEV {
				if origin == "http://localhost:5173" || origin == "http://localhost" {
					allowedOrigin = origin
				}
			} else {
				if origin == "https://myncer.hansbala.com" {
					allowedOrigin = origin
				}
			}
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
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
