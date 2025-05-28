package handlers

import (
	"context"
	"net/http"
	"slices"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/hansbala/myncer/api"
	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

const (
	cMinPasswordLen = 10
)

func NewCreateUserHandler() core.Handler {
	return &createUserHandlerImpl{}
}

type createUserHandlerImpl struct{}

var _ core.Handler = (*createUserHandlerImpl)(nil)

func (c *createUserHandlerImpl) GetRequestContainer(ctx context.Context) any {
	return &api.CreateUserRequest{}
}

func (c *createUserHandlerImpl) CheckUserPermissions(
	ctx context.Context,
	userInfo *myncer_pb.User, /*const,@nullable*/
) error {
	// No user permissions required to create user.
	return nil
}

func (c *createUserHandlerImpl) ProcessRequest(
	ctx context.Context,
	reqBody any, /*const,@nullable*/
	req *http.Request, /*const*/
	resp http.ResponseWriter,
) *core.ProcessRequestResponse {
	restReq, ok := (reqBody).(*api.CreateUserRequest)
	if !ok {
		return core.NewProcessRequestResponse_BadRequest(
			core.NewError("could not cast to create user request"),
		)
	}
	if err := c.validateRequest(restReq); err != nil {
		return core.NewProcessRequestResponse_BadRequest(
			core.WrappedError(err, "failed to validate create user request"),
		)
	}

	protoUser, err := createProtoUser(restReq)
	if err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to construct proto user"),
		)
	}
	if err := core.ToMyncerCtx(ctx).DB.UserStore.CreateUser(ctx, protoUser); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to create user"),
		)
	}

	restResp := api.NewCreateUserResponse()
	restResp.SetId(protoUser.GetId())
	if err := WriteJSONOk(resp, restResp); err != nil {
		return core.NewProcessRequestResponse_InternalServerError(
			core.WrappedError(err, "failed to write response"),
		)
	}

	return core.NewProcessRequestResponse_OK()
}

func (c *createUserHandlerImpl) validateRequest(req *api.CreateUserRequest /*const*/) error {
	if len(req.GetEmail()) == 0 {
		return core.NewError("email is required")
	}
	if len(req.GetFirstName()) == 0 {
		return core.NewError("first name is required")
	}
	if len(req.GetLastName()) == 0 {
		return core.NewError("last name is required")
	}
	if err := validatePassword(req.GetPassword()); err != nil {
		return core.WrappedError(err, "password validation failed")
	}
	return nil
}

func validatePassword(password string) error {
	// Length requirements.
	plen := len(password)
	if plen < cMinPasswordLen {
		return core.NewError("expected password to be a minimum of %d characters", cMinPasswordLen)
	}
	runeSlice := []rune(password)
	// At least one uppercase.
	if !slices.ContainsFunc(runeSlice, func(r rune) bool { return unicode.IsUpper(r)}) {
		return core.NewError("at least one uppercase letter is required")
	}
	// At least one lowercase.
	if !slices.ContainsFunc(runeSlice, func(r rune) bool { return unicode.IsLower(r)}) {
		return core.NewError("at least one lowercase letter is required")
	}
	// At least one number.
	if !slices.ContainsFunc(runeSlice, func(r rune) bool { return unicode.IsNumber(r)}) {
		return core.NewError("at least one number is required")
	}
	return nil
}

func createProtoUser(req *api.CreateUserRequest /*const*/) (*myncer_pb.User, error) {
	hashedPassword, err := hashPassword(req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &myncer_pb.User{
		Id:             uuid.New().String(),
		FirstName:      req.GetFirstName(),
		LastName:       req.GetLastName(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
	}, nil
}

func hashPassword(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return "", core.WrappedError(err, "failed to hash password using bcrypt")
	}
	return string(bytes), nil
}

