package main

import (
	"context"
	"net/http"

	connect_cors "connectrpc.com/cors"
	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/datasources"
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

	// All routes are served on a single mux.
	// We expect there is no path conflict between REST and GRPC for the time being.
	// The long term goal is to remove API entirely.
	mux := http.NewServeMux()

	// Register GRPC routes.
	userService := services.NewUserService()
	datasourceService := services.NewDatasourceService()
	syncService := services.NewSyncService()
	path, grpcHandler := myncer_pb_connect.NewUserServiceHandler(userService)
	mux.Handle(path, GetWrappedGrpcHandler(grpcHandler, myncerCtx))
	path, grpcHandler = myncer_pb_connect.NewDatasourceServiceHandler(datasourceService)
	mux.Handle(path, GetWrappedGrpcHandler(grpcHandler, myncerCtx))
	path, grpcHandler = myncer_pb_connect.NewSyncServiceHandler(syncService)
	mux.Handle(path, GetWrappedGrpcHandler(grpcHandler, myncerCtx))

	core.Printf("gRPC server listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		core.Errorf("failed: ", err)
	}
}

func GetWrappedGrpcHandler(handler http.Handler, myncerCtx *core.MyncerCtx /*const*/) http.Handler {
	// Order matters here.
	grpcWrapped := handler
	grpcWrapped = WithGRPCCors(grpcWrapped, myncerCtx)
	grpcWrapped = WithPossibleUser(grpcWrapped, myncerCtx)
	return WithMyncerCtx(grpcWrapped, myncerCtx)
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
			user, _ := auth.MaybeGetUserFromRequest(ctx, myncerCtx, r)
			if user != nil {
				ctx = auth.ContextWithUser(ctx, user)
			}
			r = r.WithContext(ctx)
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

func TestSongs(ctx context.Context) {
	n := sync_engine.NewLlmSongsNormalizer()
	protoSongs := []*myncer_pb.Song{
		{
			Name:             "Michael Jackson - Billie Jean (Live performance at 1986)",
			ArtistName:       []string{"VEVO Music"},
			AlbumName:        "Man in the mirror",
			Datasource:       myncer_pb.Datasource_DATASOURCE_YOUTUBE,
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
