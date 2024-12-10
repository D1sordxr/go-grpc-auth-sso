package vo

import "github.com/google/uuid"

type UserID struct {
	UserID       uuid.UUID
	StringUserID string
}

func NewUserID() UserID {
	return UserID{UserID: uuid.New()}
}

func StringUserID(id uuid.UUID) UserID {
	return UserID{StringUserID: id.String()}
}
