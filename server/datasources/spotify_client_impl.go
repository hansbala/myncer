package datasources

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
	"github.com/hansbala/myncer/sync_engine"
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

func (s *spotifyClientImpl) GetPlaylist(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string,
) (*myncer_pb.Playlist, error) {
	client, err := s.getClient(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get spotify client")
	}
	if len(playlistId) == 0 {
		return nil, core.NewError("invalid playlist id")
	}
	playlist, err := client.GetPlaylist(ctx, spotify.ID(playlistId))
	if err != nil {
		return nil, core.WrappedError(err, "failed to get spotify playlist with id %s", playlistId)
	}
	return spotifyPlaylistToProto(playlist), nil
}

func (s *spotifyClientImpl) GetPlaylistSongs(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string,
) ([]core.Song, error) {
	client, err := s.getClient(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get spotify client")
	}
	// Use GetPlaylistItems to fetch all songs in the playlist.
	if len(playlistId) == 0 {
		return nil, core.NewError("invalid playlist id")
	}
	allSongs := []core.Song{}
	offset := 0
	for {
		playlistTracks, err := client.GetPlaylistItems(
			ctx,
			spotify.ID(playlistId),
			spotify.Limit(cPageLimit),
			spotify.Offset(offset),
		)
		if err != nil {
			if spotifyErr, ok := err.(spotify.Error); ok &&
				spotifyErr.Status == http.StatusTooManyRequests {
				core.Printf("Spotify API rate limit hit, with message: %s", spotifyErr.Message)
			}
			return nil, core.WrappedError(
				err,
				"failed to get playlist items for playlist %s at offset %d",
				playlistId,
				offset,
			)
		}
		for _, item := range playlistTracks.Items {
			if item.Track.Track != nil {
				allSongs = append(allSongs, buildSongFromSpotifyTrack(ctx, item.Track.Track))
			}
		}
		if len(playlistTracks.Items) < cPageLimit {
			// No more items left to fetch.
			break
		}
		offset += cPageLimit
	}
	return allSongs, nil
}

func (s *spotifyClientImpl) GetPlaylists(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) ([]*myncer_pb.Playlist, error) {
	client, err := s.getClient(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get spotify client")
	}

	r := []*myncer_pb.Playlist{}
	for offset := 0; ; offset += cPageLimit {
		page, err := client.CurrentUsersPlaylists(
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
					MusicSource: createMusicSource(
						myncer_pb.Datasource_DATASOURCE_SPOTIFY,
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

func (s *spotifyClientImpl) AddToPlaylist(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string, /*const*/
	songs []core.Song, /*const*/
) error {
	client, err := s.getClient(ctx, userInfo)
	if err != nil {
		return core.WrappedError(err, "failed to get spotify client")
	}
	trackIds := []spotify.ID{}
	for _, song := range songs {
		trackIds = append(trackIds, spotify.ID(song.GetId()))
	}
	if _, err := client.AddTracksToPlaylist(ctx, spotify.ID(playlistId), trackIds...); err != nil {
		return core.WrappedError(err, "failed to add tracks to playlist %s", playlistId)
	}
	return nil
}

func (s *spotifyClientImpl) ClearPlaylist(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	playlistId string, /*const*/
) error {
	client, err := s.getClient(ctx, userInfo)
	if err != nil {
		return core.WrappedError(err, "failed to get spotify client")
	}
	// Fetch all track URIs to remove
	playlistTracks, err := client.GetPlaylistItems(ctx, spotify.ID(playlistId))
	if err != nil {
		return core.WrappedError(err, "failed to fetch playlist items")
	}

	trackIDs := []spotify.ID{}
	for _, item := range playlistTracks.Items {
		if item.Track.Track != nil {
			trackIDs = append(trackIDs, item.Track.Track.ID)
		}
	}

	if len(trackIDs) == 0 {
		return nil
	}
	_, err = client.RemoveTracksFromPlaylist(ctx, spotify.ID(playlistId), trackIDs...)
	if err != nil {
		return core.WrappedError(err, "failed to clear playlist")
	}
	return nil
}

func (s *spotifyClientImpl) Search(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
	names core.Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
	_ core.Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
	_ core.Set[string], /*const,@nullable*/ // nil, empty indicates no filtering.
) (core.Song, error) {
	client, err := s.getClient(ctx, userInfo)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get spotify client")
	}

	// Build OR-grouped qualifiers
	var clauses []string

	// track names
	if names != nil {
		raw := names.ToArray()
		filtered := filterEmpty(raw)
		if len(filtered) > 0 {
			var terms []string
			for _, n := range filtered {
				terms = append(terms, fmt.Sprintf(`track:%q`, n))
			}
			clauses = append(clauses, "("+strings.Join(terms, " OR ")+")")
		}
	}

	if len(clauses) == 0 {
		return nil, core.NewError("at least one of name, artist, or album must be provided")
	}

	query := strings.Join(clauses, " OR ")
	// Example result:
	//   (track:"Pressure") OR (artist:"Martin Garrix" OR artist:"Tove Lo")

	searchResult, err := client.Search(ctx, query, spotify.SearchTypeTrack, spotify.Limit(1))
	if err != nil {
		return nil, core.WrappedError(err, "spotify search failed for query %q", query)
	}
	if searchResult.Tracks == nil || len(searchResult.Tracks.Tracks) == 0 {
		return nil, core.NewError("no track found for query %q", query)
	}

	return buildSongFromSpotifyTrack(ctx, &searchResult.Tracks.Tracks[0]), nil
}

func (s *spotifyClientImpl) getClient(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const*/
) (*spotify.Client, error) {
	oAuthToken, err := core.ToMyncerCtx(ctx).DB.DatasourceTokenStore.GetToken(
		ctx,
		userInfo.GetId(),
		myncer_pb.Datasource_DATASOURCE_SPOTIFY,
	)
	if err != nil {
		return nil, core.WrappedError(err, "failed to get spotify token for user %s", userInfo.GetId())
	}

	tokenSource := s.getOAuthConfig(ctx).TokenSource(ctx, core.ProtoOAuthTokenToOAuth2(oAuthToken))
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return spotify.New(httpClient), nil
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

func buildSongFromSpotifyTrack(
	_ context.Context,
	track *spotify.FullTrack, /*const*/
) core.Song {
	return sync_engine.NewSong(
		&myncer_pb.Song{
			Name:             track.Name,
			ArtistName:       []string{track.Artists[0].Name},
			AlbumName:        track.Album.Name,
			Datasource:       myncer_pb.Datasource_DATASOURCE_SPOTIFY,
			DatasourceSongId: track.ID.String(),
		},
	)
}

func filterEmpty(vals []string) (out []string) {
	for _, v := range vals {
		v = strings.TrimSpace(v)
		if v != "" {
			out = append(out, v)
		}
	}
	return
}
