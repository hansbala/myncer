package core

import (
	"context"
	"database/sql"

	myncer_pb "github.com/hansbala/myncer/proto"
	"google.golang.org/protobuf/proto"
)

type SyncStore interface {
	CreateSync(ctx context.Context, sync *myncer_pb.Sync /*const*/) error
}

func NewSyncStore(db *sql.DB) SyncStore {
	return &syncStoreImpl{db: db}
}

type syncStoreImpl struct {
	db *sql.DB
}

var _ SyncStore = (*syncStoreImpl)(nil)

func (s *syncStoreImpl) CreateSync(ctx context.Context, sync *myncer_pb.Sync /*const*/) error {
	protoBytes, err := proto.Marshal(sync)
	if err != nil {
		return WrappedError(err, "failed to marshal sync proto")
	}
	if _, err := s.db.ExecContext(
		ctx,
		`INSERT INTO syncs (id, user_id, data) VALUES ($1, $2, $3)`,
		sync.GetId(),
		sync.GetUserId(),
		protoBytes,
	); err != nil {
		return WrappedError(err, "failed to create sync in sql")
	}
	return nil
}
