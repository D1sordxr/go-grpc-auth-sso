package commands

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
)

type UserDAO interface {
	Register(ctx context.Context, tx interface{}, entity entity.User) (RegisterDTO, error)
	Exists(ctx context.Context, email string) error
}
