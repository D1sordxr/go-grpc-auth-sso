package commands

import (
	"github.com/google/uuid"
)

type User struct {
	ID       int32     `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Password []byte    `json:"password"`
	AppID    int32     `json:"app_id"`
}

type RegisterDTO struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AppID    int32  `json:"app_id"`
}

type IsAdminDTO struct {
	UserID int32 `json:"user_id"`
}
