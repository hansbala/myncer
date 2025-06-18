package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewCreateSyncHandler() core.Handler {
	return &createSyncHandlerImpl{}
}

type createSyncHandlerImpl struct{}

var _ core.Handler = (*createSyncHandlerImpl)(nil)

func (cs *createSyncHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return &api.CreateSyncRequest{}
}

func (cs *createSyncHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to create a sync")
	}
	return nil
}

func (cs *createSyncHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	restReq, ok := (reqBody).(*api.CreateSyncRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("request expected to be CreateSyncRequest but got %T", reqBody),
		)
	}
	if err := cs.validateRequest(ctx, userInfo, restReq); err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "failed to validate create sync request"),
		)
	}

	// Create the sync from the request.
	sync, err := cs.createSyncFromRequest(restReq, userInfo)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to create sync from rest request"),
		)
	}

	// Persist the sync to the database.
	if err := core.ToMyncerCtx(ctx).DB.SyncStore.CreateSync(ctx, sync); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to create sync in database"),
		)
	}

	// Write a 201 back indicating we created the sync.
	return core.NewProcessRequestResponse(
		"Sync created successfully", /*msgForHttp*/
		nil, /*err*/
		http.StatusCreated,
	)
}

func (cs *createSyncHandlerImpl) validateRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	req *api.CreateSyncRequest, /*const*/
) error {
	switch v := req.GetActualInstance().(type) {
	case *api.OneWaySync:
		return validateOneWaySync(ctx, userInfo, v)
	default:
		return core.NewError("unknown sync type in validate request")
	}
}

func (cs *createSyncHandlerImpl) createSyncFromRequest(
	req *api.CreateSyncRequest, /*const*/
	userInfo *myncer_pb.User, /*const*/
) (*myncer_pb.Sync, error) {
	switch v := req.GetActualInstance().(type) {
	case *api.OneWaySync:
		return rest_helpers.RestOneWaySyncToProto(v, userInfo.GetId()), nil
	default:
		return nil, core.NewError("unknown sync type in create sync from request")
	}
}

func validateOneWaySync(
	ctx context.Context, 
	userInfo *myncer_pb.User, /*const*/
	s *api.OneWaySync, /*const*/
) error {
	// Validate the source and destination datasources are valid.
	connectedDatasources, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetConnectedDatasources(
		ctx,
		userInfo.GetId(),
	)
	if err != nil {
		return core.WrappedError(err, "failed to get connected datasources for user")
	}

	sourceProtoDs := rest_helpers.RestDatasourceToProto(s.GetSource().Datasource)
	if sourceProtoDs == myncer_pb.Datasource_DATASOURCE_UNSPECIFIED {
		return core.NewError("failed to get proto ds from %s", s.GetSource().Datasource)
	}
	destProtoDs := rest_helpers.RestDatasourceToProto(s.GetDestination().Datasource)
	if destProtoDs == myncer_pb.Datasource_DATASOURCE_UNSPECIFIED {
			return core.NewError("failed to get proto ds from %s", s.GetDestination().Datasource)
	}
	if !connectedDatasources.Contains(sourceProtoDs) {
		return core.NewError("source datasource %v is not connected for user", sourceProtoDs)
	}
	if !connectedDatasources.Contains(destProtoDs) {
		return core.NewError("destination datasource %v is not connected for user", destProtoDs)
	}
	// TODO: Make sure playlist ids are valid for the datasource.
	return nil
}
