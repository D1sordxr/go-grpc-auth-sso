package handlers

import (
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
