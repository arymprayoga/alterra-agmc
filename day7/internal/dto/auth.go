package dto

import "day7/internal/model"

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	model.User
}
