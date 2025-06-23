package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewGetSyncHandler() core.GrpcHandler[
	*myncer_pb.GetSyncRequest,
	*myncer_pb.GetSyncResponse,
] {
	return &getSyncImpl{}
}

type getSyncImpl struct{}

func (l *getSyncImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.GetSyncRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required for getting sync")
	}
	sync, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSync(ctx, reqBody.GetSyncId())
	if err != nil {
		return core.WrappedError(err, "failed to get sync from database")
	}
	if sync.GetUserId() != userInfo.GetId() {
		return core.NewError("user does not have permission to get this sync")
	}
	return nil
}

func (l *getSyncImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.GetSyncRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.GetSyncResponse] {
	sync, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSync(ctx, reqBody.GetSyncId())
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.GetSyncResponse](
			core.WrappedError(err, "failed to get sync from database"),
		)
	}
	return core.NewGrpcHandlerResponse_OK(&myncer_pb.GetSyncResponse{Sync: sync})
}
