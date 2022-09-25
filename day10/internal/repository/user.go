package repository

import (
	"day10/database"

	"day10/internal/model"
)

func GetAllUser() (interface{}, error) {
	var users []model.User

	if e := database.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(id uint) (interface{}, error) {
	user := model.User{}
	if user := database.DB.Where("id = ?", id).First(&user); user.Error != nil {
		return nil, user.Error
	}
	return user, nil
}
