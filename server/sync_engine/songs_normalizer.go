package sync_engine

import (
	"context"

	myncer_pb "github.com/hansbala/myncer/proto"
)

// SongsNormalizer is an interface for normalizing song details.
type SongsNormalizer interface {
	// Makes an LLM call to normalize details of the song.
	// Helps subsequent search quality dramatically in respective datasources much better.
	NormalizeSongs(ctx context.Context, songs []*myncer_pb.Song /*const*/) ([]*myncer_pb.Song, error)
}

type llmSongsNormalizerImpl struct{}

func (lsn *llmSongsNormalizerImpl) NormalizeSongs(
	ctx context.Context,
	songs []*myncer_pb.Song, /*const*/
) ([]*myncer_pb.Song, error) {
	// This is a stub implementation.
	// In a real implementation, you would call an LLM to normalize the song details.
	// For now, we just return the songs as-is.
	return songs, nil
}
