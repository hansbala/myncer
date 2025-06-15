package datasources

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

const (
	cYouTubeAuthURL  = "https://accounts.google.com/o/oauth2/auth"
	cYouTubeTokenURL = "https://oauth2.googleapis.com/token"
)

type YouTubeClient struct{}

func NewYouTubeClient() *YouTubeClient {
	return &YouTubeClient{}
}

func (c *YouTubeClient) ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	conf := c.getOAuthConfig(ctx)
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, core.WrappedError(err, "failed to exchange auth code with YouTube")
	}
	return tok, nil
}

func (c *YouTubeClient) GetPlaylists(ctx context.Context, token *myncer_pb.OAuthToken) ([]*core.Playlist, error) {
	httpClient := oauth2.NewClient(ctx, c.getOAuthConfig(ctx).TokenSource(ctx, core.ProtoOAuthTokenToOAuth2(token)))
	svc, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, core.WrappedError(err, "failed to create YouTube service")
	}

	call := svc.Playlists.List([]string{"snippet"}).Mine(true).MaxResults(50)
	resp, err := call.Do()
	if err != nil {
		return nil, core.WrappedError(err, "failed to fetch playlists")
	}

	var playlists []*core.Playlist
	for _, item := range resp.Items {
		playlists = append(playlists, &core.Playlist{
			ID:   item.Id,
			Name: item.Snippet.Title,
		})
	}
	return playlists, nil
}

func (c *YouTubeClient) getOAuthConfig(ctx context.Context) *oauth2.Config {
	youtubeCfg := core.ToMyncerCtx(ctx).Config.YoutubeConfig
	return &oauth2.Config{
		ClientID:     youtubeCfg.ClientId,
		ClientSecret: youtubeCfg.ClientSecret,
		RedirectURL:  youtubeCfg.RedirectUri,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
}
