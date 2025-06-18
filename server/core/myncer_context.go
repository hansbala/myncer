package core

import (
	"context"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type myncerCtxType struct{}

type MyncerCtx struct {
	DB                *Database          /*const*/
	DatasourceClients *DatasourceClients /*const*/
	Config            *myncer_pb.Config  /*const*/
	LlmClient         LlmClient          /*@nullable*/       // nil if LLM is disabled
	RequestUser       *myncer_pb.User    /*const,@nullable*/ // nil if the user is not authenticated
}

func (m *MyncerCtx) SetRequestUser(user *myncer_pb.User /*const*/) {
	m.RequestUser = user
}

type DatasourceClients struct {
	SpotifyClient DatasourceClient
	YoutubeClient DatasourceClient
}

type LlmClients struct {
	GeminiLlmClient LlmClient
	OpenAILlmClient LlmClient
}

func MustGetMyncerCtx(
	ctx context.Context,
	datasourceClients *DatasourceClients, /*const*/
	llmClients *LlmClients, /*const*/
) *MyncerCtx {
	config := MustGetConfig()
	return &MyncerCtx{
		Config: config,
		DB:     MustGetDatabase(ctx, config),
		DatasourceClients: &DatasourceClients{
			SpotifyClient: datasourceClients.SpotifyClient,
			YoutubeClient: datasourceClients.YoutubeClient,
		},
		LlmClient: MustGetLlmClient(ctx, llmClients, config),
	}
}

func WithMyncerCtx(ctx context.Context, myncerCtx *MyncerCtx) context.Context {
	return context.WithValue(ctx, myncerCtxType{}, myncerCtx)
}

func ToMyncerCtx(ctx context.Context) *MyncerCtx {
	v := ctx.Value(myncerCtxType{})
	if v == nil {
		panic("failed to get myncer ctx")
	}
	res, ok := v.(*MyncerCtx)
	if !ok {
		panic("failed to cast to myncer ctx type")
	}
	return res
}

func MustGetLlmClient(
	ctx context.Context,
	llmClients *LlmClients, /*const*/
	config *myncer_pb.Config, /*const*/
) LlmClient /*@nullable*/ {
	llmConfig := config.GetLlmConfig()
	if !llmConfig.GetEnabled() {
		return nil
	}
	provider := llmConfig.GetPreferredProvider()
	switch provider {
	case myncer_pb.LlmProvider_GEMINI:
		return llmClients.GeminiLlmClient
	case myncer_pb.LlmProvider_OPENAI:
		return llmClients.OpenAILlmClient
	default:
		panic("unsupported LLM provider: " + provider.String())
	}
}
