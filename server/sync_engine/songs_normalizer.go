package sync_engine

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

//go:embed normalizer_system.prompt
var cNormalizerSystemPrompt embed.FS

// SongsNormalizer is an interface for normalizing song details.
type SongsNormalizer interface {
	// Makes an LLM call to normalize details of the song.
	// Helps subsequent search quality dramatically in datasources.
	NormalizeSongs(ctx context.Context, songs *core.SongList) (*core.SongList, error)
}

func NewLlmSongsNormalizer() SongsNormalizer {
	return &llmSongsNormalizerImpl{}
}

var _ SongsNormalizer = (*llmSongsNormalizerImpl)(nil)

type llmSongsNormalizerImpl struct{}

func (lsn *llmSongsNormalizerImpl) NormalizeSongs(
	ctx context.Context,
	songs *core.SongList,
) (*core.SongList, error) {
	// Prompt construction.
	systemPrompt, err := lsn.getSystemPrompt()
	if err != nil {
		return nil, core.WrappedError(err, "failed to get system prompt")
	}
	userPrompt, err := lsn.getUserPrompt(songs)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get user prompt")
	}

	// Send to LLM to figure out.
	llmResponse, err := core.ToMyncerCtx(ctx).LlmClient.GetResponse(ctx, systemPrompt, userPrompt)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get normalizer llm response")
	}
	// Parse LLM response.
	normalizedSongs, err := lsn.parseLlmResponse(llmResponse)
	if err != nil {
		core.Errorf(fmt.Sprintf("failed to parse llm response: [%s]", llmResponse))
		return nil, core.WrappedError(err, "failed to parse normalizer llm response")
	}

	return normalizedSongs, nil
}

func (lsn *llmSongsNormalizerImpl) getSystemPrompt() (string, error) {
	bytes, err := cNormalizerSystemPrompt.ReadFile("normalizer_system.prompt")
	if err != nil {
		return "", core.WrappedError(err, "failed to read normalizer system prompt")
	}
	return string(bytes), nil
}

func (lsn *llmSongsNormalizerImpl) getUserPrompt(
	songs *core.SongList,
) (string, error) {
	bytes, err := songs.GetLlmJson()
	if err != nil {
		return "", core.WrappedError(err, "failed to marshal songs as JSON")
	}
	return string(bytes), nil
}

func (lsn *llmSongsNormalizerImpl) parseLlmResponse(llmResponse string) (*core.SongList, error) {
	llmResponse = cleanseJsonBeginAndEndTags(llmResponse)
	songs := []*myncer_pb.Song{}
	if err := json.Unmarshal([]byte(llmResponse), &songs); err != nil {
		return nil, core.WrappedError(err, "failed to unmarshal json from llm")
	}
	parsed := []core.Song{}
	for _, song := range songs {
		parsed = append(parsed, NewSong(song))
	}
	return core.NewSongList(parsed), nil
}

func cleanseJsonBeginAndEndTags(i string) string {
	o, _ := strings.CutPrefix(i, "```json")
	o, _ = strings.CutSuffix(o, "```")
	return o
}
