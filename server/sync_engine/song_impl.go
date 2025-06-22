package sync_engine

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewSong(
	spec *myncer_pb.Song, /*const*/
) core.Song {
	return &songImpl{
		spec: spec,
	}
}

type songImpl struct {
	spec *myncer_pb.Song
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
	case myncer_pb.Datasource_DATASOURCE_SPOTIFY:
		return s.getSpotifyId(ctx, userInfo)
	case myncer_pb.Datasource_DATASOURCE_YOUTUBE:
		return s.getYoutubeId(ctx, userInfo)
	default:
		return "", core.NewError("Unknown datasource: %v", datasource)
	}
}

func (s *songImpl) GetSpec() *myncer_pb.Song {
	return s.spec
}

func (s *songImpl) getSpotifyId(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (string, error) {
	if s.spec.GetDatasource() == myncer_pb.Datasource_DATASOURCE_SPOTIFY {
		return s.spec.GetDatasourceSongId(), nil
	}
	// Otherwise, try searching Spotify
	result, err := core.ToMyncerCtx(ctx).DatasourceClients.SpotifyClient.Search(
		ctx,
		userInfo,
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
	if s.spec.GetDatasource() == myncer_pb.Datasource_DATASOURCE_YOUTUBE {
		return s.spec.GetDatasourceSongId(), nil
	}
	// Otherwise, try searching Youtube.
	result, err := core.ToMyncerCtx(ctx).DatasourceClients.YoutubeClient.Search(
		ctx,
		userInfo,
		core.NewSet(s.GetName()),
		core.NewSet(s.GetArtistNames()...),
		core.NewSet(s.GetAlbum()),
	)
	if err != nil {
		return "", core.WrappedError(err, "youtube search failed for song: %s", s.GetName())
	}
	return result.GetId(), nil
}
