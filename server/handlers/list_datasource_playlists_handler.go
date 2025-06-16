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

func NewListDatasourcePlaylistsHandler() core.Handler {
	return &listDsPlaylistsHandlerImpl{}
}

type listDsPlaylistsHandlerImpl struct{}

var _ core.Handler = (*listDsPlaylistsHandlerImpl)(nil)

func (ldp *listDsPlaylistsHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil // No request body since it's a GET request.
}

func (ldp *listDsPlaylistsHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list datasource playlists")
	}
	return nil
}

func (ldp *listDsPlaylistsHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	ds, err := ldp.getDatasource(req)
	if err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "failed to get datasource from request path"),
		)
	}
	dsClient, err := ldp.getClientFromDatasource(ctx, ds)
	if err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "failed to get datasource client from request path"),
		)
	}
	oAuthToken, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetToken(
		ctx,
		userInfo.GetId(),
		ds,
	)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to get OAuth token for current user for datasource %v", ds),
		)
	}

	playlists, err := dsClient.GetPlaylists(ctx, oAuthToken)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to get playlists for current user"),
		)
	}

	rps := []api.Playlist{}
	for _, protoPlaylist := range playlists {
		rp, err := rest_helpers.ProtoPlaylistToRest(protoPlaylist)
		if err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to convert proto playlist to rest"),
			)
		}
		rps = append(rps, *rp)
	}

	if err := WriteJSONOk(resp, api.NewListDatasourcePlaylistsResponse(rps)); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write list datasource playlists response"),
		)
	}
	return core.NewProcessRequestResponse_OK()
}

func (ldp *listDsPlaylistsHandlerImpl) getDatasource(
	req *http.Request,
) (myncer_pb.Datasource, error) {
	// Path is expected to be like /api/v1/datasources/{datasource}/playlists/list
	// so we extract this                              ^ portion out.
	pathParts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(pathParts) < 6 {
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, core.NewError(
			"malformed path: %s",
			req.URL.Path,
		)
	}
	datasource := pathParts[len(pathParts)-3]
	return rest_helpers.RestDatasourceToProto(api.Datasource(datasource)), nil
}

func (ldp *listDsPlaylistsHandlerImpl) getClientFromDatasource(
	ctx context.Context,
	ds myncer_pb.Datasource,
) (core.DatasourceClient, error) {
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	switch ds {
	case myncer_pb.Datasource_SPOTIFY:
		return dsClients.SpotifyClient, nil
	case myncer_pb.Datasource_YOUTUBE:
		return dsClients.YoutubeClient, nil
	default:
		return nil, core.NewError("unsupported datasource: %v", ds)
	}
}
