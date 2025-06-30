package core

import (
	"context"
	"database/sql"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type SongStore interface {
	AddSong(ctx context.Context, song *myncer_pb.Song /*const*/) error
	GetSong(ctx context.Context, id string) (*myncer_pb.Song, error)
}

type songStoreImpl struct{
	db *sql.DB
}

func (s *songStoreImpl) AddSong(ctx context.Context, song *myncer_pb.Song /*const*/) error {
	return nil
}

func (s *songStoreImpl) GetSong(ctx context.Context, id string) (*myncer_pb.Song, error) {
	return nil, nil
}
