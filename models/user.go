package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Token    string `gorm:"-"`
}
