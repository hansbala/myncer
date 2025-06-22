package services

import (
	"context"

	"connectrpc.com/connect"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	myncer_pb_connect "github.com/hansbala/myncer/proto/myncer/myncer_pbconnect"
	"github.com/hansbala/myncer/rpc_handlers"
)

func NewSyncService() *SyncService {
	return &SyncService{
		createSyncHandler: rpc_handlers.NewCreateSyncHandler(),
	}
}

type SyncService struct{
	createSyncHandler  core.GrpcHandler[*myncer_pb.CreateSyncRequest, *myncer_pb.CreateSyncResponse]
}

var _ myncer_pb_connect.SyncServiceHandler = (*SyncService)(nil)

func (d *SyncService) CreateSync(
	ctx context.Context,
	req *connect.Request[myncer_pb.CreateSyncRequest], /*const*/
) (*connect.Response[myncer_pb.CreateSyncResponse], error) {
	return OrchestrateHandler(ctx, d.createSyncHandler, req.Msg)
}

func (d *SyncService) DeleteSync(
	ctx context.Context,
	req *connect.Request[myncer_pb.DeleteSyncRequest], /*const*/
) (*connect.Response[myncer_pb.DeleteSyncResponse], error) {
	return nil, core.NewError("not implemented")
}

func (d *SyncService) ListSyncs(
	ctx context.Context,
	req *connect.Request[myncer_pb.ListSyncsRequest], /*const*/
) (*connect.Response[myncer_pb.ListSyncsResponse], error) {
	return nil, core.NewError("not implemented")
}

func (d *SyncService) RunSync(
	ctx context.Context,
	req *connect.Request[myncer_pb.RunSyncRequest], /*const*/
) (*connect.Response[myncer_pb.RunSyncResponse], error) {
	return nil, core.NewError("not implemented")
}
