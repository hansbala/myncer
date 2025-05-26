package core

import (
	"context"
)

type myncerCtxType struct{}

type MyncerCtx struct {
	DB     *Database /*const*/
	Config *Config   /*const*/
}

func MustGetMyncerCtx(ctx context.Context) *MyncerCtx {
	config := MustGetConfig()
	return &MyncerCtx{
		Config: config,
		DB:     MustGetDatabase(ctx, config),
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
