package rpc_handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"golang.org/x/oauth2"
)

func NewDatasourceOAuthExchangeHandler() core.GrpcHandler[
	*myncer_pb.ExchangeOAuthCodeRequest,
	*myncer_pb.ExchangeOAuthCodeResponse,
] {
	return &datasourceOAuthExchangeImpl{}
}

type datasourceOAuthExchangeImpl struct{}

func (do *datasourceOAuthExchangeImpl) CheckPerms(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
	reqBody *myncer_pb.ExchangeOAuthCodeRequest, /*const*/
) error {
	if userInfo == nil {
		return core.NewError("user is required for oauth exchange")
	}
	return nil
}

func (do *datasourceOAuthExchangeImpl) ProcessRequest(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	reqBody *myncer_pb.ExchangeOAuthCodeRequest, /*const*/
) *core.GrpcHandlerResponse[*myncer_pb.ExchangeOAuthCodeResponse] {
	if err := do.validateRequest(reqBody); err != nil {
		return core.NewGrpcHandlerResponse_BadRequest[*myncer_pb.ExchangeOAuthCodeResponse](
			core.WrappedError(err, "request failed validation"),
		)
	}

	// Based on the datasource get the client, authenticate, and fetch the oauth2 token.
	var (
		token *oauth2.Token
		err   error
	)
	dsClients := core.ToMyncerCtx(ctx).DatasourceClients
	switch reqBody.GetDatasource() {
	case myncer_pb.Datasource_DATASOURCE_SPOTIFY:
		token, err = dsClients.SpotifyClient.ExchangeCodeForToken(ctx, reqBody.GetCode())
		if err != nil {
			return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ExchangeOAuthCodeResponse](
				core.WrappedError(err, "failed to exchange oauth code"),
			)
		}
	case myncer_pb.Datasource_DATASOURCE_YOUTUBE:
		token, err = dsClients.YoutubeClient.ExchangeCodeForToken(ctx, reqBody.GetCode())
		if err != nil {
			return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ExchangeOAuthCodeResponse](
				core.WrappedError(err, "failed to exchange oauth code"),
			)
		}
	default:
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ExchangeOAuthCodeResponse](
			core.NewError("unsuppported datasource %v", reqBody.GetDatasource()),
		)
	}

	// Convert the token to our internal format.
	oAuthToken := BuildOAuthToken(
		uuid.New().String(),
		userInfo.GetId(),
		token.AccessToken,
		token.RefreshToken,
		token.TokenType,
		time.Now().Add(time.Second*time.Duration(token.ExpiresIn)),
		reqBody.GetDatasource(),
	)

	// Save token to DB so we can use it later.
	// TODO: We should wipe the old token too (probably in a transaction too).
	if err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.AddToken(ctx, oAuthToken); err != nil {
		return core.NewGrpcHandlerResponse_InternalServerError[*myncer_pb.ExchangeOAuthCodeResponse](
			core.WrappedError(err, "failed to store oauth token to database"),
		)
	}

	return core.NewGrpcHandlerResponse_OK(
		&myncer_pb.ExchangeOAuthCodeResponse{
			OauthExchangeStatus: myncer_pb.OAuthExchangeStatus_O_AUTH_EXCHANGE_STATUS_SUCCESS,
		},
	)
}

func (do *datasourceOAuthExchangeImpl) validateRequest(
	req *myncer_pb.ExchangeOAuthCodeRequest, /*const*/
) error {
	if len(req.GetCode()) == 0 {
		return core.NewError("authorization code is required")
	}
	return nil
}
