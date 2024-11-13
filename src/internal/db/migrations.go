package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func migrate(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS ticket(
		id SERIAL PRIMARY KEY,
		passenger_name TEXT NOT NULL,
		destination TEXT NOT NULL,
		payment INTEGER NOT NULL,
		dispatch_time TIMESTAMPTZ NOT NULL,
		arrival_time TIMESTAMPTZ NOT NULL
	);
	`)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	_, err = conn.Exec(context.Background(),
		`CREATE INDEX IF NOT EXISTS idx_ticket_passenger ON ticket(passenger_name);
	`)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}

	return nil
}
