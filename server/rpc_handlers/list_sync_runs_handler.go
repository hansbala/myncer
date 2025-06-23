package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewListSyncRunsHandler() core.GrpcHandler[
	*myncer_pb.ListSyncRunsRequest,
	*myncer_pb.ListSyncRunsResponse,
] {
	return &listSyncRunsImpl{}
}

type listSyncRunsImpl struct{}

func (l *listSyncRunsImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.ListSyncRunsRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required for listing sync runs")
	}
	return nil
}

func (l *listSyncRunsImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListSyncRunsRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListSyncRunsResponse] {
	syncs, err := core.ToMyncerCtx(ctx).DB.SyncRunStore.GetSyncs(ctx, nil /*runIds*/, nil /*syncIds*/)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListSyncRunsResponse](
			core.WrappedError(err, "failed to get sync runs from database"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(&myncer_pb.ListSyncRunsResponse{SyncRuns: syncs.ToArray()})
}
