package main

import (
	"context"

	"github.com/hansbala/myncer/core"
)

func main() {
	ctx := context.Background()
	ctx = core.WithMyncerCtx(ctx, core.MustGetMyncerCtx(ctx))
	for {
	}
}
