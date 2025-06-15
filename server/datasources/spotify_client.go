package datasources

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

const (
	cPageLimit = 50
)

func NewSpotifyClient() *SpotifyClient {
	return &SpotifyClient{}
}

type SpotifyClient struct{}

// ExchangeCodeForToken makes an API request to spotify to to retrieve the access and refresh token.
// TODO: We can use the SDK's `Exchange` function to simplify this function.
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

func (s *SpotifyClient) GetPlaylists(
	ctx context.Context,
	oAuthToken *myncer_pb.OAuthToken, /*const*/
) ([]*Playlist, error) {
	clientSDK := s.getSDK(ctx, core.ProtoOAuthTokenToOAuth2(oAuthToken))

	r := []*Playlist{}
	for offset := 0; ; offset += cPageLimit {
		page, err := clientSDK.CurrentUsersPlaylists(
			ctx,
			spotify.Limit(cPageLimit),
			spotify.Offset(offset),
		)
		if err != nil {
			return nil, core.WrappedError(
				err,
				"failed to get current user playlists at offset %d",
				offset,
			)
		}

		for _, p := range page.Playlists {
			r = append(
				r,
				&Playlist{
					ID:   p.ID.String(),
					Name: p.Name,
					URI:  string(p.URI),
				},
			)
		}

		if len(page.Playlists) < cPageLimit {
			// No more pages left to get.
			break
		}
	}

	return r, nil
}

func (s *SpotifyClient) getSDK(ctx context.Context, token *oauth2.Token /*const*/) *spotify.Client {
	spotifyConfig := core.ToMyncerCtx(ctx).Config.SpotifyConfig
	authenticator := spotifyauth.New(
		spotifyauth.WithClientID(spotifyConfig.ClientId),
		spotifyauth.WithClientSecret(spotifyConfig.ClientSecret),
	)
	return spotify.New(authenticator.Client(ctx, token))
}
