package db

import (
	"context"
	"errors"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"github.com/jackc/pgx/v5"
)

type UserDAO struct {
	Conn Connection
}

func NewUserDAO(conn Connection) *UserDAO {
	return &UserDAO{Conn: conn}
}

func (dao *UserDAO) Register(ctx context.Context, tx interface{}, entity entity.User) error {
	user := ConvertEntityToModel(entity)

	conn := tx.(pgx.Tx)
	query := `INSERT INTO users (user_id, email, password, created_at)
					VALUES ($1, $2, $3, NOW())`

	_, err := conn.Exec(ctx, query, user.UserID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDAO) Exists(ctx context.Context, email string) error {
	var userEmail string
	query := "SELECT email FROM users WHERE email = $1"

	err := dao.Conn.QueryRow(ctx, query, email).Scan(&userEmail)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if userEmail == email {
		return exceptions.UserAlreadyExists
	}

	return nil
}

func (dao *UserDAO) Load(ctx context.Context, email string) (commands.User, error) {
	password := []byte("a_password")
	return commands.User{Password: password}, nil
}
