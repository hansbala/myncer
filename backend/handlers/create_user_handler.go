package handlers

import (
	"context"
	"net/http"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

func NewCreateUserHandler() core.Handler {
	return &createUserHandlerImpl{}
}

type createUserHandlerImpl struct{}

var _ core.Handler = (*createUserHandlerImpl)(nil)

func (c *createUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
) error {
	// No user permissions required to create user.
	return nil
}

func (c *createUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) error {
	return nil
}
