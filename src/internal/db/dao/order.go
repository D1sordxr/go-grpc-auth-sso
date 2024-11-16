package dao

import (
	"aviasales/src/internal/db/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

type OrderDAO struct {
	DB *pgx.Conn
}

func NewOrderDAO(conn *pgx.Conn) *OrderDAO {
	return &OrderDAO{DB: conn}
}

func (dao *OrderDAO) CreateOrder(order models.Order) error {
	clientID := uuid.New()
	addressID := uuid.New()
	serialNumber := func() int {
		now := time.Now()
		return int(now.UnixNano() % 1000000000)
	}()

	tx, err := dao.DB.Begin(context.Background())
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

	for _, v := range order.Tickets {
		_, err = tx.Exec(context.Background(), `
			UPDATE tickets 
			SET order_id = $1
			WHERE id = $2
		`, orderID, v.ID)
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
