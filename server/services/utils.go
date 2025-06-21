package services

import (
	"context"

	"connectrpc.com/connect"

	"github.com/hansbala/myncer/auth"
	"github.com/hansbala/myncer/core"
)

// OrchestrateHandler services as a compatibility layer between our internally implemented
// handlers and what connectrpc expects. It handles orchestration of our handler framework.
// Namely, makes sure user has perms to execute the request (by means of calling
// `CheckUserPermissions`) and then processing of the actual request (through `ProcessRequest`).
func OrchestrateHandler[RequestT any, ResponseT any](
	ctx context.Context,
	handler core.GrpcHandler[*RequestT, *ResponseT],
	reqBody *RequestT,
) (*connect.Response[ResponseT], error) {
	userInfo := auth.UserFromContext(ctx)
	if err := handler.CheckUserPermissions(ctx, userInfo, reqBody); err != nil {
		core.Printf("failed to check user permissions: %v", err)
		return nil, connect.NewError(
			connect.CodePermissionDenied,
			core.WrappedError(err, "failed to check user permissions"),
		)
	}
	resp := handler.ProcessRequest(ctx, userInfo, reqBody)
	if resp.Err != nil {
		return nil, connect.NewError(
			connect.Code(resp.StatusCode), // TODO: Make sure this works as expected E2E.
			core.WrappedError(resp.Err, "failed to process request"),
		)
	}
	connectResp := connect.NewResponse(resp.Response)
	if len(resp.Cookies) > 0 {
		for _, cookie := range resp.Cookies {
			connectResp.Header().Set("Set-Cookie", cookie.String())
		}
	}
	return connectResp, nil
}
