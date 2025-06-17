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
		userInfo *myncer_pb.User, /*const*/
	) ([]*myncer_pb.Playlist, error)
	GetPlaylist(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		id string, /*const*/
	) (*myncer_pb.Playlist, error)
	GetPlaylistSongs(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		playlistId string, /*const*/
	) ([]Song, error)
	AddToPlaylist(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		playlistId string, /*const*/
		songs []Song, /*const*/
	) error
	ClearPlaylist(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		playlistId string, /*const*/
	) error
	Search(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		names Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
		artistNames Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
		albumNames Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
	) (Song, error)
}
