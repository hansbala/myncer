package rpc_handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func NewCreateSyncHandler() core.GrpcHandler[
	*myncer_pb.CreateSyncRequest,
	*myncer_pb.CreateSyncResponse,
] {
	return &createSyncImpl{}
}

type createSyncImpl struct{}

func (cs *createSyncImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CreateSyncRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to create a sync")
	}
	return nil
}

func (cs *createSyncImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.CreateSyncRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.CreateSyncResponse] {
	if err := cs.validateRequest(ctx, reqBody, userInfo); err != nil {
		return core.NewGrpcHandlerResponse_BadRequest[*myncer_pb.CreateSyncResponse](
			core.WrappedError(err, "failed to validate create sync request"),
		)
	}

	// Create the sync from the request.
	sync, err := cs.createSyncFromRequest(reqBody, userInfo)
	if err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.CreateSyncResponse](
			core.WrappedError(err, "failed to create sync from request"),
		)
	}

	// Persist the sync to the database.
	if err := core.ToMyncerCtx(ctx).DB.SyncStore.CreateSync(ctx, sync); err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.CreateSyncResponse](
			core.WrappedError(err, "failed to create sync in database"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(&myncer_pb.CreateSyncResponse{Sync: sync})
}

func (cs *createSyncImpl) validateRequest(
	ctx context.Context,
	req *myncer_pb.CreateSyncRequest, /*const*/
	userInfo *myncer_pb.User, /*const*/
) error {
	syncVariant := req.GetSyncVariant()
	switch syncVariant.(type) {
	case *myncer_pb.CreateSyncRequest_OneWaySync:
		return validateOneWaySync(ctx, userInfo, req.GetOneWaySync())
	default:
		return core.NewError("unknown sync type in validate request: %T", syncVariant)
	}
}

func (cs *createSyncImpl) createSyncFromRequest(
	req *myncer_pb.CreateSyncRequest, /*const*/
	userInfo *myncer_pb.User, /*const*/
) (*myncer_pb.Sync, error) {
	syncVariant := req.GetSyncVariant()
	switch syncVariant.(type) {
	case *myncer_pb.CreateSyncRequest_OneWaySync:
		return NewSync_OneWaySync(userInfo.GetId(), req.GetOneWaySync()), nil
	default:
		return nil, core.NewError("unknown sync type in create sync from request: %T", syncVariant)
	}
}

func validateOneWaySync(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	req *myncer_pb.OneWaySync, /*const*/
) error {
	// Validate the source and destination datasources are valid.
	if req.GetSource().GetDatasource() == myncer_pb.Datasource_DATASOURCE_UNSPECIFIED {
		return core.NewError("source datasource must be specified")
	}
	if req.GetDestination().GetDatasource() == myncer_pb.Datasource_DATASOURCE_UNSPECIFIED {
		return core.NewError("destination datasource must be specified")
	}
	// Check if the user has connected the source and destination datasources.
	connectedDatasources, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetConnectedDatasources(
		ctx,
		userInfo.GetId(),
	)
	if err != nil {
		return core.WrappedError(err, "failed to get connected datasources for user")
	}
	if !connectedDatasources.Contains(req.GetSource().GetDatasource()) {
		return core.NewError("source datasource is not connected")
	}
	if !connectedDatasources.Contains(req.GetDestination().GetDatasource()) {
		return core.NewError("destination datasource is not connected")
	}
	// Basic playlist id checks.
	if len(req.GetSource().GetPlaylistId()) == 0 {
		return core.NewError("source playlist id must be specified")
	}
	if len(req.GetDestination().GetPlaylistId()) == 0 {
		return core.NewError("destination playlist id must be specified")
	}
	return nil
}

func NewSync_OneWaySync(
	userId string, /*const*/
	oneWaySync *myncer_pb.OneWaySync, /*const*/
) *myncer_pb.Sync {
	return &myncer_pb.Sync{
		Id:     uuid.NewString(),
		UserId: userId,
		SyncVariant: &myncer_pb.Sync_OneWaySync{
			OneWaySync: oneWaySync,
		},
	}
}
