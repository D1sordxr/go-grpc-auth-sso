package exceptions

import "errors"

var (
	InvalidEmailLength = errors.New("invalid email length")
	UserAlreadyExists  = errors.New("user already exists")
)
