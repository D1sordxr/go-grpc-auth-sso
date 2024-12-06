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

func (dao *UserDAO) Register(ctx context.Context, tx interface{}, entity entity.User) (commands.RegisterDTO, error) {
	user := ConvertEntityToModel(entity)

	conn := tx.(pgx.Tx)
	query := `INSERT INTO users (email, password, created_at)
					VALUES ($1, $2, NOW())
				RETURNING id`

	row := conn.QueryRow(ctx, query, user.Email, user.Password)
	err := row.Scan(
		&user.UserID,
	)
	if err != nil {
		return commands.RegisterDTO{}, err
	}

	return commands.RegisterDTO{UserID: user.UserID}, nil
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
