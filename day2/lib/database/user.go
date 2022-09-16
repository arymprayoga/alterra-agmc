package database

import (
	"day2/config"
	"day2/models"
)

func GetUsers() ([]models.Users, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return users, e
	}

	return users, nil
}

func GetSingleUser(id int) (models.Users, error) {
	users := models.Users{}

	if e := config.DB.First(&users, id).Error; e != nil {
		return users, e
	}

	return users, nil
}

func CreateUser(user models.Users) (models.Users, error) {

	if e := config.DB.Create(&user).Error; e != nil {
		return user, e
	}

	return user, nil
}

func UpdateUser(user models.Users, id int) (models.Users, error) {
	updatedUser, _ := GetSingleUser(id)
	updatedUser.Name = user.Name
	updatedUser.Email = user.Email
	updatedUser.Password = user.Password

	if err := config.DB.Save(&updatedUser).Error; err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func DeleteUser(id int) error {

	if e := config.DB.Delete(&models.Users{}, id).Error; e != nil {
		return e
	}

	return nil
}
