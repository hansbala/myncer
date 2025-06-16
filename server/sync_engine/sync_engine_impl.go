package sync_engine

import (
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewSyncEngine(spotifyClient, youtubeClient core.DatasourceClient) core.SyncEngine {
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

func (s *syncEngineImpl) RunSync(sync *myncer_pb.Sync) error {
	// TODO: Implement.
	core.Printf("Running sync in engine...")
	core.DebugPrintJson(sync)
	return nil
}
