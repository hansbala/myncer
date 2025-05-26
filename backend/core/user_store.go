package core

import (
	"database/sql"

	myncer_pb "github.com/hansbala/myncer/proto"
	"google.golang.org/protobuf/proto"
)

type UserStore interface {
	CreateUser(user *myncer_pb.User /*const*/) error
}

func NewUserStore(db *sql.DB) UserStore {
	return &userStoreImpl{db: db}
}

type userStoreImpl struct {
	db *sql.DB /*const*/
}

func (u *userStoreImpl) CreateUser(user *myncer_pb.User /*const*/) error {
	protoBytes, err := proto.Marshal(user)
	if err != nil {
		return WrappedError(err, "failed to marshal user proto")
	}
	if _, err := u.db.Exec(`
		INSERT INTO users (id, data) VALUES ($1, $2)
  `,
		user.GetId(),
		protoBytes,
	); err != nil {
		return WrappedError(err, "failed to create user in sql")
	}

	return nil
}
