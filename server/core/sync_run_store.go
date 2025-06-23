package core

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *syncRunStoreImpl) CreateSyncRun(
	ctx context.Context,
	syncRun *myncer_pb.SyncRun, /*const*/
) error {
	protoBytes, err := proto.Marshal(syncRun)
	if err != nil {
		return WrappedError(err, "failed to marshal sync run proto")
	}
	if _, err := s.db.ExecContext(
		ctx,
		`INSERT INTO sync_runs (run_id, sync_id, data) VALUES ($1, $2, $3)`,
		syncRun.GetRunId(),
		syncRun.GetSyncId(),
		protoBytes,
	); err != nil {
		return WrappedError(err, "failed to add sync run into sql")
	}
	return nil
}

func (s *syncRunStoreImpl) DeleteSyncRun(ctx context.Context, syncRunId string) error {
	if _, err := s.db.ExecContext(
		ctx,
		`DELETE FROM sync_runs WHERE run_id = $1`,
		syncRunId,
	); err != nil {
		return WrappedError(err, "failed to delete sync run from sql")
	}
	return nil
}

func (s *syncRunStoreImpl) UpdateSyncRun(
	ctx context.Context,
	syncRun *myncer_pb.SyncRun, /*const*/
) error {
	protoBytes, err := proto.Marshal(syncRun)
	if err != nil {
		return WrappedError(err, "failed to marshal sync run proto")
	}
	if _, err := s.db.ExecContext(
		ctx,
		`UPDATE sync_runs SET data = $1, updated_at = $2 WHERE run_id = $3`,
		protoBytes,
		time.Now(),
		syncRun.GetRunId(),
	); err != nil {
		return WrappedError(err, "failed to update sync run in sql")
	}
	return nil
}

func (s *syncRunStoreImpl) GetSyncs(
	ctx context.Context,
	runIds Set[string], // nil, empty indicates no filtering
	syncIds Set[string], // nil, empty indicates no filtering
) (Set[*myncer_pb.SyncRun], error) {
	conditions := []string{}
	args := []any{}
	if runIds != nil && !runIds.IsEmpty() {
		conditions = append(
			conditions,
			fmt.Sprintf("run_id IN (%s)", makePlaceholders(len(args)+1, runIds.ToArray())),
		)
		for _, runId := range runIds.ToArray() {
			args = append(args, runId)
		}
	}
	if syncIds != nil && !syncIds.IsEmpty() {
		conditions = append(
			conditions,
			fmt.Sprintf("sync_id IN (%s)", makePlaceholders(len(args)+1, syncIds.ToArray())),
		)
		for _, syncId := range syncIds.ToArray() {
			args = append(args, syncId)
		}
	}
	query := "SELECT data, created_at, updated_at FROM sync_runs"
	if len(conditions) > 0 {
		query += makeWhereAnd(conditions)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, WrappedError(err, "failed to query sync runs from sql")
	}
	defer rows.Close()

	syncRuns := NewSet[*myncer_pb.SyncRun]()
	for rows.Next() {
		var (
			protoBytes []byte
			createdAt  time.Time
			updatedAt  time.Time
			syncRun    myncer_pb.SyncRun
		)
		if err := rows.Scan(&protoBytes, &createdAt, &updatedAt); err != nil {
			return nil, WrappedError(err, "failed to scan sync run row")
		}
		if err := proto.Unmarshal(protoBytes, &syncRun); err != nil {
			return nil, WrappedError(err, "failed to unmarshal sync run proto")
		}
		syncRun.CreatedAt = timestamppb.New(createdAt)
		syncRun.UpdatedAt = timestamppb.New(updatedAt)
		syncRuns.Add(&syncRun)
	}

	return syncRuns, nil
}
