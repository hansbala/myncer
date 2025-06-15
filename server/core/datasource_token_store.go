package core

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	myncer_pb "github.com/hansbala/myncer/proto"
)

var (
	CErrTokenNotFound = NewError("oauth token not found")
)

type DatasourceTokenStore interface {
	AddToken(ctx context.Context, oAuthToken *myncer_pb.OAuthToken /*const*/) error
	GetToken(
		ctx context.Context,
		userId string,
		datasource myncer_pb.Datasource,
	) (*myncer_pb.OAuthToken, error)
}

func NewDatasourceTokenStore(db *sql.DB) DatasourceTokenStore {
	return &datasourceTokenStoreImpl{db: db}
}

type datasourceTokenStoreImpl struct {
	db *sql.DB
}

func (d *datasourceTokenStoreImpl) AddToken(
	ctx context.Context,
	oAuthToken *myncer_pb.OAuthToken, /*const*/
) error {
	protoBytes, err := proto.Marshal(oAuthToken)
	if err != nil {
		return WrappedError(err, "failed to marshal oauth token proto")
	}

	if _, err := d.db.ExecContext(
		ctx,
		`INSERT INTO datasource_tokens (id, user_id, datasource, data) VALUES ($1, $2, $3, $4)`,
		oAuthToken.GetId(),
		oAuthToken.GetUserId(),
		oAuthToken.GetDatasource().String(),
		protoBytes,
	); err != nil {
		return WrappedError(err, "failed to create oauth token in sql")
	}
	return nil
}

func (d *datasourceTokenStoreImpl) GetToken(
	ctx context.Context,
	userId string,
	datasource myncer_pb.Datasource,
) (*myncer_pb.OAuthToken, error) {
	rows := d.db.QueryRowContext(
		ctx,
		`SELECT data, created_at, updated_at
		FROM datasource_tokens
		WHERE user_id = $1 AND datasource = $2 LIMIT 1`,
		userId,
		datasource.String(),
	)

	var (
		protoBytes []byte
		createdAt  time.Time
		updatedAt  time.Time
		oAuthToken myncer_pb.OAuthToken
	)
	if err := rows.Scan(&protoBytes, &createdAt, &updatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, CErrTokenNotFound
		}
		return nil, WrappedError(err, "failed to scan datasource token row")
	}
	if err := proto.Unmarshal(protoBytes, &oAuthToken); err != nil {
		return nil, WrappedError(err, "failed to unmarshal oauth token proto")
	}
	oAuthToken.CreatedAt = timestamppb.New(createdAt)
	oAuthToken.UpdatedAt = timestamppb.New(updatedAt)
	return &oAuthToken, nil
}
