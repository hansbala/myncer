package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewTemplateHandler() core.Handler {
	return &templateHandlerImpl{}
}

type templateHandlerImpl struct{}

var _ core.Handler = (*templateHandlerImpl)(nil)

func (thi *templateHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil
}

func (thi *templateHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	return nil
}

func (thi *templateHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	return core.NewProcessRequestResponse_OK()
}
