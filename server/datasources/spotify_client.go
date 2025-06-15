package datasources

import (
	"context"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

const (
	cPageLimit       = 50
	cSpotifyAuthUrl  = "https://accounts.spotify.com/authorize"
	cSpotifyTokenUrl = "https://accounts.spotify.com/api/token"
)

func NewSpotifyClient() *SpotifyClient {
	return &SpotifyClient{}
}

type SpotifyClient struct{}

// ExchangeCodeForToken makes an API request to spotify to to retrieve the access and refresh token.
func (s *SpotifyClient) ExchangeCodeForToken(
	ctx context.Context,
	authCode string,
) (*oauth2.Token, error) {
	authenticator := s.getAuthenticator(ctx)
	token, err := authenticator.Exchange(ctx, authCode)
	if err != nil {
		return nil, core.WrappedError(err, "failed to exchange auth code with spotify")
	}
	return token, nil
}

func (s *SpotifyClient) GetPlaylists(
	ctx context.Context,
	oAuthToken *myncer_pb.OAuthToken, /*const*/
) ([]*Playlist, error) {
	clientSDK := s.getClient(ctx, core.ProtoOAuthTokenToOAuth2(oAuthToken))

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

func (s *SpotifyClient) getClient(ctx context.Context, token *oauth2.Token /*const*/) *spotify.Client {
	tokenSource := s.getOAuthConfig(ctx).TokenSource(ctx, token)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return spotify.New(httpClient)
}

func (s *SpotifyClient) getAuthenticator(ctx context.Context) *spotifyauth.Authenticator {
	spotifyConfig := core.ToMyncerCtx(ctx).Config.SpotifyConfig
	return spotifyauth.New(
		spotifyauth.WithClientID(spotifyConfig.ClientId),
		spotifyauth.WithClientSecret(spotifyConfig.ClientSecret),
		spotifyauth.WithRedirectURL(spotifyConfig.RedirectUri),
	)
}

func (s *SpotifyClient) getOAuthConfig(ctx context.Context) *oauth2.Config {
	spotifyConfig := core.ToMyncerCtx(ctx).Config.SpotifyConfig
	return &oauth2.Config{
		ClientID:     spotifyConfig.ClientId,
		ClientSecret: spotifyConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cSpotifyAuthUrl,
			TokenURL: cSpotifyTokenUrl,
		},
		RedirectURL: spotifyConfig.RedirectUri,
	}
}
