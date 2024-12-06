package db

type User struct {
	Email    string
	Password []byte
	UserID   int64
}
