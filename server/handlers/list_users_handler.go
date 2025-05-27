package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
	"github.com/hansbala/myncer/api"
)

func NewListUsersHandler() core.Handler {
	return &listUsersHandlerImpl{}
}

type listUsersHandlerImpl struct{}

var _ core.Handler = (*listUsersHandlerImpl)(nil)

func (l *listUsersHandlerImpl) GetRequestContainer(ctx context.Context) any /*@nullable*/ {
	// GET request so no request body.
	return nil
}

func (l *listUsersHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
) error {
	// Just for testing - will be removed soon.
	return nil
}

func (l *listUsersHandlerImpl) ProcessRequest(
	ctx context.Context,
	_ any, /*const,@nullable*/
	_ *http.Request, /*const*/
	resp http.ResponseWriter,
) error {
	users, err := core.ToMyncerCtx(ctx).DB.UserStore.GetUsers(ctx)
	if err != nil {
		return core.WrappedError(err, "failed to list all users")
	}
	return WriteJSONOk(resp, getJsonResponse(users))
}

func getJsonResponse(users []*myncer_pb.User /*const*/) *api.ListUsersResponse {
	restUsers := []api.User{}
	for _, puser := range users {
		restUser := api.NewUser()
		restUser.SetId(puser.GetId())
		restUser.SetFirstName(puser.GetFirstName())
		restUser.SetLastName(puser.GetLastName())
		restUser.SetEmail(puser.GetEmail())
		restUsers = append(restUsers, *restUser)
	}
	resp := api.NewListUsersResponse()
	resp.SetUsers(restUsers)
	return resp
}

