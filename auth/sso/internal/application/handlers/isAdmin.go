package handlers

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
)

type IsAdminUserHandler struct {
	UserDAO    commands.UserDAO
	UoWManager persistence.UoWManager
}

func NewIsAdminUserHandler(dao commands.UserDAO, uow persistence.UoWManager) *IsAdminUserHandler {
	return &IsAdminUserHandler{
		UserDAO:    dao,
		UoWManager: uow,
	}
}

func (h *IsAdminUserHandler) Handle(ctx context.Context, command commands.IsAdminUserCommand) (commands.IsAdminDTO, error) {
	_, _ = ctx, command
	return commands.IsAdminDTO{}, nil
}
