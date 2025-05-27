package core

import (
	"context"
	"net/http"

	myncer_pb "github.com/hansbala/myncer/proto"
)

type Handler interface {
	CheckUserPermissions(ctx context.Context, userInfo *myncer_pb.User /*const,@nullable*/) error
	ProcessRequest(ctx context.Context, req *http.Request /*const*/, resp http.ResponseWriter) error
}
