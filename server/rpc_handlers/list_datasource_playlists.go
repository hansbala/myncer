package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewListDatasourcePlaylistsHandler() core.GrpcHandler[
	*myncer_pb.ListPlaylistsRequest,
	*myncer_pb.ListPlaylistsResponse,
] {
	return &listDatasourcePlaylistsImpl{}
}

type listDatasourcePlaylistsImpl struct{}

func (l *listDatasourcePlaylistsImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListPlaylistsRequest, /*const*/
) error {
	return nil
}

func (l *listDatasourcePlaylistsImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListPlaylistsRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListPlaylistsResponse] {
	return nil
}
