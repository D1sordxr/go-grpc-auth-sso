package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func migrate(conn *pgx.Conn) error {
	err := createOrderTable(conn)
	if err != nil {
		log.Fatalf("Failed migrations: %s", err)
	}

	err = createTicketTable(conn)
	if err != nil {
		log.Fatalf("Failed migrations: %s", err)
	}

	return nil
}

func createOrderTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            client_id UUID NOT NULL,
            address_id UUID NOT NULL,
            order_status VARCHAR NOT NULL,
            payment_method VARCHAR NOT NULL,
            serial_number INTEGER NOT NULL,
            closed BOOLEAN NOT NULL,
            deleted BOOLEAN DEFAULT FALSE,
            created_at TIMESTAMP DEFAULT NOW(),
            updated_at TIMESTAMP DEFAULT NOW()
	);
	`)
	if err != nil {
		return fmt.Errorf("error creating order table: %w", err)
	}

	_, err = conn.Exec(context.Background(),
		`CREATE INDEX IF NOT EXISTS idx_order_number ON orders(serial_number);
	`)
	if err != nil {
		return fmt.Errorf("error creating order index: %w", err)
	}

	return nil
}

func createTicketTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS tickets (
			id SERIAL PRIMARY KEY,
			order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
			passenger_name TEXT NOT NULL,
			destination TEXT NOT NULL,
			payment INTEGER NOT NULL,
			dispatch_time TIMESTAMP NOT NULL,
			arrival_time TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()                    
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating ticket table: %w", err)
	}

	_, err = conn.Exec(context.Background(), `
		CREATE INDEX IF NOT EXISTS idx_ticket_passenger ON tickets(passenger_name);
	`)
	if err != nil {
		return fmt.Errorf("error creating ticket index: %w", err)
	}

	return nil
}
