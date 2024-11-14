package db

import (
	db "aviasales/src/internal/db/config"
	"context"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB *pgx.Conn
}

func NewDB(config *db.DBConfig) (*Storage, error) {
	dsn := config.ConnectionString()
	conn, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	if config.Migration {
		err = migrate(conn)
		if err != nil {
			return nil, err
		}
	}

	return &Storage{
		DB: conn,
	}, nil
}
