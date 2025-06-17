package sync_engine

import (
	"context"
	"fmt"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewSyncEngine() core.SyncEngine {
	return &syncEngineImpl{}
}

type syncEngineImpl struct{}

var _ core.SyncEngine = (*syncEngineImpl)(nil)

func (s *syncEngineImpl) RunSync(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	sync *myncer_pb.Sync, /*const*/
) error {
	core.Printf("Running sync in engine...")
	core.DebugPrintJson(sync)

	switch sync.GetSyncVariant().(type) {
	case *myncer_pb.Sync_OneWaySync:
		if err := s.runOneWaySync(
			ctx,
			userInfo,
			sync.GetOneWaySync(),
		); err != nil {
			return core.WrappedError(err, "failed to run one-way sync")
		}
		return nil
	default:
		return core.NewError(fmt.Sprintf("unknown sync variant: %T", sync.GetOneWaySync()))
	}
}

func (s *syncEngineImpl) runOneWaySync(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	sync *myncer_pb.OneWaySync, /*const*/
) error {
	sourceClient, err := s.getClient(ctx, sync.GetSource().GetDatasource())
	if err != nil {
		return err
	}
	destClient, err := s.getClient(ctx, sync.GetDestination().GetDatasource())
	if err != nil {
		return err
	}

	// Fetch songs from source playlist
	sourceSongs, err := sourceClient.GetPlaylistSongs(ctx, userInfo, sync.GetSource().GetPlaylistId())
	if err != nil {
		return core.WrappedError(err, "failed to fetch source playlist")
	}

	searchedSongs, err := s.getSearchedSongs(
		ctx,
		userInfo,
		sourceSongs,
		sync.GetDestination().GetDatasource(),
	)
	if err != nil {
		return core.WrappedError(err, "failed to get searched songs for destination datasource")
	}

	// Optionally clear destination playlist
	destPlaylistId := sync.GetDestination().GetPlaylistId()
	if sync.OverwriteExisting {
		core.Printf("Clearing destination playlist")
		if err := destClient.ClearPlaylist(ctx, userInfo, destPlaylistId); err != nil {
			return core.WrappedError(err, "failed to clear destination playlist")
		}
	}

	// Add source songs to destination
	if err := destClient.AddToPlaylist(ctx, userInfo, destPlaylistId, searchedSongs); err != nil {
		return core.WrappedError(err, "failed to add songs to destination playlist")
	}
	return nil
}

func (s *syncEngineImpl) getSearchedSongs(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	songs []core.Song, /*const*/
	datasource myncer_pb.Datasource, /*const*/
) ([]core.Song, error) {
	r := []core.Song{}
	for _, song := range songs {
		newDatasourceSongId, err := song.GetIdByDatasource(ctx, userInfo, datasource)
		if err != nil {
			return nil, core.WrappedError(err, "failed to get datasource ID for song %s", song.GetName())
		}
		r = append(
			r,
			NewSong(
				&myncer_pb.Song{
					Name:             song.GetName(),
					ArtistName:       song.GetArtistNames(),
					AlbumName:        song.GetAlbum(),
					DatasourceSongId: newDatasourceSongId,
				},
			),
		)
	}
	return r, nil
}

func (s *syncEngineImpl) getClient(
	ctx context.Context,
	datasource myncer_pb.Datasource,
) (core.DatasourceClient, error) {
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	switch datasource {
	case myncer_pb.Datasource_SPOTIFY:
		return dsClients.SpotifyClient, nil
	case myncer_pb.Datasource_YOUTUBE:
		return dsClients.YoutubeClient, nil
	default:
		return nil, core.NewError("unsupported datasource: %v", datasource)
	}
}

func getOAuthTokenForDatasource(
	ctx context.Context,
	userId string,
	datasource myncer_pb.Datasource,
) (*myncer_pb.OAuthToken, error) {
	oAuthToken, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetToken(
		ctx,
		userId,
		datasource,
	)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get OAuth token for datasource %v", datasource)
	}
	return oAuthToken, nil
}
