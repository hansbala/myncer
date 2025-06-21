package services

import (
	"context"

	"connectrpc.com/connect"

	"github.com/hansbala/myncer/core"
)

// OrchestrateHandler services as a compatibility layer between our internally implemented
// handlers and what connectrpc expects. It handles orchestration of our handler framework.
// Namely, makes sure user has perms to execute the request (by means of calling
// `CheckUserPermissions`) and then processing of the actual request (through `ProcessRequest`).
func OrchestrateHandler[RequestT any, ResponseT any](
	ctx context.Context,
	handler core.GrpcHandler[*RequestT, *ResponseT],
	req *RequestT,
) (*connect.Response[ResponseT], error) {
	userInfo := core.ToMyncerCtx(ctx).RequestUser
	if err := handler.CheckUserPermissions(ctx, userInfo, req); err != nil {
		return nil, connect.NewError(
			connect.CodePermissionDenied,
			core.WrappedError(err, "failed to check user permissions"),
		)
	}
	resp := handler.ProcessRequest(ctx, userInfo, req)
	if resp.Err != nil {
		return nil, connect.NewError(
			connect.Code(resp.StatusCode), // TODO: Make sure this works as expected E2E.
			core.WrappedError(resp.Err, "failed to process request"),
		)
	}
	return connect.NewResponse(resp.Response), nil
}
