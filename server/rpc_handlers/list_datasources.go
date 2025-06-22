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
	return nil
}

func (l *listDatasourcesImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.ListDatasourcesRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ListDatasourcesResponse] {
	return nil
}
