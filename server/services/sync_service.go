package services

import (
	"context"

	"connectrpc.com/connect"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	myncer_pb_connect "github.com/hansbala/myncer/proto/myncer/myncer_pbconnect"
	"github.com/hansbala/myncer/rpc_handlers"
	"github.com/hansbala/myncer/sync_engine"
)

func NewSyncService() *SyncService {
	return &SyncService{
		createSyncHandler:   rpc_handlers.NewCreateSyncHandler(),
		deleteSyncHandler:   rpc_handlers.NewDeleteSyncHandler(),
		listSyncsHandler:    rpc_handlers.NewListSyncsHandler(),
		getSyncHandler:      rpc_handlers.NewGetSyncHandler(),
		runSyncHandler:      rpc_handlers.NewRunSyncHandler(sync_engine.NewSyncEngine()),
		listSyncRunsHandler: rpc_handlers.NewListSyncRunsHandler(),
	}
}

type SyncService struct {
	createSyncHandler   core.GrpcHandler[*myncer_pb.CreateSyncRequest, *myncer_pb.CreateSyncResponse]
	deleteSyncHandler   core.GrpcHandler[*myncer_pb.DeleteSyncRequest, *myncer_pb.DeleteSyncResponse]
	listSyncsHandler    core.GrpcHandler[*myncer_pb.ListSyncsRequest, *myncer_pb.ListSyncsResponse]
	getSyncHandler      core.GrpcHandler[*myncer_pb.GetSyncRequest, *myncer_pb.GetSyncResponse]
	runSyncHandler      core.GrpcHandler[*myncer_pb.RunSyncRequest, *myncer_pb.RunSyncResponse]
	listSyncRunsHandler core.GrpcHandler[
		*myncer_pb.ListSyncRunsRequest,
		*myncer_pb.ListSyncRunsResponse,
	]
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
	return OrchestrateHandler(ctx, d.deleteSyncHandler, req.Msg)
}

func (d *SyncService) ListSyncs(
	ctx context.Context,
	req *connect.Request[myncer_pb.ListSyncsRequest], /*const*/
) (*connect.Response[myncer_pb.ListSyncsResponse], error) {
	return OrchestrateHandler(ctx, d.listSyncsHandler, req.Msg)
}

func (d *SyncService) GetSync(
	ctx context.Context,
	req *connect.Request[myncer_pb.GetSyncRequest], /*const*/
) (*connect.Response[myncer_pb.GetSyncResponse], error) {
	return OrchestrateHandler(ctx, d.getSyncHandler, req.Msg)
}

func (d *SyncService) RunSync(
	ctx context.Context,
	req *connect.Request[myncer_pb.RunSyncRequest], /*const*/
) (*connect.Response[myncer_pb.RunSyncResponse], error) {
	return OrchestrateHandler(ctx, d.runSyncHandler, req.Msg)
}

func (d *SyncService) ListSyncRuns(
	ctx context.Context,
	req *connect.Request[myncer_pb.ListSyncRunsRequest], /*const*/
) (*connect.Response[myncer_pb.ListSyncRunsResponse], error) {
	return OrchestrateHandler(ctx, d.listSyncRunsHandler, req.Msg)
}
