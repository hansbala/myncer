package core

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	myncer_pb "github.com/hansbala/myncer/proto"
	"google.golang.org/protobuf/proto"
)

var (
	CUserNotFoundError = NewError("user not found")
)

type UserStore interface {
	CreateUser(ctx context.Context, user *myncer_pb.User /*const*/) error
	GetUsers(ctx context.Context) ([]*myncer_pb.User, error)
	GetUserById(ctx context.Context, id string) (*myncer_pb.User, error)
	GetUserByEmail(ctx context.Context, email string) (*myncer_pb.User, error)
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
	users, err := u.getUsersInternal(ctx, nil /*ids*/, nil /*emails*/)
	if err != nil {
		return nil, WrappedError(err, "failed to get users")
	}
	return users, nil
}

func (u *userStoreImpl) GetUserById(ctx context.Context, id string) (*myncer_pb.User, error) {
	users, err := u.getUsersInternal(ctx, []string{id}, nil /*emails*/)
	if err != nil {
		return nil, WrappedError(err, "failed to get user by id")
	}
	if len(users) != 1 {
		return nil, CUserNotFoundError
	}
	return users[0], nil
}

func (u *userStoreImpl) GetUserByEmail(ctx context.Context, email string) (*myncer_pb.User, error) {
	users, err := u.getUsersInternal(ctx, nil /*ids*/, []string{email})
	if err != nil {
		return nil, WrappedError(err, "failed to get user by email")
	}
	if len(users) != 1 {
		return nil, CUserNotFoundError
	}
	return users[0], nil
}

func (u *userStoreImpl) getUsersInternal(
	ctx context.Context,
	ids []string, /*const,@nullable*/ // nil, empty indicates all.
	emails []string, /*const,@nullable*/ // nil, empty indicates all.
) ([]*myncer_pb.User, error) {
	conditions := []string{}
	args := []any{}
	if len(ids) > 0 {
		conditions = append(
			conditions,
			fmt.Sprintf("id IN (%s)", makePlaceholders(len(args), ids)),
		)
		for _, id := range ids {
			args = append(args, id)
		}
	}
	if len(emails) > 0 {
		conditions = append(
			conditions,
			fmt.Sprintf("email IN (%s)", makePlaceholders(len(args), emails)),
		)
		for _, email := range emails {
			args = append(args, email)
		}
	}
	query := "SELECT id, data, created_at, updated_at FROM users"
	if len(conditions) > 0 {
		query += makeWhereAnd(conditions)
	}
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

func makePlaceholders[T any](start int, vs []T /*const*/) string {
	r := []string{}
	for i := range vs {
		r = append(r, fmt.Sprintf("$%d", start+i+1))
	}
	return strings.Join(r, ", ")
}

func makeWhereAnd(conditions []string /*const*/) string {
	return " WHERE " + strings.Join(conditions, " AND ")
}
