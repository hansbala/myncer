package core

import (
	"context"

	myncer_pb "github.com/hansbala/myncer/proto"
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
}
