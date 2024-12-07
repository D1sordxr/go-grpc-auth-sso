package vo

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "secret"

type Token struct {
	Token string
}

func NewToken(userID string) (Token, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return Token{}, exceptions.ErrorCreatingToken
	}

	return Token{Token: token}, nil
}
