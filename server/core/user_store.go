package core

import (
	"context"
	"database/sql"
	"time"

	myncer_pb "github.com/hansbala/myncer/proto"
	"google.golang.org/protobuf/proto"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *myncer_pb.User /*const*/) error
	GetUsers(ctx context.Context) ([]*myncer_pb.User, error)
}

func NewUserStore(db *sql.DB) UserStore {
	return &userStoreImpl{db: db}
}

type userStoreImpl struct {
	db *sql.DB /*const*/
}

func (u *userStoreImpl) CreateUser(ctx context.Context, user *myncer_pb.User /*const*/) error {
	protoBytes, err := proto.Marshal(user)
	if err != nil {
		return WrappedError(err, "failed to marshal user proto")
	}
	if _, err := u.db.ExecContext(
		ctx,
		`INSERT INTO users (id, data, email) VALUES ($1, $2, $3)`,
		user.GetId(),
		protoBytes,
		user.GetEmail(),
	); err != nil {
		return WrappedError(err, "failed to create user in sql")
	}
	return nil
}

func (u *userStoreImpl) GetUsers(ctx context.Context) ([]*myncer_pb.User, error) {
	users, err := u.getUsersInternal(ctx)
	if err != nil {
		return nil, WrappedError(err, "failed to get users")
	}
	return users, nil
}

func (u *userStoreImpl) getUsersInternal(ctx context.Context) ([]*myncer_pb.User, error) {
	rows, err := u.db.QueryContext(
		ctx,
		`SELECT id, data, created_at, updated_at FROM users`,
	)
	if err != nil {
		return nil, WrappedError(err, "failed to query users")
	}
	users := []*myncer_pb.User{}
	for rows.Next() {
		var (
			id        string
			userProto myncer_pb.User
			data      []byte
			createdAt time.Time
			updatedAt time.Time
		)
		if err := rows.Scan(&id, &data, &createdAt, &updatedAt); err != nil {
			return nil, WrappedError(err, "failed to scan user row")
		}
		if err := proto.Unmarshal(data, &userProto); err != nil {
			return nil, WrappedError(err, "failed to unmarshal user proto data")
		}
		users = append(users, &userProto)
	}
	return users, nil
}
