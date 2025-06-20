package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewRunSyncHandler(
	syncEngine core.SyncEngine,
) core.Handler {
	return &runSyncHandlerImpl{
		syncEngine:    syncEngine,
	}
}

type runSyncHandlerImpl struct {
	syncEngine    core.SyncEngine
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
	sync, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSync(ctx, restReq.GetSyncId())
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "could not get sync by id"),
		)
	}
	// We want to be able to run syncs in the background.
	ctx = context.WithoutCancel(ctx)
	go func() {
		if err := rs.syncEngine.RunSync(ctx, userInfo, sync); err != nil {
			core.Errorf(core.WrappedError(err, "failed to run sync job"))
		}
	}()

	return core.NewProcessRequestResponse(
		"Sync job accepted and triggered",
		nil, /*err*/
		http.StatusAccepted,
	)
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
