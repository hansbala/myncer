package sync_engine

import (
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewSong(
	spec *myncer_pb.Song, /*const*/
	spotifyClient core.DatasourceClient,
	youtubeClient core.DatasourceClient,
) core.Song {
	return &songImpl{
		spec:          spec,
		spotifyClient: spotifyClient,
		youtubeClient: youtubeClient,
	}
}

type songImpl struct {
	spec          *myncer_pb.Song
	spotifyClient core.DatasourceClient
	youtubeClient core.DatasourceClient
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

func (s *songImpl) GetIdByDatasource(datasource myncer_pb.Datasource) (string, error) {
	switch datasource {
	case myncer_pb.Datasource_SPOTIFY:
		return s.getSpotifyId()
	case myncer_pb.Datasource_YOUTUBE:
		return s.getYoutubeId()
	default:
		return "", core.NewError("Unknown datasource: %v", datasource)
	}
}

func (s *songImpl) getSpotifyId() (string, error) {
	return "", core.NewError("Spotify ID not implemented")
}

func (s *songImpl) getYoutubeId() (string, error) {
	return "", core.NewError("YouTube ID not implemented")
}
