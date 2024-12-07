package exceptions

import "errors"

var (
	InvalidPasswordLength = errors.New("invalid password length")
	HashingError          = errors.New("invalid password")
	PasswordsDoNotMatch   = errors.New("password do not match")
)
