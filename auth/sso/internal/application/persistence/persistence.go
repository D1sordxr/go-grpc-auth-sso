package persistence

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
)

type UserDAO interface {
	Register(ctx context.Context, entity entity.User) (commands.RegisterDTO, error)
	Exists(ctx context.Context, email string) error
}

type UoW interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Begin(ctx context.Context) (interface{}, error)
}

type UoWManager interface {
	GetUoW() UoW
}
