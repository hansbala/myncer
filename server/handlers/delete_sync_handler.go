package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewDeleteSyncHandler() core.Handler {
	return &deleteSyncHandlerImpl{}
}

type deleteSyncHandlerImpl struct{}

var _ core.Handler = (*deleteSyncHandlerImpl)(nil)

func (ds *deleteSyncHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return &api.DeleteSyncRequest{}
}

func (ds *deleteSyncHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to delete sync")
	}
	return nil
}

func (ds *deleteSyncHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	restReq, ok := reqBody.(*api.DeleteSyncRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("expected DeleteSyncRequest got %T", reqBody),
		)
	}

	if err := core.ToMyncerCtx(ctx).DB.SyncStore.DeleteSync(ctx, restReq.GetSyncId()); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "could not delete sync by id"),
		)
	}
	return core.NewProcessRequestResponse_OK()
}
