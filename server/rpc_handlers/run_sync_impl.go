package rpc_handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewRunSyncHandler(syncEngine core.SyncEngine) core.GrpcHandler[
	*myncer_pb.RunSyncRequest,
	*myncer_pb.RunSyncResponse,
] {
	return &runSyncImpl{
		syncEngine: syncEngine,
	}
}

type runSyncImpl struct{
	syncEngine core.SyncEngine
}

func (rs *runSyncImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.RunSyncRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to run sync job")
	}
	if err := rs.validateRequest(ctx, userInfo, reqBody); err != nil {
		return core.WrappedError(err, "failed to validate request")
	}
	return nil
}

func (rs *runSyncImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.RunSyncRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.RunSyncResponse] {
	sync, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSync(ctx, reqBody.GetSyncId())
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.RunSyncResponse](
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

	return core.NewGrpcHandlerResponse_OK(
		&myncer_pb.RunSyncResponse{
			SyncId:       sync.GetId(),
			Status:       myncer_pb.SyncStatus_SYNC_STATUS_RUNNING,
		},
	)
}

func (rs *runSyncImpl) validateRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.RunSyncRequest, /*const*/
) error {
	if len(reqBody.GetSyncId()) == 0 {
		return core.NewError("sync id is required")
	}
	if _, err := uuid.Parse(reqBody.GetSyncId()); err != nil {
		return core.NewError("invalid sync id: %v", err)
	}
	sync, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSync(ctx, reqBody.GetSyncId())
	if err != nil {
		return core.WrappedError(err, "could not get sync with id: %s", reqBody.GetSyncId())
	}
	if userInfo.GetId() != sync.GetUserId() {
		return core.NewError(
			"user %s does not have permission to run sync %s",
			userInfo.GetId(),
			reqBody.GetSyncId(),
		)
	}
	return nil
}

