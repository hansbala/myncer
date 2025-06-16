package datasources

import (
	"context"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/rest_helpers"
)

const (
	cPageLimit       = 50
	cSpotifyAuthUrl  = "https://accounts.spotify.com/authorize"
	cSpotifyTokenUrl = "https://accounts.spotify.com/api/token"
)

func NewSpotifyClient() core.DatasourceClient {
	return &spotifyClientImpl{}
}

type spotifyClientImpl struct{}

var _ core.DatasourceClient = (*spotifyClientImpl)(nil)

// ExchangeCodeForToken makes an API request to spotify to to retrieve the access and refresh token.
func (s *spotifyClientImpl) ExchangeCodeForToken(
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

func (s *spotifyClientImpl) GetPlaylists(
	ctx context.Context,
	oAuthToken *myncer_pb.OAuthToken, /*const*/
) ([]*myncer_pb.Playlist, error) {
	clientSDK := s.getClient(ctx, core.ProtoOAuthTokenToOAuth2(oAuthToken))

	r := []*myncer_pb.Playlist{}
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
				&myncer_pb.Playlist{
					MusicSource: rest_helpers.CreateMusicSource(
						myncer_pb.Datasource_SPOTIFY,
						p.ID.String(),
					),
					Name:        p.Name,
					Description: p.Description,
					ImageUrl:    getBestSpotifyImageURL(p.Images),
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

func (s *spotifyClientImpl) getClient(ctx context.Context, token *oauth2.Token /*const*/) *spotify.Client {
	tokenSource := s.getOAuthConfig(ctx).TokenSource(ctx, token)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return spotify.New(httpClient)
}

func (s *spotifyClientImpl) getAuthenticator(ctx context.Context) *spotifyauth.Authenticator {
	spotifyConfig := core.ToMyncerCtx(ctx).Config.SpotifyConfig
	return spotifyauth.New(
		spotifyauth.WithClientID(spotifyConfig.ClientId),
		spotifyauth.WithClientSecret(spotifyConfig.ClientSecret),
		spotifyauth.WithRedirectURL(spotifyConfig.RedirectUri),
	)
}

func (s *spotifyClientImpl) getOAuthConfig(ctx context.Context) *oauth2.Config {
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

// getBestSpotifyImageURL returns the URL of the first available image from the provided images.
func getBestSpotifyImageURL(images []spotify.Image /*const*/) string {
	if len(images) > 0 {
		return images[0].URL
	}
	return ""
}
