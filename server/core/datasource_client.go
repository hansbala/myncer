package core

import (
	"context"

	"golang.org/x/oauth2"

	myncer_pb "github.com/hansbala/myncer/proto"
)

type DatasourceClient interface {
	ExchangeCodeForToken(ctx context.Context, authCode string) (*oauth2.Token, error)
	GetPlaylists(
		ctx context.Context,
		oAuthToken *myncer_pb.OAuthToken, /*const*/
	) ([]*myncer_pb.Playlist, error)
}
