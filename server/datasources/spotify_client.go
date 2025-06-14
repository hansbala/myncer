package datasources

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/hansbala/myncer/core"
)

func NewSpotifyClient() *SpotifyClient {
	return &SpotifyClient{}
}

type SpotifyClient struct {}

type SpotifyTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// ExchangeCodeForToken makes an API request to spotify to to retrieve the access and refresh token.
func (s *SpotifyClient) ExchangeCodeForToken(
	ctx context.Context, 
	authCode string,
) (*SpotifyTokenResponse, error) {
	config := core.ToMyncerCtx(ctx).Config
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", authCode)
	form.Set("redirect_uri", config.SpotifyConfig.RedirectUri)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost, 
		"https://accounts.spotify.com/api/token", 
		bytes.NewBufferString(form.Encode()),
	)
	if err != nil {
		return nil, core.WrappedError(err, "failed to create token request")
	}
	req.SetBasicAuth(config.SpotifyConfig.ClientId, config.SpotifyConfig.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, core.WrappedError(err, "failed to send token request to spotify")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// The body holds helpful context for debugging so we include it as part of the error message.
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, core.NewError("failed to make http request to spotify and failed to read body")
		}
		return nil, core.NewError("failed to make http request to spotify: %s", string(body))
	}

	tokenResponse := &SpotifyTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(tokenResponse); err != nil {
		return nil, core.WrappedError(err, "failed to decode spotify tokens response")
	}

	return tokenResponse, nil
}
