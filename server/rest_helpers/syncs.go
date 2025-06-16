package rest_helpers

import (
	"github.com/google/uuid"

	"github.com/hansbala/myncer/api"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func RestOneWaySyncToProto(
	s *api.OneWaySync, /*const*/
	userId string,
) *myncer_pb.Sync {
	return &myncer_pb.Sync{
		Id:     uuid.NewString(),
		UserId: userId,
		SyncVariant: &myncer_pb.Sync_OneWaySync{
			OneWaySync: &myncer_pb.OneWaySync{
				Source: CreateMusicSource(
					RestDatasourceToProto(s.GetSource().Datasource),
					s.GetSource().PlaylistId,
				),
				Destination: CreateMusicSource(
					RestDatasourceToProto(s.GetDestination().Datasource),
					s.GetDestination().PlaylistId,
				),
				OverwriteExisting: s.GetOverwriteExisting(),
			},
		},
	}
}

func CreateMusicSource(
	datasource myncer_pb.Datasource,
	playlistId string,
) *myncer_pb.MusicSource {
	return &myncer_pb.MusicSource{
		Datasource: datasource,
		PlaylistId: playlistId,
	}
}
