package db

import (
	"context"
	db "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/db/config"
	"github.com/jackc/pgx/v5"
)

type Connection struct {
	*pgx.Conn
}

func NewConnection(config *db.DBConfig) Connection {
	conn, err := pgx.Connect(context.Background(), config.ConnectionString())
	if err != nil {
		panic(err)
	}
	return Connection{Conn: conn}
}
