package datasources

import (
	"github.com/zmb3/spotify/v2"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func createMusicSource(
	datasource myncer_pb.Datasource,
	playlistId string,
) *myncer_pb.MusicSource {
	return &myncer_pb.MusicSource{
		Datasource: datasource,
		PlaylistId: playlistId,
	}
}

// getBestSpotifyImageURL returns the URL of the first available image from the provided images.
func getBestSpotifyImageURL(images []spotify.Image /*const*/) string {
	if len(images) > 0 {
		return images[0].URL
	}
	return ""
}

func spotifyPlaylistToProto(p *spotify.FullPlaylist /*const*/) *myncer_pb.Playlist {
	return &myncer_pb.Playlist{
		Name:        p.Name,
		Description: p.Description,
		ImageUrl:    getBestSpotifyImageURL(p.Images),
		MusicSource: &myncer_pb.MusicSource{
			Datasource: myncer_pb.Datasource_DATASOURCE_SPOTIFY,
			PlaylistId: string(p.ID),
		},
	}
}

