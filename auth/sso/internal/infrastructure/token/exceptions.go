package token

import "errors"

var (
	ErrorCreatingToken = errors.New("error creating jwt token")
	InvalidUserID      = errors.New("invalid user id")
	InvalidAppID       = errors.New("invalid app id")
)
