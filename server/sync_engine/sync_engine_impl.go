package sync_engine

import (
	"context"
	"fmt"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewSyncEngine(
	spotifyClient core.DatasourceClient,
	youtubeClient core.DatasourceClient,
) core.SyncEngine {
	return &syncEngineImpl{
		spotifyClient: spotifyClient,
		youtubeClient: youtubeClient,
	}
}

type syncEngineImpl struct {
	spotifyClient core.DatasourceClient
	youtubeClient core.DatasourceClient
}

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
	sourceClient, err := s.getClient(sync.GetSource().GetDatasource())
	if err != nil {
		return err
	}
	destClient, err := s.getClient(sync.GetDestination().GetDatasource())
	if err != nil {
		return err
	}

	// Fetch songs from source playlist
	sourceOAuthToken, err := getOAuthTokenForDatasource(
		ctx,
		userInfo.GetId(),
		sync.GetSource().GetDatasource(),
	)
	if err != nil {
		return core.WrappedError(err, "failed to get OAuth token for source datasource")
	}
	sourceSongs, err := sourceClient.GetPlaylistSongs(
		ctx,
		sync.GetSource().GetPlaylistId(),
		sourceOAuthToken,
	)
	if err != nil {
		return core.WrappedError(err, "failed to fetch source playlist")
	}

	searchedSongs, err := s.getSearchedSongs(
		ctx,
		sourceSongs,
		sync.GetDestination().GetDatasource(),
	)
	if err != nil {
		return core.WrappedError(err, "failed to get searched songs for destination datasource")
	}

	// Optionally clear destination playlist
	destOAuthToken, err := getOAuthTokenForDatasource(
		ctx,
		userInfo.GetId(),
		sync.GetDestination().GetDatasource(),
	)
	if err != nil {
		return core.WrappedError(err, "failed to get OAuth token for destination datasource")
	}
	destPlaylistId := sync.GetDestination().GetPlaylistId()
	if sync.OverwriteExisting {
		core.Printf("Clearing destination playlist")
		if err := destClient.ClearPlaylist(
			ctx,
			destOAuthToken,
			destPlaylistId,
		); err != nil {
			return core.WrappedError(err, "failed to clear destination playlist")
		}
	}

	// Add source songs to destination
	if err := destClient.AddToPlaylist(
		ctx,
		destOAuthToken,
		destPlaylistId,
		searchedSongs,
	); err != nil {
		return core.WrappedError(err, "failed to add songs to destination playlist")
	}
	return nil
}

func (s *syncEngineImpl) getSearchedSongs(
	ctx context.Context,
	songs []core.Song, /*const*/
	datasource myncer_pb.Datasource, /*const*/
) ([]core.Song, error) {
	r := []core.Song{}
	for _, song := range songs {
		newDatasourceSongId, err := song.GetDatasourceId(datasource)
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
				s.spotifyClient,
				s.youtubeClient,
			),
		)
	}
	return r, nil
}

func (s *syncEngineImpl) getClient(datasource myncer_pb.Datasource) (core.DatasourceClient, error) {
	switch datasource {
	case myncer_pb.Datasource_SPOTIFY:
		return s.spotifyClient, nil
	case myncer_pb.Datasource_YOUTUBE:
		return s.youtubeClient, nil
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
