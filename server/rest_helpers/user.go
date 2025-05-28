package rest_helpers

import (
	"github.com/hansbala/myncer/api"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func UserProtoToRest(up *myncer_pb.User /*const*/) *api.User {
	restUser := api.NewUser()
	restUser.SetId(up.GetId())
	restUser.SetFirstName(up.GetFirstName())
	restUser.SetLastName(up.GetLastName())
	restUser.SetEmail(up.GetEmail())
	return restUser
}
