package repository

import (
	"day10/database"
	"day10/internal/model"
)

func FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if user := database.DB.Where("email = ?", email).First(&user); user.Error != nil {
		return nil, user.Error
	}
	return &user, nil
}
