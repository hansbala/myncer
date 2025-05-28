package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	cJwtCookieName = "jwt"
)

func GenerateJWTToken(jwtSecret string, userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
}

func SetAuthCookie(w http.ResponseWriter, jwtToken string) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name: cJwtCookieName,
			Value: jwtToken,
			Path: "/",
			HttpOnly: true,
			// TODO: Support dev mode.
			Secure: true, // Send the cookie only over HTTPS. 
			SameSite: http.SameSiteStrictMode,
			Expires: time.Now().Add(24 * time.Hour),
		},
	)
}
