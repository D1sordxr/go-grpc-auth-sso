package db

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
)

type UserDAO struct {
	Conn Connection
}

func NewUserDAO(conn Connection) *UserDAO {
	return &UserDAO{Conn: conn}
}

func (dao *UserDAO) Register(ctx context.Context, entity entity.User) (commands.RegisterDTO, error) {
	err := dao.Conn.QueryRow(ctx, `
	
	`).Scan()
	if err != nil {
		return commands.RegisterDTO{}, err
	}
	return commands.RegisterDTO{}, nil
}
