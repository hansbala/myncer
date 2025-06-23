package core

import (
	"context"
	"database/sql"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type SyncRunStore interface {
	CreateSyncRun(ctx context.Context, syncRun *myncer_pb.SyncRun /*const*/) error
	DeleteSyncRun(ctx context.Context, syncRunId string) error
	UpdateSyncRun(ctx context.Context, syncRun *myncer_pb.SyncRun /*const*/) error
	GetSyncs(
		ctx context.Context,
		runIds Set[string], // nil, empty indicates no filtering
		syncIds Set[string], // nil, empty indicates no filtering
	) (Set[*myncer_pb.SyncRun], error)
}

func NewSyncRunStore(db *sql.DB /*const*/) SyncRunStore {
	return &syncRunStoreImpl{
		db: db,
	}
}

type syncRunStoreImpl struct {
	db *sql.DB
}

var _ SyncRunStore = (*syncRunStoreImpl)(nil)

func (s *syncRunStoreImpl) CreateSyncRun(ctx context.Context, syncRun *myncer_pb.SyncRun /*const*/) error {
	return NewError("not implemented")
}

func (s *syncRunStoreImpl) DeleteSyncRun(ctx context.Context, syncRunId string) error {
	return NewError("not implemented")
}

func (s *syncRunStoreImpl) UpdateSyncRun(ctx context.Context, syncRun *myncer_pb.SyncRun /*const*/) error {
	return NewError("not implemented")
}

func (s *syncRunStoreImpl) GetSyncs(
	ctx context.Context,
	runIds Set[string], // nil, empty indicates no filtering
	syncIds Set[string], // nil, empty indicates no filtering
) (Set[*myncer_pb.SyncRun], error) {
	return nil, NewError("not implemented")
}
