package db

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
)

func ConvertEntityToModel(user entity.User) User {
	return User{
		UserID:   user.UserID.UserID,
		Email:    user.Email.Email,
		Password: user.Password.HashedPassword,
	}
}
