package core

import (
	"context"
	"encoding/json"

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
