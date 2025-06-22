package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewListSyncsHandler() core.GrpcHandler[
	*myncer_pb.ListSyncsRequest,
	*myncer_pb.ListSyncsResponse,
] {
	return &listSyncsImpl{}
}

type listSyncsImpl struct{}

func (ls *listSyncsImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListSyncsRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list syncs")
	}
	return nil
}

func (ls *listSyncsImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListSyncsRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListSyncsResponse] {
	// Get all syncs for current user.
	syncs, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSyncs(ctx, userInfo)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListSyncsResponse](
			core.WrappedError(err, "failed to get syncs for current user"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(&myncer_pb.ListSyncsResponse{Syncs: syncs.ToArray()})
}
