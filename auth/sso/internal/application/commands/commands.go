package commands

type RegisterUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	AppID    int32  `json:"app_id" binding:"required"`
	Token    string `json:"token"`
}

type IsAdminUserCommand struct {
	UserID int64 `json:"user_id" binding:"required"`
}
