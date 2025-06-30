package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewDeleteSyncHandler() core.GrpcHandler[
	*myncer_pb.DeleteSyncRequest,
	*myncer_pb.DeleteSyncResponse,
] {
	return &deleteSyncImpl{}
}

type deleteSyncImpl struct{}

func (l *deleteSyncImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.DeleteSyncRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to delete sync")
	}
	// Makes sure the sync belongs to the user.
	sync, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSync(ctx, reqBody.GetSyncId())
	if err != nil {
		return core.WrappedError(err, "could not find sync with id: %s", reqBody.GetSyncId())
	}
	if userInfo.GetId() != sync.GetUserId() {
		return core.NewError(
			"user %s does not have permission to delete sync %s",
			userInfo.GetId(),
			reqBody.GetSyncId(),
		)
	}
	return nil
}

func (l *deleteSyncImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.DeleteSyncRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.DeleteSyncResponse] {
	if err := core.ToMyncerCtx(ctx).DB.SyncStore.DeleteSync(ctx, reqBody.GetSyncId()); err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.DeleteSyncResponse](
			core.WrappedError(err, "could not delete sync with id %s", reqBody.GetSyncId()),
		)
	}

	return core.NewGrpcHandlerResponse_OK(&myncer_pb.DeleteSyncResponse{SyncId: reqBody.GetSyncId()})
}
