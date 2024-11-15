package dao

import (
	"aviasales/src/internal/db"
	"aviasales/src/internal/db/models"
	"context"
	"github.com/google/uuid"
	"time"
)

// TODO: fix import cycle

func CreateOrder(s *db.Storage, order models.Order) error {
	clientID := uuid.New()
	addressID := uuid.New()
	serialNumber := func() int {
		now := time.Now()
		return int(now.UnixNano() % 1000000000)
	}()

	tx, err := s.DB.Begin(context.Background())
	if err != nil {
		return err
	}

	var orderID int64
	err = tx.QueryRow(context.Background(), `
		INSERT INTO orders (
                client_id, 
                address_id,
                order_status,
                payment_method,
                serial_number,
                closed,
                deleted,
                created_at,
                updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
        RETURNING id
	`, clientID, addressID, "Created", order.PaymentMethod, serialNumber, false, false).Scan(&orderID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	for range order.Tickets {
		_, err = tx.Exec(context.Background(), `
			INSERT INTO tickets (order_id) VALUES ($1)
		`, orderID)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}
