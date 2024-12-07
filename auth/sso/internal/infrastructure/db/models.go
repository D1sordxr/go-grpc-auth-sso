package db

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID    uuid.UUID
	Email     string
	Password  []byte
	CreatedAt time.Time
}
