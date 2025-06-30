package core

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type SongStore interface {
	AddSong(ctx context.Context, song *myncer_pb.Song /*const*/) error
	GetSong(ctx context.Context, id string, datasource myncer_pb.Datasource) (*myncer_pb.Song, error)
}

type songStoreImpl struct{
	db *sql.DB
}

func (s *songStoreImpl) AddSong(ctx context.Context, song *myncer_pb.Song /*const*/) error {
	protoBytes, err := proto.Marshal(song)
	if err != nil {
		return WrappedError(err, "failed to marshal proto song")
	}
	if _, err := s.db.ExecContext(
		ctx,
		`INSERT INTO songs (id, data, datasource, datasourceSongId) VALUES (?, ?, ?, ?)`,
		song.GetId(),
		protoBytes,
		song.GetDatasource().String(),
		song.GetDatasourceSongId(),
	); err != nil {
		return WrappedError(err, "failed to add song to sql database")
	}
	return nil
}

func (s *songStoreImpl) GetSong(
	ctx context.Context, 
	id string, 
	datasource myncer_pb.Datasource,
) (*myncer_pb.Song, error) {
	songs, err := s.getSongsInternal(ctx, NewSet(id), NewSet(datasource))
	if err != nil {
		return nil, WrappedError(err, "failed to get song from sql database")
	}
	if songs.IsEmpty() {
		return nil, NewError("song with id [%s] and datasource [%v] not found", id, datasource)
	}
	if len(songs) > 1 {
		return nil, NewError("multiple songs found with id [%s] and datasource [%v]", id, datasource)
	}
	return songs.ToArray()[0], nil
}

func (s *songStoreImpl) getSongsInternal(
	ctx context.Context,
	ids Set[string], /*const,@nullable*/ // nil, empty indicates no filter.
	datasources Set[myncer_pb.Datasource], /*const,@nullable*/ // nil, empty indicates no filter.
) (Set[*myncer_pb.Song], error) {
	conditions := []string{}
	args := []any{}
	if ids != nil && !ids.IsEmpty() {
		conditions = append(
			conditions, 
			fmt.Sprintf("id IN (%s)", makePlaceholders(len(args), ids.ToArray())),
		)
		for _, id := range ids.ToArray() {
			args = append(args, id)
		}
	}
	if datasources != nil && !datasources.IsEmpty() {
		conditions = append(
			conditions,
			fmt.Sprintf("datasource IN (%s)", makePlaceholders(len(args), datasources.ToArray())),
		)
		for _, ds := range datasources.ToArray() {
			args = append(args, ds.String())
		}
	}
	query := `SELECT data, created_at, updated_at FROM songs`
	if len(conditions) > 0 {
		query += makeWhereAnd(conditions)
	}
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, WrappedError(err, "failed to query songs from sql")
	}
	defer rows.Close()

	songs := NewSet[*myncer_pb.Song]()
	for rows.Next() {
		var (
			protoBytes []byte
			createdAt time.Time
			updatedAt time.Time
		)
		if err := rows.Scan(&protoBytes, createdAt, updatedAt); err != nil {
			return nil, WrappedError(err, "failed to scan song row from sql")
		}
		song := &myncer_pb.Song{}
		if err := proto.Unmarshal(protoBytes, song); err != nil {
			return nil, WrappedError(err, "failed to unmarshal song proto bytes")
		}
		songs.Add(song)
	}
	return songs, nil
}
