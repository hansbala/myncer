package datasources

import (
	"context"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/rest_helpers"
	"github.com/hansbala/myncer/sync_engine"
)

const (
	cYouTubeAuthURL  = "https://accounts.google.com/o/oauth2/auth"
	cYouTubeTokenURL = "https://oauth2.googleapis.com/token"
)

func NewYouTubeClient() core.DatasourceClient {
	return &youtubeClientImpl{}
}

type youtubeClientImpl struct{}

var _ core.DatasourceClient = (*youtubeClientImpl)(nil)

func (c *youtubeClientImpl) ExchangeCodeForToken(
	ctx context.Context,
	code string,
) (*oauth2.Token, error) {
	conf := c.getOAuthConfig(ctx)
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, core.WrappedError(err, "failed to exchange auth code with YouTube")
	}
	return tok, nil
}

func (c *youtubeClientImpl) GetPlaylists(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) ([]*myncer_pb.Playlist, error) {
	svc, err := c.getService(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get YouTube service")
	}

	call := svc.Playlists.List([]string{"snippet"}).Mine(true).MaxResults(50)
	resp, err := call.Do()
	if err != nil {
		return nil, core.WrappedError(err, "failed to fetch playlists")
	}

	var playlists []*myncer_pb.Playlist
	for _, p := range resp.Items {
		playlists = append(
			playlists,
			&myncer_pb.Playlist{
				MusicSource: rest_helpers.CreateMusicSource(
					myncer_pb.Datasource_YOUTUBE,
					p.Id,
				),
				Name:        p.Snippet.Title,
				Description: p.Snippet.Description,
				ImageUrl:    getBestThumbnailUrl(p.Snippet.Thumbnails),
			},
		)
	}
	return playlists, nil
}

func (c *youtubeClientImpl) GetPlaylist(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	id string,
) (*myncer_pb.Playlist, error) {
	svc, err := c.getService(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get YouTube service")
	}
	call := svc.Playlists.List([]string{"snippet"}).Id(id)
	resp, err := call.Do()
	if err != nil || len(resp.Items) == 0 {
		return nil, core.WrappedError(err, "failed to fetch playlist %s", id)
	}

	p := resp.Items[0]
	return &myncer_pb.Playlist{
		MusicSource: rest_helpers.CreateMusicSource(myncer_pb.Datasource_YOUTUBE, p.Id),
		Name:        p.Snippet.Title,
		Description: p.Snippet.Description,
		ImageUrl:    getBestThumbnailUrl(p.Snippet.Thumbnails),
	}, nil
}

func (c *youtubeClientImpl) GetPlaylistSongs(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string,
) ([]core.Song, error) {
	svc, err := c.getService(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get YouTube service")
	}

	songs := []core.Song{}
	nextPageToken := ""
	for {
		call := svc.PlaylistItems.
			List([]string{"snippet"}).
			PlaylistId(playlistId).
			MaxResults(50).
			PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return nil, core.WrappedError(err, "failed to fetch playlist items")
		}

		for _, item := range resp.Items {
			videoId := item.Snippet.ResourceId.VideoId
			if len(videoId) == 0 {
				continue
			}
			songs = append(songs, buildSongFromYouTubePlaylistItem(item))
		}
		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}

	return songs, nil
}

func (c *youtubeClientImpl) AddToPlaylist(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string,
	songs []core.Song,
) error {
	svc, err := c.getService(ctx, userInfo)
	if err != nil {
		return core.WrappedError(err, "failed to get YouTube service")
	}

	for _, song := range songs {
		if _, err := svc.PlaylistItems.Insert(
			[]string{"snippet"},
			&youtube.PlaylistItem{
				Snippet: &youtube.PlaylistItemSnippet{
					PlaylistId: playlistId,
					ResourceId: &youtube.ResourceId{
						Kind:    "youtube#video",
						VideoId: song.GetId(),
					},
				},
			},
		).
			Do(); err != nil {
			return core.WrappedError(err, "failed to insert video %s", song.GetName())
		}
	}
	return nil
}

