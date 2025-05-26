package main

import (
	"context"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/grpcserver"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/services"
)

func main() {
	ctx := context.Background()
	ctx = core.WithMyncerCtx(ctx, core.MustGetMyncerCtx(ctx))

	wrappedServer := grpcweb.WrapServer(getGrpcServer(ctx))
	log.Println("gRPC-Web listening on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if wrappedServer.IsGrpcWebRequest(r) || wrappedServer.IsAcceptableGrpcCorsRequest(r) {
			wrappedServer.ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
	}))
}

func getGrpcServer(ctx context.Context) *grpc.Server {
	s := grpc.NewServer()
	internalUserService := services.NewUserService(core.ToMyncerCtx(ctx).DB)
	myncer_pb.RegisterUserServiceServer(
		s,
		grpcserver.NewUserServiceServer(internalUserService),
	)
	return s
}
