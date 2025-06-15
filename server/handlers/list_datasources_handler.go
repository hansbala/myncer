package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewListDatasourcesHandler() core.Handler {
	return &listDatasourcesHandlerImpl{}
}

type listDatasourcesHandlerImpl struct{}

var _ core.Handler = (*listDatasourcesHandlerImpl)(nil)

func (ld *listDatasourcesHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil
}

func (ld *listDatasourcesHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list datasources")
	}
	return nil
}

func (ld *listDatasourcesHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {

	return nil
}