func (c *youtubeClientImpl) ClearPlaylist(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string,
) error {
	svc, err := c.getService(ctx, userInfo)
	if err != nil {
		return core.WrappedError(err, "failed to get YouTube service")
	}

	var nextPageToken string
	for {
		resp, err := svc.PlaylistItems.
			List([]string{"id"}).
			PlaylistId(playlistId).
			MaxResults(50).
			PageToken(nextPageToken).
			Do()
		if err != nil {
			return core.WrappedError(err, "failed to list playlist items")
		}

		for _, item := range resp.Items {
			if err := svc.PlaylistItems.Delete(item.Id).Do(); err != nil {
				return core.WrappedError(err, "failed to delete playlist item %s", item.Id)
			}
		}

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}

	return nil
}

func (s *youtubeClientImpl) Search(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	names core.Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
	artistNames core.Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
	albumNames core.Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
) (core.Song, error) {
	svc, err := s.getService(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get YouTube service")
	}
	// Build query string
	var queryParts []string
	if names != nil && !names.IsEmpty() {
		for name := range names {
			queryParts = append(queryParts, name)
		}
	}
	if artistNames != nil && !artistNames.IsEmpty() {
		for artist := range artistNames {
			queryParts = append(queryParts, artist)
		}
	}
	if albumNames != nil && !albumNames.IsEmpty() {
		for album := range albumNames {
			queryParts = append(queryParts, album)
		}
	}
	if len(queryParts) == 0 {
		return nil, core.NewError("at least one of name, artist, or album must be provided")
	}
	query := strings.Join(queryParts, " ")

	call := svc.Search.List([]string{"snippet"}).
		Q(query).
		Type("video").
		MaxResults(1)

	resp, err := call.Do()
	if err != nil {
		return nil, core.WrappedError(err, "failed to perform YouTube search for query %q", query)
	}

	if len(resp.Items) == 0 {
		return nil, core.NewError("no video found for query %q", query)
	}

	item := resp.Items[0]
	song, err := buildSongFormYoutubeSearchResultItem(item)
	if err != nil {
		return nil, core.WrappedError(err, "failed to build song from YouTube search result")
	}
	return song, nil
}

func (c *youtubeClientImpl) getService(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (*youtube.Service, error) {
	oAuthToken, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetToken(
		ctx,
		userInfo.GetId(),
		myncer_pb.Datasource_YOUTUBE,
	)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get youtube token for user %s", userInfo.GetId())
	}
	httpClient := oauth2.NewClient(
		ctx,
		c.getOAuthConfig(ctx).TokenSource(ctx, core.ProtoOAuthTokenToOAuth2(oAuthToken)),
	)
	svc, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, core.WrappedError(err, "failed to create YouTube service")
	}
	return svc, nil
}

func (c *youtubeClientImpl) getOAuthConfig(ctx context.Context) *oauth2.Config {
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

func buildSongFromYouTubePlaylistItem(
	pi *youtube.PlaylistItem, /*const*/
) core.Song {
	return sync_engine.NewSong(
		&myncer_pb.Song{
			Name:             pi.Snippet.Title,
			ArtistName:       []string{pi.Snippet.ChannelTitle},
			Datasource:       myncer_pb.Datasource_YOUTUBE,
			DatasourceSongId: pi.Id,
		},
	)
}

func buildSongFormYoutubeSearchResultItem(
	item *youtube.SearchResult, /*const*/
) (core.Song, error) {
	videoId := ""
	if item.Id != nil && item.Id.VideoId != "" {
		videoId = item.Id.VideoId
	} else {
		return nil, core.NewError("missing video ID in YouTube search result")
	}
	return sync_engine.NewSong(
		&myncer_pb.Song{
			Name: strings.TrimSpace(item.Snippet.Title),
			// best-effort: channel title often includes artist
			ArtistName:       []string{item.Snippet.ChannelTitle},
			Datasource:       myncer_pb.Datasource_YOUTUBE,
			DatasourceSongId: videoId,
		},
	), nil
}

// Helper to get the first available thumbnail URL from the YouTube API response.
// Prefers higher resolution thumbnails if available.
func getBestThumbnailUrl(thumbnails *youtube.ThumbnailDetails /*const*/) string {
	switch {
	case thumbnails.Maxres != nil:
		return thumbnails.Maxres.Url
	case thumbnails.Standard != nil:
		return thumbnails.Standard.Url
	case thumbnails.High != nil:
		return thumbnails.High.Url
	case thumbnails.Medium != nil:
		return thumbnails.Medium.Url
	case thumbnails.Default != nil:
		return thumbnails.Default.Url
	default:
		return ""
	}
}
