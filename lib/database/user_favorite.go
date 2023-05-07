package database

import (
	"miniproject_golang/config"
	"miniproject_golang/models"
)

func GetUserFavorite(userId int) ([]models.UserFavorite, error) {
	var userFav []models.UserFavorite

	err := config.DB.Where("user_id = ?", userId).Find(&userFav).Error

	if err != nil {
		return nil, err
	}

	return userFav, nil
}

func GetSameUserFavorite(userFav models.UserFavorite) (bool, error) {
	row := config.DB.
		Where("film_id = ? AND user_id = ?", userFav.FilmId, userFav.UserId).
		First(&userFav)
	if row.RowsAffected < 1 {
		return false, nil
	} else if row.RowsAffected >= 1 {
		return true, nil
	} else {
		return false, row.Error
	}
}

func DoFavorite(userFav models.UserFavorite) (models.UserFavorite, error) {
	err := config.DB.Create(&userFav).Error
	if err != nil {
		return models.UserFavorite{}, err
	}

	return userFav, nil
}

func UndoFavorite(userFav models.UserFavorite) (models.UserFavorite, error) {
	err := config.DB.Debug().
		Where("film_id = ? AND user_id = ?", userFav.FilmId, userFav.UserId).
		Unscoped().Delete(&userFav).Error
	if err != nil {
		return models.UserFavorite{}, err
	}

	return userFav, nil
}
