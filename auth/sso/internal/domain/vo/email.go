package vo

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/exceptions"
	"unicode/utf8"
)

type Email struct {
	Email string
}

// NewEmail creates new Email value object
func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, exceptions.InvalidEmailLength
	}
	if utf8.RuneCountInString(email) > 255 {
		return Email{}, exceptions.InvalidEmailLength
	}
	return Email{Email: email}, nil
}
