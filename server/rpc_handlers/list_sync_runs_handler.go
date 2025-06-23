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
	syncRuns, err := core.ToMyncerCtx(ctx).DB.SyncRunStore.GetSyncs(ctx, nil /*runIds*/, nil /*syncIds*/)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListSyncRunsResponse](
			core.WrappedError(err, "failed to get sync runs from database"),
		)
	}

	filteredSyncs, err := l.filterUserSyncs(ctx, syncRuns, userInfo)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListSyncRunsResponse](
			core.WrappedError(err, "failed to filter sync runs for user"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(
		&myncer_pb.ListSyncRunsResponse{SyncRuns: filteredSyncs.ToArray()},
	)
}

func (l *listSyncRunsImpl) filterUserSyncs(
	ctx context.Context,
	syncRuns core.Set[*myncer_pb.SyncRun], /*const*/
	userInfo *myncer_pb.User, /*const*/
) (core.Set[*myncer_pb.SyncRun], error) {
	userSyncIds, err := l.getUserSyncIds(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get user sync ids")
	}
	r := core.NewSet[*myncer_pb.SyncRun]()
	for _, sync := range syncRuns.ToArray() {
		if userSyncIds.Contains(sync.GetSyncId()) {
			r.Add(sync)
		}
	}
	return r, nil
}

func (l *listSyncRunsImpl) getUserSyncIds(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (core.Set[string], error) {
	userSyncs, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSyncs(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get user syncs")
	}
	r := core.NewSet[string]()
	for _, sync := range userSyncs.ToArray() {
		r.Add(sync.GetId())
	}
	return r, nil
}
