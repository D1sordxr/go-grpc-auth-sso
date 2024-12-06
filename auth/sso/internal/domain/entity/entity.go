package entity

import "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"

type User struct {
	Email    vo.Email
	Password vo.Password
}

func NewUser(email vo.Email, password vo.Password) User {
	return User{
		Email:    email,
		Password: password,
	}
}
