package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	GenreId int `json:"genre_id" form:"genre_id"`
	Name    int `json:"name" form:"name"`
}
