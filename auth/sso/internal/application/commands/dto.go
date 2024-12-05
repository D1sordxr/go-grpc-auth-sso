package commands

// TODO: all DTO's

type RegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AppID    int32  `json:"app_id"`
}

type IsAdminDTO struct {
	UserID int32 `json:"user_id"`
}
