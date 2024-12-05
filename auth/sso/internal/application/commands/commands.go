package commands

import "context"

type Auth interface {
	Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error)
	Login(ctx context.Context, dto LoginDTO) (LoginDTO, error)
	IsAdmin(ctx context.Context, dto IsAdminDTO) (IsAdminDTO, error)
}

type UserCommands struct {
	// access to storage
}

func NewUserCommands() *UserCommands {
	return &UserCommands{}
}

func (uc *UserCommands) Register(ctx context.Context, dto RegisterDTO) (RegisterDTO, error) {
	_, _ = ctx, dto
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
