package token

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/token/config"
	"time"
)

type Service struct {
	Key      string
	TokenTTL time.Duration
}

func NewTokenService(config *config.TokenConfig) *Service {
	return &Service{
		Key:      config.Key,
		TokenTTL: config.TokenTTL,
	}
}
