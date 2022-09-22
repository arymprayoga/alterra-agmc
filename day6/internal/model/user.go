package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (u *User) GenerateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"userId":     u.ID,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, err
}
