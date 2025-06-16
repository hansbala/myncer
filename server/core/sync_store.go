package core

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	myncer_pb "github.com/hansbala/myncer/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SyncStore interface {
	CreateSync(ctx context.Context, sync *myncer_pb.Sync /*const*/) error
	GetSyncs(ctx context.Context, userInfo *myncer_pb.User /*const*/) (Set[*myncer_pb.Sync], error)
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

func (s *syncStoreImpl) GetSyncs(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (Set[*myncer_pb.Sync], error) {
	syncs, err := s.getSyncsInternal(ctx, userInfo.GetId())
	if err != nil {
		return nil, WrappedError(err, "failed to get syncs from sql")
	}
	return syncs, nil
}

func (s *syncStoreImpl) getSyncsInternal(
	ctx context.Context,
	userId string, // empty indicates no filtering
) (Set[*myncer_pb.Sync], error) {
	conditions := []string{}
	args := []any{}
	if len(userId) > 0 {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, userId)
	}

	query := `SELECT data, created_at, updated_at FROM syncs`
	if len(conditions) > 0 {
		query += makeWhereAnd(conditions)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, WrappedError(err, "failed to query syncs from sql")
	}
	defer rows.Close()

	r := NewSet[*myncer_pb.Sync]()
	for rows.Next() {
		var (
			sync       myncer_pb.Sync
			protoBytes []byte
			createdAt  time.Time
			updatedAt  time.Time
		)
		if err := rows.Scan(&protoBytes, &createdAt, &updatedAt); err != nil {
			return nil, WrappedError(err, "failed to scan sync row")
		}
		if err := proto.Unmarshal(protoBytes, &sync); err != nil {
			return nil, WrappedError(err, "failed to unmarshal sync proto")
		}
		sync.CreatedAt = timestamppb.New(createdAt)
		sync.UpdatedAt = timestamppb.New(updatedAt)
		r.Add(&sync)
	}
	return r, nil
}
