package handlers

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"
)

type RegisterUserHandler struct {
	UserDAO    commands.UserDAO
	UoWManager persistence.UoWManager
}

func NewRegisterUserHandler(dao commands.UserDAO, uow persistence.UoWManager) *RegisterUserHandler {
	return &RegisterUserHandler{
		UserDAO:    dao,
		UoWManager: uow,
	}
}

func (h *RegisterUserHandler) Handle(ctx context.Context, command commands.RegisterUserCommand) (commands.RegisterDTO, error) {
	userID := vo.NewUserID()
	email, err := vo.NewEmail(command.Email)
	if err != nil {
		return commands.RegisterDTO{}, err
	}
	password, err := vo.NewPassword(command.Password)
	if err != nil {
		return commands.RegisterDTO{}, err
	}

	uow := h.UoWManager.GetUoW()

	tx, err := uow.Begin(ctx)
	if err != nil {
		return commands.RegisterDTO{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = uow.Rollback(ctx)
			panic(r)
		}
		if err != nil {
			_ = uow.Rollback(ctx)
		}
	}()

	err = h.UserDAO.Exists(ctx, email.Email)
	if err != nil {
		return commands.RegisterDTO{}, err
	}

	user := entity.NewUser(userID, email, password)

	err = h.UserDAO.Register(ctx, tx, user)
	if err != nil {
		return commands.RegisterDTO{}, err
	}
	if err = uow.Commit(ctx); err != nil {
		return commands.RegisterDTO{}, err
	}

	return commands.RegisterDTO{
		UserID: user.StringUserID(),
	}, err
}
