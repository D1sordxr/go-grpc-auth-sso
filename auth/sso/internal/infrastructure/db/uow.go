package db

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"github.com/jackc/pgx/v5"
)

type UoW struct {
	Conn Connection
	Tx   pgx.Tx
}

// Begin starts a transaction.
func (u *UoW) Begin(ctx context.Context) (interface{}, error) {
	tx, err := u.Conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	u.Tx = tx
	return u.Tx, nil
}

// Commit commits the transaction.
func (u *UoW) Commit(ctx context.Context) error {
	return u.Tx.Commit(ctx)
}

// Rollback aborts the transaction.
func (u *UoW) Rollback(ctx context.Context) error {
	return u.Tx.Rollback(ctx)
}

type UoWManager struct {
	Conn Connection
}

// GetUoW returns a new UoW.
func (u *UoWManager) GetUoW() persistence.UoW {
	return &UoW{
		Conn: u.Conn,
	}
}

func NewUoWManager(conn Connection) *UoWManager {
	return &UoWManager{
		Conn: conn,
	}
}
