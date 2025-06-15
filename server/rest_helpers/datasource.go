package rest_helpers

import (
	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func ProtoDatasourceToRest(ds myncer_pb.Datasource) (api.Datasource, error) {
	switch ds {
	case myncer_pb.Datasource_SPOTIFY:
		return api.SPOTIFY, nil
	default:
		return "", core.NewError("unknown datasource")
	}
}
