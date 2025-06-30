package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewGetPlaylistDetailsHandler() core.GrpcHandler[
	*myncer_pb.GetPlaylistDetailsRequest,
	*myncer_pb.GetPlaylistDetailsResponse,
] {
	return &getPlaylistDetailsImpl{}
}

type getPlaylistDetailsImpl struct{}

func (l *getPlaylistDetailsImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.GetPlaylistDetailsRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required for getting datasource playlist")
	}
	return nil
}

func (l *getPlaylistDetailsImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.GetPlaylistDetailsRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.GetPlaylistDetailsResponse] {
	var (
		playlist *myncer_pb.Playlist
		err      error
	)
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	switch reqBody.GetDatasource() {
	case myncer_pb.Datasource_DATASOURCE_SPOTIFY:
		playlist, err = dsClients.SpotifyClient.GetPlaylist(ctx, userInfo, reqBody.GetPlaylistId())
		if err != nil {
			return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.GetPlaylistDetailsResponse](
				core.WrappedError(err, "failed to get Spotify playlist"),
			)
		}
	case myncer_pb.Datasource_DATASOURCE_YOUTUBE:
		playlist, err = dsClients.YoutubeClient.GetPlaylist(ctx, userInfo, reqBody.GetPlaylistId())
		if err != nil {
			return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.GetPlaylistDetailsResponse](
				core.WrappedError(err, "failed to get YouTube playlist"),
			)
		}
	default:
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.GetPlaylistDetailsResponse](
			core.NewError("unsupported datasource: %s", reqBody.GetDatasource()),
		)
	}

	return core.NewGrpcHandlerResponse_OK(&myncer_pb.GetPlaylistDetailsResponse{Playlist: playlist})
}
