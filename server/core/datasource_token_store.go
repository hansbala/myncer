package core

import (
	"context"
	"database/sql"

	myncer_pb "github.com/hansbala/myncer/proto"
	"google.golang.org/protobuf/proto"
)

type DatasourceTokenStore interface {
	AddToken(ctx context.Context, oAuthToken *myncer_pb.OAuthToken /*const*/) error
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
