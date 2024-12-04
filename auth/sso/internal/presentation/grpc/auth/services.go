package auth

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	services "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyValue = 0
)

type Auth interface {
	Register(ctx context.Context, dto commands.RegisterDTO) (commands.RegisterDTO, error)
	Login(ctx context.Context, dto commands.LoginDTO) (commands.LoginDTO, error)
	//IsAdmin(ctx context.Context, dto commands.IsAdminDTO) (commands.IsAdminDTO, error)
}

type UserAuthService struct {
	services.UnimplementedAuthServer
	auth Auth
}

func NewUserAuthService(auth Auth) *UserAuthService {
	return &UserAuthService{auth: auth}
}

func (s *UserAuthService) Register(ctx context.Context, req *services.RegisterRequest) (*services.RegisterResponse, error) {
	dto := commands.RegisterDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	_ = dto

	return &services.RegisterResponse{}, nil
}

func (s *UserAuthService) Login(ctx context.Context, req *services.LoginRequest) (*services.LoginResponse, error) {
	dto := commands.LoginDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		AppID:    req.GetAppId(),
	}

	if dto.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email can not be empty")
	}
	if dto.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password can not be empty")
	}
	if dto.AppID == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app ID is required")
	}

	response, err := s.auth.Login(ctx, dto)

	return &services.LoginResponse{}, nil

}

func (s *UserAuthService) IsAdmin(ctx context.Context, req *services.IsAdminRequest) (*services.IsAdminResponse, error) {
	// TODO: Реализовать проверку админских прав
	return &services.IsAdminResponse{}, nil
}
