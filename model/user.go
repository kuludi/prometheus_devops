package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255)" form:"user" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(255)"  form:"password" binding:"required"`
	Phone    string `json:"phone" gorm:"type:varchar(255)" form:"phone" binding:"required"`
}

