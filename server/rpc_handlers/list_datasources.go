package rpc_handlers

import (
	"context"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewListDatasourcesHandler() core.GrpcHandler[
	*myncer_pb.ListDatasourcesRequest,
	*myncer_pb.ListDatasourcesResponse,
] {
	return &listDatasourcesImpl{}
}

type listDatasourcesImpl struct{}

func (l *listDatasourcesImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ListDatasourcesRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list datasources")
	}
	return nil
}

func (l *listDatasourcesImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.ListDatasourcesRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListDatasourcesResponse] {
	tokens, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetTokens(ctx, userInfo.GetId())
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ListDatasourcesResponse](
			core.WrappedError(err, "failed to get auth tokens for user"),
		)
	}

	connectedDatasources := []myncer_pb.Datasource{}
	for _, token := range tokens {
		connectedDatasources = append(connectedDatasources, token.GetDatasource())
	}

	return core.NewGrpcHandlerResponse_OK(
		&myncer_pb.ListDatasourcesResponse{
			Datasources: connectedDatasources,
		},
	)
}
