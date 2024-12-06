package db

import "time"

type User struct {
	Email     string
	Password  []byte
	UserID    int64
	CreatedAt time.Time
}
