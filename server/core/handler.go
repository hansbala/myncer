package core

import (
	"context"
	"net/http"

	myncer_pb "github.com/hansbala/myncer/proto"
)

type Handler interface {
	GetRequestContainer(ctx context.Context) any /*@nullable*/
	CheckUserPermissions(ctx context.Context, userInfo *myncer_pb.User /*const,@nullable*/) error
	ProcessRequest(
		ctx context.Context,
		reqBody any, /*const,@nullable*/
		req *http.Request, /*const*/
		resp http.ResponseWriter,
	) *ProcessRequestResponse /*const*/
}

type ProcessRequestResponse struct {
	// Plaintext message that will be sent over HTTP (if any).
	// Empty indicates nothing to be written.
	MsgForHttp string
	// Error used for internal logging on server.
	Err error
	// Use http.StatusOk, etc..
	StatusCode int
}

func NewProcessRequestResponse(
	msgForHttp string,
	err error,
	statusCode int,
) *ProcessRequestResponse {
	return &ProcessRequestResponse{
		MsgForHttp: msgForHttp,
		Err:        err,
		StatusCode: statusCode,
	}
}

func NewProcessRequestResponse_OK() *ProcessRequestResponse {
	return &ProcessRequestResponse{
		StatusCode: http.StatusOK,
	}
}

func NewProcessRequestResponse_InternalServerError(err error) *ProcessRequestResponse {
	return &ProcessRequestResponse{
		MsgForHttp: "Internal Server Error",
		Err:        err,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewProcessRequestResponse_BadRequest(err error) *ProcessRequestResponse {
	return &ProcessRequestResponse{
		MsgForHttp: "Bad Request",
		Err:        err,
		StatusCode: http.StatusBadRequest,
	}
}
