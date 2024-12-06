package db

import "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"

func ConvertRegisterDTOToUserModel(dto commands.RegisterDTO) User {
	return User{
		Email:  dto.Email,
		UserID: dto.UserID,
		// Password:   dto.Password,
	}
}
