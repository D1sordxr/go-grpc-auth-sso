package exceptions

import "errors"

var (
	InvalidPasswordLength = errors.New("invalid password length")
	HashingError          = errors.New("invalid password")
	InvalidCredentials    = errors.New("invalid credentials")
)
