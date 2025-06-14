package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	"github.com/hansbala/myncer/datasources"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewAuthExchangeHandler(spotifyClient *datasources.SpotifyClient) core.Handler {
	return &authExchangeHandlerImpl{spotifyClient: spotifyClient}
}

type authExchangeHandlerImpl struct {
	spotifyClient *datasources.SpotifyClient /*const*/
}

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
	var oAuthToken *myncer_pb.OAuthToken
	switch datasource {
	case myncer_pb.Datasource_SPOTIFY:
		tokenResponse, err := aeh.spotifyClient.ExchangeCodeForToken(ctx, restReq.GetCode())
		if err != nil {
			return core.NewProcessRequestResponse_InternalServerError(
				core.WrappedError(err, "failed to exchange oauth code"),
			)
		}
		oAuthToken = buildOAuthToken(
			tokenResponse.AccessToken,
			tokenResponse.RefreshToken,
			tokenResponse.TokenType,
			tokenResponse.Scope,
			time.Now().Add(time.Duration(tokenResponse.ExpiresIn)),
			datasource,
		)
	default:
		return core.NewProcessRequestResponse_InternalServerError(
			core.NewError("unsuppported datasource %v", datasource),
		)
	}

	// TODO: Save the `oAuthToken` to the database but just log it for now.
	core.Printf("oauth token: %v", oAuthToken)
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
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, core.NewError("malformed path")
	}
	datasource := pathParts[len(pathParts)-2]
	switch datasource {
	case string(api.SPOTIFY):
		return myncer_pb.Datasource_SPOTIFY, nil
	default:
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED, core.NewError(
			"unknown datasource %s",
			datasource,
		)
	}
}
