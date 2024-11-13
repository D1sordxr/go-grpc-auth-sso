package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func NewDB(c string) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), c)
	if err != nil {
		return nil, err
	}

	if err = migrate(conn); err != nil {
		return nil, err
	}

	return conn, nil
}
