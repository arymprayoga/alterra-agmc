package dto

import "day10/internal/model"

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	model.User
}
