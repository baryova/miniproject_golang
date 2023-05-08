package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Id   int `json:"id" form:"id"`
	Name int `json:"name" form:"name"`
}
