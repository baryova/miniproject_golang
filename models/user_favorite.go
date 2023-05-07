package models

import "gorm.io/gorm"

type UserFavorite struct {
	gorm.Model
	UserId int `json:"user_id" form:"user_id"`
	FilmId int `json:"FilmId" form:"FilmId"`
}
