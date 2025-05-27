package core

import (
	"context"
	"net/http"

	myncer_pb "github.com/hansbala/myncer/proto"
)

type Handler interface {
	GetRequestContainer(ctx context.Context) any /*@nullable*/
	CheckUserPermissions(ctx context.Context, userInfo *myncer_pb.User /*const,@nullable*/) error
	ProcessRequest(
		ctx context.Context,
		reqBody any, /*const,@nullable*/
		req *http.Request, /*const*/
		resp http.ResponseWriter,
	) error
}
