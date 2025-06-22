package services

import (
	"context"

	"connectrpc.com/connect"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	myncer_pb_connect "github.com/hansbala/myncer/proto/myncer/myncer_pbconnect"
	"github.com/hansbala/myncer/rpc_handlers"
)

func NewDatasourceService() *DatasourceService {
	return &DatasourceService{
		exchangeOAuthCodeHandler: rpc_handlers.NewDatasourceOAuthExchangeHandler(),
		listDatasourcesHandler:   rpc_handlers.NewListDatasourcesHandler(),
	}
}

type DatasourceService struct {
	exchangeOAuthCodeHandler core.GrpcHandler[
		*myncer_pb.ExchangeOAuthCodeRequest,
		*myncer_pb.ExchangeOAuthCodeResponse,
	]
	listDatasourcesHandler core.GrpcHandler[
		*myncer_pb.ListDatasourcesRequest,
		*myncer_pb.ListDatasourcesResponse,
	]
}

var _ myncer_pb_connect.DatasourceServiceHandler = (*DatasourceService)(nil)

func (d *DatasourceService) ExchangeOAuthCode(
	ctx context.Context,
	req *connect.Request[myncer_pb.ExchangeOAuthCodeRequest],
) (*connect.Response[myncer_pb.ExchangeOAuthCodeResponse], error) {
	return OrchestrateHandler(ctx, d.exchangeOAuthCodeHandler, req.Msg)
}

func (d *DatasourceService) ListDatasources(
	ctx context.Context,
	req *connect.Request[myncer_pb.ListDatasourcesRequest],
) (*connect.Response[myncer_pb.ListDatasourcesResponse], error) {
	return OrchestrateHandler(ctx, d.listDatasourcesHandler, req.Msg)
}

func (d *DatasourceService) ListPlaylists(
	context.Context,
	*connect.Request[myncer_pb.ListPlaylistsRequest],
) (*connect.Response[myncer_pb.ListPlaylistsResponse], error) {
	return nil, core.NewError("not implemented")
}

func (d *DatasourceService) GetPlaylistDetails(
	context.Context,
	*connect.Request[myncer_pb.GetPlaylistDetailsRequest],
) (*connect.Response[myncer_pb.GetPlaylistDetailsResponse], error) {
	return nil, core.NewError("not implemented")
}
