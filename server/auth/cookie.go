package auth

import (
	"net/http"
	"time"

	myncer_pb "github.com/hansbala/myncer/proto/myncer"
)

func GetAuthCookie(jwtToken string, serverMode myncer_pb.ServerMode) *http.Cookie {
	return &http.Cookie{
		Name:     cJwtCookieName,
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: isHttpOnly(serverMode),
		Secure:   true, // Send the cookie only over HTTPS.
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}
}

func GetLogoutAuthCookie(serverMode myncer_pb.ServerMode) *http.Cookie {
	return &http.Cookie{
		Name:     cJwtCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: isHttpOnly(serverMode),
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	}
}

func SetAuthCookie(w http.ResponseWriter, jwtToken string, serverMode myncer_pb.ServerMode) {
	http.SetCookie(w, GetAuthCookie(jwtToken, serverMode))
}

func ClearAuthCookie(w http.ResponseWriter, serverMode myncer_pb.ServerMode) {
	http.SetCookie(w, GetLogoutAuthCookie(serverMode))
}
