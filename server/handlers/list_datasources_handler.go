package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewListDatasourcesHandler() core.Handler {
	return &listDatasourcesHandlerImpl{}
}

type listDatasourcesHandlerImpl struct{}

var _ core.Handler = (*listDatasourcesHandlerImpl)(nil)

func (ld *listDatasourcesHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return nil
}

func (ld *listDatasourcesHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required to list datasources")
	}
	return nil
}

func (ld *listDatasourcesHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	tokens, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetTokens(ctx, userInfo.GetId())
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to get auth tokens for user"),
		)
	}

	connectedDatasources := []api.Datasource{}
	for _, token := range tokens {
		protoDs := token.GetDatasource()
		restDs, err := rest_helpers.ProtoDatasourceToRest(protoDs)
		if err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to get rest datasource from %v", protoDs),
			)
		}
		connectedDatasources = append(connectedDatasources, restDs)
	}

	if err := WriteJSONOk(resp, api.NewListDatasourcesResponse(connectedDatasources)); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.NewError("failed to write connected datasources response"),
		)
	}

	return core.NewProcessRequestResponse_OK()
}
