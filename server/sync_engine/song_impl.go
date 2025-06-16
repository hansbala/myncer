package sync_engine

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewSong(
	spec *myncer_pb.Song, /*const*/
) core.Song {
	return &songImpl{
		spec:          spec,
	}
}

type songImpl struct {
	spec          *myncer_pb.Song
}

var _ core.Song = (*songImpl)(nil)

func (s *songImpl) GetName() string {
	return s.spec.GetName()
}

func (s *songImpl) GetArtistNames() []string {
	return s.spec.GetArtistName()
}

func (s *songImpl) GetAlbum() string {
	return s.spec.GetAlbumName()
}

func (s *songImpl) GetId() string {
	return s.spec.GetDatasourceSongId()
}

func (s *songImpl) GetIdByDatasource(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	datasource myncer_pb.Datasource,
) (string, error) {
	switch datasource {
	case myncer_pb.Datasource_SPOTIFY:
		return s.getSpotifyId(ctx, userInfo)
	case myncer_pb.Datasource_YOUTUBE:
		return s.getYoutubeId(ctx, userInfo)
	default:
		return "", core.NewError("Unknown datasource: %v", datasource)
	}
}

func (s *songImpl) getSpotifyId(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (string, error) {
	if s.spec.GetDatasource() == myncer_pb.Datasource_SPOTIFY {
		return s.spec.GetDatasourceSongId(), nil
	}
	myncerCtx := core.ToMyncerCtx(ctx)
	oAuthToken, err := myncerCtx.DB.DatasourceTokenStore.GetToken(
		ctx,
		userInfo.GetId(),
		myncer_pb.Datasource_SPOTIFY,
	)
	if err != nil {
		return "", core.WrappedError(err, "failed to get OAuth token for Spotify")
	}
	// Otherwise, try searching Spotify
	result, err := myncerCtx.DatasourceClients.SpotifyClient.Search(
		ctx,
		oAuthToken,
		core.NewSet(s.GetName()),
		core.NewSet(s.GetArtistNames()...),
		core.NewSet(s.GetAlbum()),
	)
	if err != nil {
		return "", core.WrappedError(err, "spotify search failed for song: %s", s.GetName())
	}
	return result.GetId(), nil
}

func (s *songImpl) getYoutubeId(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (string, error) {
	if s.spec.GetDatasource() == myncer_pb.Datasource_YOUTUBE {
		return s.spec.GetDatasourceSongId(), nil
	}
	myncerCtx := core.ToMyncerCtx(ctx)
	oAuthToken, err := myncerCtx.DB.DatasourceTokenStore.GetToken(
		ctx,
		userInfo.GetId(),
		myncer_pb.Datasource_YOUTUBE,
	)
	if err != nil {
		return "", core.WrappedError(err, "failed to get OAuth token for Youtube")
	}
	// Otherwise, try searching Youtube.
	result, err := myncerCtx.DatasourceClients.YoutubeClient.Search(
		ctx,
		oAuthToken,
		core.NewSet(s.GetName()),
		core.NewSet(s.GetArtistNames()...),
		core.NewSet(s.GetAlbum()),
	)
	if err != nil {
		return "", core.WrappedError(err, "youtube search failed for song: %s", s.GetName())
	}
	return result.GetId(), nil
}
