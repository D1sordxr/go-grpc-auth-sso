package vo

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength = 8

type Password struct {
	Password       []byte
	HashedPassword []byte
}

func NewPassword(password string) (Password, error) {
	if len(password) < minPasswordLength {
		return Password{}, exceptions.InvalidPasswordLength
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, exceptions.HashingError
	}
	return Password{HashedPassword: hashedPassword}, nil
}

func NewBytePassword(password string) (Password, error) {
	if len(password) < minPasswordLength {
		return Password{}, exceptions.InvalidPasswordLength
	}

	bytePassword := []byte(password)
	return Password{Password: bytePassword}, nil
}

func (p Password) Matches(plainPassword []byte) error {
	if len(plainPassword) < minPasswordLength {
		return exceptions.InvalidPasswordLength
	}

	err := bcrypt.CompareHashAndPassword(p.HashedPassword, plainPassword)
	if err != nil {
		return exceptions.InvalidCredentials
	}
	return nil
}

func (p Password) GetHashedPassword() []byte {
	return p.HashedPassword
}
