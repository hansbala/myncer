package rest_helpers

import (
	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func ProtoDatasourceToRest(ds myncer_pb.Datasource) (api.Datasource, error) {
	switch ds {
	case myncer_pb.Datasource_SPOTIFY:
		return api.SPOTIFY, nil
	case myncer_pb.Datasource_YOUTUBE:
		return api.YOUTUBE, nil
	default:
		return "", core.NewError("unknown datasource")
	}
}

func RestDatasourceToProto(ds api.Datasource) myncer_pb.Datasource {
	switch ds {
	case api.SPOTIFY:
		return myncer_pb.Datasource_SPOTIFY
	case api.YOUTUBE:
		return myncer_pb.Datasource_YOUTUBE
	default:
		return myncer_pb.Datasource_DATASOURCE_UNSPECIFIED
	}
}
