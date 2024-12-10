package token

import (
	"errors"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/token/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const emptyToken = ""

type Service struct {
	Key      string
	TokenTTL time.Duration
}

type CustomClaims struct {
	jwt.RegisteredClaims
	AppID int32 `json:"app_id"`
}

func NewTokenService(config *config.TokenConfig) *Service {
	return &Service{
		Key:      config.Key,
		TokenTTL: config.TokenTTL,
	}
}

func (s *Service) GenerateToken(userID string, appID int32) (string, error) {
	if userID == "" {
		return emptyToken, InvalidUserID
	}
	if appID <= 0 {
		return emptyToken, InvalidAppID
	}

	secretKey := s.Key
	tokenExpireTime := s.TokenTTL

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AppID: appID,
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return emptyToken, ErrorCreatingToken
	}

	return token, nil
}

func (s *Service) ValidateToken(token string) (bool, error) {
	// TODO: ........
	_ = token
	return false, errors.New("not implemented")
}
