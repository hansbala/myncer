package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewListSyncsHandler() core.Handler {
	return &listSyncsHandlerImpl{}
}

type listSyncsHandlerImpl struct{}

var _ core.Handler = (*listSyncsHandlerImpl)(nil)

func (ls *listSyncsHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil // No request body since it's a GET request.
}

func (ls *listSyncsHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list syncs")
	}
	return nil
}

func (ls *listSyncsHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	// Get all syncs for current user.
	syncs, err := core.ToMyncerCtx(ctx).DB.SyncStore.GetSyncs(ctx, userInfo)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to get syncs for current user"),
		)
	}

	// Convert to rest.
	restSyncs := []api.Sync{}
	for _, protoSync := range syncs.ToArray() {
		restSync, err := rest_helpers.ProtoSyncToRest(protoSync)
		if err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to convert proto sync to rest"),
			)
		}
		restSyncs = append(restSyncs, *restSync)
	}

	// Write to response.
	if err := WriteJSONOk(resp, api.NewListSyncsResponse(restSyncs)); err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to write list syncs response"),
			)
	}
	return core.NewProcessRequestResponse_OK()
}
