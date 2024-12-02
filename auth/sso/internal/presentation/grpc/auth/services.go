package auth

import (
	"context"
	services "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/protobuf"
)

type Auth interface {
	Register(ctx context.Context, req *services.RegisterRequest) (*services.RegisterResponse, error)
	Login(ctx context.Context, req *services.LoginRequest) (*services.LoginResponse, error)
	IsAdmin(ctx context.Context, req *services.IsAdminRequest) (*services.IsAdminResponse, error)
}

type UserAuthService struct {
	services.UnimplementedAuthServer
}

func NewUserAuthService() *UserAuthService {
	return &UserAuthService{}
}

func (s *UserAuthService) Register(ctx context.Context, req *services.RegisterRequest) (*services.RegisterResponse, error) {
	// TODO: Реализовать логику регистрации
	panic("Register method not implemented")
}

func (s *UserAuthService) Login(ctx context.Context, req *services.LoginRequest) (*services.LoginResponse, error) {
	// TODO: Реализовать логику логина
	panic("Login method not implemented")
}

func (s *UserAuthService) IsAdmin(ctx context.Context, req *services.IsAdminRequest) (*services.IsAdminResponse, error) {
	// TODO: Реализовать проверку админских прав
	panic("IsAdmin method not implemented")
}
