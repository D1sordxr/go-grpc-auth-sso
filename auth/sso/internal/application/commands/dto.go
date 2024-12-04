package commands

// TODO: all DTO's

type RegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AppID    int32  `json:"app_id"`
}
