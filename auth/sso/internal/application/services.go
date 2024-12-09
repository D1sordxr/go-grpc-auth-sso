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
	RegisterUserHandler *handlers.RegisterUserHandler
	LoginUserHandler    *handlers.LoginUserHandler
	IsAdminUserHandler  *handlers.IsAdminUserHandler
}

func NewUserCommands(register *handlers.RegisterUserHandler,
	login *handlers.LoginUserHandler,
	isAdmin *handlers.IsAdminUserHandler) *UserCommands {
	return &UserCommands{
		RegisterUserHandler: register,
		LoginUserHandler:    login,
		IsAdminUserHandler:  isAdmin,
	}
}

func (uc *UserCommands) Register(ctx context.Context, command commands.RegisterUserCommand) (commands.RegisterDTO, error) {
	return uc.RegisterUserHandler.Handle(ctx, command)
}

func (uc *UserCommands) Login(ctx context.Context, command commands.LoginUserCommand) (commands.LoginDTO, error) {
	return uc.LoginUserHandler.Handle(ctx, command)
}

func (uc *UserCommands) IsAdmin(ctx context.Context, command commands.IsAdminUserCommand) (commands.IsAdminDTO, error) {
	return uc.IsAdminUserHandler.Handle(ctx, command)
}
