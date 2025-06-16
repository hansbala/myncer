package core

import myncer_pb "github.com/hansbala/myncer/proto"

type Song interface {
	GetName() string
	GetArtistNames() []string
	GetAlbum() string
	GetId() string
	GetIdByDatasource(datasource myncer_pb.Datasource) (string, error)
}
