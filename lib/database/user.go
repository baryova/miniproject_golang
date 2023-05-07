package database

import (
	"miniproject_golang/config"
	"miniproject_golang/models"
)

// Get all user
func GetUser() (users []models.User, err error) {
	err = config.DB.Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	return
}

func CreateUser(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func LoginUser(user models.User) (models.User, error) {

	err := config.DB.Where("username = ?", user.Username).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
