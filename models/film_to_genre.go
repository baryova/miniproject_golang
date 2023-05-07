package models

import "gorm.io/gorm"

type FilmToGenre struct {
	gorm.Model
	FilmId  int `json:"film_id" form:"film_id"`
	GenreId int `json:"genre_id" form:"genre_id"`
}
