package vo

import "github.com/google/uuid"

type UserID struct {
	UserID uuid.UUID
}

func NewUserID() UserID {
	return UserID{UserID: uuid.New()}
}
