package commands

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"
)

type RegisterUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	AppID    int32  `json:"app_id" binding:"required"`
	Token    string `json:"token"`
}

type IsAdminUserCommand struct {
	UserID int32 `json:"user_id" binding:"required"`
}

type Auth interface {
	Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error)
	Login(ctx context.Context, dto LoginDTO) (LoginDTO, error)
	IsAdmin(ctx context.Context, dto IsAdminDTO) (IsAdminDTO, error)
}

type UserCommands struct {
	UserDAO    UserDAO
	UoWManager persistence.UoWManager
}

func NewUserCommands(dao UserDAO, uow persistence.UoWManager) *UserCommands {
	return &UserCommands{
		UserDAO:    dao,
		UoWManager: uow,
	}
}

func (uc *UserCommands) Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error) {
	userID := vo.NewUserID()
	email, err := vo.NewEmail(dto.Email)
	if err != nil {
		return RegisterDTO{}, err
	}
	password, err := vo.NewPassword(dto.Password)
	if err != nil {
		return RegisterDTO{}, err
	}
	err = uc.UserDAO.Exists(ctx, email.Email)
	if err != nil {
		return RegisterDTO{}, err
	}

	user := entity.NewUser(userID, email, password)
	uow := uc.UoWManager.GetUoW()

	tx, err := uow.Begin(ctx)
	if err != nil {
		return RegisterDTO{}, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = uow.Rollback(ctx)
			panic(p)
		}
		if err != nil {
			_ = uow.Rollback(ctx)
		}
	}()

	err = uc.UserDAO.Register(ctx, tx, user)
	if err != nil {
		return RegisterDTO{}, err
	}
	if err = uow.Commit(ctx); err != nil {
		return RegisterDTO{}, err
	}

	return RegisterDTO{
		UserID: userID.UserID.String(),
	}, err
}

func (uc *UserCommands) Login(ctx context.Context, dto LoginDTO) (LoginDTO, error) {
	email, err := vo.NewEmail(dto.Email)
	if err != nil {
		return LoginDTO{}, err
	}
	password, err := vo.NewPassword(dto.Password)
	if err != nil {
		return LoginDTO{}, err
	}
	err = uc.UserDAO.Exists(ctx, email.Email)
	if err != nil {
		return LoginDTO{}, err
	}

	loggingUser, err := uc.UserDAO.Load(ctx, email.Email) // TODO: uc.UserDAO.Load(ctx, email.Email)
	if err != nil {
		return LoginDTO{}, err
	}

	if !password.Matches(loggingUser.Password) {
		return LoginDTO{}, exceptions.InvalidCredentials
	}

	token, err := vo.NewToken(loggingUser.UserID.String())
	if err != nil {
		return LoginDTO{}, err
	}

	return LoginDTO{Token: token.Token}, nil
}

func (uc *UserCommands) IsAdmin(ctx context.Context, dto IsAdminDTO) (IsAdminDTO, error) {
	_, _ = ctx, dto
	return IsAdminDTO{}, nil
}
