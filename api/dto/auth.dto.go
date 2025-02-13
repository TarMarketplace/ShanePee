package dto

type LoginBody struct {
	Email    string `json:"email" example:"johndoe@example.com"`
	Password string `json:"password" example:"VerySecurePassword"`
}

type ChangePasswordBody struct {
	RequestID   int64  `json:"request_id"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type RequestPasswordChangeBody struct {
	Email string `json:"email" example:"johndoe@example.com"`
}
