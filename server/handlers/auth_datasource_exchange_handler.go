package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/rest_helpers"
)

func NewAuthExchangeHandler() core.Handler {
	return &authExchangeHandlerImpl{}
}

type authExchangeHandlerImpl struct{}

var _ core.Handler = (*authExchangeHandlerImpl)(nil)

func (aeh *authExchangeHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	return &api.OAuthExchangeRequest{}
}

func (aeh *authExchangeHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody any, /*const,@nullable*/
) error {
	if userInfo == nil {
		return core.NewError("user is required for oauth exchange")
	}
	return nil
}

func (aeh *authExchangeHandlerImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse /*const*/ {
	restReq, ok := reqBody.(*api.OAuthExchangeRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("expected OAuthExchangeRequest but got %T", reqBody),
		)
	}

	if err := aeh.validateRequest(restReq); err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "request failed validation"),
		)
	}

	datasource, err := aeh.getDatasource(req)
	if err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "could not determine datasource for oauth exchange"),
		)
	}

	// Based on the datasource get the client, authenticate, and build the auth token.
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	var oAuthToken *myncer_pb.OAuthToken
	switch datasource {
	case myncer_pb.Datasource_DATASOURCE_SPOTIFY:
		token, err := dsClients.SpotifyClient.ExchangeCodeForToken(ctx, restReq.GetCode())
		if err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to exchange oauth code"),
			)
		}
		oAuthToken = BuildOAuthToken(
			uuid.New().String(),
			userInfo.GetId(),
			token.AccessToken,
			token.RefreshToken,
			token.TokenType,
			time.Now().Add(time.Second*time.Duration(token.ExpiresIn)),
			datasource,
		)
	case myncer_pb.Datasource_DATASOURCE_YOUTUBE:
		token, err := dsClients.YoutubeClient.ExchangeCodeForToken(ctx, restReq.GetCode())
		if err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to exchange oauth code"),
			)
		}
		oAuthToken = BuildOAuthToken(
			uuid.New().String(),
			userInfo.GetId(),
			token.AccessToken,
			token.RefreshToken,
			token.TokenType,
			time.Now().Add(time.Second*time.Duration(token.ExpiresIn)),
			datasource,
		)
	default:
		return core.NewProcessRequestResponse_InternalServerError(
			core.NewError("unsuppported datasource %v", datasource),
		)
	}

	// Save token to DB so we can use it later.
	// TODO: We should wipe the old token too (probably in a transaction too).
	if err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.AddToken(ctx, oAuthToken); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to store oauth token to database"),
		)
	}

	return core.NewProcessRequestResponse_OK()
}

func (aeh *authExchangeHandlerImpl) validateRequest(
	restReq *api.OAuthExchangeRequest, /*const*/
) error {
	if len(restReq.GetCode()) == 0 {
		return core.NewError("authorization code is required")
	}
	return nil
}

func (aeh *authExchangeHandlerImpl) getDatasource(
	req *http.Request, /*const*/
) (myncer_pb.Datasource, error) {
	// Path is expected to be like this pattern `/../auth/{datasource}/exchange`
	// so we extract this                                  ^ portion out.
	pathParts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(pathParts) < 4 {
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, core.NewError(
			"malformed path: %s",
			req.URL.Path,
		)
	}
	datasource := pathParts[len(pathParts)-2]
	protoDs := rest_helpers.RestDatasourceToProto(api.Datasource(datasource))
	if protoDs == myncer_pb.Datasource_DATASOURCE_UNSPECIFIED {
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, core.NewError(
			"unknown datasource %s",
			datasource,
		)
	}
	return protoDs, nil
}
