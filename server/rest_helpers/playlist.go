package rest_helpers

import (
	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
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
