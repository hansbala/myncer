package rest_helpers

import (
	"github.com/google/uuid"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func ProtoSyncToRest(
	sync *myncer_pb.Sync, /*const*/
) (*api.Sync, error) {
	restSyncVariant, err := protoSyncToSyncVariantRest(sync)
	if err != nil {
		return nil, core.WrappedError(err, "unable to determine rest sync variant")
	}
	restSyncData, err := protoSyncToRestSyncData(sync)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get rest sync data")
	}
	return api.NewSync(
		sync.GetId(),
		sync.GetCreatedAt().AsTime(),
		sync.GetUpdatedAt().AsTime(),
		restSyncVariant,
		*restSyncData,
	), nil
}

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

func ProtoMusicSourceToRest(ms *myncer_pb.MusicSource /*const*/) (*api.MusicSource, error) {
	restDs, err := ProtoDatasourceToRest(ms.GetDatasource())
	if err != nil {
		return nil, core.WrappedError(err, "failed to convert datasource to rest")
	}
	return api.NewMusicSource(restDs, ms.GetPlaylistId()), nil
}

func protoSyncToSyncVariantRest(
	sync *myncer_pb.Sync, /*const*/
) (api.SyncVariant, error) {
	switch sync.GetSyncVariant().(type) {
	case *myncer_pb.Sync_OneWaySync:
		return api.ONE_WAY, nil
	default:
		return "", core.NewError("unknown sync variant of type %T", sync.GetSyncVariant())
	}
}

func protoSyncToRestSyncData(
	sync *myncer_pb.Sync, /*const*/
) (*api.SyncSyncData, error) {
	v := sync.GetSyncVariant()
	switch v.(type) {
	case *myncer_pb.Sync_OneWaySync:
		oneWaySync, err := protoOneWaySyncToRest(sync.GetOneWaySync())
		if err != nil {
			return nil, core.WrappedError(err, "failed to convert one way sync to rest")
		}
		return &api.SyncSyncData{OneWaySync: oneWaySync}, nil
	default:
		return nil, core.NewError("unknown sync variant when converting to rest sync data")
	}
}

func protoOneWaySyncToRest(s *myncer_pb.OneWaySync /*const*/) (*api.OneWaySync, error) {
	source, err := ProtoMusicSourceToRest(s.GetSource())
	if err != nil {
		return nil, core.WrappedError(err, "failed to convert source to rest")
	}
	dest, err := ProtoMusicSourceToRest(s.GetDestination())
	if err != nil {
		return nil, core.WrappedError(err, "failed to convert destination to rest")
	}
	return api.NewOneWaySync(api.ONE_WAY, *source, *dest), nil
}
