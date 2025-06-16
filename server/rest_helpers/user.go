package rest_helpers

import (
	"github.com/hansbala/myncer/api"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func UserProtoToRest(up *myncer_pb.User /*const*/) *api.User {
	return api.NewUser(
		up.GetId(),
		up.GetFirstName(),
		up.GetLastName(),
		up.GetEmail(),
	)
}
