package commands

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"
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

// TODO: add user_id as UUID and make new value object

func (uc *UserCommands) Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error) {
	email, err := vo.NewEmail(dto.Email)
	if err != nil {
		return RegisterDTO{}, err
	}
	password, err := vo.NewPassword(dto.Password)
	if err != nil {
		return RegisterDTO{}, err
	}
	user := entity.NewUser(email, password)

	err = uc.UserDAO.Exists(ctx, email.Email)
	if err != nil {
		return RegisterDTO{}, err
	}

	response, err := uc.UserDAO.Register(ctx, user)
	if err != nil {
		return RegisterDTO{}, err
	}

	return response, nil
}

func (uc *UserCommands) Login(ctx context.Context, dto LoginDTO) (LoginDTO, error) {
	_, _ = ctx, dto
	return LoginDTO{}, nil
}

func (uc *UserCommands) IsAdmin(ctx context.Context, dto IsAdminDTO) (IsAdminDTO, error) {
	_, _ = ctx, dto
	return IsAdminDTO{}, nil
}
