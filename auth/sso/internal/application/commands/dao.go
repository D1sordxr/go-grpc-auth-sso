package commands

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
)

type UserDAO interface {
	Register(ctx context.Context, tx interface{}, entity entity.User) error
	Exists(ctx context.Context, email string) error
	Load(ctx context.Context, email string) (entity.User, error)
}
