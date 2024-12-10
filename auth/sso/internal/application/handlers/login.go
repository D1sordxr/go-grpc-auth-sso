package handlers

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"
)

type TokenService interface {
	GenerateToken(userID string, appID int32) (string, error)
	ValidateToken(token string) (bool, error)
}

type LoginUserHandler struct {
	UserDAO      commands.UserDAO
	UoWManager   persistence.UoWManager
	TokenService TokenService
}

func NewLoginUserHandler(dao commands.UserDAO, uow persistence.UoWManager, tokenService TokenService) *LoginUserHandler {
	return &LoginUserHandler{
		UserDAO:      dao,
		UoWManager:   uow,
		TokenService: tokenService,
	}
}

func (h *LoginUserHandler) Handle(ctx context.Context, command commands.LoginUserCommand) (commands.LoginDTO, error) {
	email, err := vo.NewEmail(command.Email)
	if err != nil {
		return commands.LoginDTO{}, err
	}
	password, err := vo.NewPassword(command.Password)
	if err != nil {
		return commands.LoginDTO{}, err
	}
	appID := command.AppID
	err = h.UserDAO.Exists(ctx, email.Email)
	if err != nil {
		return commands.LoginDTO{}, err
	}

	loggingUser, err := h.UserDAO.Load(ctx, email.Email) // TODO: h.UserDAO.Load(ctx, email.Email)
	if err != nil {
		return commands.LoginDTO{}, err
	}

	if !password.Matches(loggingUser.Password) {
		return commands.LoginDTO{}, exceptions.InvalidCredentials
	}

	token, err := h.TokenService.GenerateToken(loggingUser.UserID.String(), appID)
	if err != nil {
		return commands.LoginDTO{}, err
	}

	return commands.LoginDTO{Token: token}, nil
}
