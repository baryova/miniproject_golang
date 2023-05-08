package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id          int     `json:"id" form:"id"`
	Title       string  `json:"title" form:"title"`
	PosterPath  string  `json:"poster_path" form:"poster_path"`
	VoteAverage float32 `json:"vote_average" form:"vote_average"`
	VoteCount   int     `json:"vote_count" form:"vote_count"`
	Overview    string  `json:"overview" form:"overview"`
	ReleaseDate string  `json:"release_date" form:"release_date"`
}

type MovieResponse struct {
	Results []struct {
		Id          int     `json:"id" form:"id"`
		Title       string  `json:"title" form:"title"`
		PosterPath  string  `json:"poster_path" form:"poster_path"`
		VoteAverage float32 `json:"vote_average" form:"vote_average"`
		VoteCount   int     `json:"vote_count" form:"vote_count"`
		Overview    string  `json:"overview" form:"overview"`
		ReleaseDate string  `json:"release_date" form:"release_date"`
	}
}

type MovieDetailResponse struct {
	Id          int     `json:"id" form:"id"`
	Title       string  `json:"title" form:"title"`
	PosterPath  string  `json:"poster_path" form:"poster_path"`
	VoteAverage float32 `json:"vote_average" form:"vote_average"`
	VoteCount   int     `json:"vote_count" form:"vote_count"`
	Overview    string  `json:"overview" form:"overview"`
	ReleaseDate string  `json:"release_date" form:"release_date"`
	Genre       []Genre
}
