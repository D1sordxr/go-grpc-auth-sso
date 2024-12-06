package persistence

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
)

type UserDAO interface {
	Register(ctx context.Context, dto commands.RegisterDTO) (commands.RegisterDTO, error)
}

type UOW struct {
}
