package models

import "gorm.io/gorm"

type UserFavorite struct {
	gorm.Model
	UserId int `json:"user_id" form:"user_id"`
	FilmId int `json:"film_id" form:"film_id"`
}
