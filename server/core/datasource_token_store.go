package core

import (
	"context"
	"database/sql"
	"fmt"
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
	GetTokens(ctx context.Context, userId string) ([]*myncer_pb.OAuthToken, error)
	GetToken(
		ctx context.Context,
		userId string,
		datasource myncer_pb.Datasource,
	) (*myncer_pb.OAuthToken, error)
	GetConnectedDatasources(ctx context.Context, userId string) (Set[myncer_pb.Datasource], error)
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

func (d *datasourceTokenStoreImpl) GetTokens(
	ctx context.Context,
	userId string,
) ([]*myncer_pb.OAuthToken, error) {
	return d.getTokensInternal(ctx, userId, nil /*datasources*/)
}

func (d *datasourceTokenStoreImpl) GetToken(
	ctx context.Context,
	userId string,
	datasource myncer_pb.Datasource,
) (*myncer_pb.OAuthToken, error) {
	tokens, err := d.getTokensInternal(
		ctx,
		userId,
		NewSet(datasource),
	)
	if err != nil {
		return nil, WrappedError(err, "failed to get tokens")
	}
	switch len(tokens) {
	case 0:
		return nil, CErrTokenNotFound
	case 1:
		return tokens[0], nil
	default:
		return nil, NewError("expected to have found one token but multiple were found")
	}
}

func (d *datasourceTokenStoreImpl) GetConnectedDatasources(
	ctx context.Context,
	userId string,
) (Set[myncer_pb.Datasource], error) {
	tokens, err := d.getTokensInternal(ctx, userId, nil /*datasources*/)
	if err != nil {
		return nil, WrappedError(err, "failed to get tokens for user")
	}
	r := NewSet[myncer_pb.Datasource]()
	for _, token := range tokens {
		r.Add(token.GetDatasource())
	}
	return r, nil
}

func (d *datasourceTokenStoreImpl) getTokensInternal(
	ctx context.Context,
	userId string, // empty indicates no filtering
	datasources Set[myncer_pb.Datasource], /*@nullable*/ // nil, empty indicates no filtering
) ([]*myncer_pb.OAuthToken, error) {
	conditions := []string{}
	args := []any{}
	if len(userId) != 0 {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, userId)
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

	query := `SELECT data, created_at, updated_at FROM datasource_tokens`
	if len(conditions) > 0 {
		query += makeWhereAnd(conditions)
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, WrappedError(err, "failed to execute get datasource tokens sql query")
	}
	defer rows.Close()

	r := []*myncer_pb.OAuthToken{}
	for rows.Next() {
		var (
			protoBytes []byte
			createdAt  time.Time
			updatedAt  time.Time
			oAuthToken myncer_pb.OAuthToken
		)
		if err := rows.Scan(&protoBytes, &createdAt, &updatedAt); err != nil {
			return nil, WrappedError(err, "failed to scan datasource token row")
		}
		if err := proto.Unmarshal(protoBytes, &oAuthToken); err != nil {
			return nil, WrappedError(err, "failed to unmarshal oauth token proto")
		}
		oAuthToken.CreatedAt = timestamppb.New(createdAt)
		oAuthToken.UpdatedAt = timestamppb.New(updatedAt)
		r = append(r, &oAuthToken)
	}
	return r, nil
}
