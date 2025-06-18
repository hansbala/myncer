package main

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	connect_cors "connectrpc.com/cors"
	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/datasources"
	"github.com/hansbala/myncer/handlers"
	"github.com/hansbala/myncer/llm"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	myncer_pb_connect "github.com/hansbala/myncer/proto/myncer/myncer_pbconnect"
	"github.com/hansbala/myncer/services"
	"github.com/hansbala/myncer/sync_engine"
	"github.com/rs/cors"
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
	ctx = core.WithMyncerCtx(ctx, myncerCtx)
	// Any testing stuff here.
	// TestSongs(ctx)

	// REST server.
	for pattern, handler := range GetHandlersMap() {
		http.Handle(pattern, WithCors(ServerHandler(handler, myncerCtx), myncerCtx))
	}
	go func() {
		core.Printf("Myncer listening on port 8080")
		if err := http.ListenAndServe(":8080", nil /*handler*/); err != nil {
			core.Errorf("failed: ", err)
		}
	}()

	// GRPC server.
	go func() {
		userService := services.NewUserService()
		// this gives you the path pattern and the bare connect handler
		path, grpcHandler := myncer_pb_connect.NewUserServiceHandler(userService)

		// wrap it step by step:
		var h http.Handler = grpcHandler
		// 3️⃣ apply CORS for all gRPC calls
		h = WithGRPCCors(h, myncerCtx)
		// 2️⃣ attach a possible user from the JWT (if present)
		h = WithPossibleUser(h, myncerCtx)
		// 1️⃣ inject your MyncerCtx into every request context
		h = WithMyncerCtx(h, myncerCtx)

		// register the fully-wrapped handler
		mux := http.NewServeMux()
		mux.Handle(path, h)

		// and finally serve it over HTTP/2 (h2c)
		server := &http.Server{
			Addr:    "localhost:6969",
			Handler: h2c.NewHandler(mux, &http2.Server{}),
		}
		core.Printf("gRPC listening on port 6969")
		if err := server.ListenAndServe(); err != nil {
			core.Errorf(core.WrappedError(err, "gRPC server error"))
		}
	}()

	// Keep server alive.
	for {
	}
}

func GetHandlersMap() map[string]core.Handler {
	return map[string]core.Handler{
		// User handlers.
		"/api/v1/users/create": handlers.NewCreateUserHandler(),
		"/api/v1/users/login":  handlers.NewLoginUserHandler(),
		"/api/v1/users/logout": handlers.NewLogoutUserHandler(),
		"/api/v1/users/me":     handlers.NewCurrentUserHandler(),
		"/api/v1/users/edit":   handlers.NewEditUserHandler(),
		// Datasource handlers.
		"/api/v1/auth/{datasource}/exchange":                      handlers.NewAuthExchangeHandler(),
		"/api/v1/datasources/list":                                handlers.NewListDatasourcesHandler(),
		"/api/v1/datasources/{datasource}/playlists/list":         handlers.NewListDatasourcePlaylistsHandler(),
		"/api/v1/datasources/{datasource}/playlists/{playlistId}": handlers.NewGetDatasourcePlaylistHandler(),
		// Syncs handlers.
		"/api/v1/syncs/create": handlers.NewCreateSyncHandler(),
		"/api/v1/syncs/delete": handlers.NewDeleteSyncHandler(),
		"/api/v1/syncs/list":   handlers.NewListSyncsHandler(),
		"/api/v1/syncs/run":    handlers.NewRunSyncHandler(sync_engine.NewSyncEngine()),
	}
}

func WithMyncerCtx(h http.Handler, myncerCtx *core.MyncerCtx /*const*/) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Create custom ctx passing myncer ctx down with it.
			ctx := core.WithMyncerCtx(r.Context(), myncerCtx)
			// Set the context in the request.
			r = r.WithContext(ctx)
			// Serve the request.
			h.ServeHTTP(w, r)
		},
	)
}

func WithPossibleUser(h http.Handler, myncerCtx *core.MyncerCtx /*const*/) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			// Get user based on JWT auth.
			user, err := auth.MaybeGetUserFromRequest(ctx, myncerCtx, r)
			if err != nil {
				core.Warning(core.WrappedError(err, "failed to get user from request for proto handler"))
			}
			// Set user in myncerCtx if it exists.
			if user != nil {
				core.ToMyncerCtx(ctx).SetRequestUser(user)
			}
			h.ServeHTTP(w, r)
		},
	)
}

func WithGRPCCors(h http.Handler, myncerCtx *core.MyncerCtx /*const*/) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://myncer.hansbala.com"},
		AllowedMethods:   connect_cors.AllowedMethods(),
		AllowedHeaders:   connect_cors.AllowedHeaders(),
		ExposedHeaders:   connect_cors.ExposedHeaders(),
		AllowCredentials: true, // for cookies
	})
	return middleware.Handler(h)
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
		user, err := auth.MaybeGetUserFromRequest(ctx, myncerCtx, r)
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

func TestSongs(ctx context.Context) {
	n := sync_engine.NewLlmSongsNormalizer()
	protoSongs := []*myncer_pb.Song{
		{
			Name:             "Michael Jackson - Billie Jean (Live performance at 1986)",
			ArtistName:       []string{"VEVO Music"},
			AlbumName:        "Man in the mirror",
			Datasource:       myncer_pb.Datasource_YOUTUBE,
			DatasourceSongId: "abcd-12344727-2762",
		},
	}
	core.Printf("-------------")
	core.Printf("trying to normalize songs: ")
	core.DebugPrintJson(protoSongs)

	testSongs := []core.Song{}
	for _, ps := range protoSongs {
		testSongs = append(testSongs, sync_engine.NewSong(ps))
	}

	normalizedSongs, err := n.NormalizeSongs(ctx, core.NewSongList(testSongs))
	if err != nil {
		panic(err)
	}
	core.Printf("-------------")
	core.Printf("normalized songs: ")
	b, err := normalizedSongs.GetLlmJson()
	if err != nil {
		panic(err)
	}
	core.Printf(string(b))
}
