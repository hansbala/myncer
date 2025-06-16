package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewRunSyncHandler(
	spotifyClient core.DatasourceClient,
	youtubeClient core.DatasourceClient,
) core.Handler {
	return &runSyncHandlerImpl{}
}

type runSyncHandlerImpl struct{
	spotifyClient core.DatasourceClient
	youtubeClient core.DatasourceClient
}

var _ core.Handler = (*runSyncHandlerImpl)(nil)

func (rs *runSyncHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return &api.RunSyncRequest{}
}

func (rs *runSyncHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to run sync job")
	}
	if err := rs.validateRequest(ctx, userInfo, reqBody); err != nil {
		return core.WrappedError(err, "failed to validate request")
	}
	return nil
}

func (rs *runSyncHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	restReq, ok := reqBody.(*api.RunSyncRequest)
	if !ok {
		return core.NewProcessRequestResponse_InternalServerError(
			core.NewError("expected RunSyncRequest got %T", reqBody),
		)
	}
	core.Printf("running sync: ")
	core.DebugPrintJson(restReq)
	return core.NewProcessRequestResponse_OK()
}

func (rs *runSyncHandlerImpl) validateRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
) error {
	restReq, ok := reqBody.(*api.RunSyncRequest)
	if !ok {
		return core.NewError("expected RunSyncRequest got %T", reqBody)
	}
	if len(restReq.GetSyncId()) == 0 {
		return core.NewError("sync_id is required")
	}
	myncerCtx := core.ToMyncerCtx(ctx)
	sync, err := myncerCtx.DB.SyncStore.GetSync(ctx, restReq.GetSyncId())
	// Sync ID must be a valid sync.
	if err != nil {
		return core.WrappedError(err, "could not get sync by id")
	}
	// Validates the user has access to the sync.
	if sync.GetUserId() != userInfo.GetId() {
		return core.WrappedError(err, "user does not have access to run sync")
	}
	return nil
}
