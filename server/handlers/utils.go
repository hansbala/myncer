package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hansbala/myncer/core"
	myncer_pb "github.com/hansbala/myncer/proto"
)

const (
	cMinPasswordLen = 10
)

func WriteJSONOk(resp http.ResponseWriter, body any) error {
	// NOTE: Writing a body also writes the StatusOK to the header.
	if err := json.NewEncoder(resp).Encode(body); err != nil {
		return core.WrappedError(err, "failed to write to response body")
	}
	resp.Header().Add("Content-Type", "application/json")
	return nil
}

func getProtoUser(
	id string,
	firstName string,
	lastName string,
	email string,
	password string,
) (*myncer_pb.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &myncer_pb.User{
		Id:             id,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
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

func validateUserFields(
	firstName string,
	lastName string,
	email string,
	password string,
) error {
	if len(email) == 0 {
		return core.NewError("email is required")
	}
	if len(firstName) == 0 {
		return core.NewError("first name is required")
	}
	if len(lastName) == 0 {
		return core.NewError("last name is required")
	}
	if err := validatePassword(password); err != nil {
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
	if !slices.ContainsFunc(runeSlice, func(r rune) bool { return unicode.IsUpper(r) }) {
		return core.NewError("at least one uppercase letter is required")
	}
	// At least one lowercase.
	if !slices.ContainsFunc(runeSlice, func(r rune) bool { return unicode.IsLower(r) }) {
		return core.NewError("at least one lowercase letter is required")
	}
	// At least one number.
	if !slices.ContainsFunc(runeSlice, func(r rune) bool { return unicode.IsNumber(r) }) {
		return core.NewError("at least one number is required")
	}
	return nil
}

func buildOAuthToken(
	id string,
	userId string,
	accessToken string,
	refreshToken string,
	tokenType string,
	expiresAt time.Time,
	datasource myncer_pb.Datasource,
) *myncer_pb.OAuthToken {
	return &myncer_pb.OAuthToken{
		Id:           id,
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    tokenType,
		ExpiresAt:    timestamppb.New(expiresAt),
		Datasource:   datasource,
	}
}
