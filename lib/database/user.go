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

func GetUsernameById(id int) (models.User, error) {
	user := models.User{}

	err := config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUserPassword(userInput models.User) (models.User, error) {
	user := models.User{}

	err := config.DB.Model(&user).Where("username = ?", userInput.Username).
		Update("password", userInput.Password).Error
	if err != nil {
		return models.User{}, err
	}

	return userInput, err
}

func DeleteUser(user models.User) error {
	err := config.DB.Where("username = ?", user.Username).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil

}
