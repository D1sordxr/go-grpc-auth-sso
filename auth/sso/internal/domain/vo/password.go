package vo

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength = 8

type Password struct {
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

// Matches and GetHashedPassword methods for tests
func (p Password) Matches(plainPassword []byte) bool {
	if len(plainPassword) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword(p.HashedPassword, plainPassword)
	return err == nil
}

func (p Password) GetHashedPassword() []byte {
	return p.HashedPassword
}
