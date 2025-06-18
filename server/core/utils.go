package core

import (
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func WrappedError(err error, format string, a ...any) error {
	errMsg := NewError(format, a...)
	return fmt.Errorf("%s: %w", errMsg.Error(), err)
}

func NewError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func ProtoOAuthTokenToOAuth2(oAuthToken *myncer_pb.OAuthToken /*const*/) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  oAuthToken.GetAccessToken(),
		TokenType:    oAuthToken.GetTokenType(),
		RefreshToken: oAuthToken.GetRefreshToken(),
		Expiry:       oAuthToken.GetExpiresAt().AsTime(),
		// ExpiresIn field is only used for JSON marshal / unmarshal so not required to set.
	}
}

func DebugPrintJson(v any) {
	data, err := json.MarshalIndent(v, "" /*prefix*/, "  " /*indent*/)
	if err != nil {
		Errorf("could not debug print json", err)
	}
	Printf(string(data))
}
