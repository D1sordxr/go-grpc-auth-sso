package auth

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	services "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Register(ctx context.Context, dto commands.RegisterDTO) (commands.RegisterDTO, error)
	Login(ctx context.Context, dto commands.LoginDTO) (commands.LoginDTO, error)
	IsAdmin(ctx context.Context, dto commands.IsAdminDTO) (commands.IsAdminDTO, error)
}

type UserAuthService struct {
	services.UnimplementedAuthServer
	auth       Auth
	AuthServer services.AuthServer
}

func NewUserAuthService(auth Auth) *UserAuthService {
	return &UserAuthService{auth: auth}
}

func (s *UserAuthService) Register(ctx context.Context, req *services.RegisterRequest) (*services.RegisterResponse, error) {
	dto := commands.RegisterDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	response, err := s.auth.Register(ctx, dto)
	if err != nil {
		return &services.RegisterResponse{
				Message: err.Error(),
			},
			status.Error(codes.Aborted, "application error")
	}

	return &services.RegisterResponse{
		Message: "Successfully registered!",
		UserId:  response.UserID,
	}, nil
}

func (s *UserAuthService) Login(ctx context.Context, req *services.LoginRequest) (*services.LoginResponse, error) {
	dto := commands.LoginDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		AppID:    req.GetAppId(),
	}

	response, err := s.auth.Login(ctx, dto)
	if err != nil {
		return &services.LoginResponse{
			Message: err.Error(),
		}, status.Error(codes.Aborted, "application error")
	}

	return &services.LoginResponse{
		Message: "Successfully logged in!",
		Token:   response.Token,
	}, nil
}

func (s *UserAuthService) IsAdmin(ctx context.Context, req *services.IsAdminRequest) (*services.IsAdminResponse, error) {
	command := commands.IsAdminUserCommand{
		UserID: req.UserId,
	}
	response, err := s.AuthServer.IsAdmin(ctx, command)
	if err != nil {
		return &services.IsAdminResponse{
			Message: err.Error(),
		}, status.Error(codes.Aborted, "application error")
	}

	return &services.IsAdminResponse{}, nil
}
