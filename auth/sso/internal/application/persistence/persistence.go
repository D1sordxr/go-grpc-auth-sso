package persistence

import (
	"context"
)

type UoW interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Begin(ctx context.Context) (interface{}, error)
}

type UoWManager interface {
	GetUoW() UoW
}

type TokenServiceInterface interface {
	// TODO: GenerateNewToken()
	// TODO: ValidateToken() if needed
}
