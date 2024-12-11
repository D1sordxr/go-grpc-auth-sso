package entity

import "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"

type User struct {
	UserID   vo.UserID
	Email    vo.Email
	Password vo.Password
	AppID    int32
}

func NewUser(userID vo.UserID, email vo.Email, password vo.Password) User {
	return User{
		UserID:   userID,
		Email:    email,
		Password: password,
	}
}

func (u *User) StringUserID() string {
	return u.UserID.UserID.String()
}

func (u *User) ValidatePassword(plainPassword []byte) error {
	return u.Password.Matches(plainPassword)
}
