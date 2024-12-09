package application

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/handlers"
)

type UserCommandsInterface interface {
	Register(ctx context.Context, command commands.RegisterUserCommand) (commands.RegisterDTO, error)
	Login(ctx context.Context, command commands.LoginUserCommand) (commands.LoginDTO, error)
	IsAdmin(ctx context.Context, command commands.IsAdminUserCommand) (commands.IsAdminDTO, error)
}

type UserCommands struct {
	Register *handlers.RegisterUserHandler
	Login    *handlers.LoginUserHandler
	IsAdmin  *handlers.IsAdminUserHandler
}

func NewUserCommands(register *handlers.RegisterUserHandler,
	login *handlers.LoginUserHandler,
	isAdmin *handlers.IsAdminUserHandler) *UserCommands {
	return &UserCommands{
		Register: register,
		Login:    login,
		IsAdmin:  isAdmin,
	}
}
