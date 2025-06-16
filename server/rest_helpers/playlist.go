package rest_helpers

import (
	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/zmb3/spotify/v2"
)

func ProtoPlaylistToRest(
	playlist *myncer_pb.Playlist, /*const*/
) (*api.Playlist, error) {
	restDs, err := ProtoDatasourceToRest(playlist.GetMusicSource().GetDatasource())
	if err != nil {
		return nil, core.WrappedError(err, "failed to convert proto datasource to rest")
	}
	r := api.NewPlaylist(restDs, playlist.GetMusicSource().GetPlaylistId())
	r.SetName(playlist.GetName())
	r.SetDescription(playlist.GetDescription())
	r.SetImageUrl(playlist.GetImageUrl())
	return r, nil
}

func SpotifyPlaylistToProto(p *spotify.FullPlaylist /*const*/) *myncer_pb.Playlist {
	return &myncer_pb.Playlist{
		Name:        p.Name,
		Description: p.Description,
		ImageUrl:    GetBestSpotifyImageURL(p.Images),
		MusicSource: &myncer_pb.MusicSource{
			Datasource: myncer_pb.Datasource_SPOTIFY,
			PlaylistId: string(p.ID),
		},
	}
}

// GetBestSpotifyImageURL returns the URL of the first available image from the provided images.
func GetBestSpotifyImageURL(images []spotify.Image /*const*/) string {
	if len(images) > 0 {
		return images[0].URL
	}
	return ""
}
