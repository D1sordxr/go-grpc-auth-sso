package commands

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error)
	Login(ctx context.Context, dto LoginDTO) (LoginDTO, error)
	IsAdmin(ctx context.Context, dto IsAdminDTO) (IsAdminDTO, error)
}

type UserCommands struct {
	UserDAO persistence.UserDAO
}

func NewUserCommands(dao persistence.UserDAO) *UserCommands {
	return &UserCommands{UserDAO: dao}
}

func (uc *UserCommands) Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error) {
	if dto.Email == "" {
		return RegisterDTO{}, status.Error(codes.InvalidArgument, "email can not be empty")
	}
	if dto.Password == "" {
		return RegisterDTO{}, status.Error(codes.InvalidArgument, "password can not be empty")
	}

	_ = ctx

	return RegisterDTO{}, nil
}

func (uc *UserCommands) Login(ctx context.Context, dto LoginDTO) (LoginDTO, error) {
	_, _ = ctx, dto
	return LoginDTO{}, nil
}

func (uc *UserCommands) IsAdmin(ctx context.Context, dto IsAdminDTO) (IsAdminDTO, error) {
	_, _ = ctx, dto
	return IsAdminDTO{}, nil
}
