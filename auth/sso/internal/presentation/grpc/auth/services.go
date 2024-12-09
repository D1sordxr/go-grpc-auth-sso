package auth

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	services "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserAuthService struct {
	services.UnimplementedAuthServer
	Commands application.UserCommandsInterface
}

func NewUserAuthService(commands application.UserCommandsInterface) *UserAuthService {
	return &UserAuthService{Commands: commands}
}

func (s *UserAuthService) Register(ctx context.Context, req *services.RegisterRequest) (*services.RegisterResponse, error) {
	command := commands.RegisterUserCommand{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	response, err := s.Commands.Register(ctx, command)
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
	command := commands.LoginUserCommand{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		AppID:    req.GetAppId(),
	}

	response, err := s.Commands.Login(ctx, command)
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
	// TODO: ........
	command := commands.IsAdminUserCommand{
		UserID: req.UserId,
	}
	response, err := s.Commands.IsAdmin(ctx, command)
	if err != nil {
		return &services.IsAdminResponse{
			Message: err.Error(),
		}, status.Error(codes.Aborted, "application error")
	}

	return &services.IsAdminResponse{
		IsAdmin: response.UserID != 0,
		Message: "",
	}, nil
}
