package models

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ActivateUserAccountRequest struct {
	Email string `json:"email" binding:"required,email"`
	Otp   int    `json:"otp" binding:"required"`
}
