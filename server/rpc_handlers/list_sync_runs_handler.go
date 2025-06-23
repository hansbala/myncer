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
	return core.NewError("not implemented: CheckUserPermissions for ListSyncRunsHandler")
}

func (l *listSyncRunsImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListSyncRunsRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListSyncRunsResponse] {
	return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListSyncRunsResponse](
		core.NewError("not implemented: ProcessRequest for ListSyncRunsHandler"),
	)
}
