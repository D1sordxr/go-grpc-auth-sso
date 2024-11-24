package db

import (
	"context"
	db "github.com/D1sordxr/aviasales/src/internal/db/config"
	"github.com/D1sordxr/aviasales/src/internal/db/dao"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB        *pgx.Conn
	OrderDAO  *dao.OrderDAO
	TicketDAO *dao.TicketDAO
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

	orderDAO := dao.NewOrderDAO(conn)
	ticketDAO := dao.NewTicketDAO(conn)

	return &Storage{
		DB:        conn,
		OrderDAO:  orderDAO,
		TicketDAO: ticketDAO,
	}, nil
}
