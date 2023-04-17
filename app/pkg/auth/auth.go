package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

// GenerateJWT retrieves a token using a sample key and sign
// the token with the username inserting it in the claims.
func GenerateJWT() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":        time.Date(2022, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"authorized": true,
		"user":       "username",
	})

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "[ERROR] Signing Error", err
	}

	return tokenString, nil
}
