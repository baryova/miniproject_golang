package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	FilmId      int    `json:"film_id" form:"film_id"`
	Title       string `json:"title" form:"title"`
	PosterPath  string `json:"poster_path" form:"poster_path"`
	VoteAverage string `json:"vote_average" form:"vote_average"`
	VoteCount   string `json:"vote_count" form:"vote_count"`
	Overview    string `json:"overview" form:"overview"`
	ReleaseDate string `json:"release_date" form:"release_date"`
}
