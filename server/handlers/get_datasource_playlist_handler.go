package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewGetDatasourcePlaylistHandler() core.Handler {
	return &getDsPlaylistHandlerImpl{}
}

type getDsPlaylistHandlerImpl struct{}

var _ core.Handler = (*getDsPlaylistHandlerImpl)(nil)

func (dp *getDsPlaylistHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil // GET request.
}

func (dp *getDsPlaylistHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required for getting datasource playlist")
	}
	return nil
}

func (dp *getDsPlaylistHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	// Extract request info.
	datasource, playlistId, err := dp.getDatasourceAndPlaylistId(req)
	if err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "failed to get datasource and playlist ID from request path"),
		)
	}

	// Get the proto playlist.
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	playlist := &myncer_pb.Playlist{}
	switch datasource {
		case myncer_pb.Datasource_SPOTIFY:
			playlist, err = dsClients.SpotifyClient.GetPlaylist(ctx, userInfo, playlistId)
			if err != nil {
				return core.NewProcessRequestResponse_InternalServerError(
					core.WrappedError(err, "failed to get Spotify playlist"),
				)
			}
		case myncer_pb.Datasource_YOUTUBE:
			playlist, err = dsClients.YoutubeClient.GetPlaylist(ctx, userInfo, playlistId)
			if err != nil {
				return core.NewProcessRequestResponse_InternalServerError(
					core.WrappedError(err, "failed to get YouTube playlist"),
				)
			}
		default:
			return core.NewProcessRequestResponse_InternalServerError(
				core.NewError("unsupported datasource: %s", datasource),
			)
	}

	// Convert to REST and write the response.
	restPlaylist, err := rest_helpers.ProtoPlaylistToRest(playlist)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to convert proto playlist to rest"),
		)
	}
	if err := WriteJSONOk(resp, restPlaylist); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write response"),
		)
	}
	return core.NewProcessRequestResponse_OK()
}

func (dp *getDsPlaylistHandlerImpl) getDatasourceAndPlaylistId(
	req *http.Request, /*const*/
) (myncer_pb.Datasource, string, error) {
	// Path is expected to be like /api/v1/datasources/{datasource}/playlists/{playlistId}
	// so we extract this                               ^                      ^ out.
	pathParts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(pathParts) < 6 {
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, "", core.NewError(
			"malformed path: %s",
			req.URL.Path,
		)
	}

	protoDatasource := rest_helpers.RestDatasourceToProto(api.Datasource(pathParts[3]))
	if protoDatasource == myncer_pb.Datasource_DATASOURCE_UNSPECIFIED {
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, "", core.NewError(
			"invalid datasource: %s",
			pathParts[3],
		)
	}

	return protoDatasource, pathParts[5], nil
}
