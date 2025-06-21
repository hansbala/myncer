package core

import (
	"context"
	"net/http"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type GrpcHandler[Req any, Resp any] interface {
	CheckUserPermissions(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const,@nullable*/
		reqBody Req, /*const*/
	) error
	ProcessRequest(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const,@nullable*/
		reqBody Req, /*const*/
	) *GrpcHandlerResponse[Resp]
}

// TODO: Convert to builder pattern to make it more flexible and easier to use.
type GrpcHandlerResponse[T any] struct {
	// Error used for internal logging on server.
	Err error /*@nullable*/
	// Any cookies that need to be set in the response.
	Cookies []*http.Cookie /*@nullable*/
	// To support HTTP status codes.
	StatusCode int
	// Actual response.
	Response T
}

func NewGrpcHandlerResponse_BadRequest[T any](err error) *GrpcHandlerResponse[T] {
	return &GrpcHandlerResponse[T]{
		Err:        err,
		StatusCode: http.StatusBadRequest,
	}
}

func NewGrpcHandlerResponse_Unauthorized[T any](err error) *GrpcHandlerResponse[T] {
	return &GrpcHandlerResponse[T]{
		Err:        err,
		StatusCode: http.StatusUnauthorized,
	}
}

func NewGrpcHandlerResponse_InternalServerError[T any](err error) *GrpcHandlerResponse[T] {
	return &GrpcHandlerResponse[T]{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewGrpcHandlerResponse_WithCookies[T any](
	resp T,
	cookies []*http.Cookie, /*@nullable*/
) *GrpcHandlerResponse[T] {
	return &GrpcHandlerResponse[T]{
		Response:   resp,
		Cookies:    cookies,
		StatusCode: http.StatusOK,
	}
}

func NewGrpcHandlerResponse_OK[T any](resp T) *GrpcHandlerResponse[T] {
	return &GrpcHandlerResponse[T]{
		Response:   resp,
		StatusCode: http.StatusOK,
	}
}

