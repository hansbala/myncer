package core

import (
	"context"
)

type myncerCtxType struct{}

type MyncerCtx struct {
	DB                *Database          /*const*/
	DatasourceClients *DatasourceClients /*const*/
	Config            *Config            /*const*/
	LlmClient         LlmClient          /*@nullable*/ // nil if LLM is disabled
}

type DatasourceClients struct {
	SpotifyClient DatasourceClient
	YoutubeClient DatasourceClient
}

func MustGetMyncerCtx(
	ctx context.Context,
	spotifyClient DatasourceClient,
	youtubeClient DatasourceClient,
) *MyncerCtx {
	config := MustGetConfig()
	return &MyncerCtx{
		Config: config,
		DB:     MustGetDatabase(ctx, config),
		DatasourceClients: &DatasourceClients{
			SpotifyClient: spotifyClient,
			YoutubeClient: youtubeClient,
		},
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
