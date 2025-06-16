package core

import (
	"context"

	"golang.org/x/oauth2"

	myncer_pb "github.com/hansbala/myncer/proto"
)

// TODO: We can make oauth token as part of the impls (get method) so we can cleanup here.
// Callsites shouldn't have to care about passing the OAuth token around.
type DatasourceClient interface {
	ExchangeCodeForToken(ctx context.Context, authCode string) (*oauth2.Token, error)
	GetPlaylists(
		ctx context.Context,
		oAuthToken *myncer_pb.OAuthToken, /*const*/
	) ([]*myncer_pb.Playlist, error)
	GetPlaylist(
		ctx context.Context,
		id string, /*const*/
		oAuthToken *myncer_pb.OAuthToken, /*const*/
	) (*myncer_pb.Playlist, error)
	GetPlaylistSongs(
		ctx context.Context,
		playlistId string, /*const*/
		oAuthToken *myncer_pb.OAuthToken, /*const*/
	) ([]Song, error)
	AddToPlaylist(
		ctx context.Context,
		oAuthToken *myncer_pb.OAuthToken, /*const*/
		playlistId string, /*const*/
		songs []Song, /*const*/
	) error
	ClearPlaylist(
		ctx context.Context,
		oAuthToken *myncer_pb.OAuthToken, /*const*/
		playlistId string, /*const*/
	) error
	Search(
		ctx context.Context,
		oAuthToken *myncer_pb.OAuthToken, /*const*/
		names Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
		artistNames Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
		albumNames Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
	) (Song, error)
}
