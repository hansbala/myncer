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

func (l *listDatasourcePlaylistsImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListPlaylistsRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list datasource playlists")
	}
	return nil
}

func (l *listDatasourcePlaylistsImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListPlaylistsRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListPlaylistsResponse] {
	dsClient, err := l.getClientFromDatasource(ctx, reqBody.GetDatasource())
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListPlaylistsResponse](
			core.WrappedError(err, "failed to get datasource client"),
		)
	}
	playlists, err := dsClient.GetPlaylists(ctx, userInfo)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListPlaylistsResponse](
			core.WrappedError(err, "failed to get playlists for current user"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(
		&myncer_pb.ListPlaylistsResponse{
			Playlist: playlists,
		},
	)
}

func (l *listDatasourcePlaylistsImpl) getClientFromDatasource(
	ctx context.Context,
	ds myncer_pb.Datasource,
) (core.DatasourceClient, error) {
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	switch ds {
	case myncer_pb.Datasource_DATASOURCE_SPOTIFY:
		return dsClients.SpotifyClient, nil
	case myncer_pb.Datasource_DATASOURCE_YOUTUBE:
		return dsClients.YoutubeClient, nil
	default:
		return nil, core.NewError("unsupported datasource: %v", ds)
	}
}
