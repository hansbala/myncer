package core

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

type Song interface {
	GetName() string
	GetArtistNames() []string
	GetAlbum() string
	GetId() string
	GetIdByDatasource(
		ctx context.Context,
		userInfo *myncer_pb.User, /*const*/
		datasource myncer_pb.Datasource,
	) (string, error)
	GetSpec() *myncer_pb.Song
}

func NewSongList(songs []Song /*const*/) *SongList {
	return &SongList{songs: songs}
}

type SongList struct {
	songs []Song /*const*/
}

func (sl *SongList) GetLlmJson() ([]byte, error) {
	protoSongs := []*myncer_pb.Song{}
	for _, s := range sl.songs {
		protoSongs = append(protoSongs, s.GetSpec())
	}
	bytes, err := json.MarshalIndent(protoSongs, "" /*prefix*/, "  " /*indent*/)
	if err != nil {
		return nil, WrappedError(err, "failed to json marshal song list for llm")
	}
	return bytes, nil
}

func (sl *SongList) GetSongs() []Song {
	return sl.songs
}

// GetSongId returns a deterministic UUID hash of the song name, artist(s), and album name.
func GetSongId(
	songName string,
	artistNames []string, /*const*/
	albumName string,
) string {
	data := []byte(strings.ToLower(strings.TrimSpace(songName)) + "|" +
		strings.ToLower(strings.Join(artistNames, ",")) + "|" +
		strings.ToLower(strings.TrimSpace(albumName)))
	// Generate deterministic UUID using SHA-1.
	// Use a fixed namespace UUID (can be any UUID — here we use the URL namespace).
	return uuid.NewSHA1(uuid.NameSpaceURL, data).String()
}
