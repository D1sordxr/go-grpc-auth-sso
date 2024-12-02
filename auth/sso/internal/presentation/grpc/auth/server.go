package auth

import (
	"context"
	services "github.com/D1sordxr/aviasales/auth/sso/protobuf"
)

type ServerAPI struct {
	services.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int) (token string, err error)
	RegisterNewUser(ctx context.Context, email string)
}

func (s *ServerAPI) Login(ctx context.Context, lr *services.LoginRequest) (*services.LoginResponse, error) {
	// TODO:
	return nil, nil
}

func (s *ServerAPI) Register(ctx context.Context, rr *services.RegisterRequest) (*services.RegisterResponse, error) {
	// TODO
	return nil, nil
}
